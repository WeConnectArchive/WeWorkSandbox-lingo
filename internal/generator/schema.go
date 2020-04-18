package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

func GenerateSchema(schema string) (string, error) {
	f := jen.NewFile(ToPackageName(schema))
	f.HeaderComment(fmt.Sprintf(fmtSchemaHeaderComment, schema))
	f.Add(createSchemaInstance(schema))
	f.Line()
	f.Add(createGetSchema())
	f.Line()
	f.Add(createSchemaStruct())
	f.Line()
	f.Add(createGetNameStructFunc())

	return Render(f)
}

// createSchemaInstance create a private instance of the `schema` struct
//
// var instance = schema{
//    _name: "information_schema",
// }
func createSchemaInstance(schema string) *jen.Statement {
	v := jen.Var().Id("instance").Op("=").Qual("", typeSchema).Values(jen.Dict{
		jen.Id(nameField): jen.Lit(schema),
	})
	return v
}

// createGetSchema generates the method to retrieve the schema instance
//
//
// func GetSchema() *schema {
//    return &instance
// }
func createGetSchema() *jen.Statement {
	f := jen.Func().Id("GetSchema").Call().Op(ptr).Qual("", typeSchema)
	f.Block(jen.Return(jen.Op(addrOf).Id("instance")))
	return f
}

// createSchemaStruct creates the generic `schema` structure object
//
// type schema struct {
//    _name           string
// }
func createSchemaStruct() *jen.Statement {
	t := jen.Type().Id(typeSchema)
	t.StructFunc(func(g *jen.Group) {
		g.Id(nameField).String()
	})
	return t
}

// createGetNameStructFunc creates the `GetName` function for the `schema` struct
//
// func (s schema) GetName() string {
//    return s._name
// }
func createGetNameStructFunc() *jen.Statement {
	f := jen.Func().Parens(jen.Id("s").Qual("", typeSchema)).Id("GetName").Call().String()
	f.Block(jen.Return(jen.Id("s").Dot(nameField)))
	return f
}
