package query

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/sql"
)

func InsertInto(entity lingo.Table) *InsertQuery {
	insert := InsertQuery{}
	insert.table = entity
	return &insert
}

type InsertQuery struct {
	table      lingo.Table
	columns    []lingo.Expression
	values     []lingo.Expression
	selectPart lingo.Expression
}

func (i *InsertQuery) Columns(columns ...lingo.Column) *InsertQuery {
	// TODO - validate we actually want to do this with our insert columns...
	i.columns = append(i.columns, convertToStringColumns(columns)...)
	return i
}

// ValuesConstants allows inserting constant values:
//
// INSERT INTO table1 (id, name, internal_name)
// values (123456, 'name1', 'internal_name');
func (i *InsertQuery) ValuesConstants(values ...interface{}) *InsertQuery {
	for _, value := range values {
		i.values = append(i.values, expr.NewValue(value))
	}
	return i
}

// Values allows inserting expressions with SQL functions:
//
// INSERT INTO table1 (uuid, name)
// values (UNHEX("1234567891234"), 'name1');
func (i *InsertQuery) Values(values ...lingo.Expression) *InsertQuery {
	i.values = append(i.values, values...)
	return i
}

// Select allows passing in a Select Query the following insert statements:
//
// INSERT INTO table1 (uuid, name)
// SELECT UNHEX("1234567891230"), a.name
// FROM table2 as a
// LEFT JOIN table1 as b ON a.name=b.remote_name
// WHERE b.remote_name = 'other_table';
func (i *InsertQuery) Select(q *SelectQuery) *InsertQuery {
	i.selectPart = q
	return i
}

func (i InsertQuery) ToSQL(d lingo.Dialect) (sql.Data, error) {
	var s = sql.String("INSERT INTO")

	if check.IsValueNilOrBlank(i.table) {
		return nil, NewErrAroundSQL(s, errors.New("table cannot be empty"))
	}
	if i.table.GetAlias() != "" {
		return nil, NewErrAroundSQL(s, errors.New("table alias must be unset"))
	}
	tableSQL, err := i.table.ToSQL(d)
	if err != nil {
		return nil, NewErrAroundSQL(s, err)
	}
	s = s.AppendWithSpace(tableSQL)

	if check.IsValueNilOrEmpty(i.columns) {
		return nil, NewErrAroundSQL(s, errors.New("expr 'columns' cannot be empty"))
	}
	pathsSQL, err := JoinToSQL(d, sepPathComma, i.columns)
	if err != nil {
		return nil, ErrAroundSQL{err: err, sqlStr: s.String()}
	}
	s = s.SurroundAppend(" (", ")", pathsSQL) // Include space before first paren!

	if i.selectPart != nil {
		s, err = i.buildSelectFrom(d, s)
		if err != nil {
			return nil, err
		}
	} else {
		s, err = i.buildValues(d, s)
		if err != nil {
			return s, err
		}
	}

	return s, nil
}

func (i InsertQuery) buildSelectFrom(d lingo.Dialect, s sql.Data) (sql.Data, error) {
	selectSQL, err := i.selectPart.ToSQL(d)
	if err != nil {
		return nil, ErrAroundSQL{err: err, sqlStr: s.String()}
	}
	if selectSQL.String() != "" {
		s = s.AppendWithSpace(selectSQL)
	}
	return s, nil
}

func (i InsertQuery) buildValues(d lingo.Dialect, s sql.Data) (sql.Data, error) {
	colsLen := len(i.columns)
	valuesLen := len(i.values)
	if colsLen != valuesLen {
		err := fmt.Errorf("column count %d does not match values count %d", colsLen, valuesLen)
		return nil, ErrAroundSQL{err: err, sqlStr: s.String()}
	}

	valuesSQL, err := JoinToSQL(d, sepPathComma, i.values)
	if err != nil {
		return nil, ErrAroundSQL{err: err, sqlStr: s.String()}
	}
	if valuesSQL.String() != "" {
		s = s.AppendWithSpace(sql.String("VALUES")).SurroundAppend(" (", ")", valuesSQL)
	}
	return s, nil
}
