// Code generated by Lingo for table information_schema.INNODB_TEMP_TABLE_INFO - DO NOT EDIT

package qinnodbtemptableinfo

import "github.com/weworksandbox/lingo/core/path"

var instance = New()

func Q() QInnodbTempTableInfo {
	return instance
}

func TableId() path.Int64Path {
	return instance.tableId
}

func Name() path.StringPath {
	return instance.name
}

func NCols() path.IntPath {
	return instance.nCols
}

func Space() path.IntPath {
	return instance.space
}

func PerTableTablespace() path.StringPath {
	return instance.perTableTablespace
}

func IsCompressed() path.StringPath {
	return instance.isCompressed
}
