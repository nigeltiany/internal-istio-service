package main

import (
	"context"
	"github.com/nigeltiany/micro-istio-internal/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type InternalServiceServer struct {

}

func (server *InternalServiceServer) Message(ctx context.Context, request *Internal.InternalRequest) (*Internal.InternalResponse, error) {
	response := &Internal.InternalResponse{}
	response.Message = request.Message
	return response, nil
}


func main () {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	Internal.RegisterInternalServiceServer(grpcServer, &InternalServiceServer{})

	grpcServer.Serve(lis)
}