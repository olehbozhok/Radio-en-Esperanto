package Message

import (
	"bytes"
	"github.com/AlekSi/pointer"

	"github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/client/messages"
	"github.com/olebedev/go-tgbot/models"

	"github.com/Oleg-MBO/Radio-en-Esperanto/db"
	dbModels "github.com/Oleg-MBO/Radio-en-Esperanto/db/models"
	"github.com/Oleg-MBO/Radio-en-Esperanto/tgApiHandlers/helperTypes"
	"github.com/azer/logger"
	"gopkg.in/volatiletech/null.v6"
)

var Tgapi *tgbot.Router

var logMain = logger.New("tgApiMessage")

var helpmessage string = "Mi brodkastas dissendi certajn podkastojn en Esperanto de https://t.me/esperanto_radio\n" +
	"listo de komandoj:\n" +
	"/helpo - cxi tiu mesagxo\n" +
	"/aboni - se vi volas aboni\n" +
	"/malaboni - se vi volas malaboni\n" +
	"/listo - listo de podkastoj, kiujn vi abonis\n" +
	"mian kreanto estas @MrTrooper\n" +
	"bonan auxskultadon!"

// bind for /start command
func StartMessage(c *tgbot.Context) error {
	logMain.Info("c.Update.Message.Text: %s", c.Update.Message.Text)
	// send greeting message back
	resp, err := Tgapi.Messages.SendMessage(
		messages.NewSendMessageParams().WithBody(&models.SendMessageBody{
			Text:   &helpmessage,
			ChatID: c.Update.Message.Chat.ID,
		}),
	)
	if err != nil {
		logMain.Info("err StartMessage:%s", err.Error())
	}
	if resp != nil {
		logMain.Info("resp.Payload.Result.MessageID:%d", resp.Payload.Result.MessageID)
	}
	return nil
}

// bind for /add command
func AddListPodkasts(c *tgbot.Context) error {
	logMain.Info("c.Update.Message.Text: %s", c.Update.Message.Text)

	keyboard := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]*models.InlineKeyboardButton{},
	}

	tgApiHelperTypes.AddInlineKeyboardAddListPodcasts(keyboard, c.Update.Message.Chat.ID)

	resp, err := Tgapi.Messages.SendMessage(
		messages.NewSendMessageParams().WithBody(&models.SendMessageBody{
			Text:             pointer.ToString("Klaku por aboni kanalon"),
			ChatID:           c.Update.Message.Chat.ID,
			ReplyToMessageID: c.Update.Message.MessageID,
			ReplyMarkup:      keyboard,
		}),
	)
	if err != nil {
		return err
	}
	if resp != nil {
		logMain.Info("resp.Payload.Result.MessageID:%d", resp.Payload.Result.MessageID)
		// update id message AddListPodkasts for chat
		chMessId := dbModels.ChatAddRMMessagePodcast{
			ChatID:   resp.Payload.Result.Chat.ID,
			AddMSGID: null.NewInt64(resp.Payload.Result.MessageID, true),
		}
		//UPSERT
		chMessId.UpsertGP([]string{
			dbModels.ChatAddRMMessagePodcastColumns.AddMSGID,
		},
			dbModels.ChatAddRMMessagePodcastColumns.ChatID,
			dbModels.ChatAddRMMessagePodcastColumns.AddMSGID,
		)
	}
	return nil
}

// bind for /rm command
func RmListPodkasts(c *tgbot.Context) error {
	logMain.Info("c.Update.Message.Text: %s", c.Update.Message.Text)

	keyboard := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]*models.InlineKeyboardButton{},
	}

	tgApiHelperTypes.AddInlineKeyboardRmListPodcasts(keyboard, c.Update.Message.Chat.ID)

	resp, err := Tgapi.Messages.SendMessage(
		messages.NewSendMessageParams().WithBody(&models.SendMessageBody{
			Text:             pointer.ToString("Klaku por disaboni kanalon"),
			ChatID:           c.Update.Message.Chat.ID,
			ReplyToMessageID: c.Update.Message.MessageID,
			ReplyMarkup:      keyboard,
		}),
	)
	if err != nil {
		return err
	}
	if resp != nil {
		logMain.Info("resp.Payload.Result.MessageID:%d", resp.Payload.Result.MessageID)
		// update id message RmListPodkasts for chat
		chMessId := dbModels.ChatAddRMMessagePodcast{
			ChatID:  resp.Payload.Result.Chat.ID,
			RMMSGID: null.NewInt64(resp.Payload.Result.MessageID, true),
		}
		//UPSERT
		chMessId.UpsertGP([]string{
			dbModels.ChatAddRMMessagePodcastColumns.RMMSGID,
		},
			dbModels.ChatAddRMMessagePodcastColumns.ChatID,
			dbModels.ChatAddRMMessagePodcastColumns.RMMSGID)
	}

	return nil
}

//r.Message("^/listall", func (c *tgbot.Context) error {
func ListAllPodkasts(c *tgbot.Context) error {
	logMain.Info("c.Update.Message.Text: %s", c.Update.Message.Text)

	var messagebuffer bytes.Buffer

	listPodcast := botdb.GetSubscribesChannelsChat(c.Update.Message.Chat.ID)
	if len(listPodcast) != 0 {
		messagebuffer.WriteString("Listo de podkastoj kia vi abonis:\n")

		for _, ch := range listPodcast {
			messagebuffer.WriteString(ch)
			messagebuffer.WriteString("\n")
		}
	} else {
		messagebuffer.WriteString("listo estas malplena")
	}

	go Sendreply(c.Update.Message, messagebuffer.String())
	return nil
}

//r.Message("^.*", func (c *tgbot.Context) error {
func AllText(c *tgbot.Context) error {
	logMain.Info("/.* Chat.ID :%s Message.Text %s ", c.Update.Message.Chat.ID, c.Update.Message.Text)
	return nil
}
