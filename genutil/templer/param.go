package templer

import (
	"fmt"
	"path"
)

//genconstructor
type Package struct {
	name string `required:"" getter:""`
	path string `required:"" getter:""`
}

func (p Package) ImportLine() string {
	switch {
	case p.Name() == "":
		return fmt.Sprintf(`"%s"`, p.Path())
	case p.Name() == path.Base(p.Path()):
		return fmt.Sprintf(`"%s"`, p.Path())
	default:
		return fmt.Sprintf(`%s "%s"`, p.Name(), p.Path())
	}
}

//genconstructor
type Param struct {
	name string `required:"" getter:""`
	typ  Type   `required:"" getter:"Type"`
}

//genconstructor
type Field struct {
	Param   `required:""`
	isEmbed bool `getter:"" setter:""`
}

//genconstructor
type Arg struct {
	Param      `required:""`
	isVariadic bool `getter:"" setter:""`
}

//genconstructor
type Method struct {
	args      []Arg   `required:"" getter:""`
	responses []Param `required:"" getter:""`
}
