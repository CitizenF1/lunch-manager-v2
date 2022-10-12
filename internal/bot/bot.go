package bot

import (
	"lunch-manager/config"
	"lunch-manager/internal/bot/handler"

	tele "gopkg.in/telebot.v3"
)

var BotInstance *tele.Bot

func StartBot() {
	BotInstance = config.BotConfig()
	setCommands()
	BotInstance.Start()
}

func setCommands() {
	commandHandlers := []handler.CommandHandler{
		handler.NewStart(),
	}

	for _, h := range commandHandlers {
		BotInstance.Handle(h.Command(), h.Handle, h.Middleware()...)
	}

	buttonHandlers := []handler.ButtonHanders{}

	for _, b := range buttonHandlers {
		BotInstance.Handle(b.Button(), b.Handle)
	}
	BotInstance.Handle(tele.OnText, handler.OnTextRequest)
}
