package handler

import (
	tele "gopkg.in/telebot.v3"
)

type Start struct{}

func NewStart() *Start {
	return &Start{}
}

func (s *Start) Command() string {
	return "/start"
}

func (s *Start) Description() string {
	return "start bot"
}

func (s *Start) Handle(ctx tele.Context) error {
	return ctx.Send("Hello World!")
}

func (s *Start) Middleware() []tele.MiddlewareFunc {
	return nil
}
