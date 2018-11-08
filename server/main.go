package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/yuki-toida/grpc-gateway-sample/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	opts := []grpc.ServerOption{grpc.UnaryInterceptor(interceptor)}
	server := grpc.NewServer(opts...)

	pb.RegisterEchoServiceServer(server, &EchoServiceServer{})
	server.Serve(listener)
}

type EchoServiceServer struct{}

func (s *EchoServiceServer) Echo(c context.Context, m *pb.Message) (*pb.Message, error) {
	// fmt.Println(m)
	return m, nil
}

func (s *EchoServiceServer) Get(c context.Context, p *pb.Param) (*pb.Param, error) {
	// fmt.Println(p)
	return p, nil
}

func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println(info.FullMethod)
	fmt.Println(info.Server)
	fmt.Println(req)
	fmt.Println(metadata.FromIncomingContext(ctx))
	return handler(ctx, req)
}
