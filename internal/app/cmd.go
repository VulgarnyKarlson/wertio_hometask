package app

import (
	"context"
	"fmt"
)

func (w *Wrapper) Convert(ctx context.Context, args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("usage: ./app [amount] [from_currency] [to_currency]")
	}

	response := w.Controller.ConvertCurrency(ctx, args[0], args[1], args[2])
	if response.Err() != nil {
		w.Presenter.Error(response.Err())
	} else {
		w.Presenter.ConversionResult(ctx, response.Amount(), args[1], args[2], response.Result())
	}

	return nil
}
