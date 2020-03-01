package search

// SearchRequest struct
type SearchRequest struct {
	SearchWord string `json:"search_word"`
	Pagination int64  `json:"pagination"`
}

// SearchResponse struct
type SearchResponse struct {
	Search       []SearchResponseData `json:"search"`
	TotalResults string               `json:"total_results"`
	Error        string               `json:"error"`
}

// SearchResponseData struct
type SearchResponseData struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdb_id"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}
