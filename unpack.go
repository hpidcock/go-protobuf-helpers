package protobufhelpers

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

// Unpack returns a unpacked or an error
func Unpack(a *any.Any) (proto.Message, error) {
	var da ptypes.DynamicAny
	err := ptypes.UnmarshalAny(a, &da)
	if err != nil {
		return nil, err
	}
	return da.Message, nil
}
