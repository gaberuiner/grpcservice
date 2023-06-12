package service

import (
	"context"

	ps "route256/ProductService/internal/protoc/productService"
)

type Server struct {
	ps.UnimplementedProductServiceServer
}

// GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
// 	ListSkus(context.Context, *ListSkusRequest) (*ListSkusResponse, error)

func (s *Server) GetProduct(context.Context, *ps.GetProductRequest) (*ps.GetProductResponse, error) {
	return &ps.GetProductResponse{}, nil
}

func (s *Server) ListSkus(context.Context, *ps.ListSkusRequest) (*ps.ListSkusResponse, error) {
	return &ps.ListSkusResponse{}, nil
}

func NewServer() *Server {
	return &Server{}
}
