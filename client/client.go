package main

import (
	"context"
	"io"
	"log"
	decompositionpb "prime-number-decomposition/proto"

	"google.golang.org/grpc"
)

func main() {
	log.Println("[INFO] Client started ...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[ERROR] Failed to Dial: %v", err)
	}
	defer cc.Close()

	c := decompositionpb.NewPrimeNumberServiceClient(cc)

	doServerStreamingDecomposition(c)
}

func doServerStreamingDecomposition(c decompositionpb.PrimeNumberServiceClient) {
	log.Println("[INFO] doServerStreamingDecomposition invoked ...")

	req := &decompositionpb.PrimeNumberRequest{
		PrimeNumber: &decompositionpb.PrimeNumber{
			PrimeNumber: 12,
		},
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("[ERROR] Failed to stream data: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Fatalln("[ERROR] End of stream reached")
			break
		}

		if err != nil {
			log.Fatalf("[ERROR] Failed to read stream: %v", err)
		}

		log.Printf("[INFO] Response from Decomposition: %v", msg.GetPrimeNumber())
	}

}
