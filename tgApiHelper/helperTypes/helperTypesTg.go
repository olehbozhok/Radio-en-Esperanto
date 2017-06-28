package tgApiHelperTypes

import (
	"fmt"
	"github.com/Oleg-MBO/Radio-en-Esperanto/botdb"
	"github.com/olebedev/go-tgbot/models"
	"io"
	"strconv"
)

// represent ReadCloser interface in bot-api
type InputFile struct {
	io.ReadCloser
	FileName string
	OK bool
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
			&models.InlineKeyboardButton{
				Text:         ">>",
				CallbackData: "pg" + strconv.Itoa(page+1), // + " " + act,
			},
		})
	} else {
		kb.InlineKeyboard = append(kb.InlineKeyboard, []*models.InlineKeyboardButton{
			&models.InlineKeyboardButton{
				Text:         "<<",
				CallbackData: "pg" + strconv.Itoa(page-1), // + " " + act,
			},
			&models.InlineKeyboardButton{
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
			&models.InlineKeyboardButton{
				Text:         ch.ChannelName,
				CallbackData: fmt.Sprintf("add_%s", ch.ChannelName), //change Podcast
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
			&models.InlineKeyboardButton{
				Text:         ch.Channel,
				CallbackData: fmt.Sprintf("rm_%s", ch.Channel), //change Podcast
			},
		})
	}
}

func addEmpty(kb *models.InlineKeyboardMarkup) {
	kb.InlineKeyboard = append(kb.InlineKeyboard, []*models.InlineKeyboardButton{
		&models.InlineKeyboardButton{
			Text:         "malplene",
			CallbackData: "0",
		},
	})
}
