// Code generated by Lingo for table information_schema.FILES - DO NOT EDIT

package qfiles

import "github.com/weworksandbox/lingo/core/path"

var instance = New()

func Q() QFiles {
	return instance
}

func FileId() path.Int64Path {
	return instance.fileId
}

func FileName() path.StringPath {
	return instance.fileName
}

func FileType() path.StringPath {
	return instance.fileType
}

func TablespaceName() path.StringPath {
	return instance.tablespaceName
}

func TableCatalog() path.StringPath {
	return instance.tableCatalog
}

func TableSchema() path.StringPath {
	return instance.tableSchema
}

func TableName() path.StringPath {
	return instance.tableName
}

func LogfileGroupName() path.StringPath {
	return instance.logfileGroupName
}

func LogfileGroupNumber() path.Int64Path {
	return instance.logfileGroupNumber
}

func Engine() path.StringPath {
	return instance.engine
}

func FulltextKeys() path.StringPath {
	return instance.fulltextKeys
}

func DeletedRows() path.Int64Path {
	return instance.deletedRows
}

func UpdateCount() path.Int64Path {
	return instance.updateCount
}

func FreeExtents() path.Int64Path {
	return instance.freeExtents
}

func TotalExtents() path.Int64Path {
	return instance.totalExtents
}

func ExtentSize() path.Int64Path {
	return instance.extentSize
}

func InitialSize() path.Int64Path {
	return instance.initialSize
}

func MaximumSize() path.Int64Path {
	return instance.maximumSize
}

func AutoextendSize() path.Int64Path {
	return instance.autoextendSize
}

func CreationTime() path.TimePath {
	return instance.creationTime
}

func LastUpdateTime() path.TimePath {
	return instance.lastUpdateTime
}

func LastAccessTime() path.TimePath {
	return instance.lastAccessTime
}

func RecoverTime() path.Int64Path {
	return instance.recoverTime
}

func TransactionCounter() path.Int64Path {
	return instance.transactionCounter
}

func Version() path.Int64Path {
	return instance.version
}

func RowFormat() path.StringPath {
	return instance.rowFormat
}

func TableRows() path.Int64Path {
	return instance.tableRows
}

func AvgRowLength() path.Int64Path {
	return instance.avgRowLength
}

func DataLength() path.Int64Path {
	return instance.dataLength
}

func MaxDataLength() path.Int64Path {
	return instance.maxDataLength
}

func IndexLength() path.Int64Path {
	return instance.indexLength
}

func DataFree() path.Int64Path {
	return instance.dataFree
}

func CreateTime() path.TimePath {
	return instance.createTime
}

func UpdateTime() path.TimePath {
	return instance.updateTime
}

func CheckTime() path.TimePath {
	return instance.checkTime
}

func Checksum() path.Int64Path {
	return instance.checksum
}

func Status() path.StringPath {
	return instance.status
}

func Extra() path.StringPath {
	return instance.extra
}
