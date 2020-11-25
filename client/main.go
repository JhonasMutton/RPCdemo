package main

import (
	"context"
	"github.com/JhonasMutton/RPCdemo/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:8080"
)

func main() {
	//Creates Connection with server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//Generated method
	c := proto.NewNineTableServiceClient(conn)
	//Configure timeout
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	log.Printf("Client is up!")

	numbers := []int32{1,2,3,4,5,6,7,8,9,10}

	log.Printf("Sending numbers to multiply!")
	//Generated method too
	//Calling remote procedure
	r, err := c.Multiply(ctx, &proto.Data{Numbers: numbers})
	if err != nil {
		log.Fatalf("could not multiply: %v", err)
	}

	log.Print("Calculated Numbers: ")
	for _, v := range r.Numbers {
		log.Print(v)
	}
}
