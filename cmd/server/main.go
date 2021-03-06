package main

import (
	"fmt"
	"log"
	"net"

	"bitbucket.use.dom.carezen.net/grpc-example/handler"
	"bitbucket.use.dom.carezen.net/grpc-example/interceptors"
	"bitbucket.use.dom.carezen.net/grpc-example/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authMD := interceptors.AuthMD{}
	opts := make([]grpc.ServerOption, 0)
	opts = append(opts, grpc.ChainUnaryInterceptor(authMD.UnaryInterceptor()))

	grpcServer := grpc.NewServer(opts...)

	chatHandler := handler.Chat{}
	// registering specific handlers for this server
	pb.RegisterChatServiceServer(grpcServer, &chatHandler)
	log.Println("starting server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
