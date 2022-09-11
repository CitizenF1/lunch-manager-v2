package handler

import (
	"lunch-manager/internal/models"

	tele "gopkg.in/telebot.v3"
)

func OnTextRequest(ctx tele.Context) error {
	if ctx.Message().Text == "111" {
		ctx.Respond(&tele.CallbackResponse{Text: "ok"})
		return ctx.Send("ok", models.AdminMenu)
	} else {
		return nil
	}
}
