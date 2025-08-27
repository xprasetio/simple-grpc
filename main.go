package main

import (
	"context"
	"errors"
	"grpc/pb/chat"
	"grpc/pb/user"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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

type chatService struct {
	chat.UnimplementedChatServiceServer
}

func (cs *chatService) SendMessage(stream grpc.ClientStreamingServer[chat.ChatMessage, chat.ChatMessage]) error {
	for {
		res, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return status.Errorf(codes.Unknown, "failed to receive message: %v", err)
		}
		log.Printf("Receive message from user %d: %s", res.UserId, res.Content)
	}
	return stream.SendAndClose(&chat.ChatMessage{
		Content: "Thanks for your message!",
	})
}

func (cs *chatService) ReceiveMessage(req *chat.ReceiveMessageRequest, stream grpc.ServerStreamingServer[chat.ChatMessage]) error {
	log.Printf("Receive message request for user %d", req.UserId)

	err := stream.Send(&chat.ChatMessage{
		UserId:  req.UserId,
		Content: "Hello from server",
	})
	if err != nil {
		return status.Errorf(codes.Unknown, "failed to send message: %v", err)
	}
	err = stream.Send(&chat.ChatMessage{
		UserId:  req.UserId,
		Content: "How are you?",
	})
	if err != nil {
		return status.Errorf(codes.Unknown, "failed to send message: %v", err)
	}
	err = stream.Send(&chat.ChatMessage{
		UserId:  req.UserId,
		Content: "Goodbye!",
	})
	if err != nil {
		return status.Errorf(codes.Unknown, "failed to send message: %v", err)
	}
	return nil
}

// func (UnimplementedChatServiceServer) Chat(grpc.BidiStreamingServer[ChatMessage, ChatMessage]) error {
// 	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
// }

func main() {
	list, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("There is an error on listen to port 8082", err)
	}
	serv := grpc.NewServer()

	user.RegisterUserServiceServer(serv, &userService{})
	chat.RegisterChatServiceServer(serv, &chatService{})

	log.Println("gRPC server is running on port 8082")

	reflection.Register(serv) // mentok di development jangan sampai production

	if err := serv.Serve(list); err != nil {
		log.Fatal("There is an error on serve to port 8082", err)
	}
}
