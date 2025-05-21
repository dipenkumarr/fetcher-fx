package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8080", "HTTP server listen address")
	flag.Parse()

	svc := NewLoggingService(NewMetricsService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
