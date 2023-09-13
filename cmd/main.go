package main

import (
	"context"
	"os"

	"gitlab.karlson.dev/individual/wertio_hometask/internal/app"
)

func main() {
	if err := mainWithErr(); err != nil {
		panic(err)
	}
}

func mainWithErr() error {
	ctx := context.Background()
	cmd, err := app.NewWrapper()
	if err != nil {
		return err
	}

	output, err := cmd.Convert(ctx, os.Args[1:])
	if err != nil {
		return err
	}

	println(output)
	return nil
}
