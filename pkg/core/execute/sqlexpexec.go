package execute

import (
	"context"
	"database/sql"

	"github.com/weworksandbox/lingo/pkg/core"
)

type ExecSQLExpInTx = func(ctx context.Context, s ExpQuery) error

type SQLExp interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (TxSQLExp, error)
	InTx(ctx context.Context, opts *sql.TxOptions, execThis ExecSQLExpInTx) error
	ExpQuery
}

type TxSQLExp interface {
	CommitOrRollback(ctx context.Context, err error) error
	Rollback(ctx context.Context, err error) error
	ExpQuery
}

type ExpQuery interface {
	Exec(ctx context.Context, exp core.Expression) (sql.Result, error)
	Query(ctx context.Context, exp core.Expression) (RowScanner, error)
	QueryRow(ctx context.Context, exp core.Expression, valuePtrs ...interface{}) error
}

func NewSQLExp(s SQL, d core.Dialect) SQLExp {
	return sqlExpExec{
		exec: s,
		d:    d,
	}
}

type sqlExpExec struct {
	exec SQL
	d    core.Dialect
}

func (sqlExec sqlExpExec) BeginTx(ctx context.Context, opts *sql.TxOptions) (TxSQLExp, error) {
	txSQL, err := sqlExec.exec.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return NewTxSQLExp(txSQL, sqlExec.d), nil
}

func (sqlExec sqlExpExec) InTx(ctx context.Context, opts *sql.TxOptions, execThis ExecSQLExpInTx) (err error) {
	var txSQL TxSQLExp
	txSQL, err = sqlExec.BeginTx(ctx, opts)
	if err != nil {
		return err // Already Traced
	}

	defer func() {
		r := recover()
		err = txSQL.CommitOrRollback(ctx, err)
		if r != nil {
			panic(r)
		}
	}()
	err = execThis(ctx, txSQL)
	return
}

func (sqlExec sqlExpExec) Exec(ctx context.Context, exp core.Expression) (result sql.Result, err error) {
	return execExp(ctx, sqlExec.exec, sqlExec.d, exp)
}

func (sqlExec sqlExpExec) QueryRow(ctx context.Context, exp core.Expression, queryIntoPtrs ...interface{}) error {
	return queryRowExp(ctx, sqlExec.exec, sqlExec.d, exp, queryIntoPtrs)
}

func (sqlExec sqlExpExec) Query(ctx context.Context, exp core.Expression) (RowScanner, error) {
	return queryExp(ctx, sqlExec.exec, sqlExec.d, exp)
}

func NewTxSQLExp(s TxSQL, d core.Dialect) TxSQLExp {
	return sqlExpTxExec{
		exec: s,
		d:    d,
	}
}

type sqlExpTxExec struct {
	exec TxSQL
	d    core.Dialect
}

func (txExec sqlExpTxExec) CommitOrRollback(ctx context.Context, err error) error {
	return txExec.exec.CommitOrRollback(ctx, err)
}

func (txExec sqlExpTxExec) Rollback(ctx context.Context, err error) error {
	return txExec.exec.Rollback(ctx, err)
}

func (txExec sqlExpTxExec) Exec(ctx context.Context, exp core.Expression) (result sql.Result, err error) {
	return execExp(ctx, txExec.exec, txExec.d, exp)
}

func (txExec sqlExpTxExec) QueryRow(ctx context.Context, exp core.Expression, queryIntoPtrs ...interface{}) error {
	return queryRowExp(ctx, txExec.exec, txExec.d, exp, queryIntoPtrs)
}

func (txExec sqlExpTxExec) Query(ctx context.Context, exp core.Expression) (RowScanner, error) {
	return queryExp(ctx, txExec.exec, txExec.d, exp)
}

func expandSQL(dialect core.Dialect, exp core.Expression) (string, []interface{}, error) {
	lSQL, err := exp.GetSQL(dialect)
	if err != nil {
		return "", nil, err
	}
	return lSQL.String(), lSQL.Values(), nil
}

func execExp(ctx context.Context, db SQLQuery, d core.Dialect, exp core.Expression) (result sql.Result, err error) {
	var tSQL string
	var sVals []interface{}
	tSQL, sVals, err = expandSQL(d, exp)
	if err != nil {
		return nil, traceErr(ctx, err)
	}
	return db.Exec(ctx, tSQL, sVals...)
}

func queryRowExp(
	ctx context.Context, e SQLQuery, d core.Dialect, exp core.Expression, queryIntoPtrs []interface{},
) error {
	tSQL, sVals, err := expandSQL(d, exp)
	if err != nil {
		return traceErr(ctx, err)
	}

	return e.QueryRow(ctx, tSQL, sVals, queryIntoPtrs...)
}

func queryExp(ctx context.Context, db SQLQuery, d core.Dialect, exp core.Expression) (RowScanner, error) {
	tSQL, sVals, err := expandSQL(d, exp)
	if err != nil {
		return nil, traceErr(ctx, err)
	}

	return db.Query(ctx, tSQL, sVals...)
}
