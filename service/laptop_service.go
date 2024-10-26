package service

import (
	"context"
	"errors"
	"log"

	message "github.com/aniketh3014/simple-protobuf/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	laptopStore LaptopStore
}

func NewLaptopServer(laptopStore LaptopStore) *LaptopServer {
	return &LaptopServer{laptopStore}
}

func (s *LaptopServer) CreateLaptop(ctx context.Context, req *message.CreateLaptopRequest) (*message.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Panicf("recived create laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid uuid")
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop")
		}
		laptop.Id = id.String()
	}

	if ctx.Err() == context.Canceled {
		log.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Print("deadline exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	err := s.laptopStore.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop to the store: %v", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)

	res := &message.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}
