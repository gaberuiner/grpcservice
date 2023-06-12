package client

import (
	"context"
	"route256/notifications/internal/protoc/productService"

	"google.golang.org/grpc"
)

type Client struct {
	productService.ProductServiceClient
}

func (c *Client) GetProduct(ctx context.Context, in *productService.GetProductRequest, opts ...grpc.CallOption) (*productService.GetProductResponse, error) {
	return &productService.GetProductResponse{}, nil

}

func (c *Client) ListSkus(ctx context.Context, in *productService.ListSkusRequest, opts ...grpc.CallOption) (*productService.ListSkusResponse, error) {
	return &productService.ListSkusResponse{}, nil
}

// GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
// 	ListSkus(ctx context.Context, in *ListSkusRequest, opts ...grpc.CallOption) (*ListSkusResponse, error)
