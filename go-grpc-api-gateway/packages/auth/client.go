package auth

import (
	"fmt"

	pb "github.com/nuea/go-grpc-api-gateway/packages/auth/proto"
	"github.com/nuea/go-grpc-api-gateway/packages/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	// cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure()) // grpc.WithInsecure is deprecated
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(credentials.NewTLS(nil)))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
