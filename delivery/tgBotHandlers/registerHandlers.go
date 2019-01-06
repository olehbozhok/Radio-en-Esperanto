package tgbothandlers

import (
	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	channelsOnPage = 10
)

// RegisterHandlers is used to set handlers to tg bot
func RegisterHandlers(u radiobot.Usecase) {
	u.HandleTg("/start", StartMessage(u))
	u.HandleTg("/ek", StartMessage(u))
	u.HandleTg("/help", StartMessage(u))
	u.HandleTg("/helpo", StartMessage(u))

	u.HandleTg("/subscribe", SubskribeCommand(u, false))
	u.HandleTg("/aboni", SubskribeCommand(u, false))
	u.HandleTg("/malaboni", SubskribeCommand(u, false))

	// Handle an incoming callback query
	// from a callback button in an inline keyboard.
	u.HandleTg(tb.OnCallback, CallbackHandler(u))

}
