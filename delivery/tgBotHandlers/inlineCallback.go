package tgbothandlers

import (
	"log"
	"strconv"

	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
	tb "gopkg.in/tucnak/telebot.v2"
)

// CallbackHandler is called when need handle inline keyboard callback
func CallbackHandler(u radiobot.Usecase) func(*tb.Callback) {
	subscribeFunc := SubskribeCommand(u, true)

	return func(c *tb.Callback) {
		command, data, err := parseInlineData(c.Data)
		if err != nil {
			log.Printf("error in callbackHandler parseInlineData: %s", err.Error())
			u.Respond(c, &tb.CallbackResponse{
				ShowAlert: true,
				Text:      "Error",
			})
			return
		}
		switch command {
		case subscribeCmd:
			id := md5StrToByte(data)
			var chat radiobot.Chat
			chat.FromTgChat(c.Message.Chat)

			ch, err := u.FindChannelByID(id)
			if err != nil {
				log.Printf("error usecase.SubscribeChat :%s", err.Error())
				u.Respond(c, &tb.CallbackResponse{
					ShowAlert: true,
					Text:      "Error " + err.Error(),
				})
				return

			}
			err = u.SubscribeChat(&chat, ch)
			if err != nil {
				log.Printf("error usecase.SubscribeChat :%s", err.Error())
				u.Respond(c, &tb.CallbackResponse{
					ShowAlert: true,
					Text:      "Error",
				})
				return
			}
			u.Respond(c, &tb.CallbackResponse{
				Text: "Bone",
			})
			subscribeFunc(c.Message)

		case unsubscribeCmd:
			id := md5StrToByte(data)
			var chat radiobot.Chat
			chat.FromTgChat(c.Message.Chat)

			err := u.UnsubscribeChat(&chat, &radiobot.Channel{ID: id})
			if err != nil {
				log.Printf("error usecase.UnsubscribeChat :%s", err.Error())
				u.Respond(c, &tb.CallbackResponse{
					ShowAlert: true,
					Text:      "Error",
				})
				return
			}
			u.Respond(c, &tb.CallbackResponse{
				Text: "Bone",
			})
			subscribeFunc(c.Message)

		case toSubskribePage:
			u.Respond(c)
			i, err := strconv.Atoi(data)
			if err != nil {
				log.Printf("err hanlde `toSubskribePage` keyboard data callback  err:%v", err)
				u.Send(c.Message.Chat, "happened error", &tb.SendOptions{
					ReplyTo: c.Message,
				})
			}

			inlineKeyboard, err := createSubskribeInlineKeyboard{
				onPage:  channelsOnPage,
				page:    i,
				usecase: u,
				m:       c.Message,
			}.GetKeyboard()

			if err != nil {
				log.Printf("err SubskribeCommand chat:%s, err:%#v", c.Message.Chat.Recipient(), err)
				u.Send(c.Message.Chat, "happened error", &tb.SendOptions{
					ReplyTo: c.Message,
				})
			}

			u.Send(c.Message.Chat, "You can subscribe or unsubscribe using this keyboard", &tb.ReplyMarkup{
				InlineKeyboard: inlineKeyboard,
			})
		default:
			u.Respond(c)
		}

	}
}
