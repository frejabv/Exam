package main

import (
	"DisysExam/protobuf"
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
)

type server struct {
	protobuf.UnimplementedExamServer
}

func main() {
	log.Print("Welcome Server. You need to write 0, 1 or 2:") // maybe change
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	port := strings.Replace(text, "\n", "", 1)

	//Start server
	lis, err := net.Listen("tcp", ":808"+port) // here to potentially

	if err != nil { //error before listening
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //we create a new server
	protobuf.RegisterExamServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil { //error while listening
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Name(ctx context.Context, in *protobuf.NameRequest) (*protobuf.NameReply, error) {
	log.Println("Server received increment")
	return &protobuf.NameReply{}, nil
}
