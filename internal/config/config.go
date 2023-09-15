package config

import "gitlab.karlson.dev/individual/wertio_hometask/internal/repository/coinmarketcap"

type Config struct {
	Services struct {
		CoinMarketCap *coinmarketcap.Config
	}
}
