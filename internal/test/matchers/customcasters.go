package matchers

import (
	"github.com/weworksandbox/lingo/pkg/core"
)

func isSQL(a interface{}) bool {
	_, ok := a.(core.SQL)
	return ok
}

func toSQL(a interface{}) core.SQL {
	s, ok := a.(core.SQL)
	if ok {
		return s
	}
	return nil
}
