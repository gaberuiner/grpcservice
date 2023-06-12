package main

import (
	"log"
	"net/http"

	"route256/libs/srvwrapper"
	"route256/notifications/internal/handlers/getproduct"
	"route256/notifications/internal/handlers/listskus"
)

const port = ":8082"

func main() {
	hand_getproduct := &getproduct.Handler{}
	hand_listskus := &listskus.Handler{}
	http.Handle("/get_product", srvwrapper.New(hand_getproduct.Handle))
	http.Handle("/get_product", srvwrapper.New(hand_listskus.Handle))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ERR: %w", err)
	}
}
