package path_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo/pkg/core/expression/path"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/expression/json"
)

var _ = Describe("JSON", func() {

	Context("NewJSONWithAlias", func() {

		var (
			e     core.Table
			name  string
			alias string

			p path.JSON
		)

		BeforeEach(func() {
			e = NewMockTable()
			name = "name"
			alias = "alias"
		})

		JustBeforeEach(func() {
			p = path.NewJSONWithAlias(e, name, alias)
		})

		It("Returns a `JSON`", func() {
			Expect(p).To(BeAssignableToTypeOf(path.JSON{}))
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

	Context("NewJSON", func() {

		var (
			e    core.Table
			name string

			p path.JSON
		)

		BeforeEach(func() {
			e = NewMockTable()
			name = "name"
		})

		JustBeforeEach(func() {
			p = path.NewJSON(e, name)
		})

		It("Returns a `JSON`", func() {
			Expect(p).To(BeAssignableToTypeOf(path.JSON{}))
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

		Context("Extract", func() {

			var (
				paths  []string
				result core.Expression
			)

			BeforeEach(func() {
				paths = []string{
					"$.path1",
					"$.path2",
				}
			})

			JustBeforeEach(func() {
				result = p.Extract(paths...)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(expression.NewJSONOperation(p, json.Extract, expression.NewValue(paths))))
			})
		})
	})
})
