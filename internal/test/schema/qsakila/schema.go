// Code generated by Lingo for schema sakila - DO NOT EDIT

// +build !nolingo

package qsakila

import (
	"github.com/weworksandbox/lingo"
)

var instance = schema{_name: "sakila"}

func GetSchema() lingo.Name {
	return instance
}

type schema struct {
	_name string
}

func (s schema) GetName() string {
	return s._name
}
