package auth

import (
	"fmt"

	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/nuea/go-grpc-api-gateway/packages/auth/proto"
	"github.com/nuea/go-grpc-api-gateway/packages/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	// cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure()) // grpc.WithInsecure is deprecated
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewAuthServiceClient(cc)
}
