package client

import (
	"context"
	"log"

	stockbitClient "github.com/michaelhendraw/stockbit_microservice/grpc/client"
	stockbitProto "github.com/michaelhendraw/stockbit_microservice/grpc/proto"
)

// StockbitClientInterface interface
type StockbitClientInterface interface {
	Search(ctx context.Context, req *stockbitProto.SearchRequest) (*stockbitProto.SearchResponse, error)
}

// NewStockbitClient func
func NewStockbitClient(grpcEndpoint string) StockbitClientInterface {
	options := stockbitClient.Options{
		Address: grpcEndpoint,
	}
	client, err := stockbitClient.GetClient(&options)
	if err != nil {
		log.Println("func NewStockbitClient", err)
	}
	return client
}
