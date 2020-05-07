package execute_test

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/dialect"
	"github.com/weworksandbox/lingo/pkg/core/execute"
)

var _ = PDescribe("sqlexpexec.go", func() {

	Context("#NewSQLExp", func() {
		var (
			// Input
			db *sql.DB
			d  core.Dialect

			// Output
			execSQL execute.SQLExp
		)
		BeforeEach(func() {
			d = dialect.MySQL{}
		})
		JustBeforeEach(func() {
			execSQL = execute.NewSQLExp(execute.NewSQL(db), d)
		})

		It("Creates a SQLExp", func() {
			Expect(execSQL).ToNot(BeNil())
		})
	})
})
