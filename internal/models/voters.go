package models

import (
	"encoding/json"
	"log"
	"os"
)

type Voter struct {
	UserName string `json:"user_name"`
	UserID   int64  `json:"user_id"`
}

type Voters struct {
	User map[int64]string `json:"user"`
	// Voters []Voter `json:"voters"`
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
