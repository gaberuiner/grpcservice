package orderpayed

import (
	"context"
	"errors"
	"log"
	"time"
)

type Handler struct{}

var ErrOrderNotSet = errors.New("order not set")

type Request struct {
	OrderID int64 `json:"odredID"`
}

type Response struct {
	// pusto
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
	return Response{}, nil
}
