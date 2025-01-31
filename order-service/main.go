package main

import (
	"log"
	"net/http"

	"github.com/charlie-goldenowl/go_labs/order-service/handler"
)

func main() {
	http.HandleFunc("/place-order", handler.PlaceOrder)
	http.HandleFunc("/cancel-order", handler.CancelOrder)

	log.Println("Order Service is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
