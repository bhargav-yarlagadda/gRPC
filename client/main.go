package main

import (
	"log"
	pb "gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port,grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close();
	if(err!=nil){
		log.Fatalf("Error in connecting to the server: %v",err)

	}
	client := pb.NewGreetServiceClient(conn)
	// names := &pb.NameList{
	// 	Names: []string{"a","b","c","d"},
	// }
	callSayHello(client)

}