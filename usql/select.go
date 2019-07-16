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
	sqlStr := selectSql(item, table, condition, "")
	row := stub.QueryRowContext(newQueryCtx(), sqlStr, args...)
	err := item.ScanRow(row)
	if err == sql.ErrNoRows {
		return NotFound
	}
	return err
}

// Query items.
//
// - `item` serves as a "type parameter". Use any non-nil pointer of item. Its value is unchanged.
// - `orderBy` goes to the ORDER BY clause. Empty string means not ordered.
// - `condition` goes to the WHERE clause. Empty string means no condition.
// Returns `NotFound` if no such an record in DB.
//
// Example:
// ```
//   items, err := SelectItems(db, &Student{}, "student_table", "id DESC", "score >= ?", 60)
//   if err != nil {
//     ...handles err
//   }
//   var s []*Student
//   for _, item := range items {
//     s = append(s, item.(*Student))
//   }
// ```
func SelectItems(stub Querier, item DBItem, table string, orderBy string,
	condition string, args ...interface{}) ([]DBItem, error) {

	sqlStr := selectSql(item, table, condition, orderBy)
	rows, err := stub.QueryContext(newQueryCtx(), sqlStr, args...)
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

func SelectCount(stub Querier, table string, condition string, args ...interface{}) (int64, error) {
	sql := "SELECT COUNT(*) FROM " + table
	if condition != "" {
		sql += fmt.Sprintf(" WHERE %s", condition)
	}
	row := stub.QueryRowContext(newQueryCtx(), sql, args...)

	var count int64
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
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

func selectSql(item DBItem, table string, condition string, orderBy string) string {
	sql := fmt.Sprintf("SELECT %s FROM %s", colStr(item), table)
	if condition != "" {
		sql += fmt.Sprintf(" WHERE %s", condition)
	}
	if orderBy != "" {
		sql += " ORDER BY " + orderBy
	}
	return sql
}

// Creates a new item instance.
// `item` is a pointer to the item. Only the type is significant. Any non-nil value will work.
func newItem(item DBItem) DBItem {
	return reflect.New(reflect.ValueOf(item).Elem().Type()).Interface().(DBItem)
}

func newQueryCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	return ctx
}
