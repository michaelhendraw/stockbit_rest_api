package stockbit

import (
	"context"
	"errors"
	"log"

	stockbitProto "github.com/michaelhendraw/stockbit_microservice/grpc/proto"
	"github.com/michaelhendraw/stockbit_rest_api/pkg/search"
	stockbitclient "github.com/michaelhendraw/stockbit_rest_api/pkg/stockbit/client"
)

// S var
var S stockbit

// Init func
func Init(stockbitClient stockbitclient.StockbitClientInterface) {
	S.client = stockbitClient
}

// Search func
func (s *stockbit) Search(req search.SearchRequest) (*search.SearchResponse, error) {
	var result search.SearchResponse

	// build grpc request
	grpcReq := &stockbitProto.SearchRequest{
		SearchWord: req.SearchWord,
		Pagination: req.Pagination,
	}

	// hit grpc
	grpcSearchResponse, err := S.client.Search(context.Background(), grpcReq)
	if err != nil {
		log.Println("func Search", err)
		return nil, err
	}
	if grpcSearchResponse == nil {
		log.Println("func Search Response nil")
		return nil, errors.New("Response nil")
	}

	// grpcSearchResponse
	searchResponseDatas := []search.SearchResponseData{}
	if grpcSearchResponse.Search != nil {
		for _, data := range grpcSearchResponse.Search {
			searchResponseDatas = append(searchResponseDatas, search.SearchResponseData{
				Title:  data.Title,
				Year:   data.Year,
				ImdbID: data.ImdbID,
				Type:   data.Type,
				Poster: data.Poster,
			})
		}
	}
	result.Search = searchResponseDatas
	result.TotalResults = grpcSearchResponse.TotalResults
	result.Error = grpcSearchResponse.Error

	return &result, nil
}
