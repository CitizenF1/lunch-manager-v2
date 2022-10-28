package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"lunch-manager/config"
	"lunch-manager/internal/bot/alert"
	"lunch-manager/internal/bot/buttons"
	"lunch-manager/internal/bot/cafe"
	"lunch-manager/internal/bot/handler"
	"lunch-manager/internal/models"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

var (
	BotInstance *tele.Bot

	PollSendHour   = 17
	PollSendMinute = 0

	StopPollHour   = 9
	StopPollMinute = 0

	NotificationHour   = 21
	NotificationMinute = 0
)

func StartBot() {
	BotInstance = config.BotConfig()
	setCommands()
	// cafe.NewCafe()
	err := models.SetGroupID()
	if err != nil {
		log.Println(err)
	}
	go Timer()
	BotInstance.Start()
}

func Timer() {
	location, _ := time.LoadLocation("UTC")
	for {
		hour := time.Now().In(location).Add(time.Hour * 6).Hour()
		minute := time.Now().In(location).Add(time.Hour * 6).Minute()

		if hour == PollSendHour && minute == PollSendMinute && !handler.Sendet {
			tomorrow := time.Now().Add(time.Hour * 23)
			if tomorrow.Weekday() != 6 || tomorrow.Weekday() != 0 {
				poll := cafe.CreatePoll("Обед на:" + tomorrow.Format("02.01") + " Меню: " + tomorrow.Weekday().String())
				pollMessage, err := BotInstance.Send(models.BotGroup, poll)
				if err != nil {
					log.Println(err)
				}
				err = BotInstance.Pin(pollMessage)
				if err != nil {
					log.Println(err)
				}
				handler.CurrentPoll = pollMessage
				handler.Sendet = true
			}
		}

		if hour == StopPollHour && minute == StopPollMinute && handler.Sendet {
			StopPoll()
			handler.Sendet = false
		}
		if hour == NotificationHour && minute == NotificationMinute && handler.Sendet {
			nonVoted := alert.Notification()
			if nonVoted != "" {
				message := "Кандидаты на голодовку 😅: "
				BotInstance.Send(models.BotGroup, message+nonVoted)
			}
		}
		fmt.Println(hour, ":", minute)
		time.Sleep(time.Minute)
	}
}

func setCommands() {
	commandHandlers := []handler.CommandHandler{
		handler.NewStart(),
		handler.NewMenu(),
		handler.NewPoll(),
		handler.NewDelete(),
		handler.NewSubscribe(),
		handler.NewUnSubscribe(),
	}

	for _, h := range commandHandlers {
		BotInstance.Handle(h.Command(), h.Handle, h.Middleware()...)
	}

	buttonHandlers := []handler.ButtonHanders{
		buttons.NewAuthButton(BotInstance),
	}

	for _, b := range buttonHandlers {
		BotInstance.Handle(b.Button(), b.Handle)
	}

	BotInstance.Handle(tele.OnAddedToGroup, models.BotAddedToGroup)
	BotInstance.Handle(tele.OnPollAnswer, models.UserPollAnswer)
}

func StopPoll() {
	cafes := cafe.Cafe{}
	b, err := os.ReadFile("./jsons/cafe.json")
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(b, &cafes)
	if err != nil {
		log.Println(err)
	}
	pollresult, err := BotInstance.StopPoll(handler.CurrentPoll)
	if err != nil {
		fmt.Println(err)
	}
	message := ""
	for _, value := range pollresult.Options {
		message += value.Text + " : " + fmt.Sprintf("%v", value.VoterCount) + " шт\n"
	}

	if cafes.Admin != nil {
		BotInstance.Send(cafes.Admin, "Результаты: \n"+message)
	} else {
		BotInstance.Send(models.BotGroup, "Результаты: \n"+message)
	}
}
