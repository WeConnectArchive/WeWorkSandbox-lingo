package operator_test

import (
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/core/operator"
)

var _ = DescribeTable("Operator",

	func(op operator.Operand, sqlStr string) {
		opStr := op.String()
		Expect(opStr).To(Equal(sqlStr))

		sql, err := op.GetSQL(NewMockDialect())
		Expect(err).ToNot(HaveOccurred())
		Expect(sql).To(matchers.MatchSQLString(sqlStr))
	},

	Entry("And", operator.And, "AND"),
	Entry("Or", operator.Or, "OR"),
	Entry("Equal", operator.Eq, "="),
	Entry("Not Equal", operator.NotEq, "<>"),
	Entry("Less Than", operator.LessThan, "<"),
	Entry("Less Than or Equal", operator.LessThanOrEqual, "<="),
	Entry("Greater Than", operator.GreaterThan, ">"),
	Entry("Greater Than or Equal", operator.GreaterThanOrEqual, ">="),
	Entry("Like", operator.Like, "LIKE"),
	Entry("Not Like", operator.NotLike, "NOT LIKE"),
	Entry("Is Null", operator.Null, "IS NULL"),
	Entry("Is Not Null", operator.NotNull, "IS NOT NULL"),
	Entry("In", operator.In, "IN"),
	Entry("Not In", operator.NotIn, "NOT IN"),
	Entry("Between", operator.Between, "BETWEEN"),
	Entry("Not Between", operator.NotBetween, "NOT BETWEEN"),
)
