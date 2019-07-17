package genutil

import (
	"bytes"
	"go/ast"
	"go/printer"
	"path"
	"strings"
	"unicode"
)

type TypePrinter struct {
	FullString string
}

func (w AstPkgWalker) ToTypePrinter(typeExpr ast.Expr) (TypePrinter, error) {
	b := new(bytes.Buffer)
	err := printer.Fprint(b, w.FileSet, typeExpr)
	if err != nil {
		return TypePrinter{}, err
	}

	fullStr, err := resolveTypeFullString(AstFileToImportMap(w.ToFile(typeExpr)), w.PkgPath, b.String())
	if err != nil {
		return TypePrinter{}, err
	}
	return TypePrinter{
		FullString: fullStr,
	}, nil
}

func ToTypePrinter(importMap map[string]string, currentPkg string, typ string) (TypePrinter, error) {
	fullStr, err := resolveTypeFullString(importMap, currentPkg, typ)
	if err != nil {
		return TypePrinter{}, err
	}
	return TypePrinter{
		FullString: fullStr,
	}, nil
}

func resolveTypeFullString(importMap map[string]string, currentPkg string, typeName string) (string, error) {
	if typeName == "" {
		return "", nil
	}

	ind := strings.IndexAny(typeName, "-<> []")
	switch ind {
	case 0:
		ind = 1
	case -1:
		ind = len(typeName)
	}

	word := typeName[:ind]

	// convert pkg to full path
	switch {
	case !unicode.IsLetter(rune(word[0])):
		// noop
	case word == "chan", word == "map":
		// noop
	case isPrimitiveType(word):
		// noop
	case !strings.ContainsRune(word, '.'):
		word = currentPkg + "." + word
	default:
		ss := strings.SplitN(word, ".", 2)
		if pkgPath, ok := importMap[ss[0]]; ok {
			word = pkgPath + "." + ss[1]
		}
	}

	t, err := resolveTypeFullString(importMap, currentPkg, typeName[ind:])
	if err != nil {
		return "", err
	}
	return word + t, nil
}

func isPrimitiveType(t string) bool {
	switch t {
	case "string", "rune", "byte",
		"bool",
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64":
		return true
	}
	return false
}

func (p TypePrinter) Print(currentPkg string) string {
	return p.print(p.FullString, currentPkg)
}

func (p TypePrinter) print(fullPath string, currentPkg string) string {
	if fullPath == "" {
		return ""
	}

	ind := strings.IndexAny(fullPath, "-<> []")
	switch ind {
	case 0:
		ind = 1
	case -1:
		ind = len(fullPath)
	}

	word := fullPath[:ind]

	if path.Dir(word) == currentPkg {
		word = strings.TrimPrefix(path.Ext(word), ".")
	}

	return path.Base(word) + p.print(fullPath[ind:], currentPkg)
}

func (p TypePrinter) ImportPkgs(currentPkg string) []string {
	pkgs := make([]string, 0, 2)

	ss := strings.FieldsFunc(p.FullString, func(c rune) bool {
		return !unicode.IsLetter(c) && c != '.' && c != '/'
	})
	for _, s := range ss {
		if !strings.ContainsRune(s, '.') {
			continue
		}
		pkgs = append(pkgs, strings.SplitN(s, ".", 2)[0])
	}
	return pkgs
}

func (p TypePrinter) ImportPkgMap(currentPkg string) map[string]string {
	pkgs := p.ImportPkgs(currentPkg)
	m := make(map[string]string, len(pkgs))
	for _, pkg := range pkgs {
		m[path.Base(pkg)] = pkg
	}
	return m
}
