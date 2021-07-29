package main

import (
	"log"

	"bitbucket.use.dom.carezen.net/grpc-example/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &pb.Message{
		Id:           1,
		Body:         "Hello From Client!",
		PhoneNumbers: []string{"111", "222"},
		PersonInfo: &pb.Person{
			Name:     "Roman",
			LastName: "Kosyi",
		},
	})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Body)
}
