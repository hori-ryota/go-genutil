package genutil

import (
	"go/ast"
	"path"
	"strings"
)

func AstFileToImportMap(file *ast.File) map[string]string {
	m := make(map[string]string, len(file.Imports))
	for _, imp := range file.Imports {
		pkgPath := strings.Trim(imp.Path.Value, `"`)
		if imp.Name == nil {
			m[path.Base(pkgPath)] = pkgPath
			continue
		}
		m[imp.Name.Name] = pkgPath
	}
	return m
}
