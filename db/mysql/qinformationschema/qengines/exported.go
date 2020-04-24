// Code generated by Lingo for table information_schema.ENGINES - DO NOT EDIT

package qengines

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QEngines {
	return instance
}

func Engine() path.String {
	return instance.engine
}

func Support() path.String {
	return instance.support
}

func Comment() path.String {
	return instance.comment
}

func Transactions() path.String {
	return instance.transactions
}

func Xa() path.String {
	return instance.xa
}

func Savepoints() path.String {
	return instance.savepoints
}
