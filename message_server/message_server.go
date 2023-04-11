package messageserver

import (
	"fmt"
	"log"
	"net"

	message "github.com/2751997nam/go-helpers/message"

	"google.golang.org/grpc"
)

const (
	gRPCPort = "50001"
)

func GRPCListen(messageServer message.MessageServiceServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen to gRPC: %v", err)
	}
	s := grpc.NewServer()

	message.RegisterMessageServiceServer(s, messageServer)

	log.Printf("gRPC server started on port %s", gRPCPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen to gRPC: %v", err)
	}
}
