package main

import (
	"context"
	"log"
	"net"

	"github.com/go-redis/redis"
	pb "github.com/morimolymoly/user-db-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	redisAddress    = "user-db-grpc-redis:6379"
	port            = ":8100"
	success_message = "正常に登録されました"
	// failed_message  = "なんかだめだわ……"
)

var redisClient *redis.Client

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
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "",
		DB:       0,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to ping for redis: %v", err)
	}

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
