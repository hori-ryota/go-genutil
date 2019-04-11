package genutil_test

import (
	"testing"

	"github.com/hori-ryota/go-genutil/genutil"
	"github.com/stretchr/testify/assert"
)

func TestLocalPathToPackagePath(t *testing.T) {
	for _, tt := range []struct {
		s string
		t string
	}{
		{
			s: ".",
			t: "github.com/hori-ryota/go-genutil/genutil",
		},
		{
			s: "..",
			t: "github.com/hori-ryota/go-genutil",
		},
		{
			s: "../genutil",
			t: "github.com/hori-ryota/go-genutil/genutil",
		},
	} {
		tt := tt
		t.Run(tt.s, func(t *testing.T) {
			got, err := genutil.LocalPathToPackagePath(tt.s)
			assert.NoError(t, err)
			assert.Equal(t, tt.t, got)
		})
	}
}
