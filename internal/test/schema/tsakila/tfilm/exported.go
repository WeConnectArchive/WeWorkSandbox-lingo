// Code generated by Lingo for table sakila.film - DO NOT EDIT

// +build !nolingo

package tfilm

import (
	"github.com/weworksandbox/lingo/expr"
)

var instance = New()

func T() TFilm {
	return instance
}

func FilmId() expr.Int16 {
	return instance.FilmId()
}

func Title() expr.String {
	return instance.Title()
}

func Description() expr.String {
	return instance.Description()
}

func ReleaseYear() expr.UnsupportedType {
	return instance.ReleaseYear()
}

func LanguageId() expr.Int8 {
	return instance.LanguageId()
}

func OriginalLanguageId() expr.Int8 {
	return instance.OriginalLanguageId()
}

func RentalDuration() expr.Int8 {
	return instance.RentalDuration()
}

func RentalRate() expr.Binary {
	return instance.RentalRate()
}

func Length() expr.Int16 {
	return instance.Length()
}

func ReplacementCost() expr.Binary {
	return instance.ReplacementCost()
}

func Rating() expr.String {
	return instance.Rating()
}

func SpecialFeatures() expr.String {
	return instance.SpecialFeatures()
}

func LastUpdate() expr.Time {
	return instance.LastUpdate()
}
