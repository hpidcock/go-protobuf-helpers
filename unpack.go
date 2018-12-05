package protobufhelpers

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

// Unpack returns a unpacked or an error (will panic if eout is nil on error)
func Unpack(a *any.Any, eout *error) proto.Message {
	var da ptypes.DynamicAny
	err := ptypes.UnmarshalAny(a, &da)
	if err != nil {
		if eout != nil {
			*eout = err
		} else {
			panic(err)
		}
		return nil
	}
	return da.Message
}
