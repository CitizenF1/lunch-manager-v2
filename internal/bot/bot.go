package bot

import (
	"fmt"
	"log"
	"lunch-manager/config"
	"lunch-manager/internal/bot/alert"
	"lunch-manager/internal/bot/buttons"
	"lunch-manager/internal/bot/cafe"
	"lunch-manager/internal/bot/handler"
	"lunch-manager/internal/models"
	"time"

	tele "gopkg.in/telebot.v3"
)

var (
	BotInstance *tele.Bot

	StopPollHour   = 9
	StopPollMinute = 0

	NotificationHour   = 21
	NotificationMinute = 0
)

func StartBot() {
	BotInstance = config.BotConfig()
	setCommands()
	cafe.NewCafe()
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

		if hour == StopPollHour && minute == StopPollMinute && handler.Sendet {
			StopPoll()
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
	// BotInstance.Handle(tele.OnText, handler.OnTextRequest)

	// BotInstance.Handle(tele.OnText, func(ctx tele.Context) error {

	// 	return nil
	// })
}

func StopPoll() {
	p, err := BotInstance.StopPoll(handler.CurrentPoll)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
