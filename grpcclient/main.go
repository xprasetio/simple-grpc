package main

import (
	"context"
	"grpc/pb/user"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cliendConn, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect: ", err)
	}
	userClient := user.NewUserServiceClient(cliendConn)
	response, err := userClient.CreateUser(context.Background(), &user.User{
		Id:      1,
		Age:     18,
		Balance: 14000,
		Address: &user.Address{
			Province: "shanghai",
			City:     "shanghai",
			Country:  "china",
		},
	})
	if err != nil {
		log.Fatal("failed to create user: ", err)
	}
	log.Println("create user response: ", response.Message)
}
