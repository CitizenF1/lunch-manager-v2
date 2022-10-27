package handler

import (
	"log"
	"lunch-manager/internal/bot/cafe"
	"lunch-manager/internal/models"

	tele "gopkg.in/telebot.v3"
)

type PollHandler struct {
}

var (
	Poll        *tele.Poll
	CurrentPoll *tele.Message
	Sendet      bool
)

func NewPoll() *PollHandler {
	return &PollHandler{}
}

func (s *PollHandler) Command() string {
	return "/poll"
}

func (s *PollHandler) Description() string {
	return "poll"
}

func (s *PollHandler) Handle(ctx tele.Context) error {
	if Sendet {
		ctx.Send("Опрос уже создан")
	} else if ctx.Message().Payload != "" {
		payload := ctx.Message().Payload
		limit := 254
		if len(payload) > limit {
			cutString := payload[:limit]
			Poll = cafe.CreatePoll(cutString)
		} else {
			Poll = cafe.CreatePoll(ctx.Message().Payload)
		}
		Sendet = true
		Pollmessage, err := Poll.Send(ctx.Bot(), models.BotGroup, nil)
		if err != nil {
			log.Println(err)
		}
		CurrentPoll = Pollmessage

		ctx.Send("Опрос создан")
	} else {
		ctx.Send("Введите payload /poll {payload}")
	}
	// currentPoll, err := poll.Send(ctx.Bot(), models.BotGroup, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// CurrentPoll = currentPoll

	// Sendet = true
	return nil
}

func (s *PollHandler) Middleware() []tele.MiddlewareFunc {
	return nil
}
