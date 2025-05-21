package main

import (
	"context"
	"fmt"
)

type metricsService struct {
	next PriceFetcher
}

func NewMetricsService(next PriceFetcher) PriceFetcher {
	return &metricsService{
		next: next,
	}
}

func (s *metricsService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	// metrics storage. Placeholder for actual metrics logic, push to prometheus (gauge, counters, etc.)
	fmt.Println("Pushing metrics to prometheus")
	return s.next.FetchPrice(ctx, ticker)
}
