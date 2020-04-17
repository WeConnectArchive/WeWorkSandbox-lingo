package acceptance

import (
	"strings"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"

	"github.com/weworksandbox/lingo/db/mysql/qinformationschema/qcharactersets"
	"github.com/weworksandbox/lingo/db/mysql/qinformationschema/qcollations"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core/dialect"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/query"
	"github.com/weworksandbox/lingo/pkg/core/sort"
)

type SelectTest struct {
	Query        *query.SelectQuery
	SQLAssert    types.GomegaMatcher
	ValuesAssert []types.GomegaMatcher
	ErrAssert    types.GomegaMatcher
}

var _ = ginkgo.Describe("Select", func() {
	table.DescribeTable("Select",
		func(s SelectTest) {
			// Sanity check
			Expect(s.Query).ToNot(BeNil())
			Expect(s.SQLAssert).ToNot(BeNil())
			Expect(s.ErrAssert).ToNot(BeNil())

			// Convert to []interface{} for MatchSQLValues
			var convertedValues = make([]interface{}, len(s.ValuesAssert))
			for idx := range s.ValuesAssert {
				convertedValues[idx] = s.ValuesAssert[idx]
			}

			d := dialect.Default{}
			sql, err := s.Query.GetSQL(d)
			Expect(sql).To(MatchSQLString(s.SQLAssert))
			Expect(sql).To(MatchSQLValues(AllInSlice(convertedValues...)))
			Expect(err).To(s.ErrAssert)
		},
		table.Entry(SelectLeftJoinWhereOrderBy()),
	)
})

func stripNT(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", "")
	return strings.TrimSpace(s)
}

func SelectLeftJoinWhereOrderBy() (string, SelectTest) {
	cs := qcharactersets.As("cs")
	col := qcollations.As("col")

	return "SelectLeftJoinWhereOrderBy", SelectTest{
		Query: query.Select(cs.Description(), cs.CharacterSetName()).
			From(cs).
			Join(col, expression.LeftJoin, cs.CharacterSetName().EqPath(col.CharacterSetName())).
			Where(cs.Maxlen().GT(60)).
			OrderBy(cs.Maxlen(), sort.Descending).
			Where(cs.DefaultCollateName().Eq("DefaultName")),
		SQLAssert: ContainSubstring(stripNT(`
					SELECT cs.DESCRIPTION, cs.CHARACTER_SET_NAME
					FROM information_schema.CHARACTER_SETS AS cs
					LEFT JOIN information_schema.COLLATIONS AS col ON cs.CHARACTER_SET_NAME = col.CHARACTER_SET_NAME
					WHERE (cs.MAXLEN > ? AND cs.DEFAULT_COLLATE_NAME = ?)
					ORDER BY cs.MAXLEN DESC`)),
		ValuesAssert: []types.GomegaMatcher{
			BeEquivalentTo(60),
			BeEquivalentTo("DefaultName"),
		},
		ErrAssert: BeNil(),
	}
}
