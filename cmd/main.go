package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/irootpro/chat-auth/pkg/servers/grpc/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 3010

type server struct {
	desc.UnimplementedUserV1Server
}

func (s server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Printf("email: %s, name: %s", req.Email, req.Name)
	return &desc.CreateResponse{}, nil
}

func main() {

	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterUserV1Server(s, &server{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc server %v", err)
	}
}
