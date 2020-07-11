package path_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/path"
	"github.com/weworksandbox/lingo/expr/set"
)

var _ = Describe("Float32", func() {

	Context("NewFloat32WithAlias", func() {

		var (
			e     lingo.Table
			name  string
			alias string

			p path.Float32
		)

		BeforeEach(func() {
			e = NewMockTable()
			name = "name"
			alias = "alias"
		})

		JustBeforeEach(func() {
			p = path.NewFloat32WithAlias(e, name, alias)
		})

		It("Returns a `Float32`", func() {
			Expect(p).To(BeAssignableToTypeOf(path.Float32{}))
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

	Context("NewFloat32", func() {

		var (
			e    lingo.Table
			name string

			p path.Float32
		)

		BeforeEach(func() {
			e = NewMockTable()
			name = "name"
		})

		JustBeforeEach(func() {
			p = path.NewFloat32(e, name)
		})

		It("Returns a `Float32`", func() {
			Expect(p).To(BeAssignableToTypeOf(path.Float32{}))
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
				value float32
				s     lingo.Set
			)

			BeforeEach(func() {
				value = 1.5555
			})

			JustBeforeEach(func() {
				s = p.To(value)
			})

			It("Returns a valid `set.Set`", func() {
				Expect(s).ToNot(BeNil())
				Expect(s).To(Equal(set.NewSet(p, expr.NewValue(value))))
			})
		})

		Context("ToExpression", func() {

			var (
				value lingo.Expression
				s     lingo.Set
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				s = p.ToExpr(value)
			})

			It("Returns a valid `set.Set`", func() {
				Expect(s).ToNot(BeNil())
				Expect(s).To(Equal(set.NewSet(p, value)))
			})
		})

		Context("Eq", func() {

			var (
				value  float32
				result lingo.Expression
			)

			BeforeEach(func() {
				value = 1.5555
			})

			JustBeforeEach(func() {
				result = p.Eq(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.Eq, expr.NewValue(value))))
			})
		})

		Context("EqPath", func() {

			var (
				value  lingo.Expression
				result lingo.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.EqPath(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.Eq, value)))
			})
		})

		Context("NotEq", func() {

			var (
				value  float32
				result lingo.Expression
			)

			BeforeEach(func() {
				value = 1.5555
			})

			JustBeforeEach(func() {
				result = p.NotEq(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.NotEq, expr.NewValue(value))))
			})
		})

		Context("NotEqPath", func() {

			var (
				value  lingo.Expression
				result lingo.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.NotEqPath(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.NotEq, value)))
			})
		})

		Context("LT", func() {

			var (
				value  float32
				result lingo.Expression
			)

			BeforeEach(func() {
				value = 1.5555
			})

			JustBeforeEach(func() {
				result = p.LT(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.LessThan, expr.NewValue(value))))
			})
		})

		Context("LTPath", func() {

			var (
				value  lingo.Expression
				result lingo.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.LTPath(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.LessThan, value)))
			})
		})

		Context("LTOrEq", func() {

			var (
				value  float32
				result lingo.Expression
			)

			BeforeEach(func() {
				value = 1.5555
			})

			JustBeforeEach(func() {
				result = p.LTOrEq(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.LessThanOrEqual, expr.NewValue(value))))
			})
		})

		Context("LTOrEqPath", func() {

			var (
				value  lingo.Expression
				result lingo.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.LTOrEqPath(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.LessThanOrEqual, value)))
			})
		})

		Context("GT", func() {

			var (
				value  float32
				result lingo.Expression
			)

			BeforeEach(func() {
				value = 1.5555
			})

			JustBeforeEach(func() {
				result = p.GT(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.GreaterThan, expr.NewValue(value))))
			})
		})

		Context("GTPath", func() {

			var (
				value  lingo.Expression
				result lingo.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.GTPath(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.GreaterThan, value)))
			})
		})

		Context("GTOrEq", func() {

			var (
				value  float32
				result lingo.Expression
			)

			BeforeEach(func() {
				value = 1.5555
			})

			JustBeforeEach(func() {
				result = p.GTOrEq(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.GreaterThanOrEqual, expr.NewValue(value))))
			})
		})

		Context("GTOrEqPath", func() {

			var (
				value  lingo.Expression
				result lingo.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.GTOrEqPath(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.GreaterThanOrEqual, value)))
			})
		})

		Context("IsNull", func() {

			var (
				result lingo.Expression
			)

			JustBeforeEach(func() {
				result = p.IsNull()
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewUnary(p, operator.Null)))
			})
		})

		Context("IsNotNull", func() {

			var (
				result lingo.Expression
			)

			JustBeforeEach(func() {
				result = p.IsNotNull()
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewUnary(p, operator.NotNull)))
			})
		})

		Context("In", func() {

			var (
				value  []float32
				result lingo.Expression
			)

			BeforeEach(func() {
				value = []float32{
					1.555,
					6.222,
				}
			})

			JustBeforeEach(func() {
				// Doing it this way to make it more apparent we are passing in multiple params
				result = p.In(value[0], value[1])
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.In, expr.NewParens(expr.NewValue(value)))))
			})
		})

		Context("InPath", func() {

			var (
				value  []lingo.Expression
				result lingo.Expression
			)

			BeforeEach(func() {
				value = []lingo.Expression{
					NewMockExpression(),
					NewMockExpression(),
				}
			})

			JustBeforeEach(func() {
				result = p.InPaths(value[0], value[1])
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.In, expr.NewParens(expr.ToList(value)))))
			})
		})

		Context("NotIn", func() {

			var (
				value  []float32
				result lingo.Expression
			)

			BeforeEach(func() {
				value = []float32{
					1.555,
					6.222,
				}
			})

			JustBeforeEach(func() {
				// Doing it this way to make it more apparent we are passing in multiple params
				result = p.NotIn(value[0], value[1])
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.NotIn, expr.NewParens(expr.NewValue(value)))))
			})
		})

		Context("NotInPath", func() {

			var (
				value  []lingo.Expression
				result lingo.Expression
			)

			BeforeEach(func() {
				value = []lingo.Expression{
					NewMockExpression(),
					NewMockExpression(),
				}
			})

			JustBeforeEach(func() {
				result = p.NotInPaths(value[0], value[1])
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.NotIn, expr.NewParens(expr.ToList(value)))))
			})
		})

		Context("Between", func() {

			var (
				firstValue  float32
				secondValue float32
				result      lingo.Expression
			)

			BeforeEach(func() {
				firstValue = 1.555
				secondValue = 6.222
			})

			JustBeforeEach(func() {
				result = p.Between(firstValue, secondValue)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.Between, expr.NewParens(expr.NewValue(firstValue).And(expr.NewValue(secondValue))))))
			})
		})

		Context("BetweenPaths", func() {

			var (
				firstValue  lingo.Expression
				secondValue lingo.Expression
				result      lingo.Expression
			)

			BeforeEach(func() {
				firstValue = NewMockExpression()
				secondValue = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.BetweenPaths(firstValue, secondValue)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.Between, expr.NewParens(operator.NewBinary(firstValue, operator.And, secondValue)))))
			})
		})

		Context("NotBetween", func() {

			var (
				firstValue  float32
				secondValue float32
				result      lingo.Expression
			)

			BeforeEach(func() {
				firstValue = 1.555
				secondValue = 6.222
			})

			JustBeforeEach(func() {
				result = p.NotBetween(firstValue, secondValue)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.NotBetween, expr.NewParens(expr.NewValue(firstValue).And(expr.NewValue(secondValue))))))
			})
		})

		Context("NotBetweenPaths", func() {

			var (
				firstValue  lingo.Expression
				secondValue lingo.Expression
				result      lingo.Expression
			)

			BeforeEach(func() {
				firstValue = NewMockExpression()
				secondValue = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.NotBetweenPaths(firstValue, secondValue)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewBinary(p, operator.NotBetween, expr.NewParens(operator.NewBinary(firstValue, operator.And, secondValue)))))
			})
		})
	})
})
