// Code generated by Lingo for table sakila.staff_list - DO NOT EDIT

// +build !nolingo

package tstafflist

import (
	"github.com/weworksandbox/lingo/expr"
)

var instance = New()

func T() TStaffList {
	return instance
}

func Id() expr.Int8 {
	return instance.Id()
}

func Name() expr.String {
	return instance.Name()
}

func Address() expr.String {
	return instance.Address()
}

func ZipCode() expr.String {
	return instance.ZipCode()
}

func Phone() expr.String {
	return instance.Phone()
}

func City() expr.String {
	return instance.City()
}

func Country() expr.String {
	return instance.Country()
}

func SID() expr.Int8 {
	return instance.SID()
}
