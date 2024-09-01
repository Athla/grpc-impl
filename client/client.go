package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/Athla/grpc-impl/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	ADDR = "localhost:8080"
)

type MdTask struct {
	Name    string
	Content string
	Done    bool
}

func readMd(filename string) string {
	cnt, _ := os.ReadFile(filename)
	return string(cnt)
}

func main() {
	conn, err := grpc.NewClient(ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewMdServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	mds := []MdTask{
		{Name: "Microservices in Go", Content: readMd("~/git/golang-dsa/Concepts/Microservice_Go.md"), Done: false},
		{Name: "Recursion", Content: readMd("~/git/golang-dsa/Concepts/Recursion.md"), Done: false},
		{Name: "REST", Content: readMd("~/git/golang-dsa/Concepts/REST.md"), Done: false},
	}

	for _, md := range mds {
		res, err := c.CreateMd(ctx, &pb.NewMd{Name: md.Name, Description: md.Content, Done: md.Done})
		if err != nil {
			log.Fatalf("Could not create md reference due: %v", err)
		}

		log.Printf(`
			ID: %v
			Name: %v
			Description: %v
			Done: %v
			`,
			res.GetId(),
			res.GetName(),
			res.GetContent(),
			res.GetDone())

	}
}
