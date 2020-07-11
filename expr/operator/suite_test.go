package operator_test

import (
	"errors"
	"testing"

	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/internal/test/runner"
	"github.com/weworksandbox/lingo/sql"
)

//go:generate pegomock generate github.com/weworksandbox/lingo Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo -m Expression
func TestOperator(t *testing.T) {
	runner.SetupAndRunUnit(t, "operator", "unit")
}

type operatorDialectSuccess struct{}

func (operatorDialectSuccess) GetName() string { return "operator success" }

func (s operatorDialectSuccess) UnaryOperator(sql.Data, operator.Operator) (sql.Data, error) {
	return sql.New("operator sql", []interface{}{5555}), nil
}

func (s operatorDialectSuccess) BinaryOperator(sql.Data, operator.Operator, sql.Data) (sql.Data, error) {
	return sql.New("operator sql", []interface{}{5555}), nil
}

func (s operatorDialectSuccess) VariadicOperator(sql.Data, operator.Operator, []sql.Data) (sql.Data, error) {
	return sql.New("operator sql", []interface{}{5555}), nil
}

type operatorDialectFailure struct{ operatorDialectSuccess }

func (f operatorDialectFailure) UnaryOperator(sql.Data, operator.Operator) (sql.Data, error) {
	return nil, errors.New("operator failure")
}

func (f operatorDialectFailure) BinaryOperator(sql.Data, operator.Operator, sql.Data) (sql.Data, error) {
	return nil, errors.New("operator failure")
}

func (f operatorDialectFailure) VariadicOperator(sql.Data, operator.Operator, []sql.Data) (sql.Data, error) {
	return nil, errors.New("operator failure")
}
