package coinmarketcap

import "time"

type response struct {
	Data   map[string]coinInfo `json:"data"`
	Status struct {
		Timestamp    time.Time `json:"timestamp"`
		ErrorCode    int       `json:"error_code"`
		ErrorMessage string    `json:"error_message"`
		Elapsed      int       `json:"elapsed"`
		CreditCount  int       `json:"credit_count"`
		Notice       string    `json:"notice"`
	} `json:"status"`
}

type coinInfo struct {
	ID     int               `json:"id"`
	Name   string            `json:"name"`
	Symbol string            `json:"symbol"`
	Quote  map[string]*quote `json:"quote"`
}

type quote struct {
	Price                 float64   `json:"price"`
	Volume24H             float64   `json:"volume_24h"`
	VolumeChange24H       float64   `json:"volume_change_24h"`
	PercentChange1H       float64   `json:"percent_change_1h"`
	PercentChange24H      float64   `json:"percent_change_24h"`
	PercentChange7D       float64   `json:"percent_change_7d"`
	PercentChange30D      float64   `json:"percent_change_30d"`
	MarketCap             float64   `json:"market_cap"`
	MarketCapDominance    int       `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
	LastUpdated           time.Time `json:"last_updated"`
}
