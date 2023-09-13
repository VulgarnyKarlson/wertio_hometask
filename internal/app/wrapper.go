package app

import (
	"gitlab.karlson.dev/individual/wertio_hometask/internal/config"
	"gitlab.karlson.dev/individual/wertio_hometask/internal/services/coinmarketcap"
)

type Wrapper struct {
	config   *config.Config
	Services struct {
		CoinMarketCap *coinmarketcap.Wrapper
	}
}

func NewWrapper() (*Wrapper, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	var w = &Wrapper{}
	w.config = cfg

	cmcAPI := coinmarketcap.New(cfg.Services.CoinMarketCap)
	w.Services.CoinMarketCap = cmcAPI

	return w, nil
}
