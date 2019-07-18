package strtype

import "fmt"

func ToConverter(paramType string, paramName string) string {
	switch paramType {
	case "map[string]string":
		return fmt.Sprintf(`func(s map[string]string) []string {
					keys := make([]string, 0, len(s))
					for k := range s {
						keys = append(keys, k)
					}
					sort.Strings(keys)
					dst := make([]string, len(s))
					for i, k := range keys {
						dst[i] = k + "=" + s[k]
					}
					return dst
				}(%s)`,
			paramName,
		)
	case "bool":
		return fmt.Sprintf(`strconv.FormatBool(%s)`, paramName)
	case "int", "int8", "int16", "int32":
		return fmt.Sprintf(`strconv.FormatInt(int64(%s), 10)`, paramName)
	case "int64":
		return fmt.Sprintf(`strconv.FormatInt(%s, 10)`, paramName)
	case "uint", "uint8", "uint16", "uint32":
		return fmt.Sprintf(`strconv.FormatUint(uint64(%s), 10)`, paramName)
	case "uint64":
		return fmt.Sprintf(`strconv.FormatUint(%s, 10)`, paramName)
	case "float32", "float64":
		return fmt.Sprintf(`fmt.Sprint(%s)`, paramName)
	case "[]string":
		return paramName
	case "[]int", "[]int8", "[]int16", "[]int32":
		return fmt.Sprintf(`func(ss %s) []string {
					dst := make([]string, len(ss))
					for i := range ss {
						dst[i] = strconv.FormatInt(int64(ss[i]), 10)
					}
					return dst
				}("%s")`,
			paramType,
			paramName,
		)
	case "[]int64":
		return fmt.Sprintf(`func(ss %s) []string {
					dst := make([]string, len(ss))
					for i := range ss {
						dst[i] = strconv.FormatInt(ss[i], 10)
					}
					return dst
				}("%s")`,
			paramType,
			paramName,
		)
	case "[]uint", "[]uint8", "[]uint16", "[]uint32":
		return fmt.Sprintf(`func(ss %s) []string {
					dst := make([]string, len(ss))
					for i := range ss {
						dst[i] = strconv.FormatUint(uint64(ss[i]), 10)
					}
					return dst
				}("%s")`,
			paramType,
			paramName,
		)
	case "[]uint64":
		return fmt.Sprintf(`func(ss %s) []string {
					dst := make([]string, len(ss))
					for i := range ss {
						dst[i] = strconv.FormatUint(ss[i], 10)
					}
					return dst
				}("%s")`,
			paramType,
			paramName,
		)
	case "string":
		return paramName
	default:
		return fmt.Sprintf(`fmt.Sprint(%s)`, paramName)
	}
}

func ImportsForConverter(paramType string) []string {
	switch paramType {
	case "map[string]string":
		return []string{"sort"}
	case "bool":
		return []string{"strconv"}
	case "int", "int8", "int16", "int32":
		return []string{"strconv"}
	case "int64":
		return []string{"strconv"}
	case "uint", "uint8", "uint16", "uint32":
		return []string{"strconv"}
	case "uint64":
		return []string{"strconv"}
	case "float32", "float64":
		return []string{"fmt"}
	case "[]string":
		return []string{"fmt"}
	case "[]int", "[]int8", "[]int16", "[]int32":
		return []string{"strconv"}
	case "[]int64":
		return []string{"strconv"}
	case "[]uint", "[]uint8", "[]uint16", "[]uint32":
		return []string{"strconv"}
	case "[]uint64":
		return []string{"strconv"}
	case "string":
		return []string{}
	default:
		return []string{"fmt"}
	}
}
