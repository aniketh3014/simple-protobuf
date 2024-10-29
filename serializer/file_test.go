package serializer_test

import (
	"testing"

	message "github.com/aniketh3014/simple-protobuf/pb"
	"github.com/aniketh3014/simple-protobuf/sample"
	"github.com/aniketh3014/simple-protobuf/serializer"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	assert.Nil(t, err)

	err = serializer.WriteProtobufToJsonFile(laptop1, jsonFile)
	assert.Nil(t, err)

	laptop2 := &message.Laptop{}

	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	assert.Nil(t, err)

	assert.True(t, proto.Equal(laptop1, laptop2))
}
