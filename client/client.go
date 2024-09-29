package main

import (
	"context"
	"fmt"
	proto "grpc_intro/gRPC"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the server
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}

	client := proto.NewITUDatabaseClient(conn)

	println("Connected to the server")
	for {
		// Wait for a new line
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatalf("failed to read input: %v", err)
		}

		if input == "exit" {
			break
		} else if input == "get" {
			// Get the list of students
			students, err := client.GetStudents(context.Background(), &proto.Empty{})
			if err != nil {
				log.Fatalf("failed to get students: %v", err)
			}

			for _, student := range students.Students {
				println(student)
			}
		} else {
			println("Unknown command")
		}

	}

}
