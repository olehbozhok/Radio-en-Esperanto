package tgbothandlers

import (
	"log"

	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
	tb "gopkg.in/tucnak/telebot.v2"
)

var helpMessage = "Mi brodkastas dissendi certajn podkastojn en Esperanto de https://t.me/esperanto_radio\n" +
	"listo de komandoj:\n" +
	"/helpo - cxi tiu mesagxo\n" +
	"/aboni - se vi volas aboni au malaboni\n" +
	"mian kreanto estas @MrTrooper\n" +
	"bonan auxskultadon!"

// StartMessage must call on /start command
func StartMessage(u radiobot.Usecase) func(*tb.Message) {
	return func(m *tb.Message) {
		u.SendTgMessage(m.Chat, helpMessage)
	}
}

// SubskribeCommand must call on /subscribe command
func SubskribeCommand(u radiobot.Usecase, editMessage bool) func(*tb.Message) {
	return func(m *tb.Message) {

		inlineKeyboard, err := createSubskribeInlineKeyboard{
			onPage:  channelsOnPage,
			page:    1,
			usecase: u,
			m:       m,
		}.GetKeyboard()

		if err != nil {
			log.Printf("err SubskribeCommand chat:%s, err:%#v", m.Chat.Recipient(), err)
			u.SendTgMessage(m.Chat, "happened error", &tb.SendOptions{
				ReplyTo: m,
			})
		}

		text := "Se vi volas aboni au malaboni al podcastoj vi povas uzi ĉi tiun klavaron"
		if editMessage {
			u.EditTgMessage(m, text, &tb.ReplyMarkup{InlineKeyboard: inlineKeyboard})
			return
		}
		u.SendTgMessage(m.Chat, text,
			&tb.ReplyMarkup{InlineKeyboard: inlineKeyboard})
	}
}
