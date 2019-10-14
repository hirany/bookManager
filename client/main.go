package main

import (
	pb "../pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {

	var bn int32
	var sn int32

	for {
		fmt.Print("student number: ")
		fmt.Scan(&sn)
		fmt.Print("book number: ")
		fmt.Scan(&bn)

		connection, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		defer connection.Close()

		client := pb.NewLendServiceClient(connection)

		context, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := client.Lend(context, &pb.LendRequest{BookNumber: bn, StudentNumber: sn})
		if err != nil {
			fmt.Println("NG")
			log.Fatal(err)
		} else {
			fmt.Println("OK")
			fmt.Println("book number    : ", response.GetBookNumber())
			fmt.Println("student number : ", response.GetStudentNumber())
		}
	}
}