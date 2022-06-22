package main

import (
	"goProtogRPC/database"
	server2 "goProtogRPC/server"
	"goProtogRPC/studentpb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")

	server := server2.NewStudentServer(repo)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)

	reflection.Register(s)
	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
