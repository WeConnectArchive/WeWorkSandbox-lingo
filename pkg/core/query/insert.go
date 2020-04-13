package query

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/helpers"
)

func InsertInto(entity core.Table) *InsertQuery {
	insert := InsertQuery{}
	insert.table = entity
	return &insert
}

type InsertQuery struct {
	table      core.Table
	columns    []core.Expression
	values     []core.Expression
	selectPart core.Expression
}

func (i *InsertQuery) Columns(columns ...core.Column) *InsertQuery {
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
		i.values = append(i.values, expression.NewValue(value))
	}
	return i
}

// Values allows inserting expressions with SQL functions:
//
// INSERT INTO table1 (uuid, name)
// values (UNHEX("1234567891234"), 'name1');
func (i *InsertQuery) Values(values ...core.Expression) *InsertQuery {
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
func (i *InsertQuery) Select(s *SelectQuery) *InsertQuery {
	i.selectPart = s
	return i
}

func (i *InsertQuery) GetSQL(d core.Dialect) (core.SQL, error) {
	var sql = core.NewSQL("INSERT INTO", nil)

	if helpers.IsValueNilOrBlank(i.table) {
		return nil, expression.ErrorAroundSql(expression.ExpressionIsNil("table"), sql.String())
	}
	if i.table.GetAlias() != "" {
		return nil, expression.ErrorAroundSql(errors.New("table alias must be unset"), sql.String())
	}
	if tableSql, err := i.table.GetSQL(d); err != nil {
		return nil, expression.ErrorAroundSql(err, sql.String())
	} else {
		sql = sql.AppendSqlWithSpace(tableSql)
	}

	if helpers.IsValueNilOrEmpty(i.columns) {
		return nil, expression.ErrorAroundSql(expression.ExpressionCannotBeEmpty("columns"), sql.String())
	}
	if pathsSql, err := CombinePathSQL(d, i.columns); err != nil {
		return nil, expression.ErrorAroundSql(err, sql.String())
	} else {
		sql = sql.AppendSqlWithSpace(pathsSql.SurroundWithParens())
	}

	if i.selectPart != nil {
		if selectSql, err := i.selectPart.GetSQL(d); err != nil {
			return nil, expression.ErrorAroundSql(err, sql.String())
		} else if selectSql.String() != "" {
			sql = sql.AppendSqlWithSpace(selectSql)
		}
	} else {
		colsLen := len(i.columns)
		valuesLen := len(i.values)
		if colsLen != valuesLen {
			err := fmt.Errorf("column count %d does not match values count %d", colsLen, valuesLen)
			return nil, expression.ErrorAroundSql(err, sql.String())
		}

		if valuesSql, err := CombinePathSQL(d, i.values); err != nil {
			return nil, expression.ErrorAroundSql(err, sql.String())
		} else if valuesSql.String() != "" {
			sql = sql.AppendStringWithSpace("VALUES").AppendSqlWithSpace(valuesSql.SurroundWithParens())
		}
	}

	return sql, nil
}
