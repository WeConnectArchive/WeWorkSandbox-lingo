// Code generated by Lingo for table sakila.nicer_but_slower_film_list - DO NOT EDIT

// +build !nolingo

package tnicerbutslowerfilmlist

import (
	"github.com/weworksandbox/lingo/expr"
)

var instance = New()

func T() TNicerButSlowerFilmList {
	return instance
}

func FID() expr.Int16 {
	return instance.FID()
}

func Title() expr.String {
	return instance.Title()
}

func Description() expr.String {
	return instance.Description()
}

func Category() expr.String {
	return instance.Category()
}

func Price() expr.Binary {
	return instance.Price()
}

func Length() expr.Int16 {
	return instance.Length()
}

func Rating() expr.String {
	return instance.Rating()
}

func Actors() expr.String {
	return instance.Actors()
}
