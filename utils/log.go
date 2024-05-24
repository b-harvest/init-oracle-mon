package utils

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// Logger setup
	output := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC1123,
	}
	log.Logger = log.Output(output)
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Error(err error) {
	stack := string(debug.Stack())
	log.Error().Err(err).Msg("\n" + stack)
}

func Debug(msg any) {
	message := fmt.Sprint(msg)
	log.Debug().Msg(message)
}
