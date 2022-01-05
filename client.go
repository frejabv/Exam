package main

import (
	"Exam/protobuf"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

var frontendPorts = []string{":8070", ":8071"}
var Cclient, client00, client01 protobuf.ExamClient

func main() {
	//Set output to log file
	LOG_FILE := "log.txt"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	//connect to frontend
	conn, err := grpc.Dial(frontendPorts[0], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client00 = protobuf.NewExamClient(conn)
	Cclient = client00

	//connect to 2nd frontend
	conn1, err1 := grpc.Dial(frontendPorts[1], grpc.WithInsecure(), grpc.WithBlock())
	if err1 != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err1)
	}

	defer conn1.Close()

	client01 = protobuf.NewExamClient(conn1)

	//start client execution
	fmt.Println("Write put followed by the key and value (space separated) to insert a value \n Write get followed by the desired key to get the associated value \n Results will be written to log file")
	go TakeInput()

	go pingFrontend()

	//Keep program alive
	time.Sleep(1000 * time.Second)
}

func TakeInput() {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		inputParsed := strings.Replace(input, "\n", "", 1)
		inputSplit := strings.Split(inputParsed, " ")

		if inputSplit[0] == "put" {
			key, err1 := strconv.ParseInt(inputSplit[1], 10, 64)
			value, err2 := strconv.ParseInt(inputSplit[2], 10, 64)
			if quickPingFrontend() {
				continue
			}
			response, err3 := Cclient.Put(context.Background(), &protobuf.PutRequest{Key: int32(key), Value: int32(value), LamportTimestamp: 0})

			if err1 != nil || err2 != nil {
				fmt.Println("Wrong input, try again")
			} else if err3 != nil {
				log.Fatal("Could not put")
			} else if err1 == nil && err2 == nil {
				log.Println("Succes:", response.Succes)
			}
		} else if inputSplit[0] == "get" {
			key, err := strconv.ParseInt(inputSplit[1], 10, 32)
			if quickPingFrontend() {
				continue
			}
			response, err1 := Cclient.Get(context.Background(), &protobuf.GetRequest{Key: int32(key), LamportTimestamp: 0})

			if err != nil {
				fmt.Println("Wrong input, try again")
			} else if err1 != nil {
				log.Fatal("Could not get")
			} else {
				log.Printf("Value for key: %d is: %d", key, response.Value)
			}
		}
	}
}

func pingFrontend() {
	for {
		if quickPingFrontend() {
			break
		}
		time.Sleep(5 * time.Second)
	}
}

func quickPingFrontend() bool {
	var changed = false
	response, err := Cclient.Ping(context.Background(), &protobuf.PingRequest{})
	if err != nil || !response.Alive {
		Cclient = client01
		changed = true
	}
	return changed
}
