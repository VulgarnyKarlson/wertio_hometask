package controller

import (
	"context"
	"fmt"
	"strconv"

	"gitlab.karlson.dev/individual/wertio_hometask/internal/usecase"
)

type CLIController struct {
	outputPort usecase.OutputPort
	interactor *usecase.Interactor
}

func NewCLIController(
	interactor *usecase.Interactor,
	outputPort usecase.OutputPort,
) *CLIController {
	return &CLIController{interactor: interactor, outputPort: outputPort}
}

func (c *CLIController) ConvertCurrency(ctx context.Context, amount, from, to string) {
	if amount == "" || from == "" || to == "" {
		c.outputPort.Error(fmt.Errorf("usage: ./app [amount] [from_currency] [to_currency]"))
		return
	}

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		c.outputPort.Error(fmt.Errorf("invalid amount: %w", err))
		return
	}
	result, err := c.interactor.ConvertCurrency(ctx, amount, from, to)
	if err != nil {
		c.outputPort.Error(err)
	} else {
		c.outputPort.ConversionResult(ctx, amountFloat, from, to, result)
	}
}
