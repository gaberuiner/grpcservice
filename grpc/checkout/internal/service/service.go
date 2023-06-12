package service

import (
	"context"
	"route256/checkout/internal/protoc/checkout"
)

type Service struct {
	checkout.UnimplementedCheckoutServer
}

// type CheckoutClient interface {
// 	AddToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*AddToCartResponse, error)
// 	DeleteFromCart(ctx context.Context, in *DeleteFromCartRequest, opts ...grpc.CallOption) (*DeleteFromCartResponse, error)
// 	ListCart(ctx context.Context, in *ListCartRequest, opts ...grpc.CallOption) (*ListCartResponse, error)
// 	Purchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error)
// }

func (s *Service) AddToCart(context.Context, *checkout.AddToCartRequest) (*checkout.AddToCartResponse, error) {
	return &checkout.AddToCartResponse{}, nil
}

func (s *Service) DeleteFromCart(context.Context, *checkout.DeleteFromCartRequest) (*checkout.DeleteFromCartResponse, error) {
	return &checkout.DeleteFromCartResponse{}, nil
}

func (s *Service) ListCart(context.Context, *checkout.ListCartRequest) (*checkout.ListCartResponse, error) {
	return &checkout.ListCartResponse{
		Items: []*checkout.Items{{Sku: 1, Count: 2, Name: "Name", Price: 12}, {Sku: 2, Count: 1, Name: "Name2", Price: 4}},
	}, nil
}

func (s *Service) Purchase(context.Context, *checkout.PurchaseRequest) (*checkout.PurchaseResponse, error) {
	return &checkout.PurchaseResponse{
		OrderID: 1,
	}, nil
}

func NewServer() *Service {
	return &Service{}
}
