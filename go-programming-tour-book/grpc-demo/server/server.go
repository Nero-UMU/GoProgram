package main

import (
	"context"
	"flag"
	pb "grpc-demo/proto"
)

var port string

type GreeterServer struct{}

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口")
	flag.Parse()
}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest)
