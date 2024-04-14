package main

import (
	"log"
	"net"

	"github.com/lgustavopalmieri/grpc-microservices/serviceB/internal/pbCategory"
	"github.com/lgustavopalmieri/grpc-microservices/serviceB/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Não foi possível conectar: %v", err)
	}
	defer conn.Close()

	categoryService := service.NewCategoryService(conn)
	grpcServer := grpc.NewServer()
	pbCategory.RegisterCategoryServiceServer(grpcServer, categoryService)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
