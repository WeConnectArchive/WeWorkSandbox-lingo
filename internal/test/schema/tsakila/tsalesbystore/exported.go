// Code generated by Lingo for table sakila.sales_by_store - DO NOT EDIT

// +build !nolingo

package tsalesbystore

import (
	"github.com/weworksandbox/lingo/expr"
)

var instance = New()

func T() TSalesByStore {
	return instance
}

func Store() expr.String {
	return instance.Store()
}

func Manager() expr.String {
	return instance.Manager()
}

func TotalSales() expr.Binary {
	return instance.TotalSales()
}
