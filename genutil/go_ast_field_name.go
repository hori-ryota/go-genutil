package genutil

import (
	"fmt"
	"go/ast"
	"reflect"
)

func ParseFieldName(field *ast.Field) string {
	if field.Names == nil {
		// embedding
		switch t := field.Type.(type) {
		case *ast.Ident:
			return t.Name
		case *ast.SelectorExpr:
			return t.Sel.Name
		default:
			panic(fmt.Errorf(
				"failed to parse FieldType: unknown %s",
				reflect.TypeOf(t),
			))
		}
	}

	return field.Names[0].Name
}
