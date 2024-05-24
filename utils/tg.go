package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/rs/zerolog/log"
)

type TG struct {
	enable  bool
	title   string
	token   string
	chat_id string
}
type Body struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

var tg TG
var tgQueue chan func()

func SetTg(enable bool, title string, token string, chat_id string) {
	if !enable {
		return
	}

	// Set TG
	// It is singleton
	tg = TG{
		enable,
		title,
		token,
		chat_id,
	}
	tgQueue = make(chan func())

	// For thread safe
	go func() {
		for tg := range tgQueue {
			tg()
		}
	}()
}

func enqueue(tg func()) {
	tgQueue <- tg
}

func SendTg(msg string) {
	if !tg.enable {
		return
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tg.token)

	msg = fmt.Sprintf("%s\n%s", tg.title, msg)
	body := Body{
		tg.chat_id,
		msg,
		"markdown",
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Error().Err(err).Msg("\n" + string(debug.Stack()))
		return
	}
	buff := bytes.NewBuffer(bodyBytes)

	req, err := http.NewRequest(
		"POST",
		url,
		buff,
	)
	if err != nil {
		log.Error().Err(err).Msg("\n" + string(debug.Stack()))
		return
	}
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}

	tg := func() {
		resp, err := client.Do(req)
		if err != nil {
			log.Error().Err(err).Msg("\n" + string(debug.Stack()))
			return
		}
		defer resp.Body.Close()

		if !(200 <= resp.StatusCode && resp.StatusCode <= 299) {
			log.Error().Msg("Fail to seding msg from tg module")
			return
		}
	}
	enqueue(tg)
}
