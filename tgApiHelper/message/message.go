package Message

import (
	"bytes"
	"github.com/AlekSi/pointer"
	"github.com/Oleg-MBO/Radio-en-Esperanto/botdb"
	"github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/client/messages"
	"github.com/olebedev/go-tgbot/models"

	"github.com/Oleg-MBO/Radio-en-Esperanto/tgApiHelper/helperTypes"
	"github.com/azer/logger"
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

//r.Message("^/list", func (c *tgbot.Context) error {
// /add
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
		botdb.ChatAddRmMessagePodcasts{
			ChatID:   resp.Payload.Result.Chat.ID,
			AddMsgID: &resp.Payload.Result.MessageID,
		}.UpdateAddMsgID()
	}
	return nil
}

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
		botdb.ChatAddRmMessagePodcasts{
			ChatID:  resp.Payload.Result.Chat.ID,
			RmMsgID: &resp.Payload.Result.MessageID,
		}.UpdateRmMsgID()
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
			messagebuffer.WriteString(ch.Channel)
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

//
//	tgApiHelperTypes.AddKeyboardPaginage(keyboard, 0)
//	//
//	//keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []*models.InlineKeyboardButton{
//	//	&models.InlineKeyboardButton{
//	//		Text:         "test2",
//	//		CallbackData: "testData",
//	//	},
//	//})
//
//	resp, err := Tgapi.Messages.SendMessage(
//		messages.NewSendMessageParams().WithBody(&models.SendMessageBody{
//			Text:        pointer.ToString(messagebuffer.String()),
//			ChatID:      c.Update.Message.Chat.ID,
//			ReplyMarkup: keyboard,
//		}),
//	)
//	if err != nil {
//		//return err
//		logMain.Info(" send msg err %s", err.Error())
//	}
//	if resp != nil {
//		logMain.Info("esp.Payload.Result.MessageID %s", resp.Payload.Result.MessageID)
//	}
//	return nil
//}

//func SubscribeChannel(c *tgbot.Context) error {
//	logMain.Info("Update.Message.Text  ", c.Update.Message.Text)
//	txt := strings.TrimSpace(c.Update.Message.Text)
//	indexToSplit := strings.Index(txt, " ")
//	if indexToSplit == -1 {
//		Tgapi.Messages.SendMessage(
//			messages.NewSendMessageParams().WithBody(&models.SendMessageBody{
//				Text:   pointer.ToString("you forget to write channel name"),
//				ChatID: c.Update.Message.Chat.ID,
//			}),
//		)
//		return nil
//	}
//
//	chName := txt[indexToSplit+1:]
//	q := botdb.PodcastChannelType{ChannelName: chName}.IsExist()
//	if q {
//		botdb.ChatPodcasts{
//			ChatID:  c.Update.Message.Chat.ID,
//			Channel: chName,
//		}.Save()
//		Tgapi.Messages.SendMessage(
//			messages.NewSendMessageParams().WithBody(&models.SendMessageBody{
//				Text:             pointer.ToString("done"),
//				ChatID:           c.Update.Message.Chat.ID,
//				ReplyToMessageID: c.Update.Message.MessageID,
//			}),
//		)
//		return nil
//	}
//
//	Tgapi.Messages.SendMessage(
//		messages.NewSendMessageParams().WithBody(&models.SendMessageBody{
//			Text:             pointer.ToString(chName + " not exist"),
//			ChatID:           c.Update.Message.Chat.ID,
//			ReplyToMessageID: c.Update.Message.MessageID,
//		}),
//	)
//	return nil
//}
