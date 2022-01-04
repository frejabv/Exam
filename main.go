package main

import (
	"Exam/protobuf"
	"bufio"
	"context"
	"fmt"
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
	//Set output to log file
	LOG_FILE := "log.txt" 
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	//Welcome message to register correct ports
	fmt.Print("Welcome Server. You need to write 0, 1 or 2:") // maybe change
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
	log.Println("Server received -Name- request")
	return &protobuf.NameReply{}, nil
}
