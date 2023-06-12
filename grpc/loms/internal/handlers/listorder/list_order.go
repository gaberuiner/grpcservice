package listorder

import (
	"context"
	"errors"
	"log"
	"time"
)

var ErrOrderNotSet = errors.New("order not set")

type Handler struct{}

type Request struct {
	OrderID int64 `json:"orderID"`
}

type Response struct {
	Status string `json:"status"` // (new | awaiting payment | failed | payed | cancelled)
	User   int64  `json:"user"`
	Items  Items  `json:"items"`
}

type Items struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

func (r Request) Validate() error {
	if r.OrderID == 0 {
		return ErrOrderNotSet
	}
	return nil
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("%+v", req)
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	select {
	case <-time.After(5 * time.Second):
	case <-ctx.Done():
		log.Println(ctx.Err())
		return Response{}, ctx.Err()
	}
	return Response{
		Status: "awaiting payment",
		User:   1,
		Items:  Items{Sku: 1, Count: 2},
	}, nil
}
