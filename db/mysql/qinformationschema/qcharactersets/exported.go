// Code generated by Lingo for table information_schema.CHARACTER_SETS - DO NOT EDIT

package qcharactersets

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QCharacterSets {
	return instance
}

func CharacterSetName() path.String {
	return instance.characterSetName
}

func DefaultCollateName() path.String {
	return instance.defaultCollateName
}

func Description() path.String {
	return instance.description
}

func Maxlen() path.Int64 {
	return instance.maxlen
}
