package models

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	tele "gopkg.in/telebot.v3"
)

type Voter struct {
	UserName string `json:"user_name"`
	UserID   int64  `json:"user_id"`
}

type Voters struct {
	TotalUser map[string]bool `json:"total_users"`
}

func UserPollAnswer(ctx tele.Context) error {
	if ctx.PollAnswer().Sender.Username != "" {
		err := MarkVoter(ctx.PollAnswer().Sender.Username, true)
		if err != nil {
			log.Println(err, "Error markVoter")
		}
	}
	if len(ctx.PollAnswer().Options) == 0 {
		err := MarkVoter(ctx.PollAnswer().Sender.Username, false)
		if err != nil {
			log.Println(err, "Error markVoter")
		}
	}
	return nil
}

func MarkVoter(userName string, votet bool) error {
	voters, err := SetVoterJson()
	if err != nil {
		return err
	}

	voters.TotalUser[userName] = votet

	err = WriteVoters(voters)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func WriteVoters(voters Voters) error {
	b, err := json.Marshal(voters)
	if err != nil {
		return err
	}
	err = os.WriteFile("./jsons/voters.json", b, 0666)
	if err != nil {
		return err
	}
	return nil
}

func SetVoterJson() (Voters, error) {
	file, err := os.ReadFile("./jsons/voters.json")
	voters := Voters{}
	if err != nil {
		log.Println(err)
		return voters, err
	}
	err = json.Unmarshal(file, &voters)
	if err != nil {
		log.Println(err)
		return voters, err
	}
	return voters, nil
}
