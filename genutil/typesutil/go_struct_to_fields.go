package typesutil

import "go/types"

func StructToFields(s *types.Struct) []*types.Var {
	fields := make([]*types.Var, s.NumFields())
	for i := range fields {
		fields[i] = s.Field(i)
	}
	return fields
}

func TypeToFields(s types.Type) []*types.Var {
	t, ok := s.Underlying().(*types.Struct)
	if !ok {
		return nil
	}
	return StructToFields(t)
}

func ObjectToFields(s types.Object) []*types.Var {
	return TypeToFields(s.Type())
}

func StructToTags(s *types.Struct) []string {
	tags := make([]string, s.NumFields())
	for i := range tags {
		tags[i] = s.Tag(i)
	}
	return tags
}

func TypeToTags(s types.Type) []string {
	t, ok := s.Underlying().(*types.Struct)
	if !ok {
		return nil
	}
	return StructToTags(t)
}

func ObjectToTags(s types.Object) []string {
	return TypeToTags(s.Type())
}

func FuncToArgs(f *types.Func) []*types.Var {
	t, ok := f.Type().Underlying().(*types.Signature)
	if !ok {
		return nil
	}
	args := make([]*types.Var, t.Params().Len())
	for i := range args {
		args[i] = t.Params().At(i)
	}
	return args
}
