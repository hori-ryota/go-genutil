package templer

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

//genconstructor
type RootParam struct {
	types []Type `required:"" getter:""`
}

func (p RootParam) Structs(filterFunc ...func(t StructType) bool) []StructType {
	results := make([]StructType, 0, len(p.Types()))
	types := p.Types()
TYPE_LOOP:
	for i := range types {
		t, ok := p.types[i].(StructType)
		if !ok {
			continue
		}
		for _, f := range filterFunc {
			if !f(t) {
				continue TYPE_LOOP
			}
		}
		results = append(results, t)
	}
	return results
}

func ParseDir(targetDir string, fileFilter func(os.FileInfo) bool) (map[string]RootParam, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	list, err := fd.Readdir(-1)
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()
	pkgMap, err := parser.ParseDir(
		fset,
		filepath.FromSlash(targetDir),
		fileFilter,
		0,
	)
	ast.NewPackage(fset)

	if err != nil {
		return nil, err
	}
	return nil, nil
}
