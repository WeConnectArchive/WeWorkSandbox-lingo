package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

func GeneratePackageMembers(info TableInfo, columns []*column) (string, error) {
	tableName := info.Name()

	f := jen.NewFile(ToPackageName(tableName))
	f.HeaderComment(fmt.Sprintf(fmtTableHeaderComment, info.Schema(), tableName))
	f.HeaderComment(createBuildTag(tagNoLingo))
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
// func UUID() path.Binary {
//    return q.uuid
// }
//
// func Name() path.String {
//    return q.name
// }
func createExportedColumnFunctions(cols []*column) <-chan *jen.Statement {
	var response = make(chan *jen.Statement)
	go func() {
		defer close(response)
		for _, col := range cols {
			f := jen.Func().Id(col.MethodName()).Call().Add(jen.Qual(col.PathTypeName()))
			f.Block(jen.Return(jen.Id(packageInstance).Dot(col.MemberName())))
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
func createExportedQFunc(tableName string) *jen.Statement {
	structName := ToTableStruct(tableName)
	f := jen.Func().Id("Q").Call().Add(jen.Qual("", structName))
	f.Block(jen.Return(jen.Id(packageInstance)))
	return f
}
