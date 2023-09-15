package controller

import (
	"context"
	"fmt"
	"strconv"

	"gitlab.karlson.dev/individual/wertio_hometask/internal/domain"
	"gitlab.karlson.dev/individual/wertio_hometask/internal/usecase"
)

type CLIController struct {
	interactor *usecase.Interactor
}

func NewCLIController(
	interactor *usecase.Interactor,
) *CLIController {
	return &CLIController{interactor: interactor}
}

func (c *CLIController) ConvertCurrency(ctx context.Context, amount, from, to string) *domain.Response {
	if amount == "" || from == "" || to == "" {
		return domain.NewResponse(0, 0, fmt.Errorf("usage: ./app [amount] [from_currency] [to_currency]"))
	}

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return domain.NewResponse(0, 0, err)
	}
	result, err := c.interactor.ConvertCurrency(ctx, amount, from, to)
	if err != nil {
		return domain.NewResponse(0, 0, err)
	} else {
		return domain.NewResponse(amountFloat, result, nil)
	}
}
