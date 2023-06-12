package main

import (
	"context"
	"log"
	check "route256/checkout/internal/protoc/checkout"

	"google.golang.org/grpc"
)

const grpcPort = ":50051"

func main() {
	// Устанавливаем соединение с сервером gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при установлении соединения: %v", err)
	}
	defer conn.Close()

	// Создаем клиент gRPC
	client := check.NewCheckoutClient(conn)

	// Вызываем методы сервера

	// Пример вызова метода AddToCart
	addToCartReq := &check.AddToCartRequest{
		User:  1,
		Sku:   2,
		Count: 3,
	}
	addToCartRes, err := client.AddToCart(context.Background(), addToCartReq)
	if err != nil {
		log.Fatalf("Error AddToCart: %v", err)
	}
	log.Printf("AddToCart success: %v", addToCartRes)

	// Пример вызова метода DeleteFromCart
	deleteFromCartReq := &check.DeleteFromCartRequest{
		User:  1,
		Sku:   2,
		Count: 3,
	}
	deleteFromCartRes, err := client.DeleteFromCart(context.Background(), deleteFromCartReq)
	if err != nil {
		log.Fatalf("Error DeleteFromCart: %v", err)
	}
	log.Printf("DeleteFromCart success: %v", deleteFromCartRes)

	// Пример вызова метода ListCart
	listCartReq := &check.ListCartRequest{
		User: 1,
	}
	listCartRes, err := client.ListCart(context.Background(), listCartReq)
	if err != nil {
		log.Fatalf("Error ListCart: %v", err)
	}
	log.Printf("ListCart success: %v", listCartRes)

	// Пример вызова метода Purchase
	purchaseReq := &check.PurchaseRequest{
		User: 1,
	}
	purchaseRes, err := client.Purchase(context.Background(), purchaseReq)
	if err != nil {
		log.Fatalf("Error Purchase: %v", err)
	}
	log.Printf("Purchase sucsess: %v", purchaseRes)
}
