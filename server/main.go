package main

import (
	"context"
	"log"
	"net"

	pb "github.com/morimolymoly/user-db-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8100"
)

// UserServiceServer ... implement UserServiceServer
type UserServiceServer struct {
}

// GetAllUser ... implement UserServiceServer
func (s *UserServiceServer) GetAllUser(ctx context.Context, req *pb.GetAllUserRequest) (*pb.Users, error) {
	return &pb.Users{U: []*pb.User{&pb.User{UserID: 0, Username: "もり"}}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &UserServiceServer{})

	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
