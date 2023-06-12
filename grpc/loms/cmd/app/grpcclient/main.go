package main

import (
	"context"
	"log"
	"route256/loms/internal/protoc/loms"

	"google.golang.org/grpc"
)

const grpcPort = ":50052"

func main() {

	conn, err := grpc.Dial(grpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := loms.NewLomsClient(conn)

	ctx := context.Background()

	createOrderRequest := &loms.CreateOrderRequest{}
	createOrderResponse, err := client.CreateOrder(ctx, createOrderRequest)
	if err != nil {
		log.Fatalf("CreateOrder failed: %v", err)
	}
	log.Printf("CreateOrder response: %v", createOrderResponse)

	listOrderRequest := &loms.ListOrderRequest{}
	listOrderResponse, err := client.ListOrder(ctx, listOrderRequest)
	if err != nil {
		log.Fatalf("ListOrder failed: %v", err)
	}
	log.Printf("ListOrder response: %v", listOrderResponse)

	cancelOrderRequest := &loms.CancelOrderRequest{}
	cancelOrderResponse, err := client.CancelOrder(ctx, cancelOrderRequest)
	if err != nil {
		log.Fatalf("CancelOrder failed: %v", err)
	}
	log.Printf("CancelOrder response: %v", cancelOrderResponse)

	orderPayedRequest := &loms.OrderPayedRequest{}
	orderPayedResponse, err := client.OrderPayed(ctx, orderPayedRequest)
	if err != nil {
		log.Fatalf("OrderPayed failed: %v", err)
	}
	log.Printf("OrderPayed response: %v", orderPayedResponse)

	stocksRequest := &loms.StocksRequest{}
	stocksResponse, err := client.Stocks(ctx, stocksRequest)
	if err != nil {
		log.Fatalf("Stocks failed: %v", err)
	}
	log.Printf("Stocks response: %v", stocksResponse)
}
