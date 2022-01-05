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
	"sync"

	"google.golang.org/grpc"
)

type server struct {
	protobuf.UnimplementedExamServer
}

type Lamport struct {
	mu      sync.Mutex
	lamport int32
}

var hashTable = make([]int32, 100) 
var port string
var lamport = Lamport{}

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
	fmt.Println("Welcome Server. You need to write 0, 1 or 2 and press enter:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	port = strings.Replace(text, "\n", "", 1)

	//Start server
	lis, err := net.Listen("tcp", ":808"+port)

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

func (s *server) Put(ctx context.Context, in *protobuf.PutRequest) (*protobuf.PutReply, error) {
	log.Printf("Server %s received Put request for key: %d with value: %d", port, in.Key, in.Value)
	if in.LamportTimestamp > lamport.getLamport() {
		lamport.replaceLamport(in.LamportTimestamp)
	}
	lamport.Inc()
	hashTable[in.Key] = in.Value
	return &protobuf.PutReply{Succes: true}, nil
}

func (s *server) Get(ctx context.Context, in *protobuf.GetRequest) (*protobuf.GetReply, error) {
	log.Printf("Server %s received Get request for key: %d", port, in.Key)
	if in.LamportTimestamp > lamport.getLamport() {
		lamport.replaceLamport(in.LamportTimestamp)
	}
	lamport.Inc()
	return &protobuf.GetReply{Value: hashTable[in.Key], LamportTimestamp: 0}, nil
}

func (l *Lamport) Inc() {
	l.mu.Lock()
	l.lamport++
	l.mu.Unlock()
}

func (l *Lamport) getLamport() int32 {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.lamport
}

func (l *Lamport) replaceLamport(newLamport int32) {
	l.mu.Lock()
	l.lamport = newLamport
	defer l.mu.Unlock()
}
