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

	for _, srcDir := range build.Default.SrcDirs() {
		prefix := srcDir + "/"
		if strings.HasPrefix(s, prefix) {
			return filepath.ToSlash(strings.TrimPrefix(s, prefix)), nil
		}
	}
	return "", errors.New("failed to resolve package path")
}
