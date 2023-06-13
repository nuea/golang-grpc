package order

import (
	"github.com/gin-gonic/gin"
	"github.com/nuea/go-grpc-api-gateway/packages/auth"
	"github.com/nuea/go-grpc-api-gateway/packages/config"
	"github.com/nuea/go-grpc-api-gateway/packages/order/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
