package usecase

import (
	"context"
	"strconv"

	"gitlab.karlson.dev/individual/wertio_hometask/internal/repository/coinmarketcap"
)

type Interactor struct {
	Repo *coinmarketcap.Wrapper
}

func NewInteractor(repo *coinmarketcap.Wrapper) *Interactor {
	return &Interactor{Repo: repo}
}

func (c *Interactor) ConvertCurrency(ctx context.Context, amount, from, to string) (result float64, err error) {
	ticker, err := c.Repo.GetTicker(ctx, from, to)
	if err != nil {
		return result, err
	}

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return result, err
	}
	result = amountFloat * ticker.Price()
	return result, err
}
