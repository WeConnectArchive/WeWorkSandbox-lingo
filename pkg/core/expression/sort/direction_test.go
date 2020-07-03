package sort_test

import (
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core/expression/sort"
)

var _ = DescribeTable("Direction",

	func(dir sort.Direction, sqlStr string) {
		opStr := dir.String()
		Expect(opStr).To(Equal(sqlStr))

		sql, err := dir.ToSQL(NewMockDialect())
		Expect(err).ToNot(HaveOccurred())
		Expect(sql).To(matchers.MatchSQLString(sqlStr))
		Expect(sql).To(matchers.MatchSQLValues(BeEmpty()))
	},

	Entry("Ascending", sort.Ascending, "ASC"),
	Entry("Descending", sort.Descending, "DESC"),
)
