package models

import (
	"lunch-manager/internal/utils"

	"gopkg.in/telebot.v3"
)

// ПРИМЕР
var (
	BtnAuth        = telebot.Btn{Text: "Авторизация 🚪"}
	BtnCafe        = telebot.Btn{Text: "🔔 Subscribe"}
	BtnUnsubscribe = telebot.Btn{Text: "🔕 Unsubscribe"}
	AdminBtnUsers  = telebot.Btn{Text: "👤 Users"}
	AdminBtnBlogs  = telebot.Btn{Text: "📝 Blogs"}
	AdminBtnUpdate = telebot.Btn{Text: "📣 Update"}
	menuButtons    = telebot.Row{BtnAuth, BtnCafe, BtnUnsubscribe}
	Menu           = &telebot.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: utils.CreateReplyMarkup(
			menuButtons,
		),
	}
	AdminMenu = &telebot.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: utils.CreateReplyMarkup(
			// menuButtons,
			telebot.Row{AdminBtnUsers, AdminBtnBlogs, AdminBtnUpdate},
		),
	}
)

// var (
// 	MainMenu = &tele.ReplyMarkup{ResizeKeyboard: true}
// 	BtnAuth  = MainMenu.Text("Авторизация 🚪")

// 	MainMaster = &tele.ReplyMarkup{ResizeKeyboard: true}

// 	MenuCafe = &tele.ReplyMarkup{ResizeKeyboard: true}
// 	BtnCafe  = MenuCafe.Text("КАФЕ")
// )

// func InitMenuButtons() {
// 	MainMenu.Reply(
// 		MainMenu.Row(BtnAuth),
// 	)
// 	MenuCafe.Reply(
// 		MenuCafe.Row(BtnCafe),
// 	)
// }
