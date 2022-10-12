package cafe

import (
	tele "gopkg.in/telebot.v3"
)

type Cafe struct {
	Name         string         `json:"name"`
	Admin        *tele.User     `json:"admin"`
	TodayMenu    map[string]int `json:"today_menu"`
	CafePassword string         `json:"admin_password"`
}
