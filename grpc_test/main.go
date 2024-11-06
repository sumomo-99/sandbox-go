package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/sumomo-99/sandbox-go/grpc_test/grpc_test"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTestServer
}

func (s *server) GetNode(ctx context.Context, req *pb.AppVersion) (*pb.Node, error) {
	node, ok := os.LookupEnv("Battle-Node")
	message := ""

	if !ok {
		message = "not defined"
	} else {
		message = node
	}

	reply := &pb.Node{
		Node: fmt.Sprintf("Node is %s!", message),
	}
	return reply, nil
}

func (s *server) Health(ctx context.Context, req *pb.AppVersion) (*pb.Live, error) {
	live := &pb.Live{
		Live: "Ok",
	}
	return live, nil
}

func main() {
	port := flag.Int("port", 80, "The server port")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("[::]:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTestServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
