package bot

import (
	"log"
	"lunch-manager/config"
	"lunch-manager/internal/bot/cafe"
	"lunch-manager/internal/bot/handler"
	"lunch-manager/internal/models"

	tele "gopkg.in/telebot.v3"
)

var (
	BotInstance *tele.Bot
)

func StartBot() {
	BotInstance = config.BotConfig()
	setCommands()
	cafe.NewCafe()
	err := models.SetGroupID()
	if err != nil {
		log.Println(err)
	}
	BotInstance.Start()
}

func setCommands() {
	commandHandlers := []handler.CommandHandler{
		handler.NewStart(),
		handler.NewMenu(),
	}

	for _, h := range commandHandlers {
		BotInstance.Handle(h.Command(), h.Handle, h.Middleware()...)
	}

	buttonHandlers := []handler.ButtonHanders{
		handler.NewAuthButton(BotInstance),
	}

	for _, b := range buttonHandlers {
		BotInstance.Handle(b.Button(), b.Handle)
	}
	BotInstance.Handle(tele.OnAddedToGroup, models.BotAddedToGroup)
	// BotInstance.Handle(tele.OnText, handler.OnTextRequest)
	// BotInstance.Handle(tele.OnAddedToGroup)

	// BotInstance.Handle(tele.OnText, func(ctx tele.Context) error {

	// 	return nil
	// })
}
