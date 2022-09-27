package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	o := protojson.MarshalOptions{}
	o.Multiline = true
	e, err := o.Marshal(message)
	if err != nil {
		return "", nil
	}

	return string(e), nil
}
