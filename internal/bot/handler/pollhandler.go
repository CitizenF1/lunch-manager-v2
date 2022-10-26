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
		Poll = cafe.CreatePoll(ctx.Message().Payload)
		Sendet = true
		message, err := Poll.Send(ctx.Bot(), models.BotGroup, nil)
		if err != nil {
			log.Println(err)
		}
		CurrentPoll = message
		ctx.Send("Опрос создан")
	} else {
		ctx.Send("Введите payload")
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
