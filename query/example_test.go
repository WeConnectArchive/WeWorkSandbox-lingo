package query_test

import (
	"fmt"

	"github.com/weworksandbox/lingo/dialect"
	"github.com/weworksandbox/lingo/internal/test/schema/tsakila/tfilmactor"
	"github.com/weworksandbox/lingo/query"
)

func ExampleSelectQuery_From_where() {
	d, _ := dialect.NewDialect()

	fa := tfilmactor.As("fa")
	s, _ := query.Select(fa.FilmId()).From(fa).Where(fa.ActorId().Between(1, 10)).ToSQL(d)

	fmt.Println(s.String())
	fmt.Println(s.Values())
	// Output:
	//SELECT fa.film_id FROM film_actor AS fa WHERE fa.actor_id BETWEEN ? AND ?
	//[1 10]
}

func ExampleSelectQuery_Restrict() {
	d, _ := dialect.NewDialect()

	const maxPageNum = 1 // To limit output for example
	pageSize := uint64(150)
	fa := tfilmactor.As("fa")
	q := query.SelectFrom(fa)

	for page := uint64(0); page < maxPageNum; page++ {
		s, _ := q.Restrict(query.Page(pageSize, page*pageSize)).ToSQL(d)

		fmt.Println(s.String())
		fmt.Println(s.Values())
		// Output:
		//SELECT fa.actor_id, fa.film_id, fa.last_update FROM film_actor AS fa LIMIT ? OFFSET ?
		//[150 0]
	}
}
