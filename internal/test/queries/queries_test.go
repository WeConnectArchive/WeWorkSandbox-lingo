package queries_test

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/internal/test/runner"
)

func BenchmarkQueries(b *testing.B) {
	b.ReportAllocs()

	for _, query := range allQueries {
		if !query.Benchmark {
			b.Skip("Benchmark turned off for query ", query.Name)
		}

		if query.Dialect == nil {
			b.Errorf("Query '%s' does not have a Dialect", query.Name)
		}

		b.Run(query.Name, func(parallel *testing.B) {
			parallel.ReportAllocs()
			parallel.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, _ = query.SQL.GetSQL(query.Dialect)
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
				Expect(p.SQL).ToNot(BeNil())
				Expect(p.SQLAssert).ToNot(BeNil())
				Expect(p.ErrAssert).ToNot(BeNil())

				sql, err := p.SQL.GetSQL(p.Dialect)
				Expect(err).To(p.ErrAssert)
				Expect(sql).To(MatchSQLString(p.SQLAssert))
				Expect(sql).To(MatchSQLValues(p.ValuesAssert))
			},
			acceptanceEntries...,
		)
	})

	runner.SetupAndRunFunctional(t, "Queries")
}

var (
	allQueries        = aggregateQueries(selectQueries)
	acceptanceEntries = queriesToEntries(allQueries)
)

func aggregateQueries(q ...[]Query) []Query {
	var result []Query
	for idx := range q {
		result = append(result, q[idx]...)
	}
	return result
}

func queriesToEntries(queries []Query) []table.TableEntry {
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
