// Code generated by Lingo for table information_schema.ENGINES - DO NOT EDIT

package qengines

import "github.com/weworksandbox/lingo/core/path"

var instance = New()

func Q() QEngines {
	return instance
}

func Engine() path.StringPath {
	return instance.engine
}

func Support() path.StringPath {
	return instance.support
}

func Comment() path.StringPath {
	return instance.comment
}

func Transactions() path.StringPath {
	return instance.transactions
}

func Xa() path.StringPath {
	return instance.xa
}

func Savepoints() path.StringPath {
	return instance.savepoints
}
