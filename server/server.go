package main

import (
	"log"
	"net"
	decompositionpb "prime-number-decomposition/proto"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeNumberDecomposition(req *decompositionpb.PrimeNumberRequest, stream decompositionpb.PrimeNumberService_PrimeNumberDecompositionServer) error {
	log.Println("[INFO] PrimeNumberDecomposition function invoked")
	number := req.PrimeNumber.GetPrimeNumber()
	var divisor = int32(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&decompositionpb.PrimeNumberResponse{
				PrimeNumber: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			log.Printf("[INFO] Divisor has increased to: %v", divisor)
		}
	}
	return nil
}

func main() {

	log.Println("[INFO] Server started ...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("[ERROR] Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	decompositionpb.RegisterPrimeNumberServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("[ERROR] Failed to serve: %v", err)
	}
}
