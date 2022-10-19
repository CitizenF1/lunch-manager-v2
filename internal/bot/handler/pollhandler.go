package handler

import (
	tele "gopkg.in/telebot.v3"
)

type Poll struct{}

func NewPoll() *Poll {
	return &Poll{}
}

func (s *Poll) Command() string {
	return "/menus"
}

func (s *Poll) Description() string {
	return "menu"
}

func (s *Poll) Handle(ctx tele.Context) error {
	return ctx.Send("опрос")
}

func (s *Poll) Middleware() []tele.MiddlewareFunc {
	return nil
}
