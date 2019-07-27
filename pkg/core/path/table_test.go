package path_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
	"golang.org/x/xerrors"
)

var _ = Describe("Table", func() {

	Context("ExpandTableWithDialect", func() {

		var (
			d     core.Dialect
			table core.Table

			sql core.SQL
			err error
		)

		BeforeEach(func() {
			d = expandTableDialectSuccess{}
			table = NewMockTable()
		})

		JustBeforeEach(func() {
			sql, err = path.ExpandTableWithDialect(d, table)
		})

		It("Returns valid SQL", func() {
			Expect(sql).To(matchers.MatchSQLString("expand table sql"))
		})

		It("Returns no error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		Context("`Dialect` does not support `ExpandTable`", func() {

			BeforeEach(func() {
				d = NewMockDialect()
			})

			It("Returns a nil SQL", func() {
				Expect(sql).To(BeNil())
			})

			It("Returns a Dialect not supported error", func() {
				Expect(err).To(MatchError(matchers.EqString("dialect function '%s' not supported", "ExpandTable")))
			})
		})

		Context("Dialect returns an error", func() {

			BeforeEach(func() {
				d = expandTableDialectFailure{}
			})

			It("Returns a nil SQL", func() {
				Expect(sql).To(BeNil())
			})

			It("Returns the `Dialect` `ExpandTable` error", func() {
				Expect(err).To(MatchError("expand table error"))
			})
		})
	})
})

type expandTableDialectSuccess struct{}

func (expandTableDialectSuccess) GetName() string { return "expand table dialect" }
func (expandTableDialectSuccess) ExpandTable(entity core.Table) (core.SQL, error) {
	return core.NewSQLf("expand table sql"), nil
}

type expandTableDialectFailure struct{ expandTableDialectSuccess }

func (expandTableDialectFailure) ExpandTable(entity core.Table) (core.SQL, error) {
	return nil, xerrors.New("expand table error")
}
