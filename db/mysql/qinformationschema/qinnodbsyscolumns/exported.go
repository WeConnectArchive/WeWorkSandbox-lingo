// Code generated by Lingo for table information_schema.INNODB_SYS_COLUMNS - DO NOT EDIT

package qinnodbsyscolumns

import "github.com/weworksandbox/lingo/core/path"

var instance = New()

func Q() QInnodbSysColumns {
	return instance
}

func TableId() path.Int64Path {
	return instance.tableId
}

func Name() path.StringPath {
	return instance.name
}

func Pos() path.Int64Path {
	return instance.pos
}

func Mtype() path.IntPath {
	return instance.mtype
}

func Prtype() path.IntPath {
	return instance.prtype
}

func Len() path.IntPath {
	return instance.len
}
