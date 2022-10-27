package handler

import (
	"log"
	"lunch-manager/internal/bot/cafe"

	tele "gopkg.in/telebot.v3"
)

type Subscribe struct{}

func NewSubscribe() *Subscribe {
	return &Subscribe{}
}

func (s *Subscribe) Command() string {
	return "/subscribe"
}

func (s *Subscribe) Description() string {
	return "Subscribe"
}

func (s *Subscribe) Handle(ctx tele.Context) error {
	user := ctx.Sender()
	err := cafe.WriteToJson(*user)
	if err != nil {
		log.Println(err)
	}
	return ctx.Send("Вы подписаны на рассылку")
}

func (s *Subscribe) Middleware() []tele.MiddlewareFunc {
	return nil
}
