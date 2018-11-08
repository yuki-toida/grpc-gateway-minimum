package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/yuki-toida/grpc-gateway-sample/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}

	opts := []grpc.ServerOption{grpc.UnaryInterceptor(authentication)}
	server := grpc.NewServer(opts...)

	pb.RegisterAuthServiceServer(server, &AuthServiceServer{})
	server.Serve(listener)
}

func authentication(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println(info.FullMethod)
	fmt.Println(req)
	fmt.Println(metadata.FromIncomingContext(ctx))

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "not found metadata")
	}
	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Error(codes.Unauthenticated, "not found metadata")
	}

	fmt.Println(values[0])

	return handler(ctx, req)
}

type AuthServiceServer struct{}

func (s *AuthServiceServer) Get(c context.Context, p *pb.Param) (*pb.Param, error) {
	// fmt.Println(p)
	return p, nil
}
