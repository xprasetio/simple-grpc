package main

import (
	"context"
	"grpc/pb/user"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type userService struct {
	user.UnimplementedUserServiceServer
}

func (us *userService) CreateUser(ctx context.Context, userReq *user.User) (*user.CreateResponse, error) {
	log.Println("Receive Create User Request")
	return &user.CreateResponse{
		Message: "User Created",
	}, nil
}

func main() {
	list, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("There is an error on listen to port 8082", err)
	}
	serv := grpc.NewServer()

	user.RegisterUserServiceServer(serv, &userService{})
	log.Println("gRPC server is running on port 8082")

	reflection.Register(serv) // mentok di development jangan sampai production

	if err := serv.Serve(list); err != nil {
		log.Fatal("There is an error on serve to port 8082", err)
	}
}
