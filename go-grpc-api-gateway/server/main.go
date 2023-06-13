package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nuea/go-grpc-api-gateway/packages/auth"
	"github.com/nuea/go-grpc-api-gateway/packages/config"
	"github.com/nuea/go-grpc-api-gateway/packages/order"
	"github.com/nuea/go-grpc-api-gateway/packages/product"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config: ", err)
	}

	r := gin.Default()

	// r.Use(cors.Default())
	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
