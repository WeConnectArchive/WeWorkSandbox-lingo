// Code generated by Lingo for table sakila.inventory - DO NOT EDIT

// +build !nolingo

package tinventory

import (
	"github.com/weworksandbox/lingo/expr"
)

var instance = New()

func T() TInventory {
	return instance
}

func InventoryId() expr.Int32 {
	return instance.InventoryId()
}

func FilmId() expr.Int16 {
	return instance.FilmId()
}

func StoreId() expr.Int8 {
	return instance.StoreId()
}

func LastUpdate() expr.Time {
	return instance.LastUpdate()
}
