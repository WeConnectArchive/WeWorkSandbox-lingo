package sql_test

import (
	"github.com/weworksandbox/lingo/pkg/core/sql"
	"testing"
)

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := sql.String("sqlString")
		if s.String() != "sqlString" {
			b.Errorf("Expected sqlString but got %s instead", s.String())
		}
	}
}
