package main

import (
	"context"
	pb "gRPC/proto"
	"io"
	"log"
)
func callSayHelloServerStream(client pb.GreetServiceClient,names *pb.NameList){
	log.Printf("streaming started")
	stream,err := client.SayHelloServerStream(context.Background(),names)
	if(err!=nil){
		log.Fatal(err,"cannot stream")
	}
	for{
		msg,err:=stream.Recv()
		if err == io.EOF{
			break
		}
		if err!=nil{
			log.Fatal("err in streaming: ",err)
		}
		log.Print(msg)
	}
	log.Print("Streaming is completed")
}