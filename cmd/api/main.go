package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/michaelhendraw/stockbit_rest_api/pkg/config"
	"github.com/michaelhendraw/stockbit_rest_api/pkg/search"
	"github.com/michaelhendraw/stockbit_rest_api/pkg/stockbit"
	stockbitclient "github.com/michaelhendraw/stockbit_rest_api/pkg/stockbit/client"
	"github.com/michaelhendraw/stockbit_rest_api/pkg/util/generate"
)

// SearchResponse struct
type SearchResponse struct {
	Search       []SearchResponseData `json:"search,omitempty"`
	TotalResults string               `json:"total_results,omitempty"`
	Error        string               `json:"error,omitempty"`
}

// SearchResponseData struct
type SearchResponseData struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdb_id"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

// GetSearch func
func GetSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &SearchResponse{}

	searchWord := r.FormValue("searchword")
	paginationStr := r.FormValue("pagination")

	if searchWord == "" && paginationStr == "" {
		response.Error = "Searchword and pagination cannot be empty"
		json.NewEncoder(w).Encode(response)
		return
	}

	if searchWord == "" {
		response.Error = "Searchword cannot be empty"
		json.NewEncoder(w).Encode(response)
		return
	}

	if paginationStr == "" {
		response.Error = "Pagination cannot be empty"
		json.NewEncoder(w).Encode(response)
		return
	}

	pagination, err := strconv.ParseInt(paginationStr, 10, 64)
	if err != nil {
		response.Error = "Pagination should be numeric"
		json.NewEncoder(w).Encode(response)
		return
	}

	req := search.SearchRequest{
		SearchWord: searchWord,
		Pagination: pagination,
	}

	apiData, err := stockbit.S.Search(req)
	if err != nil {
		response.Error = "System error, please try again later"
		json.NewEncoder(w).Encode(response)
		return
	}
	if apiData == nil {
		response.Error = "System error, please try again later"
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(apiData)
}

func main() {
	log.Println("Starting: Stockbit API")

	// init config
	cfg := config.Init()

	// init stockbit grpc microservice
	stockbitClient := stockbitclient.NewStockbitClient(generate.CleanGRPCAddress(cfg.GRPC.StockbitEndpoint))
	stockbit.Init(stockbitClient)

	port := cfg.Server.Port
	log.Println("Running : Stockbit API PORT", port)

	router := mux.NewRouter()

	router.HandleFunc("/stockbit/search", GetSearch).Methods("GET")
	router.Path("/stockbit/search").Queries("searchword", "{searchword}", "pagination", "{pagination}").HandlerFunc(GetSearch).Methods("GET")

	http.ListenAndServe(":"+port, router)
}
