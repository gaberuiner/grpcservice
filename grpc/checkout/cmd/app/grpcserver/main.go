package main

import (
	"log"
	"net"
	"route256/checkout/internal/mw/panic"
	"route256/checkout/internal/protoc/checkout"
	"route256/checkout/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = ":50051"

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("can't create listener: %s", err)
	}

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(panic.Interceptor))
	reflection.Register(s)
	checkout.RegisterCheckoutServer(s, service.NewServer())
	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %w", err)
	}

}
