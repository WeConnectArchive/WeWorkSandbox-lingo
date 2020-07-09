package path_test

import (
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expr"
	"github.com/weworksandbox/lingo/pkg/core/expr/operator"
	"github.com/weworksandbox/lingo/pkg/core/expr/path"
)

var _ = Describe("Int64", func() {

	Context("NewInt64PathWithAlias", func() {

		var (
			e     core.Table
			name  string
			alias string

			p path.Int64
		)

		BeforeEach(func() {
			e = NewMockTable()
			name = "name"
			alias = "alias"
		})

		JustBeforeEach(func() {
			p = path.NewInt64PathWithAlias(e, name, alias)
		})

		It("Returns a `Int64`", func() {
			Expect(p).To(BeAssignableToTypeOf(path.Int64{}))
		})

		It("Has the given parent table", func() {
			Expect(p.GetParent()).To(Equal(e))
		})

		It("Has the given name", func() {
			Expect(p.GetName()).To(Equal(name))
		})

		It("Has the given alias", func() {
			Expect(p.GetAlias()).To(Equal(alias))
		})
	})

	Context("NewInt64Path", func() {

		var (
			e    core.Table
			name string

			p path.Int64
		)

		BeforeEach(func() {
			e = NewMockTable()
			name = "name"
		})

		JustBeforeEach(func() {
			p = path.NewInt64Path(e, name)
		})

		It("Returns a `Int64`", func() {
			Expect(p).To(BeAssignableToTypeOf(path.Int64{}))
		})

		It("Has the given parent table", func() {
			Expect(p.GetParent()).To(Equal(e))
		})

		It("Has the given name", func() {
			Expect(p.GetName()).To(Equal(name))
		})

		It("Has a blank alias", func() {
			Expect(p.GetAlias()).To(BeEmpty())
		})

		Context("As", func() {

			var (
				alias string
			)

			BeforeEach(func() {
				alias = "new_name"
			})

			JustBeforeEach(func() {
				p = p.As(alias)
			})

			It("Adds the alias", func() {
				Expect(p.GetAlias()).To(Equal(alias))
			})
		})

		Context("To", func() {

			var (
				value int64
				set   core.Set
			)

			BeforeEach(func() {
				value = math.MinInt64
			})

			JustBeforeEach(func() {
				set = p.To(value)
			})

			It("Returns a valid `core.SetDialect`", func() {
				Expect(set).ToNot(BeNil())
				Expect(set).To(Equal(expr.NewSet(p, expr.NewValue(value))))
			})
		})

		Context("ToExpression", func() {

			var (
				value core.Expression
				set   core.Set
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				set = p.ToExpression(value)
			})

			It("Returns a valid `core.SetDialect`", func() {
				Expect(set).ToNot(BeNil())
				Expect(set).To(Equal(expr.NewSet(p, value)))
			})
		})

		Context("Eq", func() {

			var (
				value  int64
				result core.Expression
			)

			BeforeEach(func() {
				value = math.MinInt64
			})

			JustBeforeEach(func() {
				result = p.Eq(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Eq, expr.NewValue(value))))
			})
		})

		Context("EqPath", func() {

			var (
				value  core.Expression
				result core.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.EqPath(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Eq, value)))
			})
		})

		Context("NotEq", func() {

			var (
				value  int64
				result core.Expression
			)

			BeforeEach(func() {
				value = math.MinInt64
			})

			JustBeforeEach(func() {
				result = p.NotEq(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotEq, expr.NewValue(value))))
			})
		})

		Context("NotEqPath", func() {

			var (
				value  core.Expression
				result core.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.NotEqPath(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotEq, value)))
			})
		})

		Context("LT", func() {

			var (
				value  int64
				result core.Expression
			)

			BeforeEach(func() {
				value = math.MinInt64
			})

			JustBeforeEach(func() {
				result = p.LT(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.LessThan, expr.NewValue(value))))
			})
		})

		Context("LTPath", func() {

			var (
				value  core.Expression
				result core.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.LTPath(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.LessThan, value)))
			})
		})

		Context("LTOrEq", func() {

			var (
				value  int64
				result core.Expression
			)

			BeforeEach(func() {
				value = math.MinInt64
			})

			JustBeforeEach(func() {
				result = p.LTOrEq(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.LessThanOrEqual, expr.NewValue(value))))
			})
		})

		Context("LTOrEqPath", func() {

			var (
				value  core.Expression
				result core.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.LTOrEqPath(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.LessThanOrEqual, value)))
			})
		})

		Context("GT", func() {

			var (
				value  int64
				result core.Expression
			)

			BeforeEach(func() {
				value = math.MinInt64
			})

			JustBeforeEach(func() {
				result = p.GT(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.GreaterThan, expr.NewValue(value))))
			})
		})

		Context("GTPath", func() {

			var (
				value  core.Expression
				result core.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.GTPath(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.GreaterThan, value)))
			})
		})

		Context("GTOrEq", func() {

			var (
				value  int64
				result core.Expression
			)

			BeforeEach(func() {
				value = math.MinInt64
			})

			JustBeforeEach(func() {
				result = p.GTOrEq(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.GreaterThanOrEqual, expr.NewValue(value))))
			})
		})

		Context("GTOrEqPath", func() {

			var (
				value  core.Expression
				result core.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.GTOrEqPath(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.GreaterThanOrEqual, value)))
			})
		})

		Context("IsNull", func() {

			var (
				result core.Expression
			)

			JustBeforeEach(func() {
				result = p.IsNull()
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Null)))
			})
		})

		Context("IsNotNull", func() {

			var (
				result core.Expression
			)

			JustBeforeEach(func() {
				result = p.IsNotNull()
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotNull)))
			})
		})

		Context("In", func() {

			var (
				value  []int64
				result core.Expression
			)

			BeforeEach(func() {
				value = []int64{
					math.MinInt64,
					math.MaxInt64,
				}
			})

			JustBeforeEach(func() {
				// Doing it this way to make it more apparent we are passing in multiple params
				result = p.In(value[0], value[1])
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.In, expr.NewValue(value))))
			})
		})

		Context("InPath", func() {

			var (
				value  []core.Expression
				result core.Expression
			)

			BeforeEach(func() {
				value = []core.Expression{
					NewMockExpression(),
					NewMockExpression(),
				}
			})

			JustBeforeEach(func() {
				result = p.InPaths(value[0], value[1])
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.In, value[0], value[1])))
			})
		})

		Context("NotIn", func() {

			var (
				value  []int64
				result core.Expression
			)

			BeforeEach(func() {
				value = []int64{
					math.MinInt64,
					math.MaxInt64,
				}
			})

			JustBeforeEach(func() {
				// Doing it this way to make it more apparent we are passing in multiple params
				result = p.NotIn(value[0], value[1])
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotIn, expr.NewValue(value))))
			})
		})

		Context("NotInPath", func() {

			var (
				value  []core.Expression
				result core.Expression
			)

			BeforeEach(func() {
				value = []core.Expression{
					NewMockExpression(),
					NewMockExpression(),
				}
			})

			JustBeforeEach(func() {
				result = p.NotInPaths(value[0], value[1])
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotIn, value[0], value[1])))
			})
		})

		Context("Between", func() {

			var (
				firstValue  int64
				secondValue int64
				result      core.Expression
			)

			BeforeEach(func() {
				firstValue = math.MinInt64
				secondValue = math.MaxInt64
			})

			JustBeforeEach(func() {
				result = p.Between(firstValue, secondValue)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Between, expr.NewValue(firstValue).And(expr.NewValue(secondValue)))))
			})
		})

		Context("BetweenPaths", func() {

			var (
				firstValue  core.Expression
				secondValue core.Expression
				result      core.Expression
			)

			BeforeEach(func() {
				firstValue = NewMockExpression()
				secondValue = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.BetweenPaths(firstValue, secondValue)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Between, operator.NewOperator(firstValue, operator.And, secondValue))))
			})
		})

		Context("NotBetween", func() {

			var (
				firstValue  int64
				secondValue int64
				result      core.Expression
			)

			BeforeEach(func() {
				firstValue = math.MinInt64
				secondValue = math.MaxInt64
			})

			JustBeforeEach(func() {
				result = p.NotBetween(firstValue, secondValue)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotBetween, expr.NewValue(firstValue).And(expr.NewValue(secondValue)))))
			})
		})

		Context("NotBetweenPaths", func() {

			var (
				firstValue  core.Expression
				secondValue core.Expression
				result      core.Expression
			)

			BeforeEach(func() {
				firstValue = NewMockExpression()
				secondValue = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.NotBetweenPaths(firstValue, secondValue)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotBetween, operator.NewOperator(firstValue, operator.And, secondValue))))
			})
		})
	})
})
