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
	defer func() (*message.MessageResponse, error) {
		if r := recover(); r != nil {
			return &message.MessageResponse{
				Status:     "fail",
				Message:    "recovered from the panic",
				StatusCode: 500,
			}, fmt.Errorf("recovered from the panic")
		}

		return nil, nil
	}()
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
	handle, params := m.Router.GetRoute(input.Type, input.Method)
	if handle != nil {
		for key, value := range params {
			data[key] = value
		}
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
