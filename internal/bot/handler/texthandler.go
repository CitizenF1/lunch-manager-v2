package handler

import (
	"fmt"
	"log"
	"lunch-manager/internal/models"
	"lunch-manager/internal/utils"

	tele "gopkg.in/telebot.v3"
)

func OnTextRequest(ctx tele.Context) error {
	cafes, err := utils.GetAllCafe()
	if err != nil {
		log.Printf("Error get all cafe %v", err)
		return err
	}
	for i := 0; i < len(cafes); i++ {
		fmt.Println(ctx.Message().Text)
		if ctx.Message().Text == cafes[i].CafePassword {
			fmt.Println("ok")
			ctx.Respond(&tele.CallbackResponse{Text: "ok"})
			return ctx.Send("ok", models.AdminMenu)
		} else {
			return nil
		}
	}
	return nil
}
