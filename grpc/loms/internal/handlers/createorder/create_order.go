package createorder

import (
	"context"
	"errors"
	"log"
)

type Handler struct{}

var ErrNoUserSet = errors.New("no user set")

type Request struct {
	User  int64 `json:"user"`
	Items Items `json:"items"`
}

type Items struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type Response struct {
	OrderID int64 `json:"orderID"`
}

func (r Request) Validate() error {
	if r.User == 0 {
		return ErrNoUserSet
	}
	return nil
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("%+v", req)

	return Response{
		OrderID: 1,
	}, nil
}
