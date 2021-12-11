package transport

import (
	"context"
	"github.com/status-owl/auth-service/proto"
)

//go:generate protoc --go_out=../../proto --go-grpc_out=../../proto  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative  --proto_path=../../proto ../../proto/authsvc.proto

type grpcServer struct {
	proto.UnimplementedAuthServiceServer
}

func (s *grpcServer) CreateAuthMethod(ctx context.Context, request *proto.CreateClientRequest) (*proto.CreateClientReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *grpcServer) mustEmbedUnimplementedAuthServiceServer() {
	//TODO implement me
	panic("implement me")
}
