package utils

import (
	"encoding/json"
	"lunch-manager/internal/bot/cafe"
	"os"

	"gopkg.in/telebot.v3"
)

func CreateReplyMarkup(rows ...telebot.Row) [][]telebot.ReplyButton {
	replyKeys := make([][]telebot.ReplyButton, 0, len(rows))
	for _, row := range rows {
		keys := make([]telebot.ReplyButton, 0, len(row))
		for _, btn := range row {
			btn := btn.Reply()
			if btn != nil {
				keys = append(keys, *btn)
			}
		}
		replyKeys = append(replyKeys, keys)
	}
	return replyKeys
}

func GetAllCafe() ([]cafe.Cafe, error) {
	cafes := []cafe.Cafe{}
	file, err := os.ReadFile("./jsons/cafe.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &cafes)
	if err != nil {
		return nil, err
	}
	return cafes, nil
}
