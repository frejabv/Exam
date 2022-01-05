package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"Exam/protobuf"

	"google.golang.org/grpc"
)

type server struct {
	protobuf.UnimplementedExamServer
}

var client, client1, client2 protobuf.ExamClient
var lamportTime int32 = 0
var amountOfServers = 3

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
	fmt.Println("Welcome Frontend. You need to write 0 or 1 and press enter:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	port := strings.Replace(text, "\n", "", 1)

	//Start frontend server
	go startServer(port)

	//Start frontend client(s)
	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err)
	}

	conn1, err1 := grpc.Dial(":8081", grpc.WithInsecure(), grpc.WithBlock())
	if err1 != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err1)
	}

	conn2, err2 := grpc.Dial(":8082", grpc.WithInsecure(), grpc.WithBlock())
	if err2 != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err2)
	}

	defer conn.Close()
	defer conn1.Close()
	defer conn2.Close()

	client = protobuf.NewExamClient(conn)
	client1 = protobuf.NewExamClient(conn1)
	client2 = protobuf.NewExamClient(conn2)

	fmt.Println("Frontend is running")

	//Keep program alive
	time.Sleep(1000 * time.Second)
}

func startServer(port string) {
	lis, err := net.Listen("tcp", ":807"+port)

	if err != nil { //error before listening
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //we create a new server
	protobuf.RegisterExamServer(s, &server{})

	fmt.Println("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil { //error while listening
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Put(ctx context.Context, in *protobuf.PutRequest) (*protobuf.PutReply, error) {
	lamportTime++

	message, err := client.Put(context.Background(), &protobuf.PutRequest{Key: in.Key, Value: in.Value, LamportTimestamp: lamportTime})
	message1, err1 := client1.Put(context.Background(), &protobuf.PutRequest{Key: in.Key, Value: in.Value, LamportTimestamp: lamportTime})
	message2, err2 := client2.Put(context.Background(), &protobuf.PutRequest{Key: in.Key, Value: in.Value, LamportTimestamp: lamportTime})

	if err != nil && err1 != nil && err2 != nil {
		log.Fatal("Put to servers did not succeed")
	}

	//check if a majority has returned true on succes
	var messages []*protobuf.PutReply
	if err == nil {
		messages = append(messages, message)
	}
	if err1 == nil {
		messages = append(messages, message1)
	}
	if err2 == nil {
		messages = append(messages, message2)
	}
	

	var succesCount = 0
	for i := 0; i < len(messages); i++ {
		if messages[i].Succes {
			succesCount++
		}
	}

	if succesCount >= (amountOfServers/2)+1 {
		return &protobuf.PutReply{Succes: true}, nil
	} else {
		return &protobuf.PutReply{Succes: false}, nil 
	}
}

func (s *server) Get(ctx context.Context, in *protobuf.GetRequest) (*protobuf.GetReply, error) {
	message, err := client.Get(context.Background(), &protobuf.GetRequest{Key: in.Key, LamportTimestamp: lamportTime})
	message1, err1 := client1.Get(context.Background(), &protobuf.GetRequest{Key: in.Key, LamportTimestamp: lamportTime})
	message2, err2 := client2.Get(context.Background(), &protobuf.GetRequest{Key: in.Key, LamportTimestamp: lamportTime})

	//check what the newest value is, by getting lamport timestamp
	type serverResponse struct {
		serverId int32
		lamport  int32
		value    int32
	}

	var responses []serverResponse

	if err == nil {
		response := serverResponse{
			serverId: 0,
			lamport:  message.LamportTimestamp,
			value:    message.Value,
		}
		responses = append(responses, response)
	}
	if err1 == nil {
		response := serverResponse{
			serverId: 1,
			lamport:  message1.LamportTimestamp,
			value:    message1.Value,
		}
		responses = append(responses, response)
	}
	if err2 == nil {
		response := serverResponse{
			serverId: 2,
			lamport:  message2.LamportTimestamp,
			value:    message2.Value,
		}
		responses = append(responses, response)
	}

	var serverWithHighestTime int32
	for i := 0; i < len(responses); i++ {
		if responses[i].lamport > responses[serverWithHighestTime].lamport {
			serverWithHighestTime = responses[i].serverId
		}
	}

	highestValue := responses[serverWithHighestTime].value

	//synchronise values of servers
	if err == nil && message.Value != highestValue {
		client.Put(context.Background(), &protobuf.PutRequest{Key: in.Key, Value: highestValue, LamportTimestamp: lamportTime})
	}
	if err1 == nil && message1.Value != highestValue {
		client1.Put(context.Background(), &protobuf.PutRequest{Key: in.Key, Value: highestValue, LamportTimestamp: lamportTime})
	}
	if err2 == nil && message2.Value != highestValue {
		client2.Put(context.Background(), &protobuf.PutRequest{Key: in.Key, Value: highestValue, LamportTimestamp: lamportTime})
	}

	return &protobuf.GetReply{Value: highestValue, LamportTimestamp: lamportTime}, nil
}

func (s *server) Ping(ctx context.Context, in *protobuf.PingRequest) (*protobuf.PingReply, error) {
	return &protobuf.PingReply{Alive: true}, nil
}
