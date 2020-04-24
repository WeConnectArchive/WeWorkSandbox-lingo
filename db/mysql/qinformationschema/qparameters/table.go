// Code generated by Lingo for table information_schema.PARAMETERS - DO NOT EDIT

package qparameters

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QParameters {
	return newQParameters(alias)
}

func New() QParameters {
	return newQParameters("")
}

func newQParameters(alias string) QParameters {
	q := QParameters{_alias: alias}
	q.specificCatalog = path.NewStringPath(q, "SPECIFIC_CATALOG")
	q.specificSchema = path.NewStringPath(q, "SPECIFIC_SCHEMA")
	q.specificName = path.NewStringPath(q, "SPECIFIC_NAME")
	q.ordinalPosition = path.NewIntPath(q, "ORDINAL_POSITION")
	q.parameterMode = path.NewStringPath(q, "PARAMETER_MODE")
	q.parameterName = path.NewStringPath(q, "PARAMETER_NAME")
	q.dataType = path.NewStringPath(q, "DATA_TYPE")
	q.characterMaximumLength = path.NewIntPath(q, "CHARACTER_MAXIMUM_LENGTH")
	q.characterOctetLength = path.NewIntPath(q, "CHARACTER_OCTET_LENGTH")
	q.numericPrecision = path.NewInt64Path(q, "NUMERIC_PRECISION")
	q.numericScale = path.NewIntPath(q, "NUMERIC_SCALE")
	q.datetimePrecision = path.NewInt64Path(q, "DATETIME_PRECISION")
	q.characterSetName = path.NewStringPath(q, "CHARACTER_SET_NAME")
	q.collationName = path.NewStringPath(q, "COLLATION_NAME")
	q.dtdIdentifier = path.NewStringPath(q, "DTD_IDENTIFIER")
	q.routineType = path.NewStringPath(q, "ROUTINE_TYPE")
	return q
}

type QParameters struct {
	_alias                 string
	specificCatalog        path.String
	specificSchema         path.String
	specificName           path.String
	ordinalPosition        path.Int
	parameterMode          path.String
	parameterName          path.String
	dataType               path.String
	characterMaximumLength path.Int
	characterOctetLength   path.Int
	numericPrecision       path.Int64
	numericScale           path.Int
	datetimePrecision      path.Int64
	characterSetName       path.String
	collationName          path.String
	dtdIdentifier          path.String
	routineType            path.String
}

// core.Table Functions

func (q QParameters) GetColumns() []core.Column {
	return []core.Column{
		q.specificCatalog,
		q.specificSchema,
		q.specificName,
		q.ordinalPosition,
		q.parameterMode,
		q.parameterName,
		q.dataType,
		q.characterMaximumLength,
		q.characterOctetLength,
		q.numericPrecision,
		q.numericScale,
		q.datetimePrecision,
		q.characterSetName,
		q.collationName,
		q.dtdIdentifier,
		q.routineType,
	}
}

func (q QParameters) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QParameters) GetAlias() string {
	return q._alias
}

func (q QParameters) GetName() string {
	return "PARAMETERS"
}

func (q QParameters) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QParameters) SpecificCatalog() path.String {
	return q.specificCatalog
}

func (q QParameters) SpecificSchema() path.String {
	return q.specificSchema
}

func (q QParameters) SpecificName() path.String {
	return q.specificName
}

func (q QParameters) OrdinalPosition() path.Int {
	return q.ordinalPosition
}

func (q QParameters) ParameterMode() path.String {
	return q.parameterMode
}

func (q QParameters) ParameterName() path.String {
	return q.parameterName
}

func (q QParameters) DataType() path.String {
	return q.dataType
}

func (q QParameters) CharacterMaximumLength() path.Int {
	return q.characterMaximumLength
}

func (q QParameters) CharacterOctetLength() path.Int {
	return q.characterOctetLength
}

func (q QParameters) NumericPrecision() path.Int64 {
	return q.numericPrecision
}

func (q QParameters) NumericScale() path.Int {
	return q.numericScale
}

func (q QParameters) DatetimePrecision() path.Int64 {
	return q.datetimePrecision
}

func (q QParameters) CharacterSetName() path.String {
	return q.characterSetName
}

func (q QParameters) CollationName() path.String {
	return q.collationName
}

func (q QParameters) DtdIdentifier() path.String {
	return q.dtdIdentifier
}

func (q QParameters) RoutineType() path.String {
	return q.routineType
}
