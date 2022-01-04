package main

import (
	"DisysExam/protobuf"
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

func main() {
	//Set output to log file
	LOG_FILE := "log.txt" //maybe?
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	//connect to server
	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := protobuf.NewExamClient(conn)

	//make grpc call
	message, error := client.Name(context.Background(), &protobuf.NameRequest{}) //change this
	if error != nil {
		log.Fatal("Something went wrong")
	}
	log.Println(message)

	//Keep program alive
	time.Sleep(1000 * time.Second)
}
