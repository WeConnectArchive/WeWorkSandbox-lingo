// Code generated by Lingo for table sakila.nicer_but_slower_film_list - DO NOT EDIT

package qnicerbutslowerfilmlist

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QNicerButSlowerFilmList {
	return instance
}

func Fid() path.Int16 {
	return instance.fid
}

func Title() path.String {
	return instance.title
}

func Description() path.String {
	return instance.description
}

func Category() path.String {
	return instance.category
}

func Price() path.Binary {
	return instance.price
}

func Length() path.Int16 {
	return instance.length
}

func Rating() path.String {
	return instance.rating
}

func Actors() path.String {
	return instance.actors
}