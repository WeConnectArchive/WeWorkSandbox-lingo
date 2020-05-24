package execute

// QueryType is used as the trace Event name but can be used for other cases.
type QueryType int

const (
	QTUnknown QueryType = iota
	QTRow
	QTRows
	QTExec
)

func (qt QueryType) String() string {
	switch qt {
	case QTRow:
		return "QueryRow"
	case QTRows:
		return "QueryRows"
	case QTExec:
		return "QueryExec"
	}
	return "UnknownQuery"
}
