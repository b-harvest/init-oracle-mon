package utils

type TgConfig struct {
	Enable bool   `toml:"enable"`
	Token  string `toml:"token"`
	ChatID string `toml:"chat_id"`
}
