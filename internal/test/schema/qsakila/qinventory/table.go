// Code generated by Lingo for table sakila.inventory - DO NOT EDIT

// +build !nolingo

package qinventory

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/path"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) QInventory {
	return newQInventory(alias)
}

func New() QInventory {
	return newQInventory("")
}

func newQInventory(alias string) QInventory {
	q := QInventory{
		_alias: alias,
	}
	q.inventoryId = path.NewInt32(q, "inventory_id")
	q.filmId = path.NewInt16(q, "film_id")
	q.storeId = path.NewInt8(q, "store_id")
	q.lastUpdate = path.NewTime(q, "last_update")
	return q
}

type QInventory struct {
	_alias string

	inventoryId path.Int32
	filmId      path.Int16
	storeId     path.Int8
	lastUpdate  path.Time
}

// lingo.Table Functions

func (q QInventory) GetColumns() []lingo.Column {
	return []lingo.Column{
		q.inventoryId,
		q.filmId,
		q.storeId,
		q.lastUpdate,
	}
}

func (q QInventory) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QInventory) GetAlias() string {
	return q._alias
}

func (q QInventory) GetName() string {
	return "inventory"
}

func (q QInventory) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QInventory) InventoryId() path.Int32 {
	return q.inventoryId
}

func (q QInventory) FilmId() path.Int16 {
	return q.filmId
}

func (q QInventory) StoreId() path.Int8 {
	return q.storeId
}

func (q QInventory) LastUpdate() path.Time {
	return q.lastUpdate
}
