package execute_test

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/exporters/trace/stdout"
	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/weworksandbox/lingo/internal/test/schema/qsakila/qactor"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/dialect"
	"github.com/weworksandbox/lingo/pkg/core/execute"
	"github.com/weworksandbox/lingo/pkg/core/query"
	"github.com/weworksandbox/lingo/pkg/core/sort"
)

var _ = Describe("executor.go", func() {

	Context("#NewSQL", func() {
		var (
			// Input
			db *sql.DB
			d  core.Dialect

			// Output
			execSQL execute.SQL
		)
		BeforeEach(func() {
			exporter, err := stdout.NewExporter(stdout.Options{})
			provider, err := trace.NewProvider(
				trace.WithSyncer(exporter),
			)
			Expect(err).ToNot(HaveOccurred())
			global.SetTraceProvider(provider)

			d = dialect.MySQL{}

			db, err = sql.Open("mysql", "root:lingo@tcp(localhost:9999)/?maxAllowedPacket=0")
			Expect(err).ToNot(HaveOccurred())
		})
		JustBeforeEach(func() {
			execSQL = execute.NewSQL(db, d)
		})

		It("Creates a SQL", func() {
			Expect(execSQL).ToNot(BeNil())
		})

		Context("#QueryRow", func() {
			It("Asdf", func() {
				execSQL = execute.NewSQL(db, d)

				ctx := context.Background()
				q := query.Select(qactor.FirstName()).From(qactor.Q()).OrderBy(qactor.LastUpdate(), sort.Descending)

				var a actor
				err := execSQL.QueryRow(ctx, q, &a.FirstName)
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})

type actor struct {
	ActorId int16
	FirstName string
	LastName string
	LastUpdate time.Time
}
