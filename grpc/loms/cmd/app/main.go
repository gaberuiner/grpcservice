package main

import (
	"log"
	"net/http"

	"route256/libs/srvwrapper"
	"route256/loms/internal/handlers/cancelorder"
	"route256/loms/internal/handlers/createorder"
	"route256/loms/internal/handlers/listorder"
	"route256/loms/internal/handlers/orderpayed"
	"route256/loms/internal/handlers/stocks"
)

const port = ":8081"

func main() {
	hand_stocks := &stocks.Handler{}
	hand_createorder := &createorder.Handler{}
	hand_listorder := &listorder.Handler{}
	hand__orderpayed := &orderpayed.Handler{}
	hand_cancelorder := &cancelorder.Handler{}
	http.Handle("/stocks", srvwrapper.New(hand_stocks.Handle))
	http.Handle("/createOrder", srvwrapper.New(hand_createorder.Handle))
	http.Handle("/listOrder", srvwrapper.New(hand_listorder.Handle))
	http.Handle("/orderPayed", srvwrapper.New(hand__orderpayed.Handle))
	http.Handle("/cancelOrder", srvwrapper.New(hand_cancelorder.Handle))

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ERR: ", err)
	}
}
