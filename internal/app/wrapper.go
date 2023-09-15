package app

import (
	"gitlab.karlson.dev/individual/wertio_hometask/internal/config"
	"gitlab.karlson.dev/individual/wertio_hometask/internal/controller"
	"gitlab.karlson.dev/individual/wertio_hometask/internal/presenter"
	"gitlab.karlson.dev/individual/wertio_hometask/internal/repository/coinmarketcap"
	"gitlab.karlson.dev/individual/wertio_hometask/internal/usecase"
)

type Wrapper struct {
	config     *config.Config
	Presenter  usecase.OutputPort
	Controller usecase.InputPort
	Interactor *usecase.Interactor
}

func NewWrapper() (*Wrapper, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	var w = &Wrapper{}
	w.config = cfg

	repository := coinmarketcap.New(cfg.Services.CoinMarketCap)
	w.Interactor = usecase.NewInteractor(repository)
	w.Presenter = new(presenter.Console)
	w.Controller = controller.NewCLIController(w.Interactor, w.Presenter)

	return w, nil
}
