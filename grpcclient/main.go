package main

import (
	"context"
	"grpc/pb/chat"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect: ", err)
	}
	// userClient := user.NewUserServiceClient(cliendConn)
	// response, err := userClient.CreateUser(context.Background(), &user.User{
	// 	Id:      1,
	// 	Age:     18,
	// 	Balance: 14000,
	// 	Address: &user.Address{
	// 		Province: "shanghai",
	// 		City:     "shanghai",
	// 		Country:  "china",
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal("failed to create user: ", err)
	// }
	// log.Println("create user response: ", response.Message)

	//======client streaming======
	// chatClient := chat.NewChatServiceClient(clientConn)
	// stream, err := chatClient.SendMessage(context.Background())
	// if err != nil {
	// 	log.Fatal("failed to send message: ", err)
	// }
	// err =stream.Send(&chat.ChatMessage{
	// 	UserId:  1,
	// 	Content: "hello from client",
	// })
	// if err != nil {
	// 	log.Fatal("failed to send via stream: ", err)
	// }
	// err =stream.Send(&chat.ChatMessage{
	// 	UserId:  2,
	// 	Content: "hello from again",
	// })
	// if err != nil {
	// 	log.Fatal("failed to send via stream: ", err)
	// }

	// res, err := stream.CloseAndRecv()
	// if err != nil {
	// 	log.Fatal("failed to receive message: ", err)

	// }
	// log.Println("connection closed: ", res.Content)
	//======end of client streaming======

	chatClient := chat.NewChatServiceClient(clientConn)
	stream, err := chatClient.ReceiveMessage(context.Background(), &chat.ReceiveMessageRequest{
		UserId: 1,
	})
	if err != nil {
		log.Fatal("failed to receive message: ", err)
	}
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatal("failed to receive message from stream: ", err)
		}
		log.Printf("message from user %d: %s", msg.UserId, msg.Content)
	}
}
