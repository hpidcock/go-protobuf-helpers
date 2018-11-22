package protobufhelpers

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

// Pack returns an any for the message. If m is not marshalable, panics.
func Pack(m proto.Message) *any.Any {
	a, err := ptypes.MarshalAny(m)
	if err != nil {
		panic(err)
	}
	return a
}
