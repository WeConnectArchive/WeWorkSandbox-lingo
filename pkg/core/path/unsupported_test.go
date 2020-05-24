package path_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/weworksandbox/lingo/internal/test/matchers"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

var _ = Describe("Unsupported", func() {

	Context("NewUnsupportedWithAlias", func() {

		var (
			e     core.Table
			name  string
			alias string

			p path.Unsupported
		)

		BeforeEach(func() {
			e = NewMockTable()
			name = "name"
			alias = "alias"
		})

		JustBeforeEach(func() {
			p = path.NewUnsupportedWithAlias(e, name, alias)
		})

		It("Returns a `Unsupported`", func() {
			Expect(p).To(BeAssignableToTypeOf(path.Unsupported{}))
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

	Context("NewUnsupported", func() {

		var (
			e    core.Table
			name string

			p path.Unsupported
		)

		BeforeEach(func() {
			e = NewMockTable()
			name = "name"
		})

		JustBeforeEach(func() {
			p = path.NewUnsupported(e, name)
		})

		It("Returns a `Unsupported`", func() {
			Expect(p).To(BeAssignableToTypeOf(path.Unsupported{}))
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

		It("Has empty SQL for GetSQL", func() {
			Expect(p.GetSQL(nil)).To(matchers.MatchSQLString(""))
		})
	})
})
