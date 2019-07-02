package templer_test

import "github.com/hori-ryota/go-genutil/genutil/templer"

var _ templer.Type = templer.PrimitiveType{}
var _ templer.Type = templer.StructType{}
var _ templer.Type = templer.SliceType{}
var _ templer.Type = templer.MapType{}
var _ templer.Type = templer.ChannelType{}
