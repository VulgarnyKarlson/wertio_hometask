package app

import (
	"context"
	"fmt"
)

func (w *Wrapper) Convert(ctx context.Context, args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("usage: ./app [amount] [from_currency] [to_currency]")
	}
	w.Controller.ConvertCurrency(ctx, args[0], args[1], args[2])
	return nil
}
