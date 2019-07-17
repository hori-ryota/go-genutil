package genutil

import (
	"go/ast"
)

func AllDeclsFromAstPkg(pkg *ast.Package) []ast.Decl {
	decls := make([]ast.Decl, 0, 100)
	for _, file := range ToSortedFileListFromFileMapOfAst(pkg.Files) {
		decls = append(decls, file.Decls...)
	}
	return decls
}
