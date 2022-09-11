package config

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func BotConfig() *tele.Bot {
	pref := tele.Settings{
		Token:  os.Getenv("BOTTOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return &tele.Bot{}
	}
	return b
}
