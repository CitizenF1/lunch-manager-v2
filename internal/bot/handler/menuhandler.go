package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"lunch-manager/internal/bot/cafe"
	"os"

	tele "gopkg.in/telebot.v3"
)

type Menu struct{}

func NewMenu() *Menu {
	return &Menu{}
}

func (s *Menu) Command() string {
	return "/menu"
}

func (s *Menu) Description() string {
	return "menu"
}

func (s *Menu) Handle(ctx tele.Context) error {
	payload := ctx.Message().Payload

	// scafe.TodayMenu = nil
	// cafe.BCafe.TodayMenu[payload] = 0
	fmt.Println(payload)

	bytes, err := json.Marshal(cafe.BCafe)
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile("./jsons/cafe.json", bytes, 0666)
	if err != nil {
		fmt.Println(err)
	}
	return ctx.Send("Принято")
}

func (s *Menu) Middleware() []tele.MiddlewareFunc {
	return nil
}
