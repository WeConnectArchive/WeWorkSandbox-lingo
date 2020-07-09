package query

import (
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func NewErrAroundSQL(s sql.Data, err error) error {
	return ErrAroundSQL{
		err:    err,
		sqlStr: s.String(),
	}
}

type ErrAroundSQL struct {
	err    error
	sqlStr string
}

func (e ErrAroundSQL) Error() string {
	return fmt.Sprintf("an error occurred around sqlStr '%s': %s", e.lastChars(), e.Unwrap().Error())
}

func (e ErrAroundSQL) SQL() string   { return e.sqlStr }
func (e ErrAroundSQL) Unwrap() error { return e.err }

func (e ErrAroundSQL) lastChars() string {
	const length = 20
	var sqlLen = len(e.SQL())
	if sqlLen <= length {
		return e.SQL()
	}
	return "..." + e.SQL()[sqlLen-length:]
}
