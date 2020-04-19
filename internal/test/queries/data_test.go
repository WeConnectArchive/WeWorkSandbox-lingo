package queries_test

import (
	"strings"

	"github.com/onsi/gomega/types"

	"github.com/weworksandbox/lingo/db/mysql/qinformationschema/qcharactersets"
	"github.com/weworksandbox/lingo/db/mysql/qinformationschema/qcollations"
	"github.com/weworksandbox/lingo/pkg/core"
)

const (
	maxLen      = 60
	defCollName = "DefaultName"
)

var (
	charSetNameIn = []string{
		"NAME1",
		"NAME2",
		"NAME3",
	}

	cs  = qcharactersets.As("CS")
	col = qcollations.As("COL")
)

// Query is used by Functional tests, along with benchmark tests. They are used for setting up common data to
// ensure performance and code quality.
type Query struct {
	Name      string
	Focus     bool
	Benchmark bool

	// Params used during the test
	Params
}

type Params struct {
	Dialect      core.Dialect
	SQL          core.Expression
	SQLAssert    types.GomegaMatcher
	ValuesAssert types.GomegaMatcher
	ErrAssert    types.GomegaMatcher
}

// trimQuery replaces newlines with spaces, and removing any tabs. This way, SQL.SQL can use backticks.
func trimQuery(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", "")
	return strings.TrimSpace(s)
}
