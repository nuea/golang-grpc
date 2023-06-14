package order

import (
	"fmt"

	"github.com/nuea/go-grpc-api-gateway/packages/config"
	pb "github.com/nuea/go-grpc-api-gateway/packages/order/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewOrderServiceClient(cc)
}
