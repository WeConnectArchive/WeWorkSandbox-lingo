// Code generated by Lingo for table sakila.city - DO NOT EDIT

// +build !nolingo

package tcity

import (
	"github.com/weworksandbox/lingo/expr/path"
)

var instance = New()

func T() TCity {
	return instance
}

func CityId() path.Int16 {
	return instance.cityId
}

func City() path.String {
	return instance.city
}

func CountryId() path.Int16 {
	return instance.countryId
}

func LastUpdate() path.Time {
	return instance.lastUpdate
}