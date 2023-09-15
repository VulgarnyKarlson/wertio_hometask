package usecase

import (
	"context"
)

type InputPort interface {
	ConvertCurrency(ctx context.Context, amount string, from string, to string)
}

type OutputPort interface {
	ConversionResult(ctx context.Context, amount float64, from, to string, result float64)
	Error(err error)
}
