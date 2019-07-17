package genutil

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

type AstPkgWalker struct {
	FileSet *token.FileSet
	Pkg     *ast.Package
	Files   []*ast.File
	Decls   []ast.Decl

	PkgPath string
}

func DirToAstWalker(targetDir string, fileFilter func(finfo os.FileInfo) bool) (map[string]AstPkgWalker, error) {
	fset := token.NewFileSet()
	pkgMap, err := parser.ParseDir(
		fset,
		filepath.FromSlash(targetDir),
		fileFilter,
		parser.ParseComments,
	)
	if err != nil {
		return nil, err
	}

	m := make(map[string]AstPkgWalker, len(pkgMap))
	for k, v := range pkgMap {
		m[k], err = ParseAstPkg(fset, v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

func ParseAstPkg(fset *token.FileSet, pkg *ast.Package) (AstPkgWalker, error) {
	var aFilePath string
	for _, file := range pkg.Files {
		aFilePath = fset.File(file.Package).Name()
	}

	pkgPath, err := LocalPathToPackagePath(filepath.Dir(aFilePath))
	if err != nil {
		return AstPkgWalker{}, err
	}

	return AstPkgWalker{
		FileSet: fset,
		Pkg:     pkg,
		Files:   ToSortedFileListFromFileMapOfAst(pkg.Files),
		Decls:   AllDeclsFromAstPkg(pkg),

		PkgPath: pkgPath,
	}, nil
}

func (w AstPkgWalker) AllGenDecls() []*ast.GenDecl {
	decls := w.Decls
	l := make([]*ast.GenDecl, 0, len(decls))
	for _, decl := range decls {
		decl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		l = append(l, decl)
	}
	return l
}

func (w AstPkgWalker) AllTypeSpecs() []*ast.TypeSpec {
	decls := w.AllGenDecls()
	l := make([]*ast.TypeSpec, 0, len(decls))
	for _, decl := range decls {
		if decl.Tok != token.TYPE {
			continue
		}
		for _, spec := range decl.Specs {
			l = append(l, spec.(*ast.TypeSpec))
		}
	}
	return l
}

func (w AstPkgWalker) AllStructSpecs() []*ast.TypeSpec {
	specs := w.AllTypeSpecs()
	l := make([]*ast.TypeSpec, 0, len(specs))
	for _, spec := range specs {
		_, ok := spec.Type.(*ast.StructType)
		if !ok {
			continue
		}
		l = append(l, spec)
	}
	return l
}

func (w AstPkgWalker) TypeSpecToGenDecl(spec *ast.TypeSpec) *ast.GenDecl {
	for _, decl := range w.AllGenDecls() {
		for _, s := range decl.Specs {
			if s == spec {
				return decl
			}
		}
	}
	return nil
}

func (w AstPkgWalker) ToFile(node ast.Node) *ast.File {
	fileName := w.FileSet.File(node.Pos()).Name()
	for _, file := range w.Files {
		if fileName == w.FileSet.File(file.Package).Name() {
			return file
		}
	}
	return nil
}
