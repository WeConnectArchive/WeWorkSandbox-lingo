// Code generated by Lingo for table sakila.rental - DO NOT EDIT

// +build !nolingo

package qrental

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expr/path"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func As(alias string) QRental {
	return newQRental(alias)
}

func New() QRental {
	return newQRental("")
}

func newQRental(alias string) QRental {
	q := QRental{_alias: alias}
	q.rentalId = path.NewInt(q, "rental_id")
	q.rentalDate = path.NewTime(q, "rental_date")
	q.inventoryId = path.NewInt32(q, "inventory_id")
	q.customerId = path.NewInt16(q, "customer_id")
	q.returnDate = path.NewTime(q, "return_date")
	q.staffId = path.NewInt8(q, "staff_id")
	q.lastUpdate = path.NewTime(q, "last_update")
	return q
}

type QRental struct {
	_alias      string
	rentalId    path.Int
	rentalDate  path.Time
	inventoryId path.Int32
	customerId  path.Int16
	returnDate  path.Time
	staffId     path.Int8
	lastUpdate  path.Time
}

// core.Table Functions

func (q QRental) GetColumns() []core.Column {
	return []core.Column{
		q.rentalId,
		q.rentalDate,
		q.inventoryId,
		q.customerId,
		q.returnDate,
		q.staffId,
		q.lastUpdate,
	}
}

func (q QRental) ToSQL(d core.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QRental) GetAlias() string {
	return q._alias
}

func (q QRental) GetName() string {
	return "rental"
}

func (q QRental) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QRental) RentalId() path.Int {
	return q.rentalId
}

func (q QRental) RentalDate() path.Time {
	return q.rentalDate
}

func (q QRental) InventoryId() path.Int32 {
	return q.inventoryId
}

func (q QRental) CustomerId() path.Int16 {
	return q.customerId
}

func (q QRental) ReturnDate() path.Time {
	return q.returnDate
}

func (q QRental) StaffId() path.Int8 {
	return q.staffId
}

func (q QRental) LastUpdate() path.Time {
	return q.lastUpdate
}
