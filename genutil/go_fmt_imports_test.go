package genutil_test

import (
	"testing"

	"github.com/hori-ryota/go-genutil/genutil"
	"github.com/stretchr/testify/assert"
)

func TestGoFmtImports(t *testing.T) {
	for _, tt := range []struct {
		name string
		src  map[string]string
		want string
	}{
		{
			name: "empty",
			src:  map[string]string{},
			want: ``,
		},
		{
			name: "only standard pacakge",
			src: map[string]string{
				"fmt":   "fmt",
				"bytes": "bytes",
				"json":  "encoding/json",
			},
			want: `import (
"bytes"
"encoding/json"
"fmt"
)`,
		},
		{
			name: "with alias",
			src: map[string]string{
				"afmt":  "fmt",
				"bytes": "bytes",
				"ajson": "encoding/json",
			},
			want: `import (
"bytes"
ajson "encoding/json"
afmt "fmt"
)`,
		},
		{
			name: "with non-standard package",
			src: map[string]string{
				"fmt":        "fmt",
				"bytes":      "bytes",
				"json":       "encoding/json",
				"go-genutil": "github.com/hori-ryota/go-genutil",
			},
			want: `import (
"bytes"
"encoding/json"
"fmt"

"github.com/hori-ryota/go-genutil"
)`,
		},
		{
			name: "with aliased non-standard package",
			src: map[string]string{
				"fmt":     "fmt",
				"genutil": "github.com/hori-ryota/go-genutil",
			},
			want: `import (
"fmt"

genutil "github.com/hori-ryota/go-genutil"
)`,
		},
		{
			name: "only non-standard package",
			src: map[string]string{
				"go-genutil": "github.com/hori-ryota/go-genutil",
			},
			want: `import (
"github.com/hori-ryota/go-genutil"
)`,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, genutil.GoFmtImports(tt.src))
		})
	}
}
