package presenter

import (
	"context"
	"fmt"
	"strconv"
)

type Console struct{}

func (p *Console) ConversionResult(_ context.Context, amount float64, from, to string, result float64) {
	amountFormatted := strconv.FormatFloat(amount, 'f', -1, 64)
	resultFormatted := strconv.FormatFloat(result, 'f', -1, 64)
	fmt.Printf(
		"%s %s = %s %s \n",
		amountFormatted, from, resultFormatted, to,
	)
}

func (p *Console) Error(err error) {
	fmt.Println("Error:", err)
}
