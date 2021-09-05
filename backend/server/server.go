package server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// Server defines a gRPC server.
type Server struct {
	S *grpc.Server
}

// New inits a Server object.
func New(opt ...grpc.ServerOption) *Server {
	return &Server{grpc.NewServer(opt...)}
}

// RegisterService registres a service as defined in .proto file.
func (s *Server) RegisterService(serviceDesc *grpc.ServiceDesc, logic interface{}) *Server {
	s.S.RegisterService(serviceDesc, logic)
	return s
}

// Start starts the server and listens on the specified port.
func (s *Server) Start(port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}
	return s.S.Serve(lis)
}
