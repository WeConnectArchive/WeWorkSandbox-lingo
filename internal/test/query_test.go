package test_test

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/internal/test/runner"
)

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

func BenchmarkAcceptance(acceptanceB *testing.B) {
	for _, query := range allQueries {
		if !query.Benchmark {
			continue
		}

		acceptanceB.ResetTimer()
		acceptanceB.Run(query.Name, func(b *testing.B) {
			if query.Dialect == nil {
				b.Errorf("Query '%s' does not have a Dialect", query.Name)
			}

			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, _ = query.SQL.GetSQL(query.Dialect)
				}
			})
		})
	}
}

func TestAcceptance(t *testing.T) {
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

	runner.SetupAndRunAcceptance(t, "acceptance")
}
