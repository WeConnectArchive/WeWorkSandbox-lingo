// Code generated by Lingo for table sakila.actor - DO NOT EDIT

// +build !nolingo

package qactor

import "github.com/weworksandbox/lingo/pkg/core/expr/path"

var instance = New()

func Q() QActor {
	return instance
}

func ActorId() path.Int16 {
	return instance.actorId
}

func FirstName() path.String {
	return instance.firstName
}

func LastName() path.String {
	return instance.lastName
}

func LastUpdate() path.Time {
	return instance.lastUpdate
}
