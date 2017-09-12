package tgApiHelperTypes

import (
	"fmt"
	"io"
	"strconv"

	"github.com/olebedev/go-tgbot/models"

	"github.com/Oleg-MBO/Radio-en-Esperanto/db"
)

// represent ReadCloser interface in bot-api
type InputFile struct {
	io.ReadCloser
	FileName string
	OK       bool
}

func (i *InputFile) Name() string {
	return i.FileName
}

type CanSendToTelegramType struct {
	Status     bool
	LengthIsOk bool
}

func AddKeyboardPaginage(kb *models.InlineKeyboardMarkup, page int) {
	if page == 0 {
		kb.InlineKeyboard = append(kb.InlineKeyboard, []*models.InlineKeyboardButton{
			{
				Text:         ">>",
				CallbackData: "pg" + strconv.Itoa(page+1), // + " " + act,
			},
		})
	} else {
		kb.InlineKeyboard = append(kb.InlineKeyboard, []*models.InlineKeyboardButton{
			{
				Text:         "<<",
				CallbackData: "pg" + strconv.Itoa(page-1), // + " " + act,
			},
			{
				Text:         ">>",
				CallbackData: "pg" + strconv.Itoa(page+1), // + " " + act,
			},
		})
	}
}

// add keyboard
func AddInlineKeyboardAddListPodcasts(kb *models.InlineKeyboardMarkup, chatid int64) {
	channels := botdb.GetUnsubscribesChannelsChat(chatid)
	if len(channels) == 0 {
		addEmpty(kb)
		return
	}
	for _, ch := range channels {
		kb.InlineKeyboard = append(kb.InlineKeyboard, []*models.InlineKeyboardButton{
			{
				Text:         ch,
				CallbackData: fmt.Sprintf("add_%s", ch), //change Podcast
			},
		})
	}
}

// remove keyboard
func AddInlineKeyboardRmListPodcasts(kb *models.InlineKeyboardMarkup, chatid int64) {
	channels := botdb.GetSubscribesChannelsChat(chatid)
	if len(channels) == 0 {
		addEmpty(kb)
		return
	}
	for _, ch := range channels {
		kb.InlineKeyboard = append(kb.InlineKeyboard, []*models.InlineKeyboardButton{
			{
				Text:         ch,
				CallbackData: fmt.Sprintf("rm_%s", ch), //change Podcast
			},
		})
	}
}

func addEmpty(kb *models.InlineKeyboardMarkup) {
	kb.InlineKeyboard = append(kb.InlineKeyboard, []*models.InlineKeyboardButton{
		{
			Text:         "malplene",
			CallbackData: "0",
		},
	})
}
