package handler

import tele "gopkg.in/telebot.v3"

type CommandHandler interface {
	Command() string
	Description() string
	Handle(ctx tele.Context) error
	Middleware() []tele.MiddlewareFunc
}

type InlineBtnHandler interface {
	tele.CallbackEndpoint
	Description() string
	Handle(ctx tele.Context) error
	Middleware() []tele.MiddlewareFunc
}

type ButtonHanders interface {
	Button() *tele.Btn
	Handle(ctx tele.Context) error
}
