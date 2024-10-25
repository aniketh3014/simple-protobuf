package serializer

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

func WriteProtobufToBinaryFile(msg proto.Message, filename string) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("cannot marshl proto message to binary: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}

	return nil
}

func ReadProtobufFromBinaryFile(filename string, msg proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}

	err = proto.Unmarshal(data, msg)
	if err != nil {
		return fmt.Errorf("cannot unarshal binary to  proto: %w", err)
	}
	return nil
}

func WriteProtobufToJsonFile(msg proto.Message, filename string) error {
	data, err := ProtobufToJson(msg)
	if err != nil {
		return fmt.Errorf("cannot marshal proto to json: %w", err)
	}

	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write json file: %w", err)
	}

	return nil
}
