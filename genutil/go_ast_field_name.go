package genutil

import "go/ast"

func ParseFieldName(field *ast.Field) string {
	if field.Names == nil {
		// embedding
		return field.Type.(*ast.Ident).Name
	}

	return field.Names[0].Name
}
