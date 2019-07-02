package templer

type Type interface {
	TypeKind() TypeKind
	IsPointer() bool
}

type TypeKind string

const (
	TypeKindPrimitive TypeKind = "primitive"
	TypeKindStruct    TypeKind = "struct"
	TypeKindSlice     TypeKind = "slice"
	TypeKindMap       TypeKind = "map"
	TypeKindChannel   TypeKind = "channel"
)

//genconstructor
type PrimitiveType struct {
	typeKind  TypeKind `required:"TypeKindPrimitive" getter:""`
	typeName  string   `required:"" getter:""`
	isPointer bool     `getter:"" setter:""`
}

//genconstructor
type StructType struct {
	typeKind  TypeKind `required:"TypeKindStruct" getter:""`
	typeName  string   `required:"" getter:""`
	pkg       Package  `required:"" getter:""`
	fields    []Field  `required:"" getter:""`
	methods   []Method `required:"" getter:""`
	isPointer bool     `getter:"" setter:""`
}

//genconstructor
type SliceType struct {
	typeKind  TypeKind `required:"TypeKindSlice" getter:""`
	itemType  Type     `required:"" getter:""`
	isPointer bool     `getter:"" setter:""`
}

//genconstructor
type MapType struct {
	typeKind  TypeKind `required:"TypeKindMap" getter:""`
	keyType   Type     `required:"" getter:""`
	valueType Type     `required:"" getter:""`
	isPointer bool     `getter:"" setter:""`
}

//genconstructor
type ChannelType struct {
	typeKind  TypeKind `required:"TypeKindChannel" getter:""`
	itemType  Type     `required:"" getter:""`
	isPointer bool     `getter:"" setter:""`
}
