package main

import (
	"context"
	"fmt"
	"os"

	"bharvest.io/init-oracle-mon/app"
	"bharvest.io/init-oracle-mon/server"
	"bharvest.io/init-oracle-mon/utils"
	"github.com/pelletier/go-toml/v2"
)

func main() {
	ctx := context.Background()

	f, err := os.ReadFile("config.toml")
	if err != nil {
		utils.Error(err)
		panic(err)
	}
	cfg := app.Config{}
	err = toml.Unmarshal(f, &cfg)
	if err != nil {
		utils.Error(err)
		panic(err)
	}

	tgTitle := fmt.Sprintf("🤖 init-oracle-mon 🤖")
	utils.SetTg(cfg.Tg.Enable, tgTitle, cfg.Tg.Token, cfg.Tg.ChatID)

	go server.Run(cfg.General.ListenPort)

	baseapp := app.NewBaseApp(&cfg)
	baseapp.Run(ctx)
}
