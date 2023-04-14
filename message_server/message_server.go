package messageserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	message "github.com/2751997nam/go-helpers/message"

	"google.golang.org/grpc"
)

const (
	gRPCPort = "50001"
)

type MessageServer struct {
	message.UnimplementedMessageServiceServer
	Router MessageRouter
}

func (m *MessageServer) HandleMessage(ctx context.Context, req *message.MessageRequest) (*message.MessageResponse, error) {
	input := req.GetMessageEntry()
	data := map[string]any{}
	err := json.Unmarshal([]byte(input.Data), &data)
	if err != nil {
		log.Println(err)
		res := &message.MessageResponse{
			Status:  "fail",
			Message: "parse data error",
		}

		return res, err
	}

	if handle, ok := m.Router.Handles[fmt.Sprintf("%s_%s", input.Type, input.Method)]; ok {
		res, err := handle.Handle(data)
		return &res, err
	}

	res := &message.MessageResponse{
		Status:  "fail",
		Message: "404 not found",
	}

	return res, fmt.Errorf("404 not found")
}

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
