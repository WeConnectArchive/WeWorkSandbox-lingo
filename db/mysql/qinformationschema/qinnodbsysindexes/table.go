// Code generated by Lingo for table information_schema.INNODB_SYS_INDEXES - DO NOT EDIT

package qinnodbsysindexes

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/path"
)

func As(alias string) QInnodbSysIndexes {
	return newQInnodbSysIndexes(alias)
}

func New() QInnodbSysIndexes {
	return newQInnodbSysIndexes("")
}

func newQInnodbSysIndexes(alias string) QInnodbSysIndexes {
	q := QInnodbSysIndexes{_alias: alias}
	q.indexId = path.NewInt64Path(q, "INDEX_ID")
	q.name = path.NewStringPath(q, "NAME")
	q.tableId = path.NewInt64Path(q, "TABLE_ID")
	q.__type = path.NewIntPath(q, "TYPE")
	q.nFields = path.NewIntPath(q, "N_FIELDS")
	q.pageNo = path.NewIntPath(q, "PAGE_NO")
	q.space = path.NewIntPath(q, "SPACE")
	q.mergeThreshold = path.NewIntPath(q, "MERGE_THRESHOLD")
	return q
}

type QInnodbSysIndexes struct {
	_alias         string
	indexId        path.Int64Path
	name           path.StringPath
	tableId        path.Int64Path
	__type         path.IntPath
	nFields        path.IntPath
	pageNo         path.IntPath
	space          path.IntPath
	mergeThreshold path.IntPath
}

// core.Table Functions

func (q QInnodbSysIndexes) GetColumns() []core.Column {
	return []core.Column{
		q.indexId,
		q.name,
		q.tableId,
		q.__type,
		q.nFields,
		q.pageNo,
		q.space,
		q.mergeThreshold,
	}
}

func (q QInnodbSysIndexes) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QInnodbSysIndexes) GetAlias() string {
	return q._alias
}

func (q QInnodbSysIndexes) GetName() string {
	return "INNODB_SYS_INDEXES"
}

func (q QInnodbSysIndexes) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbSysIndexes) IndexId() path.Int64Path {
	return q.indexId
}

func (q QInnodbSysIndexes) Name() path.StringPath {
	return q.name
}

func (q QInnodbSysIndexes) TableId() path.Int64Path {
	return q.tableId
}

func (q QInnodbSysIndexes) Type() path.IntPath {
	return q.__type
}

func (q QInnodbSysIndexes) NFields() path.IntPath {
	return q.nFields
}

func (q QInnodbSysIndexes) PageNo() path.IntPath {
	return q.pageNo
}

func (q QInnodbSysIndexes) Space() path.IntPath {
	return q.space
}

func (q QInnodbSysIndexes) MergeThreshold() path.IntPath {
	return q.mergeThreshold
}
