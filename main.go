package main

import (
	"context"
	"log"
)

func main() {
	svc := NewLoggingService(NewMetricsService(&priceFetcher{}))

	price, err := svc.FetchPrice(context.Background(), "BTC")
	if err != nil {
		log.Fatal(err)
	}
	println("Price:", price)
}
