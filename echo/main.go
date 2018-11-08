package main

import (
	"context"
	"net"
	"time"

	pb "github.com/yuki-toida/grpc-gateway-sample/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterEchoServiceServer(server, &EchoServiceServer{})
	server.Serve(listener)
}

type EchoServiceServer struct{}

func (s *EchoServiceServer) Post(c context.Context, m *pb.Message) (*pb.Message, error) {
	return m, nil
}

func (s *EchoServiceServer) Get(c context.Context, p *pb.Param) (*pb.Param, error) {
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := pb.NewAuthServiceClient(conn)
	res, err := client.Get(ctx, &pb.Param{Id: p.Id})
	if err != nil {
		panic(err)
	}
	return res, nil
}
