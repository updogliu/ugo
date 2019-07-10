package usql

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

type Student struct {
	Name     string         `sqlcol:"name"`
	Score    int64          `sqlcol:"score"`
	HomeAddr sql.NullString `sqlcol:"home_addr"`
}

// Implement the `DBItem` interface.
func (s *Student) ScanRow(row RowScanner) error {
	// Must list the `sqlcol` tagged fields in order .
	return row.Scan(&s.Name, &s.Score, &s.HomeAddr)
}

func TestColStr(t *testing.T) {
	expected := "name,score,home_addr"
	require.Equal(t, expected, colStr(&Student{}))
}

func TestSelectSql(t *testing.T) {
	expected := "SELECT name,score,home_addr FROM students WHERE score > ?"
	require.Equal(t, expected, selectSql(&Student{}, "students", "score > ?"))
}
