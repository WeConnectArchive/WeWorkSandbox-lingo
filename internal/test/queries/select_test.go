package queries_test

import (
	. "github.com/onsi/gomega"
	"github.com/weworksandbox/lingo/pkg/core/expression"

	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/internal/test/schema/qsakila/qcategory"
	"github.com/weworksandbox/lingo/internal/test/schema/qsakila/qfilmactor"
	"github.com/weworksandbox/lingo/internal/test/schema/qsakila/qfilmcategory"
	"github.com/weworksandbox/lingo/internal/test/schema/qsakila/qfilmtext"
	"github.com/weworksandbox/lingo/internal/test/schema/qsakila/qinventory"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/dialect"
	"github.com/weworksandbox/lingo/pkg/core/expressions"
	"github.com/weworksandbox/lingo/pkg/core/join"
	"github.com/weworksandbox/lingo/pkg/core/query"
	"github.com/weworksandbox/lingo/pkg/core/sort"
)

var selectQueries = []Query{
	{
		Name:      "InventoryIDAndFilmID_ForStoreID",
		Benchmark: true,
		Params: Params{
			Dialect: dialect.Default{},
			SQL: func() core.Expression {
				const (
					storeId = 2
				)
				return query.Select(
					qinventory.InventoryId(),
					qinventory.FilmId(),
				).From(
					qinventory.Q(),
				).Where(
					qinventory.StoreId().Eq(storeId),
				)
			},
			SQLAssert: ContainSubstring(trimQuery(`
					SELECT inventory.inventory_id, inventory.film_id
					FROM sakila.inventory`)),
			ValuesAssert: AllInSlice(
				BeEquivalentTo(2),
			),
			ErrAssert: BeNil(),
		},
	},
	{
		Name:      "NumFilms_ForActorID",
		Benchmark: true,
		Params: Params{
			Dialect: dialect.Default{},
			SQL: func() core.Expression {
				const (
					actorID = 10
				)
				return query.Select(
					expressions.Count(qfilmactor.FilmId()),
				).From(
					qfilmactor.Q(),
				).Where(
					qfilmactor.ActorId().Eq(actorID),
				)
			},
			SQLAssert: ContainSubstring(trimQuery(`
					SELECT COUNT(film_actor.film_id)
					FROM sakila.film_actor
					WHERE film_actor.actor_id = ?
			`)),
			ValuesAssert: AllInSlice(
				BeEquivalentTo(10),
			),
			ErrAssert: BeNil(),
		},
	},
	{
		// Note this Query uses a pointer to actorID
		Name:      "MovieTitlesByCategory_ForActor_CategoryAsc",
		Benchmark: true,
		Params: Params{
			Dialect: dialect.Default{},
			SQL: func() core.Expression {
				var (
					actorID = int16(10)
				)

				fa := qfilmactor.As("fa")
				fc := qfilmcategory.As("fc")
				ft := qfilmtext.As("ft")
				cat := qcategory.As("cat")

				return query.Select(
					ft.Title(), cat.Name(),
				).From(
					fa,
				).Join(
					fc, join.Inner, fc.FilmId().EqPath(fa.FilmId()),
				).Join(
					ft, join.Inner, ft.FilmId().EqPath(fa.FilmId()),
				).Join(
					cat, join.Inner, fc.CategoryId().EqPath(cat.CategoryId()),
				).Where(
					fa.ActorId().EqPath(expression.NewValue(&actorID)),
				).OrderBy(
					cat.Name(), sort.Ascending,
				)
			},
			SQLAssert: ContainSubstring(trimQuery(`
					SELECT ft.title, cat.name
					FROM sakila.film_actor AS fa
					INNER JOIN sakila.film_category AS fc
						ON fc.film_id = fa.film_id
					INNER JOIN sakila.film_text AS ft
						ON ft.film_id = fa.film_id
					INNER JOIN sakila.category AS cat
						ON fc.category_id = cat.category_id
					WHERE fa.actor_id = ?
					ORDER BY cat.name ASC`)),
			ValuesAssert: AllInSlice(
				BeEquivalentTo(10),
			),
			ErrAssert: BeNil(),
		},
	},
}
