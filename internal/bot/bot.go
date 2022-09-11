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

	buttonHandlers := []handler.ButtonHanders{
		handler.NewAuthButton(BotInstance),
	}

	for _, b := range buttonHandlers {
		BotInstance.Handle(b.Button(), b.Handle)
	}
	// BotInstance.Handle(tele.OnText, handler.OnTextRequest)
	// BotInstance.Handle(tele.OnText, func(ctx tele.Context) error {
	// 	logger.
	// 		WithField("chat_id", ctx.Chat().ID).
	// 		Printf("[%s] %ds", ctx.Message().OriginalSenderName, ctx.Message().Text)

	// 	for _, handler := range messageHandlers {
	// 		err := handler.Handle(ctx)
	// 		if err != nil {
	// 			return errors.Wrap(err, "failed to handle message")
	// 		}
	// 	}

	// 	return nil
	// })
}
