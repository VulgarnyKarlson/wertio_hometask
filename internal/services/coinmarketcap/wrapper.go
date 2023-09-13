package coinmarketcap

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.karlson.dev/individual/wertio_hometask/internal/common/httpclient"
	"gitlab.karlson.dev/individual/wertio_hometask/internal/domain"
)

type Wrapper struct {
	cfg        *Config
	httpClient *httpclient.Wrapper
}

func New(cfg *Config) *Wrapper {
	httpClient := httpclient.NewWrapper(cfg.URL, map[string]string{
		"X-CMC_PRO_API_KEY": cfg.APIKey,
	})
	return &Wrapper{
		cfg:        cfg,
		httpClient: httpClient,
	}
}

func (w *Wrapper) GetTicker(ctx context.Context, from, to string) (*domain.Ticker, error) {
	opts := &httpclient.Request{
		Path:   "v1/cryptocurrency/quotes/latest",
		Method: "GET",
		Query: map[string]string{
			"symbol":  from,
			"convert": to,
		},
	}
	resp, err := w.httpClient.NewRequest(ctx, opts)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error responseCode from coinmarketcap: %d", resp.StatusCode)
	}

	var resData response
	if err := json.Unmarshal(resp.Body, &resData); err != nil {
		return nil, err
	}

	if resData.Status.ErrorCode != 0 {
		return nil, fmt.Errorf("errorCode from coinmarketcap: %s", resData.Status.ErrorMessage)
	}

	if len(resData.Data) == 0 {
		return nil, fmt.Errorf("empty response from coinmarketcap, data is empty")
	}

	if _, ok := resData.Data[from]; !ok {
		return nil, fmt.Errorf("empty response from coinmarketcap, from is empty")
	}

	if _, ok := resData.Data[from].Quote[to]; !ok {
		return nil, fmt.Errorf("empty response from coinmarketcap, to is empty")
	}

	var ticker *domain.Ticker
	for _, v := range resData.Data {
		ticker = domain.NewTicker(v.ID, v.Symbol, v.Quote[to].Price)
	}

	return ticker, nil
}
