package generator

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func GeneratePackageMembers(settings Settings, info TableInfo, dbToPath DBToPathType) (string, error) {
	tableName := info.Name()
	columns := convertCols(info.Columns(), settings.ReplaceFieldName, dbToPath)

	f := NewFile(ToPackageName(tableName))
	f.HeaderComment(fmt.Sprintf(fmtTableHeaderComment, info.Schema(), tableName))
	f.ImportName(pkgCorePath, "path")
	f.Var().Id(packageInstance).Op("=").Id("New").Call()
	f.Line()
	f.Add(createExportedQFunc(tableName))
	f.Line()
	for colFunc := range createExportedColumnFunctions(columns) {
		f.Add(colFunc)
		f.Line()
	}

	return Render(f)
}

// createColumnFunctions makes each columns Global / Exported Path function
//
// func UUID() path.BinaryPath {
//    return q.uuid
// }
//
// func Name() path.StringPath {
//    return q.name
// }
func createExportedColumnFunctions(cols []*column) <-chan *Statement {
	var response = make(chan *Statement)
	go func() {
		defer close(response)
		for _, col := range cols {
			f := Func().Id(col.MethodName()).Call().Add(Qual(col.PathTypeName()))
			f.Block(Return(Id(packageInstance).Dot(col.MemberName())))
			response <- f
		}
	}()
	return response
}

// createExportedQFunc makes the generic Q Get function
//
//
// func Q() QCharacterSets {
//    return instance
// }
func createExportedQFunc(tableName string) *Statement {
	structName := ToTableStruct(tableName)
	f := Func().Id("Q").Call().Add(Qual("", structName))
	f.Block(Return(Id(packageInstance)))
	return f
}
