// Code generated by Lingo for table information_schema.ROUTINES - DO NOT EDIT

package qroutines

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/path"
)

func As(alias string) QRoutines {
	return newQRoutines(alias)
}

func New() QRoutines {
	return newQRoutines("")
}

func newQRoutines(alias string) QRoutines {
	q := QRoutines{_alias: alias}
	q.specificName = path.NewStringPath(q, "SPECIFIC_NAME")
	q.routineCatalog = path.NewStringPath(q, "ROUTINE_CATALOG")
	q.routineSchema = path.NewStringPath(q, "ROUTINE_SCHEMA")
	q.routineName = path.NewStringPath(q, "ROUTINE_NAME")
	q.routineType = path.NewStringPath(q, "ROUTINE_TYPE")
	q.dataType = path.NewStringPath(q, "DATA_TYPE")
	q.characterMaximumLength = path.NewIntPath(q, "CHARACTER_MAXIMUM_LENGTH")
	q.characterOctetLength = path.NewIntPath(q, "CHARACTER_OCTET_LENGTH")
	q.numericPrecision = path.NewInt64Path(q, "NUMERIC_PRECISION")
	q.numericScale = path.NewIntPath(q, "NUMERIC_SCALE")
	q.datetimePrecision = path.NewInt64Path(q, "DATETIME_PRECISION")
	q.characterSetName = path.NewStringPath(q, "CHARACTER_SET_NAME")
	q.collationName = path.NewStringPath(q, "COLLATION_NAME")
	q.dtdIdentifier = path.NewStringPath(q, "DTD_IDENTIFIER")
	q.routineBody = path.NewStringPath(q, "ROUTINE_BODY")
	q.routineDefinition = path.NewStringPath(q, "ROUTINE_DEFINITION")
	q.externalName = path.NewStringPath(q, "EXTERNAL_NAME")
	q.externalLanguage = path.NewStringPath(q, "EXTERNAL_LANGUAGE")
	q.parameterStyle = path.NewStringPath(q, "PARAMETER_STYLE")
	q.isDeterministic = path.NewStringPath(q, "IS_DETERMINISTIC")
	q.sqlDataAccess = path.NewStringPath(q, "SQL_DATA_ACCESS")
	q.sqlPath = path.NewStringPath(q, "SQL_PATH")
	q.securityType = path.NewStringPath(q, "SECURITY_TYPE")
	q.created = path.NewTimePath(q, "CREATED")
	q.lastAltered = path.NewTimePath(q, "LAST_ALTERED")
	q.sqlMode = path.NewStringPath(q, "SQL_MODE")
	q.routineComment = path.NewStringPath(q, "ROUTINE_COMMENT")
	q.definer = path.NewStringPath(q, "DEFINER")
	q.characterSetClient = path.NewStringPath(q, "CHARACTER_SET_CLIENT")
	q.collationConnection = path.NewStringPath(q, "COLLATION_CONNECTION")
	q.databaseCollation = path.NewStringPath(q, "DATABASE_COLLATION")
	return q
}

type QRoutines struct {
	_alias                 string
	specificName           path.StringPath
	routineCatalog         path.StringPath
	routineSchema          path.StringPath
	routineName            path.StringPath
	routineType            path.StringPath
	dataType               path.StringPath
	characterMaximumLength path.IntPath
	characterOctetLength   path.IntPath
	numericPrecision       path.Int64Path
	numericScale           path.IntPath
	datetimePrecision      path.Int64Path
	characterSetName       path.StringPath
	collationName          path.StringPath
	dtdIdentifier          path.StringPath
	routineBody            path.StringPath
	routineDefinition      path.StringPath
	externalName           path.StringPath
	externalLanguage       path.StringPath
	parameterStyle         path.StringPath
	isDeterministic        path.StringPath
	sqlDataAccess          path.StringPath
	sqlPath                path.StringPath
	securityType           path.StringPath
	created                path.TimePath
	lastAltered            path.TimePath
	sqlMode                path.StringPath
	routineComment         path.StringPath
	definer                path.StringPath
	characterSetClient     path.StringPath
	collationConnection    path.StringPath
	databaseCollation      path.StringPath
}

// core.Table Functions

func (q QRoutines) GetColumns() []core.Column {
	return []core.Column{
		q.specificName,
		q.routineCatalog,
		q.routineSchema,
		q.routineName,
		q.routineType,
		q.dataType,
		q.characterMaximumLength,
		q.characterOctetLength,
		q.numericPrecision,
		q.numericScale,
		q.datetimePrecision,
		q.characterSetName,
		q.collationName,
		q.dtdIdentifier,
		q.routineBody,
		q.routineDefinition,
		q.externalName,
		q.externalLanguage,
		q.parameterStyle,
		q.isDeterministic,
		q.sqlDataAccess,
		q.sqlPath,
		q.securityType,
		q.created,
		q.lastAltered,
		q.sqlMode,
		q.routineComment,
		q.definer,
		q.characterSetClient,
		q.collationConnection,
		q.databaseCollation,
	}
}

func (q QRoutines) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QRoutines) GetAlias() string {
	return q._alias
}

func (q QRoutines) GetName() string {
	return "ROUTINES"
}

func (q QRoutines) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QRoutines) SpecificName() path.StringPath {
	return q.specificName
}

func (q QRoutines) RoutineCatalog() path.StringPath {
	return q.routineCatalog
}

func (q QRoutines) RoutineSchema() path.StringPath {
	return q.routineSchema
}

func (q QRoutines) RoutineName() path.StringPath {
	return q.routineName
}

func (q QRoutines) RoutineType() path.StringPath {
	return q.routineType
}

func (q QRoutines) DataType() path.StringPath {
	return q.dataType
}

func (q QRoutines) CharacterMaximumLength() path.IntPath {
	return q.characterMaximumLength
}

func (q QRoutines) CharacterOctetLength() path.IntPath {
	return q.characterOctetLength
}

func (q QRoutines) NumericPrecision() path.Int64Path {
	return q.numericPrecision
}

func (q QRoutines) NumericScale() path.IntPath {
	return q.numericScale
}

func (q QRoutines) DatetimePrecision() path.Int64Path {
	return q.datetimePrecision
}

func (q QRoutines) CharacterSetName() path.StringPath {
	return q.characterSetName
}

func (q QRoutines) CollationName() path.StringPath {
	return q.collationName
}

func (q QRoutines) DtdIdentifier() path.StringPath {
	return q.dtdIdentifier
}

func (q QRoutines) RoutineBody() path.StringPath {
	return q.routineBody
}

func (q QRoutines) RoutineDefinition() path.StringPath {
	return q.routineDefinition
}

func (q QRoutines) ExternalName() path.StringPath {
	return q.externalName
}

func (q QRoutines) ExternalLanguage() path.StringPath {
	return q.externalLanguage
}

func (q QRoutines) ParameterStyle() path.StringPath {
	return q.parameterStyle
}

func (q QRoutines) IsDeterministic() path.StringPath {
	return q.isDeterministic
}

func (q QRoutines) SqlDataAccess() path.StringPath {
	return q.sqlDataAccess
}

func (q QRoutines) SqlPath() path.StringPath {
	return q.sqlPath
}

func (q QRoutines) SecurityType() path.StringPath {
	return q.securityType
}

func (q QRoutines) Created() path.TimePath {
	return q.created
}

func (q QRoutines) LastAltered() path.TimePath {
	return q.lastAltered
}

func (q QRoutines) SqlMode() path.StringPath {
	return q.sqlMode
}

func (q QRoutines) RoutineComment() path.StringPath {
	return q.routineComment
}

func (q QRoutines) Definer() path.StringPath {
	return q.definer
}

func (q QRoutines) CharacterSetClient() path.StringPath {
	return q.characterSetClient
}

func (q QRoutines) CollationConnection() path.StringPath {
	return q.collationConnection
}

func (q QRoutines) DatabaseCollation() path.StringPath {
	return q.databaseCollation
}
