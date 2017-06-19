package encoding

import (
	"encoding/json"
	"github.com/gogo/protobuf/proto"
	"io"
)

type Codec struct {
	Name       string
	NewEncoder func(io.Writer) Encoder
}

type Encoder interface {
	Encode(Marshaler) error
}

type Marshaler interface {
	proto.Marshaler
	json.Marshaler
}

var (
	ProtobufCodec = Codec{
		Name: "protobuf",
	}
	JsonCodec = Codec{
		Name: "json",
	}
)
