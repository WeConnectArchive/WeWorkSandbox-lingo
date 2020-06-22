package query_test

import (
	"fmt"

	"github.com/weworksandbox/lingo/internal/test/schema/qsakila/qfilmactor"
	"github.com/weworksandbox/lingo/pkg/core/dialect"
	"github.com/weworksandbox/lingo/pkg/core/query"
)

func ExampleSelectFrom() {
	d, _ := dialect.NewDefault()

	pageNum := int64(2)
	pageSize := int64(5)

	modify := query.Page(pageSize, pageNum * pageSize)

	fa := qfilmactor.As("fa")
	s, _ := query.SelectFrom(fa).Restrict(modify).ToSQL(d)
	fmt.Println(s)
	// Output: SELECT fa.actor_id, fa.film_id, fa.last_update FROM film_actor AS fa LIMIT ? OFFSET ?
}
