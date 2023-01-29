package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/seg491X-team36/vsn-backend/schema"
	"google.golang.org/grpc"
)

type Server struct {
	schema.UnimplementedBackendServer
}

func (s *Server) Counter(stream schema.Backend_CounterServer) error {
	counterValue := 0
	for {
		// receive messages
		message, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&schema.CounterResponse{
				Value: int64(counterValue),
			})
		}
		counterValue += int(message.Increment)
		fmt.Println(counterValue)
	}
}

func main() {
	go func() {
		lis, _ := net.Listen("tcp", ":3000")
		srv := grpc.NewServer()
		schema.RegisterBackendServer(srv, &Server{})

		fmt.Println("grpc recorder listening...")
		log.Fatal(srv.Serve(lis))
	}()

	<-context.Background().Done()
}
