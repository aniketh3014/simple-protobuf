package service_test

import (
	"context"
	"fmt"
	"net"
	"testing"

	message "github.com/aniketh3014/simple-protobuf/pb"
	"github.com/aniketh3014/simple-protobuf/sample"
	"github.com/aniketh3014/simple-protobuf/serializer"
	"github.com/aniketh3014/simple-protobuf/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopStore := service.NewInmemoryLaptopStore()
	_, addr := startTestLaptopServer(t, laptopStore)
	laptopClient := newtestLaptopClient(t, addr)

	laptop := sample.NewLaptop()

	expectedID := laptop.Id
	req := &message.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	other, err := laptopStore.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	requireSameLaptop(t, laptop, other)

}

func requireSameLaptop(t *testing.T, laptop1 *message.Laptop, laptop2 *message.Laptop) {
	json1, err := serializer.ProtobufToJson(laptop1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJson(laptop2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}

func startTestLaptopServer(t *testing.T, store service.LaptopStore) (*service.LaptopServer, string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 0))
	require.NoError(t, err)

	laptopServer := service.NewLaptopServer(store)

	grpcServer := grpc.NewServer()
	message.RegisterLaptopServiceServer(grpcServer, laptopServer)
	go grpcServer.Serve(lis)

	return laptopServer, lis.Addr().String()
}

func newtestLaptopClient(t *testing.T, serverAddr string) message.LaptopServiceClient {
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	return message.NewLaptopServiceClient(conn)
}
