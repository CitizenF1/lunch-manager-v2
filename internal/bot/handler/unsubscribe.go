package handler

import (
	"encoding/json"
	"log"
	"lunch-manager/internal/bot/cafe"
	"os"

	tele "gopkg.in/telebot.v3"
)

type UnSubscribe struct{}

func NewUnSubscribe() *UnSubscribe {
	return &UnSubscribe{}
}

func (s *UnSubscribe) Command() string {
	return "/unsubscribe"
}

func (s *UnSubscribe) Description() string {
	return "UnSubscribe"
}

func (s *UnSubscribe) Handle(ctx tele.Context) error {
	cafes := cafe.Cafe{}
	b, err := os.ReadFile("./jsons/cafe.json")
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(b, &cafes)
	if err != nil {
		log.Println(err)
	}

	if cafes.Admin != nil {
		if cafes.Admin.ID == ctx.Sender().ID {
			cafes.Admin = nil
			b, err := json.Marshal(cafes)
			if err != nil {
				log.Println(err)
			}
			err = os.WriteFile("./jsons/cafe.json", b, 0666)
			if err != nil {
				log.Println(err)
			}
			ctx.Send("Вы удалены из рассылки")
		} else {
			ctx.Send("Что-то пошло не так (!id)")
		}
	} else {
		ctx.Send("Что-то пошло не так (nil pointer)")
	}

	return nil
}

func (s *UnSubscribe) Middleware() []tele.MiddlewareFunc {
	return nil
}
