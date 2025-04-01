package main

import (
	"log"
	"net"
	pb "gRPC/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)


type HelloServer struct{
	pb.GreetServiceServer
}
func main() {
	lis, err := net.Listen("tcp",port)
	if(err!=nil){
		log.Fatal("Failed to start server.")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer,&HelloServer{})
	if err:= grpcServer.Serve(lis);err!=nil{
		log.Fatalf("Failed to start: %v",err)
	}

}	