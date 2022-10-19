package models

import (
	"encoding/json"
	"log"
	"os"

	tele "gopkg.in/telebot.v3"
)

var BotGroup *tele.Chat

func BotAddedToGroup(ctx tele.Context) error {
	b, err := json.Marshal(ctx.Chat())
	if err != nil {
		log.Println(err, "Error marshal json")
		return err
	}
	err = os.WriteFile("./jsons/group.json", b, 0666)
	if err != nil {
		log.Println(err, "Error write json GROUP")
		return err
	}
	return nil
}

func SetGroupID() error {
	file, err := os.ReadFile("./jsons/group.json")
	if err != nil {
		return err
	}
	botgroup := tele.Chat{}
	err = json.Unmarshal(file, &botgroup)
	if err != nil {
		return err
	}
	BotGroup = &botgroup
	return nil
}
