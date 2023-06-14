package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nuea/go-grpc-order-svc/packages/client"
	"github.com/nuea/go-grpc-order-svc/packages/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/nuea/go-grpc-order-svc/packages/database"

	"github.com/nuea/go-grpc-order-svc/packages/config"
	pb "github.com/nuea/go-grpc-order-svc/packages/proto"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config: ", err)
	}

	db := database.Connect(c)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listen: ", err)
	}
	fmt.Printf("Order Service Listening on %s\n", c.Port)

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	s := &services.Service{
		DB:         db,
		ProductSvc: productSvc,
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterOrderServiceServer(grpcServer, s)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
