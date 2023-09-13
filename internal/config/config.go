package config

import "gitlab.karlson.dev/individual/wertio_hometask/internal/services/coinmarketcap"

type Config struct {
	Services struct {
		CoinMarketCap *coinmarketcap.Config
	}
}
