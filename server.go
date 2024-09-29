package main

import (
	"context"
	"fmt"
	proto "grpc_intro/gRPC"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ITU_databaseServer struct {
	proto.UnimplementedITUDatabaseServer
	students []string
}

func (s *ITU_databaseServer) GetStudents(ctx context.Context, in *proto.Empty) (*proto.Students, error) {
	return &proto.Students{Students: s.students}, nil
}

func main() {
	// Create a new gRPC server
	server := &ITU_databaseServer{students: []string{}}

	server.start_server()

	for {
		// Wait for a new line
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatalf("failed to read input: %v", err)
		}

		if input == "exit" {
			break
		}

		// Add the new student to the list
		server.students = append(server.students, input)
		println("Added student: ", input)
	}

}

func (server *ITU_databaseServer) start_server() {
	println("Starting server on port 5050")
	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer.Serve(listener)
	// Create a new gRPC server
	proto.RegisterITUDatabaseServer(grpcServer, server)

	serveError := grpcServer.Serve(listener)

	println("Server started")
	if serveError != nil {
		log.Fatalf("failed to serve: %v", serveError)
	}

}
