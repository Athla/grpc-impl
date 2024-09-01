package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Athla/grpc-impl/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	PORT = ":69420"
)

type MdServer struct {
	pb.UnimplementedMdServiceServer
}

func (s *MdServer) CreateMd(ctx context.Context, in *pb.NewMd) (*pb.Md, error) {
	log.Printf("Received: %v", in.GetName())
	md := &pb.Md{
		Name:    in.GetName(),
		Content: in.GetDescription(),
		Done:    false,
		Id:      uuid.New().String(),
	}

	return md, nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Unable to listen due: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterMdServiceServer(s, &MdServer{})
	log.Printf("Server listening at: %v ", lis.Addr())
	if err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
