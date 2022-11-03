package alert

import (
	"log"
	"lunch-manager/internal/models"
)

func Notification() string {
	voter, err := models.SetVoterJson()
	nonVotet := ""
	if err != nil {
		log.Println(err)
	}
	for username, value := range voter.TotalUser {
		if !value {
			nonVotet += "" + username + " "
		}
	}
	return nonVotet
}
