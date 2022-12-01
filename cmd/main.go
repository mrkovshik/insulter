package main

import (
	"github.com/mrkovshik/insulter/grpc/proto"
    "github.com/mrkovshik/insulter/pkg/server"

	"log"
	"net"

	"google.golang.org/grpc"
)


func main() {
    // Create new gRPC server instance
    s := grpc.NewServer()


    srv := &insulter.GRPCServer{
      
    }

    // Register gRPC server for handle
    proto.RegisterInsulterServer(s, srv)

    // Listen on port 8080
    l, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal(err)
    }

    // Start gRPC server
    if err := s.Serve(l); err != nil {
        log.Fatal(err)
    }
}
