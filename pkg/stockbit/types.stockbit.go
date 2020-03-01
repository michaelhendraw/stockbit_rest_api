package stockbit

import (
	stockbitclient "github.com/michaelhendraw/stockbit_rest_api/pkg/stockbit/client"
)

type stockbit struct {
	client stockbitclient.StockbitClientInterface
}
