// Code generated by Lingo for table information_schema.CHARACTER_SETS - DO NOT EDIT

package qcharactersets

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/path"
)

func As(alias string) QCharacterSets {
	return newQCharacterSets(alias)
}

func New() QCharacterSets {
	return newQCharacterSets("")
}

func newQCharacterSets(alias string) QCharacterSets {
	q := QCharacterSets{_alias: alias}
	q.characterSetName = path.NewStringPath(q, "CHARACTER_SET_NAME")
	q.defaultCollateName = path.NewStringPath(q, "DEFAULT_COLLATE_NAME")
	q.description = path.NewStringPath(q, "DESCRIPTION")
	q.maxlen = path.NewInt64Path(q, "MAXLEN")
	return q
}

type QCharacterSets struct {
	_alias             string
	characterSetName   path.StringPath
	defaultCollateName path.StringPath
	description        path.StringPath
	maxlen             path.Int64Path
}

// core.Table Functions

func (q QCharacterSets) GetColumns() []core.Column {
	return []core.Column{
		q.characterSetName,
		q.defaultCollateName,
		q.description,
		q.maxlen,
	}
}

func (q QCharacterSets) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QCharacterSets) GetAlias() string {
	return q._alias
}

func (q QCharacterSets) GetName() string {
	return "CHARACTER_SETS"
}

func (q QCharacterSets) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QCharacterSets) CharacterSetName() path.StringPath {
	return q.characterSetName
}

func (q QCharacterSets) DefaultCollateName() path.StringPath {
	return q.defaultCollateName
}

func (q QCharacterSets) Description() path.StringPath {
	return q.description
}

func (q QCharacterSets) Maxlen() path.Int64Path {
	return q.maxlen
}
