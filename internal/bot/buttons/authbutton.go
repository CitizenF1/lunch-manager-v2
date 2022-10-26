package buttons

import (
	"lunch-manager/internal/models"

	tele "gopkg.in/telebot.v3"
)

type AuthButton struct {
	bot *tele.Bot
}

func NewAuthButton(bot *tele.Bot) *AuthButton {
	return &AuthButton{bot: bot}
}

func (a *AuthButton) Button() *tele.Btn {
	return &models.BtnAuth
}

func (a *AuthButton) Handle(ctx tele.Context) error {
	return ctx.Send("Введите пароль")
}
