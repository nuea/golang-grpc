package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/nuea/go-grpc-auth-svc/packages/services"

	"github.com/nuea/go-grpc-auth-svc/packages/config"
	"github.com/nuea/go-grpc-auth-svc/packages/database"
	pb "github.com/nuea/go-grpc-auth-svc/packages/proto"
	"github.com/nuea/go-grpc-auth-svc/packages/utils"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config: ", err)
	}

	db := database.Connect(c)

	jwt := utils.JwtWrapper{
		SecretKey:            c.JWTSecretKey,
		Issuer:               "go-grpc-auth-svc",
		AccessTokenMinuteTTL: 60 * c.AccessTokenMinuteTTL,
	}

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Printf("Auth Service Listening on %s\n", c.Port)

	s := &services.Service{
		DB:  db,
		JWT: jwt,
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAuthServiceServer(grpcServer, s)
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
