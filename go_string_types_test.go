package genutil_test

import (
	"testing"

	genutil "github.com/hori-ryota/go-genutil"
	"github.com/stretchr/testify/assert"
)

func TestGoStringTypes(t *testing.T) {
	for _, tt := range []struct {
		paramType string
		want      string
	}{
		{
			paramType: "string",
			want:      "string",
		},
		{
			paramType: "bool",
			want:      "string",
		},
		{
			paramType: "int",
			want:      "string",
		},
		{
			paramType: "map[string]string",
			want:      "[]string",
		},
		{
			paramType: "[]bool",
			want:      "[]string",
		},
		{
			paramType: "[]int",
			want:      "[]string",
		},
	} {
		tt := tt
		t.Run(tt.paramType, func(t *testing.T) {
			assert.Equal(t, tt.want, genutil.GoStringTypes(tt.paramType))
		})
	}
}
