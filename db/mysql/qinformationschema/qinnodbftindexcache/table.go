// Code generated by Lingo for table information_schema.INNODB_FT_INDEX_CACHE - DO NOT EDIT

package qinnodbftindexcache

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QInnodbFtIndexCache {
	return newQInnodbFtIndexCache(alias)
}

func New() QInnodbFtIndexCache {
	return newQInnodbFtIndexCache("")
}

func newQInnodbFtIndexCache(alias string) QInnodbFtIndexCache {
	q := QInnodbFtIndexCache{_alias: alias}
	q.word = path.NewStringPath(q, "WORD")
	q.firstDocId = path.NewInt64Path(q, "FIRST_DOC_ID")
	q.lastDocId = path.NewInt64Path(q, "LAST_DOC_ID")
	q.docCount = path.NewInt64Path(q, "DOC_COUNT")
	q.docId = path.NewInt64Path(q, "DOC_ID")
	q.position = path.NewInt64Path(q, "POSITION")
	return q
}

type QInnodbFtIndexCache struct {
	_alias     string
	word       path.String
	firstDocId path.Int64
	lastDocId  path.Int64
	docCount   path.Int64
	docId      path.Int64
	position   path.Int64
}

// core.Table Functions

func (q QInnodbFtIndexCache) GetColumns() []core.Column {
	return []core.Column{
		q.word,
		q.firstDocId,
		q.lastDocId,
		q.docCount,
		q.docId,
		q.position,
	}
}

func (q QInnodbFtIndexCache) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QInnodbFtIndexCache) GetAlias() string {
	return q._alias
}

func (q QInnodbFtIndexCache) GetName() string {
	return "INNODB_FT_INDEX_CACHE"
}

func (q QInnodbFtIndexCache) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbFtIndexCache) Word() path.String {
	return q.word
}

func (q QInnodbFtIndexCache) FirstDocId() path.Int64 {
	return q.firstDocId
}

func (q QInnodbFtIndexCache) LastDocId() path.Int64 {
	return q.lastDocId
}

func (q QInnodbFtIndexCache) DocCount() path.Int64 {
	return q.docCount
}

func (q QInnodbFtIndexCache) DocId() path.Int64 {
	return q.docId
}

func (q QInnodbFtIndexCache) Position() path.Int64 {
	return q.position
}
