package service_test

import (
	"context"
	"fmt"
	"net"
	"testing"

	message "github.com/aniketh3014/simple-protobuf/pb"
	"github.com/aniketh3014/simple-protobuf/sample"
	"github.com/aniketh3014/simple-protobuf/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopServer, addr := startTestLaptopServer(t)
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

}

func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 0))
	require.NoError(t, err)

	laptopServer := service.NewLaptopServer(service.NewInmemoryLaptopStore())

	grpcServer := grpc.NewServer()
	message.RegisterLaptopServiceServer(grpcServer, laptopServer)
	go grpcServer.Serve(lis)

	return laptopServer, lis.Addr().String()
}

func newtestLaptopClient(t *testing.T, serverAddr string) message.LaptopServiceClient {
	conn, err := grpc.NewClient(serverAddr)
	require.NoError(t, err)

	message.NewLaptopServiceClient(conn)
}
