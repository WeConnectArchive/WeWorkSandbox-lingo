// Code generated by Lingo for table sakila.customer_list - DO NOT EDIT

package qcustomerlist

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QCustomerList {
	return instance
}

func Id() path.Int16 {
	return instance.id
}

func Name() path.String {
	return instance.name
}

func Address() path.String {
	return instance.address
}

func ZipCode() path.String {
	return instance.zipCode
}

func Phone() path.String {
	return instance.phone
}

func City() path.String {
	return instance.city
}

func Country() path.String {
	return instance.country
}

func Notes() path.String {
	return instance.notes
}

func Sid() path.Int8 {
	return instance.sid
}
