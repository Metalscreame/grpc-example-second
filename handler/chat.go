package handler

import (
	"context"
	"fmt"
	"log"

	"bitbucket.use.dom.carezen.net/grpc-example/pb"
)

type Chat struct {
	pb.UnimplementedChatServiceServer
}

func (s *Chat) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message from client: %v", in)

	user, _ := ctx.Value("user").(string)
	return &pb.Message{Body: fmt.Sprintf("Hello From %s the Server:ChatHandler!", user)}, nil
}