package main

import (
	"context"
	"log"
	"os"

	pb "backend/proto"

	"backend/logic/factorial"
	"backend/logic/fibonacci"

	"backend/server"
)

type logic struct {
	pb.UnimplementedCalculatorServer
}

func (l *logic) Fibonacci(ctx context.Context, r *pb.Req) (*pb.Resp, error) {
	return &pb.Resp{Result: fibonacci.Fibonacci(r.GetN())}, nil
}

func (l *logic) Factorial(ctx context.Context, r *pb.Req) (*pb.Resp, error) {
	return &pb.Resp{Result: factorial.Factorial(r.GetN())}, nil
}

var s *server.Server

func init() {
	s = server.New()
	s.RegisterService(&pb.Calculator_ServiceDesc, &logic{})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	log.Fatalln(s.Start(port))
}
