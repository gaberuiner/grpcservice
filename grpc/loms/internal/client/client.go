package client

import (
	"context"
	"route256/loms/internal/protoc/loms"

	"google.golang.org/grpc"
)

type Client struct {
	loms.LomsClient
}

// type LomsClient interface {
// 	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
// 	ListOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderResponse, error)
// 	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
// 	OrderPayed(ctx context.Context, in *OrderPayedRequest, opts ...grpc.CallOption) (*OrderPayedResponse, error)
// 	Stocks(ctx context.Context, in *StocksRequest, opts ...grpc.CallOption) (*StocksResponse, error)
// }

func (c *Client) CreateOrder(ctx context.Context, in *loms.CreateOrderRequest, opts ...grpc.CallOption) (*loms.CreateOrderResponse, error) {
	return &loms.CreateOrderResponse{}, nil

}
func (c *Client) ListOrder(ctx context.Context, in *loms.ListOrderRequest, opts ...grpc.CallOption) (*loms.ListOrderResponse, error) {
	return &loms.ListOrderResponse{}, nil

}
func (c *Client) CancelOrder(ctx context.Context, in *loms.CancelOrderRequest, opts ...grpc.CallOption) (*loms.CancelOrderResponse, error) {
	return &loms.CancelOrderResponse{}, nil

}
func (c *Client) OrderPayed(ctx context.Context, in *loms.OrderPayedRequest, opts ...grpc.CallOption) (*loms.OrderPayedResponse, error) {
	return &loms.OrderPayedResponse{}, nil

}
func (c *Client) Stocks(ctx context.Context, in *loms.StocksRequest, opts ...grpc.CallOption) (*loms.StocksResponse, error) {
	return &loms.StocksResponse{}, nil
}
