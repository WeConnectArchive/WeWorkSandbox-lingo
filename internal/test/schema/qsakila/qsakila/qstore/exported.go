// Code generated by Lingo for table sakila.store - DO NOT EDIT

package qstore

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QStore {
	return instance
}

func StoreId() path.BoolPath {
	return instance.storeId
}

func ManagerStaffId() path.BoolPath {
	return instance.managerStaffId
}

func AddressId() path.Int16Path {
	return instance.addressId
}

func LastUpdate() path.TimePath {
	return instance.lastUpdate
}