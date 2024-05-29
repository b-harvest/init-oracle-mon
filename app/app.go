package app

import (
	"context"
	"os"

	"bharvest.io/init-oracle-mon/utils"
)

func NewBaseApp(cfg *Config) *BaseApp {
	return &BaseApp{
		cfg:       cfg,
		chErr:     make(chan error, 10),
		chTimeout: make(chan struct{}, 1),
		pid:       os.Getpid(),
	}
}

func (app *BaseApp) Run(ctx context.Context) {
	appCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go app.SubVoteExtension(appCtx)

	for {
		select {
		case <-app.chTimeout:
			// Restart the subscription
			go app.SubVoteExtension(appCtx)
		case err := <-app.chErr:
			utils.Error(err)
			utils.SendTg(err.Error()) // For debugging
			return
		}
	}
}
