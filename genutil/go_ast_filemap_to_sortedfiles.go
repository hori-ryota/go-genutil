package genutil

import (
	"go/ast"
	"sort"
)

func ToSortedFileListFromFileMapOfAst(s map[string]*ast.File) []*ast.File {
	names := make([]string, 0, len(s))
	for name := range s {
		names = append(names, name)
	}
	sort.Strings(names)

	l := make([]*ast.File, len(names))
	for i, name := range names {
		l[i] = s[name]
	}
	return l
}
