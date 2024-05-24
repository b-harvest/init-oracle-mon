package app

import (
	"bharvest.io/init-oracle-mon/utils"
)

type Config struct {
	General struct {
		ListenPort int  `toml:"listen_port"`
		ConsAddr string `toml:"cons_addr"`
		RPC string `toml:"rpc"`
	} `toml:"general"`
	Tg utils.TgConfig `toml:"tg"`
}

type BaseApp struct {
	cfg      *Config
	chVEHash chan string
	chErr    chan error
}
