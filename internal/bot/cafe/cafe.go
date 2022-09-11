package cafe

import (
	tele "gopkg.in/telebot.v3"
)

type Cafe struct {
	Name         string         `json:"name"`
	Admin        *tele.User     `json:"admin"`
	Menu         map[string]int `json:"menu"`
	CafePassword string         `json:"cafe_password"`
}
