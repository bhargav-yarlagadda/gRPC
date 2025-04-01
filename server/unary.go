package main

import (
	"context"
	pb "gRPC/proto"
)

func (s *HelloServer) SayHello(ctx context.Context,req *pb.NoParam) (*pb.HelloResponse,error){
	return &pb.HelloResponse{
		Message: "hello for unary server",
	},nil
}