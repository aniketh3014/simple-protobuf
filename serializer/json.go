package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJson(msg proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		UseEnumNumbers:    false,
		EmitDefaultValues: true,
		Indent:            "	",
	}
	jsonBytes, err := marshaler.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
