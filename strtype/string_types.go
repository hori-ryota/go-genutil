package strtype

import "strings"

func ToStringTypes(paramType string) string {
	if strings.HasPrefix(paramType, "map") || strings.HasPrefix(paramType, "[]") {
		return "[]string"
	}
	return "string"
}
