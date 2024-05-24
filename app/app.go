package app

import (
	"context"

	"bharvest.io/init-oracle-mon/utils"
)

func NewBaseApp(cfg *Config) *BaseApp {
	return &BaseApp{
		cfg:   cfg,
		chErr: make(chan error, 10),
	}
}

func (app *BaseApp) Run(ctx context.Context) {
	appCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go app.SubVoteExtension(appCtx)
	for {
		select {
		case err := <-app.chErr:
			utils.Error(err)
			return
		}
	}
}
