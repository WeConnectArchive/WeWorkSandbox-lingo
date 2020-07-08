package path_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression/path"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var _ = Describe("Column", func() {

	Context("ExpandColumnWithDialect", func() {

		var (
			d   core.Dialect
			col core.Column

			sql sql.Data
			err error
		)

		BeforeEach(func() {
			d = expandColumnDialectSuccess{}
			col = NewMockColumn()
		})

		JustBeforeEach(func() {
			sql, err = path.ExpandColumnWithDialect(d, col)
		})

		It("Returns valid SQL", func() {
			Expect(sql).To(matchers.MatchSQLString("expand column sql"))
		})

		It("Returns no error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		Context("`Dialect` does not support `ExpandColumnDialect`", func() {

			BeforeEach(func() {
				d = NewMockDialect()
			})

			It("Returns a nil SQL", func() {
				Expect(sql).To(BeNil())
			})

			It("Returns a Dialect not supported error", func() {
				Expect(err).To(MatchError(matchers.EqString("dialect function '%s' not supported", "ExpandColumnDialect")))
			})
		})

		Context("Dialect returns an error", func() {

			BeforeEach(func() {
				d = expandColumnDialectFailure{}
			})

			It("Returns a nil SQL", func() {
				Expect(sql).To(BeNil())
			})

			It("Returns the `Dialect` `ExpandColumnDialect` error", func() {
				Expect(err).To(MatchError("expand column error"))
			})
		})
	})
})

type expandColumnDialectSuccess struct{}

func (expandColumnDialectSuccess) GetName() string { return "expand column dialect" }
func (expandColumnDialectSuccess) ExpandColumn(_ core.Column) (sql.Data, error) {
	return sql.String("expand column sql"), nil
}

type expandColumnDialectFailure struct{ expandColumnDialectSuccess }

func (expandColumnDialectFailure) ExpandColumn(_ core.Column) (sql.Data, error) {
	return nil, errors.New("expand column error")
}
