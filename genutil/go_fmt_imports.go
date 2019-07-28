package genutil

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func GoFmtImports(pkgs map[string]string) string {
	if len(pkgs) == 0 {
		return ""
	}
	pathToNameMap := make(map[string]string, len(pkgs))
	for name, pat := range pkgs {
		pathToNameMap[pat] = name
	}

	groups := make([][]string, 2)

	for _, pkg := range pkgs {
		if len(strings.Split(pkg, "/")) < 3 && !strings.Contains(pkg, ".") {
			groups[0] = append(groups[0], pkg)
			continue
		}
		groups[1] = append(groups[1], pkg)
	}

	b := new(bytes.Buffer)
	for _, group := range groups {
		group := group
		sort.Slice(group, func(i, j int) bool {
			return group[i] < group[j]
		})
		for _, pkg := range group {
			_, err := b.WriteString(pathToNameMap[pkg])
			if err != nil {
				panic(err)
			}
			_, err = b.WriteRune(' ')
			if err != nil {
				panic(err)
			}
			_, err = b.WriteString(strconv.Quote(pkg))
			if err != nil {
				panic(err)
			}
			_, err = b.WriteRune('\n')
			if err != nil {
				panic(err)
			}
		}
		_, err := b.WriteRune('\n')
		if err != nil {
			panic(err)
		}
	}

	return fmt.Sprintf(`import (
%s
)`,
		strings.TrimSpace(b.String()),
	)
}
