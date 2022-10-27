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

func CreatePoll(question string) *tele.Poll {
	poll := &tele.Poll{
		Question: question,
	}
	poll.AddOptions("первое+второе+компот", "салат+Второе+компот", "первое+салат+компот", "не буду")
	return poll
}

func WriteToJson(user tele.User) error {
	cafe := Cafe{
		Name:  "WorkNeat",
		Admin: &user,
	}
	b, err := json.Marshal(&cafe)
	if err != nil {
		log.Println(err)
		return err
	}
	err = os.WriteFile("./jsons/cafe.json", b, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
