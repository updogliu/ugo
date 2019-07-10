package usql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

var (
	NotFound = errors.New("Not Found")
)

// Implemented by `*sql.Row`, `*sql.Rows`.
type RowScanner interface {
	Scan(dest ...interface{}) error
}

// Implemented by `*sql.DB`, `*sql.Stmt`, `*sql.Tx`.
type Querier interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// Used to tag a public struct field with a corresponding DB table column.
// See the `select_test.go` for an example.
const SQLCOL = "sqlcol"

// An item that can be constructed by scanning a row in a DB query result.
type DBItem interface {
	ScanRow(row RowScanner) error
}

// Query a single item.
//
// `item` is a pointer to the item struct to hold the result.
// `condition` goes to the WHERE clause. Empty string means querying all.
// Returns `NotFound` if no such an record in DB.
//
// Example:
// ```
//   var s Student
//   err := SelectItem(db, &s, "student_table", "score >= ?", 60)
// ```
func SelectItem(stub Querier, item DBItem, table string, condition string, args ...interface{}) error {
	row := stub.QueryRowContext(newQueryCtx(), selectSql(item, table, condition), args...)
	err := item.ScanRow(row)
	if err == sql.ErrNoRows {
		return NotFound
	}
	return err
}

// Query items.
//
// `item` serves as a "type parameter". Use any non-nil pointer of item. Its value is unchanged.
// `condition` goes to the WHERE clause.
// Returns `NotFound` if no such an record in DB.
//
// Example:
// ```
//   items, err := SelectItems(db, &Student{}, "student_table", "score >= ?", 60)
//   if err != nil {
//     ...handles err
//   }
//   var s []*Student
//   for _, item := range items {
//     s = append(s, item.(*Student))
//   }
// ```
func SelectItems(stub Querier, item DBItem, table string, condition string, args ...interface{}) (
	[]DBItem, error) {
	rows, err := stub.QueryContext(newQueryCtx(), selectSql(item, table, condition), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []DBItem
	for rows.Next() {
		v := newItem(item)
		if err := v.ScanRow(rows); err != nil {
			return nil, err
		}
		items = append(items, v)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// Extract column names from struct tag `sqlcol`
// `item` is a pointer to the item. Only the type is significant. Any non-nil value will work.
func colStr(item DBItem) string {
	t := reflect.TypeOf(item).Elem()
	var cols []string
	for i := 0; i < t.NumField(); i++ {
		col := t.Field(i).Tag.Get(SQLCOL)
		if col != "" {
			cols = append(cols, col)
		}
	}

	if len(cols) != 0 {
		return strings.Join(cols, ",")
	}
	return ""
}

func selectSql(item DBItem, table string, condition string) string {
	sql := fmt.Sprintf("SELECT %s FROM %s", colStr(item), table)
	if condition != "" {
		sql += fmt.Sprintf(" WHERE %s", condition)
	}
	return sql
}

// Creates a new item instance.
// `item` is a pointer to the item. Only the type is significant. Any non-nil value will work.
func newItem(item DBItem) DBItem {
	return reflect.New(reflect.ValueOf(item).Elem().Type()).Interface().(DBItem)
}

func newQueryCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}
