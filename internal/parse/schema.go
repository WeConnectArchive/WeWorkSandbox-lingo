package parse

import (
	"database/sql"
)

type Column struct {
	columnType *sql.ColumnType
	table      string
}

func (c Column) Name() string          { return c.columnType.Name() }
func (c Column) Table() string         { return c.table }
func (c Column) Type() *sql.ColumnType { return c.columnType }

type ForeignKey struct {
	name string
}

func (c ForeignKey) Name() string { return c.name }

// func (c *ForeignKey) Columns() []*Column { return c.columns }
