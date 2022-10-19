package cafe

import (
	"encoding/json"
	"log"
	"os"

	tele "gopkg.in/telebot.v3"
)

var BCafe *Cafe

type Cafe struct {
	Name         string         `json:"name"`
	Admin        *tele.User     `json:"admin"`
	TodayMenu    map[string]int `json:"today_menu"`
	CafePassword string         `json:"admin_password"`
}

func NewCafe() {
	file, err := os.ReadFile("./jsons/cafe.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(file, &BCafe)
	if err != nil {
		log.Println(err, "NEW CAFE")
	}
}

func WriteToJson() {

}
