package queries_test

import (
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/execute"
	"github.com/weworksandbox/lingo/expr"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/internal/test/schema/tsakila/trental"
	"github.com/weworksandbox/lingo/query"
)

var insertQueries = []QueryTest{
	{
		Name:          "InsertIntoRental_CurTSAndDateAddCurrTS",
		PendingReason: "Insert is currently broken",
		Params: Params{
			Dialect: DefaultDialect,
			SQL: func() lingo.Expression {
				const (
					inventoryID = int32(470)
				)
				return query.InsertInto(
					trental.T(),
				).Columns(
					trental.RentalDate(), trental.InventoryId(), trental.CustomerId(),
					trental.ReturnDate(), trental.StaffId(),
				).Values(
					expr.CurrentTimestamp(), inventoryID, int16(482), expr.CurrentTimestamp(), int8(2),
				)
			},
			SQLStrAssert: EqString(trimQuery(`
					INSERT INTO rental (rental_date, inventory_id, customer_id, return_date, staff_id)
					VALUES (CURRENT_TIMESTAMP, ?, ?, CURRENT_TIMESTAMP, ?)`,
			)),
			SQLValuesAssert: AllInSlice(
				Equal(int32(470)),
				Equal(int16(482)),
				Equal(int8(2)),
			),
			ExecuteParams: ExecuteParams{
				Type:         execute.QTExec,
				AssertValues: executeResult(16050, 1),
			},
		},
	},
}
