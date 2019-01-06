package tgbothandlers

import (
	"fmt"

	"github.com/pkg/errors"

	tb "gopkg.in/tucnak/telebot.v2"

	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
)

type createSubskribeInlineKeyboard struct {
	onPage  int
	page    int
	usecase radiobot.Usecase
	m       *tb.Message
}

func (k createSubskribeInlineKeyboard) GetKeyboard() ([][]tb.InlineButton, error) {

	rbChat := new(radiobot.Chat)
	rbChat.FromTgChat(k.m.Chat)
	rbChat.SetID()
	err := k.usecase.FindOrRegisterChat(rbChat)
	if err != nil {
		return nil, errors.Wrap(err, "usecase.FindOrRegisterChat ")
	}
	subscribedChID := rbChat.SubscribedChannelsID

	// get first page
	channels, err := k.usecase.GetChannels(k.onPage, k.onPage*(k.page-1))
	if err != nil {
		return nil, errors.Wrap(err, "usecase.GetChannels ")
	}
	InlineKeyboard := [][]tb.InlineButton{}
	for _, ch := range channels {
		var unsubskribed bool
		for _, subskrCh := range subscribedChID {

			if ch.ID == subskrCh {
				unsubskribed = true
			}
		}
		var line []tb.InlineButton
		if !unsubskribed {
			line = []tb.InlineButton{tb.InlineButton{
				Text: ch.Name,
				Data: subscribeOn(ch.ID),
			}}
		} else {
			line = []tb.InlineButton{tb.InlineButton{
				Text: "❌ " + ch.Name,
				Data: unsubscribeOn(ch.ID),
			}}
		}
		InlineKeyboard = append(InlineKeyboard, line)
	}

	// add pagination
	line := []tb.InlineButton{}
	if k.page > 1 {
		line = append(line, tb.InlineButton{
			Text: "<<",
			Data: moveToPageSubsribe(k.page - 1),
		})
	}

	if !(len(channels) < k.onPage && k.page == 1) {
		line = append(line, tb.InlineButton{
			Text: fmt.Sprint("Paĝo ", k.page),
			Data: "_",
		})
	}
	if len(InlineKeyboard) == k.onPage {
		line = append(line, tb.InlineButton{
			Text: ">>",
			Data: moveToPageSubsribe(k.page + 1),
		})
	}
	InlineKeyboard = append(InlineKeyboard, line)

	return InlineKeyboard, nil

}
