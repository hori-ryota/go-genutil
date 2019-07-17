package genutil

import (
	"errors"
	"go/build"
	"path/filepath"
	"strings"
)

func LocalPathToPackagePath(s string) (string, error) {
	s, err := filepath.Abs(s)
	if err != nil {
		return "", err
	}

	s = filepath.ToSlash(s)

	for _, srcDir := range build.Default.SrcDirs() {
		srcDir = filepath.ToSlash(srcDir)
		prefix := srcDir + "/"
		if strings.HasPrefix(s, prefix) {
			return strings.TrimPrefix(s, prefix), nil
		}
	}
	return "", errors.New("failed to resolve package path")
}
