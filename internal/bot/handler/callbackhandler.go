package handler

import (
	tele "gopkg.in/telebot.v3"
)

const (
	// cunfirm callbacks
	CallbackConfirm = "confirm_call"
	CallbackCancel  = "cancel_all"
)

func OnCallBack(ctx tele.Context) error {

	return nil
}
