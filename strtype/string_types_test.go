package strtype_test

import (
	"testing"

	"github.com/hori-ryota/go-genutil/strtype"
	"github.com/stretchr/testify/assert"
)

func TestToStringTypes(t *testing.T) {
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
			assert.Equal(t, tt.want, strtype.ToStringTypes(tt.paramType))
		})
	}
}
