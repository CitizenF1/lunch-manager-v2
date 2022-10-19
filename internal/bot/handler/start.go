package handler

import (
	"fmt"
	"lunch-manager/internal/models"

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

	voters, err := models.SetVoterJson()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(voters.User)

	return ctx.Send("Hello World!", models.Menu)
}

func (s *Start) Middleware() []tele.MiddlewareFunc {
	return nil
}
