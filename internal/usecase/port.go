package usecase

import (
	"context"

	"gitlab.karlson.dev/individual/wertio_hometask/internal/domain"
)

type InputPort interface {
	ConvertCurrency(ctx context.Context, amount string, from string, to string)
}

type OutputPort interface {
	ConversionResult(ctx context.Context, amount float64, from, to string, result float64)
	Error(err error)
}

type Controller interface {
	ConvertCurrency(ctx context.Context, amount, from, to string) *domain.Response
}
