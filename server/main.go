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
	port            = ":8100"
	success_message = "正常に登録されました"
	// failed_message  = "なんかだめだわ……"
)

// UserServiceServer ... implement UserServiceServer
type UserServiceServer struct {
}

// GetAllUser ... implement UserServiceServer
func (s *UserServiceServer) GetAllUser(ctx context.Context, req *pb.GetAllUserRequest) (*pb.Users, error) {
	return &pb.Users{U: []*pb.User{&pb.User{UserID: 0, Username: "もり"}}}, nil
}

// AddUser ... implement UserServiceServer
func (s *UserServiceServer) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.Result, error) {
	log.Printf("ADDUSER: Username: %s, UserID: %v\n", req.Username, req.UserID)
	return &pb.Result{Success: true, Message: success_message}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &UserServiceServer{})
	log.Println("RegisterUserServiceServer success!")

	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
