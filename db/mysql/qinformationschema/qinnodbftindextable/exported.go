// Code generated by Lingo for table information_schema.INNODB_FT_INDEX_TABLE - DO NOT EDIT

package qinnodbftindextable

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QInnodbFtIndexTable {
	return instance
}

func Word() path.String {
	return instance.word
}

func FirstDocId() path.Int64 {
	return instance.firstDocId
}

func LastDocId() path.Int64 {
	return instance.lastDocId
}

func DocCount() path.Int64 {
	return instance.docCount
}

func DocId() path.Int64 {
	return instance.docId
}

func Position() path.Int64 {
	return instance.position
}
