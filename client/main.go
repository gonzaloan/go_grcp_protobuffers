package main

import (
	"context"
	"goProtogRPC/testpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := testpb.NewTestServiceClient(cc)
	//DoUnary(c)
	//DoClientStreaming(c)
	//DoServerStreaming(c)
	DoBidirectionalStreaming(c)
}

func DoUnary(c testpb.TestServiceClient) {
	req := &testpb.GetTestRequest{Id: "t1"}
	res, err := c.GetTest(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling GetTest: %v", err)
	}
	log.Printf("response from server: %v", res)
}

func DoClientStreaming(c testpb.TestServiceClient) {
	questions := []*testpb.Question{
		{
			Id:       "q8t1",
			Answer:   "Azul",
			Question: "Color related to Go",
			TestId:   "t1",
		},
		{
			Id:       "q9t1",
			Answer:   "Google",
			Question: "Company that created Go",
			TestId:   "t1",
		},
		{
			Id:       "q10t1",
			Answer:   "Backend",
			Question: "Go is Back or Front language",
			TestId:   "t1",
		},
	}

	stream, err := c.SetQuestions(context.Background())
	if err != nil {
		log.Fatalf("error while calling SetQuestions: %v", err)
	}
	for _, question := range questions {
		log.Println("sending question: ", question.Id)
		stream.Send(question)
		time.Sleep(2 * time.Second)
	}
	msg, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response: %v", err)
	}
	log.Printf("response from server: %v", msg)
}

func DoServerStreaming(c testpb.TestServiceClient) {
	req := &testpb.GetStudentsPerTestRequest{TestId: "t1"}
	stream, err := c.GetStudentsPerTest(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GetStudents: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while receiving response: %v", err)
		}
		log.Printf("response from server: %v", msg)
	}
}

func DoBidirectionalStreaming(c testpb.TestServiceClient) {
	answer := testpb.TakeTestRequest{Answer: "42"}
	numberOfQuestions := 4

	//Go routing for every question to simulate
	waitChannel := make(chan struct{})
	stream, err := c.TakeTest(context.Background())
	if err != nil {
		log.Fatalf("error while receiving response: %v", err)
	}

	//Send questions
	go func() {
		for i := 0; i < numberOfQuestions; i++ {
			stream.Send(&answer)
			time.Sleep(2 * time.Second)
		}
	}()
	//Receive
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while reading stream: %v", err)
				break
			}
			log.Printf("response from server: %v", res)
		}
	}()

	<-waitChannel
}
