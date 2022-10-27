package handler

import (
	log "fmt"
	"lunch-manager/internal/models"

	tele "gopkg.in/telebot.v3"
)

type Delete struct{}

func NewDelete() *Delete {
	return &Delete{}
}

func (s *Delete) Command() string {
	return "/delete"
}

func (s *Delete) Description() string {
	return "delete uset"
}

func (s *Delete) Handle(ctx tele.Context) error {
	if ctx.Message().Payload != "" {
		payload := ctx.Message().Payload
		voters, err := models.SetVoterJson()
		if err != nil {
			log.Println(err)
		}
		for username := range voters.TotalUser {
			if username == payload {
				delete(voters.TotalUser, payload)
			} else {
				ctx.Send("Пользователь не найден")
			}
		}
		if models.WriteVoters(voters); err != nil {
			log.Println(err)
		}
		ctx.Send("Пользователь удален")
	} else {
		ctx.Send("Введите username пользователся которого хотите удалить после команды /delete {UserName}")
	}

	return nil
}

func (s *Delete) Middleware() []tele.MiddlewareFunc {
	return nil
}
