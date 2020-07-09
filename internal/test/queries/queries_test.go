package queries_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/spf13/viper"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/execute"
	"github.com/weworksandbox/lingo/internal/config"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/internal/test/runner"
)

// QueryTest is used by Functional tests, along with benchmark tests. They are used for setting up common data to
// ensure performance and code quality.
type QueryTest struct {
	Name      string
	Focus     bool
	Benchmark bool

	// Params used during the test
	Params Params
}

type Params struct {
	Dialect         func() (lingo.Dialect, error)
	SQL             func() lingo.Expression
	SQLStrAssert    types.GomegaMatcher
	SQLValuesAssert types.GomegaMatcher

	// Params for executing this test
	ExecuteParams ExecuteParams
}

func (p Params) Validate() {
	ExpectWithOffset(1, p.Dialect).ToNot(BeNil(), "Dialect was nil")
	ExpectWithOffset(1, p.SQL).ToNot(BeNil(), "SQL was nil")
	ExpectWithOffset(1, p.SQLStrAssert).ToNot(BeNil(), "SQLStrAssert was nil")
}

const DefaultTimeout = 10 * time.Millisecond

type ExecuteParams struct {
	Type     execute.QueryType
	Timeout  time.Duration
	ScanData []interface{}
	Assert   [][]interface{}
}

func (e ExecuteParams) Validate() {
	ExpectWithOffset(1, e.Type).To(BeElementOf(execute.QTRow, execute.QTRows, execute.QTExec), "must be a valid QT")
	// Timeout now has a default. Check out DefaultTimeout.
	//ExpectWithOffset(1, e.Timeout).To(BeNumerically(">", time.Duration(0)))
	ExpectWithOffset(1, e.ScanData).ToNot(BeEmpty(), "Requires at ScanData pointers / values")
}

func (e ExecuteParams) WithTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, time.Duration(math.Max(float64(e.Timeout), float64(DefaultTimeout))))
}

func BenchmarkQueries(b *testing.B) {
	b.ReportAllocs()

	hasFocus := false
	for _, query := range allQueries {
		if query.Benchmark && query.Focus {
			hasFocus = true
			break
		}
	}

	for _, query := range allQueries {
		b.Run(query.Name, func(parallel *testing.B) {
			if !query.Benchmark {
				b.Skip("Benchmark turned off for query ", query.Name)
			}
			if hasFocus && !query.Focus {
				b.Skip("Focus not enabled for query ", query.Name)
			}
			if query.Params.Dialect == nil {
				b.Errorf("QueryTest '%s' does not have a Dialect", query.Name)
			}

			d, err := query.Params.Dialect()
			if err != nil {
				b.Errorf("unable to create dialect for query %s: %w", query.Name, err)
			}

			parallel.ReportAllocs()
			parallel.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, _ = query.Params.SQL().ToSQL(d)
				}
			})
		})
	}
}

func TestQueries(t *testing.T) {
	var _ = ginkgo.Describe("Queries", func() {
		table.DescribeTable("query.go",
			func(p Params) {
				// Sanity check
				Expect(p).ToNot(BeNil())
				p.Validate()
				d, err := p.Dialect()
				Expect(err).ToNot(HaveOccurred())

				sqlStr, err := p.SQL().ToSQL(d)
				Expect(err).To(Not(HaveOccurred()))
				Expect(sqlStr).To(MatchSQLString(p.SQLStrAssert))
				Expect(sqlStr).To(MatchSQLValues(p.SQLValuesAssert))
			},
			acceptanceEntries...,
		)
	})

	runner.SetupAndRunUnit(t, "Queries", "functional")
}

// Run with `-- --config ../testdata/sakila/lingo-config.yml` flags if running from this function.
// If running from root directory, update the path.
func TestExecute(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	loadFunctionalConfigOrFatal()

	dsn := viper.GetString("dsn")
	if dsn == "" {
		t.Fatalf("'dsn' in config file is not set")
	}

	conf, err := mysql.ParseDSN(dsn)
	if err != nil {
		t.Fatalf("unable to parse found dsn: %s", err)
	}
	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		t.Fatalf("unable to connect to database: %s", err)
	}
	t.Cleanup(func() {
		if closeErr := db.Close(); closeErr != nil {
			t.Logf("error while cleaning up database: %s", closeErr)
		}
	})

	// Setup 1 connection for tests while also checking for db availability.
	db.SetMaxIdleConns(1)

	pingCtx, pingCf := context.WithTimeout(context.Background(), DefaultTimeout*5) // Give extra time for DB connect
	pingErr := db.PingContext(pingCtx)
	pingCf()
	if pingErr != nil {
		t.Fatalf("unable to ping database: %s", pingErr)
	}

	var _ = ginkgo.Describe("Queries", func() {
		table.DescribeTable("query.go",
			func(p Params) {
				// Sanity check
				Expect(p).ToNot(BeNil())
				Expect(p.ExecuteParams).ToNot(BeNil(), "Requires ExecuteParams")
				p.Validate()
				p.ExecuteParams.Validate()

				ctx, cf := p.ExecuteParams.WithTimeout(context.Background())
				defer cf()

				d, err := p.Dialect()
				Expect(err).ToNot(HaveOccurred())

				sqlExec := execute.NewSQLExp(execute.NewSQL(db), d)
				sqlExp := p.SQL()

				switch p.ExecuteParams.Type {
				case execute.QTExec:
					Expect(true).To(BeFalse(), "QTExec tests not implemented")

				case execute.QTRow:
					Expect(p.ExecuteParams.Assert).To(HaveLen(1),
						"Must assert values from result of QueryRow")
					Expect(p.ExecuteParams.Assert[0]).To(HaveLen(len(p.ExecuteParams.ScanData)),
						"Number of columns asserting should match number of columns scanning")

					queryErr := sqlExec.QueryRow(ctx, sqlExp, p.ExecuteParams.ScanData...)
					Expect(queryErr).ToNot(HaveOccurred())
					Expect(p.ExecuteParams.ScanData).To(Equal(p.ExecuteParams.Assert[0]))

				case execute.QTRows:
					Expect(p.ExecuteParams.Assert).To(EachElementMust(HaveLen(len(p.ExecuteParams.ScanData))),
						"Number of columns asserting should match number of columns scanning")

					scanner, queryErr := sqlExec.Query(ctx, sqlExp)
					Expect(queryErr).ToNot(HaveOccurred(), "unable to query for scanner")
					defer scanner.Close(ctx)

					for idx := 0; idx < len(p.ExecuteParams.Assert) && scanner.ScanRow(p.ExecuteParams.ScanData...); idx++ {

						Expect(p.ExecuteParams.ScanData).To(Equal(p.ExecuteParams.Assert[idx]),
							fmt.Sprintf("row %d did not match", idx))
					}
					Expect(scanner.Err(ctx)).ToNot(HaveOccurred(),
						"An error occurred while scanning or after scanning")
				}
			},
			acceptanceEntries...,
		)
	})

	runner.SetupAndRunFunctional(t, "Execute")
}

func loadFunctionalConfigOrFatal() {
	// Init Functional Test Args
	for idx := len(os.Args) - 1; idx >= 0; idx-- {
		if os.Args[idx] == "--" {
			subArgs := os.Args[idx+1:]
			if err := config.FileFlag.Parse(subArgs); err != nil {
				log.Fatalf("unable to parse the os.Args={%v}: %s", subArgs, err)
			}
			if err := config.ReadConfig(); err != nil {
				log.Fatalf("unable to read config from %s: %s", config.File, err)
			}
			break
		}
	}
}

var (
	allQueries        = aggregateQueries(selectQueries)
	acceptanceEntries = queriesToEntries(allQueries)
)

func aggregateQueries(q ...[]QueryTest) []QueryTest {
	var result []QueryTest
	for idx := range q {
		result = append(result, q[idx]...)
	}
	return result
}

func queriesToEntries(queries []QueryTest) []table.TableEntry {
	var entries = make([]table.TableEntry, len(queries))
	for idx, query := range queries {
		entries[idx] = table.TableEntry{
			Description: query.Name,
			Parameters:  []interface{}{query.Params},
			Pending:     false,
			Focused:     query.Focus,
		}
	}
	return entries
}

// trimQuery replaces newlines with spaces, and removing any tabs. This way, SQL.SQL can use backticks.
func trimQuery(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", "")
	return strings.TrimSpace(s)
}
