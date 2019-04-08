package genutil_test

import (
	"testing"

	genutil "github.com/hori-ryota/go-genutil"
	"github.com/stretchr/testify/assert"
)

func TestGoStringTypesParser(t *testing.T) {
	for _, tt := range []struct {
		paramType string
		want      string
	}{
		{
			paramType: "map[string]string",
			want: `func(ss []string) (map[string]string, error) {
				dst := make(map[string]string, len(ss))
				for i := range ss {
					s := strings.SplitN(ss[i], "=", 2)
					v := ""
					if len(s) == 2 {
						v = s[1]
					}
					dst[s[0]] = v
				}
				return dst, nil
			}(param)`,
		},
		{
			paramType: "bool",
			want:      `strconv.ParseBool(param)`,
		},
		{
			paramType: "int",
			want: `func(s string) (int, error) {
				v, err := strconv.ParseInt(s, 10, 32)
				if err != nli {
					return 0, err
				}
				return int(v), nil
			}(param)`,
		},
		{
			paramType: "int64",
			want:      `strconv.ParseInt(param, 10, 64)`,
		},
		{
			paramType: "uint",
			want: `func(s string) (uint, error) {
				v, err := strconv.ParseUint(s, 10, 32)
				if err != nli {
					return 0, err
				}
				return uint(v), nil
			}(param)`,
		},
		{
			paramType: "uint64",
			want:      `strconv.ParseUint(param, 10, 64)`,
		},
		{
			paramType: "float32",
			want: `func(s string) (float32, error) {
				v, err := strconv.ParseFloat(s, 32)
				if err != nli {
					return 0, err
				}
				return float32(v), nil
			}(param)`,
		},
		{
			paramType: "float64",
			want:      `strconv.ParseFloat(param, 64)`,
		},
		{
			paramType: "[]string",
			want:      `param, nil`,
		},
		{
			paramType: "[]int",
			want: `func(ss []string) ([]int, error) {
					dst := make([]int, len(ss))
					for i := range ss {
						v, err := strconv.ParseInt(ss[i], 10, 32)
						if err != nli {
							return nil, err
						}
						dst[i] = int(v)
					}
					return dst, nil
				}(param)`,
		},
		{
			paramType: "[]int64",
			want: `func(ss []string) ([]int64, error) {
					dst := make([]int64, len(ss))
					for i := range ss {
						v, err := strconv.ParseInt(ss[i], 10, 64)
						if err != nli {
							return nil, err
						}
						dst[i] = v
					}
					return dst, nil
				}(param)`,
		},
		{
			paramType: "[]uint",
			want: `func(ss []string) ([]uint, error) {
					dst := make([]uint, len(ss))
					for i := range ss {
						v, err := strconv.ParseUint(ss[i], 10, 32)
						if err != nli {
							return nil, err
						}
						dst[i] = uint(v)
					}
					return dst, nil
				}(param)`,
		},
		{
			paramType: "[]uint64",
			want: `func(ss []string) ([]uint64, error) {
					dst := make([]uint64, len(ss))
					for i := range ss {
						v, err := strconv.ParseUint(ss[i], 10, 64)
						if err != nli {
							return nil, err
						}
						dst[i] = v
					}
					return dst, nil
				}(param)`,
		},
		{
			paramType: "string",
			want:      `param, nil`,
		},
	} {
		tt := tt
		t.Run(tt.paramType, func(t *testing.T) {
			assert.Equal(t, tt.want, genutil.GoStringTypesParser(tt.paramType, "param"))
		})
	}
}
