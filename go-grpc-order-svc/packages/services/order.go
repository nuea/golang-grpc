package services

import (
	"context"
	"net/http"

	"github.com/nuea/go-grpc-order-svc/packages/client"
	"github.com/nuea/go-grpc-order-svc/packages/database"
	"github.com/nuea/go-grpc-order-svc/packages/models"
	pb "github.com/nuea/go-grpc-order-svc/packages/proto"
)

type Service struct {
	pb.UnimplementedOrderServiceServer
	pb.UnimplementedProductServiceServer
	DB         database.Handler
	ProductSvc client.ProductServiceClient
}

func (s *Service) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	database := s.DB.GetDatabase()
	product, err := s.ProductSvc.FindOne(req.ProductId)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	} else if product.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{Status: product.Status, Error: product.Error}, nil
	} else if product.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: "Stock too less"}, nil
	}

	order := models.Order{
		Price:     product.Data.Price,
		ProductId: product.Data.Id,
		UserId:    req.UserId,
	}

	database.Create(&order)

	res, err := s.ProductSvc.DecreaseStock(req.ProductId, order.Id)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	} else if res.Status == http.StatusConflict {
		database.Delete(&models.Order{}, order.Id)

		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: res.Error}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}
