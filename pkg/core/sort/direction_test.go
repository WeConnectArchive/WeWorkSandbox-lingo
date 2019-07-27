package sort_test

import (
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core/sort"
)

var _ = DescribeTable("Direction",

	func(dir sort.Direction, sqlStr string) {
		opStr := dir.String()
		Expect(opStr).To(Equal(sqlStr))

		sql, err := dir.GetSQL(NewMockDialect())
		Expect(err).ToNot(HaveOccurred())
		Expect(sql).To(matchers.MatchSQLString(sqlStr))
	},

	Entry("Ascending", sort.Ascending, "ASC"),
	Entry("Descending", sort.Descending, "DESC"),
)
