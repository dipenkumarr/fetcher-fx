package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

// Business logic for fetching price
func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)

}

var priceMocks = map[string]float64{
	"BTC": 5000,
	"ETH": 400,
	"SOL": 200.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return 0, fmt.Errorf("the given ticker %s is not found", ticker)
	}
	return price, nil
}
