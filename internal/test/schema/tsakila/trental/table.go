// Code generated by Lingo for table sakila.rental - DO NOT EDIT

// +build !nolingo

package trental

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) TRental {
	t := New()
	if alias != "" {
		t.alias = expr.Lit(alias)
	}
	return t
}

func New() TRental {
	return TRental{}
}

type TRental struct {
	alias lingo.Expression
}

// lingo.Table Functions

func (t TRental) GetTableName() string {
	return "rental"
}

func (t TRental) GetColumns() []lingo.Expression {
	return []lingo.Expression{
		t.RentalId(),
		t.RentalDate(),
		t.InventoryId(),
		t.CustomerId(),
		t.ReturnDate(),
		t.StaffId(),
		t.LastUpdate(),
	}
}

func (t TRental) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return expr.Table(t).ToSQL(d)
}

func (t TRental) GetName() lingo.Expression {
	return expr.TableName(t)
}

func (t TRental) GetAlias() lingo.Expression {
	return t.alias
}

// Column Functions

func (t TRental) RentalId() expr.Int {
	return expr.Column(t, expr.Lit("rental_id")).ToSQL
}

func (t TRental) RentalDate() expr.Time {
	return expr.Column(t, expr.Lit("rental_date")).ToSQL
}

func (t TRental) InventoryId() expr.Int32Column {
	return expr.Int32Column{
		Table:  t,
		Column: expr.Lit("inventory_id"),
	}
}

func (t TRental) CustomerId() expr.Int16 {
	return expr.Column(t, expr.Lit("customer_id")).ToSQL
}

func (t TRental) ReturnDate() expr.Time {
	return expr.Column(t, expr.Lit("return_date")).ToSQL
}

func (t TRental) StaffId() expr.Int8 {
	return expr.Column(t, expr.Lit("staff_id")).ToSQL
}

func (t TRental) LastUpdate() expr.Time {
	return expr.Column(t, expr.Lit("last_update")).ToSQL
}
