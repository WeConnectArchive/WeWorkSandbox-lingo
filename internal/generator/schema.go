package generator

import (
	"fmt"

	//revive:disable-next-line
	. "github.com/dave/jennifer/jen"
)

func GenerateSchema(schema string) (string, error) {
	f := NewFile(ToPackageName(schema))
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
func createSchemaInstance(schema string) *Statement {
	v := Var().Id("instance").Op("=").Qual("", typeSchema).Values(Dict{
		Id(nameField): Lit(schema),
	})
	return v
}

// createGetSchema generates the method to retrieve the schema instance
//
//
// func GetSchema() *schema {
//    return &instance
// }
func createGetSchema() *Statement {
	f := Func().Id("GetSchema").Call().Op(ptr).Qual("", typeSchema)
	f.Block(Return(Op(addrOf).Id("instance")))
	return f
}

// createSchemaStruct creates the generic `schema` structure object
//
// type schema struct {
//    _name           string
// }
func createSchemaStruct() *Statement {
	t := Type().Id(typeSchema)
	t.StructFunc(func(g *Group) {
		g.Id(nameField).String()
	})
	return t
}

// createGetNameStructFunc creates the `GetName` function for the `schema` struct
//
// func (s schema) GetName() string {
//    return s._name
// }
func createGetNameStructFunc() *Statement {
	f := Func().Parens(Id("s").Qual("", typeSchema)).Id("GetName").Call().String()
	f.Block(Return(Id("s").Dot(nameField)))
	return f
}
