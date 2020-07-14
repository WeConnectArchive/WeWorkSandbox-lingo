package queries_test

import (
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/execute"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/join"
	"github.com/weworksandbox/lingo/expr/sort"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/internal/test/schema/tsakila/tcategory"
	"github.com/weworksandbox/lingo/internal/test/schema/tsakila/tfilmactor"
	"github.com/weworksandbox/lingo/internal/test/schema/tsakila/tfilmcategory"
	"github.com/weworksandbox/lingo/internal/test/schema/tsakila/tfilmtext"
	"github.com/weworksandbox/lingo/internal/test/schema/tsakila/tinventory"
	"github.com/weworksandbox/lingo/query"
)

var selectQueries = []QueryTest{
	{
		Name:      "CountInventoryID_ForStoreID",
		Benchmark: true,
		Params: Params{
			Dialect: DefaultDialect,
			SQL: func() lingo.Expression {
				const (
					storeID = 2
				)
				return query.Select(
					expr.Count(tinventory.InventoryId()),
				).From(
					tinventory.T(),
				).Where(
					tinventory.StoreId().Eq(storeID),
				)
			},
			SQLStrAssert: EqString(trimQuery(`
					SELECT COUNT(inventory.inventory_id)
					FROM inventory
					WHERE inventory.store_id = ?`,
			)),
			SQLValuesAssert: AllInSlice(
				BeEquivalentTo(2),
			),
			ExecuteParams: ExecuteParams{
				Type:     execute.QTRow,
				ScanData: row(ptrI16(0)),
				Assert:   rows(row(ptrI16(2311))),
			},
		},
	},
	{
		Name:      "CountSakilaInventoryID_ForStoreID", // Includes schema in output
		Benchmark: true,
		Params: Params{
			Dialect: DefaultDialectWithSchema,
			SQL: func() lingo.Expression {
				const (
					storeID = 2
				)
				return query.Select(
					expr.Count(tinventory.InventoryId()),
				).From(
					tinventory.T(),
				).Where(
					tinventory.StoreId().Eq(storeID),
				)
			},
			SQLStrAssert: EqString(trimQuery(`
					SELECT COUNT(inventory.inventory_id)
					FROM sakila.inventory
					WHERE inventory.store_id = ?`,
			)),
			SQLValuesAssert: AllInSlice(
				BeEquivalentTo(2),
			),
			ExecuteParams: ExecuteParams{
				Type:     execute.QTRow,
				ScanData: row(ptrI16(0)),
				Assert:   rows(row(ptrI16(2311))),
			},
		},
	},
	{
		Name:      "CountFilms_ForActorID",
		Benchmark: true,
		Params: Params{
			Dialect: DefaultDialect,
			SQL: func() lingo.Expression {
				const (
					actorID = 10
				)
				return query.Select(
					expr.Count(tfilmactor.FilmId()),
				).From(
					tfilmactor.T(),
				).Where(
					tfilmactor.ActorId().Eq(actorID),
				)
			},
			SQLStrAssert: EqString(trimQuery(`
					SELECT COUNT(film_actor.film_id)
					FROM film_actor
					WHERE film_actor.actor_id = ?
			`)),
			SQLValuesAssert: AllInSlice(
				BeEquivalentTo(10),
			),
			ExecuteParams: ExecuteParams{
				Type:     execute.QTRow,
				ScanData: row(ptrI32(0)),
				Assert:   rows(row(ptrI32(22))),
			},
		},
	},
	{
		// Note this QueryTest uses a pointer to actorID
		Name:      "MovieTitlesByCategory_ForActorIdPtr_CategoryAsc",
		Benchmark: true,
		Params: Params{
			Dialect: DefaultDialect,
			SQL: func() lingo.Expression {
				var (
					actorID = int16(10)
				)

				fa := tfilmactor.As("fa")
				fc := tfilmcategory.As("fc")
				ft := tfilmtext.As("ft")
				cat := tcategory.As("cat")

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
					fa.ActorId().EqPath(expr.NewValue(&actorID)),
				).OrderBy(
					cat.Name(), sort.Ascending,
				)
			},
			SQLStrAssert: ContainSubstring(trimQuery(`
					SELECT ft.title, cat.name
					FROM film_actor AS fa
					INNER JOIN film_category AS fc
						ON fc.film_id = fa.film_id
					INNER JOIN film_text AS ft
						ON ft.film_id = fa.film_id
					INNER JOIN category AS cat
						ON fc.category_id = cat.category_id
					WHERE fa.actor_id = ?
					ORDER BY cat.name ASC`)),
			SQLValuesAssert: AllInSlice(
				BeEquivalentTo(10),
			),
			ExecuteParams: ExecuteParams{
				Type:     execute.QTRows,
				ScanData: row(ptrStr(""), ptrStr("")),
				Assert: rows(
					row(ptrStr("WATERFRONT DELIVERANCE"), ptrStr("Action")),
					row(ptrStr("LORD ARIZONA"), ptrStr("Action")),
					row(ptrStr("PUNK DIVORCE"), ptrStr("Animation")),
					row(ptrStr("CROOKED FROGMEN"), ptrStr("Children")),
					row(ptrStr("JEEPERS WEDDING"), ptrStr("Classics")),
					row(ptrStr("PREJUDICE OLEANDER"), ptrStr("Classics")),
					row(ptrStr("LIFE TWISTED"), ptrStr("Comedy")),
					row(ptrStr("ACADEMY DINOSAUR"), ptrStr("Documentary")),
					row(ptrStr("WEDDING APOLLO"), ptrStr("Documentary")),
					row(ptrStr("MOD SECRETARY"), ptrStr("Documentary")),
					row(ptrStr("GOLDFINGER SENSIBILITY"), ptrStr("Drama")),
					row(ptrStr("USUAL UNTOUCHABLES"), ptrStr("Foreign")),
					row(ptrStr("DIVINE RESURRECTION"), ptrStr("Games")),
					row(ptrStr("ALABAMA DEVIL"), ptrStr("Horror")),
					row(ptrStr("REAP UNFAITHFUL"), ptrStr("Horror")),
					row(ptrStr("JAWBREAKER BROOKLYN"), ptrStr("Music")),
					row(ptrStr("WIZARD COLDBLOODED"), ptrStr("Music")),
					row(ptrStr("WON DARES"), ptrStr("Music")),
					row(ptrStr("DRAGONFLY STRANGERS"), ptrStr("New")),
					row(ptrStr("VACATION BOONDOCK"), ptrStr("Sci-Fi")),
					row(ptrStr("SHAKESPEARE SADDLE"), ptrStr("Sports")),
					row(ptrStr("TROUBLE DATE"), ptrStr("Travel")),
				),
			},
		},
	},
}
