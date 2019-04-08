package genutil_test

import (
	"testing"

	genutil "github.com/hori-ryota/go-genutil"
	"github.com/stretchr/testify/assert"
)

func TestGoStringTypesConverter(t *testing.T) {
	for _, tt := range []struct {
		paramType string
		want      string
	}{
		{
			paramType: "bool",
			want:      "strconv.FormatBool(param)",
		},
		{
			paramType: "map[string]string",
			want: `func(s map[string]string) []string {
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
				}(param)`,
		},
		{
			paramType: "bool",
			want:      `strconv.FormatBool(param)`,
		},
		{
			paramType: "int",
			want:      `strconv.FormatInt(int64(param), 10)`,
		},
		{
			paramType: "int64",
			want:      `strconv.FormatInt(param, 10)`,
		},
		{
			paramType: "uint",
			want:      `strconv.FormatUint(uint64(param), 10)`,
		},
		{
			paramType: "uint64",
			want:      `strconv.FormatUint(param, 10)`,
		},
		{
			paramType: "float32",
			want:      `fmt.Sprint(param)`,
		},
		{
			paramType: "[]string",
			want:      `param`,
		},
		{
			paramType: "[]int",
			want: `func(ss []int) []string {
					dst := make([]string, len(ss))
					for i := range ss {
						dst[i] = strconv.FormatInt(int64(ss[i]), 10)
					}
					return dst
				}("param")`,
		},
		{
			paramType: "[]int64",
			want: `func(ss []int64) []string {
					dst := make([]string, len(ss))
					for i := range ss {
						dst[i] = strconv.FormatInt(ss[i], 10)
					}
					return dst
				}("param")`,
		},
		{
			paramType: "[]uint",
			want: `func(ss []uint) []string {
					dst := make([]string, len(ss))
					for i := range ss {
						dst[i] = strconv.FormatUint(uint64(ss[i]), 10)
					}
					return dst
				}("param")`,
		},
		{
			paramType: "[]uint64",
			want: `func(ss []uint64) []string {
					dst := make([]string, len(ss))
					for i := range ss {
						dst[i] = strconv.FormatUint(ss[i], 10)
					}
					return dst
				}("param")`,
		},
		{
			paramType: "string",
			want:      `param`,
		},
		{
			paramType: "unknown",
			want:      `fmt.Sprint(param)`,
		},
	} {
		tt := tt
		t.Run(tt.paramType, func(t *testing.T) {
			assert.Equal(t, tt.want, genutil.GoStringTypesConverter(tt.paramType, "param"))
		})
	}
}
