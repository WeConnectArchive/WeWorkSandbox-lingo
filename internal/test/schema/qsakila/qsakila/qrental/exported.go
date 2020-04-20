// Code generated by Lingo for table sakila.rental - DO NOT EDIT

package qrental

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QRental {
	return instance
}

func RentalId() path.IntPath {
	return instance.rentalId
}

func RentalDate() path.TimePath {
	return instance.rentalDate
}

func InventoryId() path.Int32Path {
	return instance.inventoryId
}

func CustomerId() path.Int16Path {
	return instance.customerId
}

func ReturnDate() path.TimePath {
	return instance.returnDate
}

func StaffId() path.BoolPath {
	return instance.staffId
}

func LastUpdate() path.TimePath {
	return instance.lastUpdate
}
