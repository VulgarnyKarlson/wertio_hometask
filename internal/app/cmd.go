package app

import (
	"context"
	"fmt"
	"os"
	"strconv"
)

func (w *Wrapper) Convert(ctx context.Context, args []string) (string, error) {
	if len(args) != 3 {
		return "", fmt.Errorf("usage: ./app [amount] [from_currency] [to_currency]")
	}

	amount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		return "", fmt.Errorf("invalid amount: %w", err)
	}

	from, to := os.Args[2], os.Args[3]

	ticker, err := w.Services.CoinMarketCap.GetTicker(ctx, from, to)
	if err != nil {
		return "", err
	}

	amountFormatted := strconv.FormatFloat(amount, 'f', -1, 64)

	return fmt.Sprintf(
		"%s %s = %s %s",
		amountFormatted, from, strconv.FormatFloat(amount*ticker.Price(), 'f', -1, 64), to,
	), nil
}
