package genutil

import "strings"

func GoStringTypes(paramType string) string {
	if strings.HasPrefix(paramType, "map") || strings.HasPrefix(paramType, "[]") {
		return "[]string"
	}
	return "string"
}
