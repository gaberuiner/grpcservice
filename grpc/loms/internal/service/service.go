package service

import (
	"context"
	"route256/loms/internal/protoc/loms"
)

type Server struct {
	loms.UnimplementedLomsServer
}

//	type LomsServer interface {
//		CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
//		ListOrder(context.Context, *ListOrderRequest) (*ListOrderResponse, error)
//		CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
//		OrderPayed(context.Context, *OrderPayedRequest) (*OrderPayedResponse, error)
//		Stocks(context.Context, *StocksRequest) (*StocksResponse, error)
//		mustEmbedUnimplementedLomsServer()
//	}
func (s *Server) CreateOrder(context.Context, *loms.CreateOrderRequest) (*loms.CreateOrderResponse, error) {
	return &loms.CreateOrderResponse{
		OrderID: 1,
	}, nil
}
func (s *Server) ListOrder(context.Context, *loms.ListOrderRequest) (*loms.ListOrderResponse, error) {
	return &loms.ListOrderResponse{
		Status: "new",
		User:   12,
		Items:  []*loms.Items{{Sku: 1, Count: 2}, {Sku: 3, Count: 4}},
	}, nil
}
func (s *Server) CancelOrder(context.Context, *loms.CancelOrderRequest) (*loms.CancelOrderResponse, error) {
	return &loms.CancelOrderResponse{}, nil
}
func (s *Server) OrderPayed(context.Context, *loms.OrderPayedRequest) (*loms.OrderPayedResponse, error) {
	return &loms.OrderPayedResponse{}, nil

}
func (s *Server) Stocks(context.Context, *loms.StocksRequest) (*loms.StocksResponse, error) {
	return &loms.StocksResponse{
		Stocks: []*loms.Stock{{WarehouseID: 1, Count: 2}, {WarehouseID: 3, Count: 4}},
	}, nil
}

func NewServer() *Server {
	return &Server{}
}
