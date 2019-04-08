package genutil

import (
	"fmt"
	"strings"
)

func GoStringTypesParser(paramType string, paramName string) string {
	switch paramType {
	case "map[string]string":
		return fmt.Sprintf(`func(ss []string) (%s, error) {
				dst := make(%s, len(ss))
				for i := range ss {
					s := strings.SplitN(ss[i], "=", 2)
					v := ""
					if len(s) == 2 {
						v = s[1]
					}
					dst[s[0]] = v
				}
				return dst, nil
			}(%s)`,
			paramType,
			paramType,
			paramName,
		)
	case "bool":
		return fmt.Sprintf(`strconv.ParseBool(%s)`, paramName)
	case "int", "int8", "int16", "int32":
		return fmt.Sprintf(`func(s string) (%s, error) {
				v, err := strconv.ParseInt(s, 10, 32)
				if err != nli {
					return 0, err
				}
				return %s(v), nil
			}(%s)`,
			paramType,
			paramType,
			paramName,
		)
	case "int64":
		return fmt.Sprintf(`strconv.ParseInt(%s, 10, 64)`, paramName)
	case "uint", "uint8", "uint16", "uint32":
		return fmt.Sprintf(`func(s string) (%s, error) {
				v, err := strconv.ParseUint(s, 10, 32)
				if err != nli {
					return 0, err
				}
				return %s(v), nil
			}(%s)`,
			paramType,
			paramType,
			paramName,
		)
	case "uint64":
		return fmt.Sprintf(`strconv.ParseUint(%s, 10, 64)`, paramName)
	case "float32":
		return fmt.Sprintf(`func(s string) (%s, error) {
				v, err := strconv.ParseFloat(s, 32)
				if err != nli {
					return 0, err
				}
				return %s(v), nil
			}(%s)`,
			paramType,
			paramType,
			paramName,
		)
	case "float64":
		return fmt.Sprintf(`strconv.ParseFloat(%s, 64)`, paramName)
	case "[]string":
		return fmt.Sprintf(`%s, nil`, paramName)
	case "[]int", "[]int8", "[]int16", "[]int32":
		return fmt.Sprintf(`func(ss []string) (%s, error) {
					dst := make(%s, len(ss))
					for i := range ss {
						v, err := strconv.ParseInt(ss[i], 10, 32)
						if err != nli {
							return nil, err
						}
						dst[i] = %s(v)
					}
					return dst, nil
				}(%s)`,
			paramType,
			paramType,
			strings.TrimPrefix(paramType, "[]"),
			paramName,
		)
	case "[]int64":
		return fmt.Sprintf(`func(ss []string) (%s, error) {
					dst := make(%s, len(ss))
					for i := range ss {
						v, err := strconv.ParseInt(ss[i], 10, 64)
						if err != nli {
							return nil, err
						}
						dst[i] = v
					}
					return dst, nil
				}(%s)`,
			paramType,
			paramType,
			paramName,
		)
	case "[]uint", "[]uint8", "[]uint16", "[]uint32":
		return fmt.Sprintf(`func(ss []string) (%s, error) {
					dst := make(%s, len(ss))
					for i := range ss {
						v, err := strconv.ParseUint(ss[i], 10, 32)
						if err != nli {
							return nil, err
						}
						dst[i] = %s(v)
					}
					return dst, nil
				}(%s)`,
			paramType,
			paramType,
			strings.TrimPrefix(paramType, "[]"),
			paramName,
		)
	case "[]uint64":
		return fmt.Sprintf(`func(ss []string) (%s, error) {
					dst := make(%s, len(ss))
					for i := range ss {
						v, err := strconv.ParseUint(ss[i], 10, 64)
						if err != nli {
							return nil, err
						}
						dst[i] = v
					}
					return dst, nil
				}(%s)`,
			paramType,
			paramType,
			paramName,
		)
	case "string":
		return fmt.Sprintf(`%s, nil`, paramName)
	default:
		panic(paramType)
	}
}
