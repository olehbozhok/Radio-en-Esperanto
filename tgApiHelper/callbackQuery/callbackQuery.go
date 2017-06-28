package callbackQuery

import (
	"github.com/AlekSi/pointer"
	"github.com/Oleg-MBO/Radio-en-Esperanto/botdb"
	"github.com/Oleg-MBO/Radio-en-Esperanto/tgApiHelper/helperTypes"
	"github.com/azer/logger"
	"github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/client/callbacks"
	"github.com/olebedev/go-tgbot/client/messages"
	"github.com/olebedev/go-tgbot/models"
)

var Tgapi *tgbot.Router
var logQuery = logger.New("tgApiCallbackQuery")

//"^add_.+$"
func AddPodkastId(c *tgbot.Context) error {
	logQuery.Info("got callbackQuery data:%s", c.Update.CallbackQuery.Data)
	logQuery.Info("got callbackQuery ID:%s", c.Update.CallbackQuery.ID)
	logQuery.Info("got callbackQuery c.Update.CallbackQuery.Message.Chat.ID:%s", c.Update.CallbackQuery.Message.Chat.ID)
	//logQuery.Info("got callbackQuery Message:%s", c.Update.CallbackQuery.Message.Text)
	logQuery.Info("got callbackQuery Message ID:%d", c.Update.CallbackQuery.Message.MessageID)
	logQuery.Info("got callbackQuery Message .From.ID:%d", c.Update.CallbackQuery.Message.From.ID)

	answer := callbacks.NewAnswerCallbackQueryParams()
	answer.SetCallbackQueryID(c.Update.CallbackQuery.ID)
	//answer.WithText(c.Update.CallbackQuery.Data[4:])
	//answer.WithShowAlert(true)
	answer.WithText("OK")

	err := botdb.ChatPodcasts{
		Chat:    c.Update.CallbackQuery.Message.Chat.ID,
		Channel: c.Update.CallbackQuery.Data[4:],
	}.Save()

	if err != nil {
		logQuery.Error("AddPodkastId error:", err.Error())
	}
	ACQOK, err := Tgapi.Callbacks.AnswerCallbackQuery(answer)
	if err != nil {
		logQuery.Error("Tgapi.Callbacks.AnswerCallbackQuery error: %s", err.Error())
	}
	if ACQOK != nil && *ACQOK.Payload.Ok != true {
		logQuery.Error("ACQOK error: %s", ACQOK.Error())
	}

	go updateAddAndRmPodkastMSG(c.Update.CallbackQuery.Message.Chat.ID)
	return nil
}

//"^rm_.+$"
func RmPodkastId(c *tgbot.Context) error {
	logQuery.Info("got callbackQuery data:%s", c.Update.CallbackQuery.Data)
	logQuery.Info("got callbackQuery ID:%s", c.Update.CallbackQuery.ID)
	logQuery.Info("got callbackQuery c.Update.CallbackQuery.Message.Chat.ID:%s", c.Update.CallbackQuery.Message.Chat.ID)
	//logQuery.Info("got callbackQuery Message:%s", c.Update.CallbackQuery.Message.Text)
	logQuery.Info("got callbackQuery Message ID:%d", c.Update.CallbackQuery.Message.MessageID)
	logQuery.Info("got callbackQuery Message .From.ID:%d", c.Update.CallbackQuery.Message.From.ID)

	answer := callbacks.NewAnswerCallbackQueryParams()
	answer.SetCallbackQueryID(c.Update.CallbackQuery.ID)
	answer.WithText("OK")

	err := botdb.ChatPodcasts{
		Chat:    c.Update.CallbackQuery.Message.Chat.ID,
		Channel: c.Update.CallbackQuery.Data[3:],
	}.Delete()

	if err != nil {
		logQuery.Error("RmPodkastId error:", err.Error())
	}

	ACQOK, err := Tgapi.Callbacks.AnswerCallbackQuery(answer)
	if err != nil {
		logQuery.Error("Tgapi.Callbacks.AnswerCallbackQuery error: %s", err.Error())
	}
	if ACQOK != nil && *ACQOK.Payload.Ok != true {
		logQuery.Error("ACQOK error: %s", ACQOK.Error())
	}

	go updateAddAndRmPodkastMSG(c.Update.CallbackQuery.Message.Chat.ID)
	return nil
}

func updateAddAndRmPodkastMSG(chaitd int64) {
	addrmID := botdb.ChatAddRmMessagePodcasts{ChatID: chaitd}.Find()
	if addrmID.RmMsgID != nil {
		go func() {
			keyboard := &models.InlineKeyboardMarkup{
				InlineKeyboard: [][]*models.InlineKeyboardButton{},
			}

			tgApiHelperTypes.AddInlineKeyboardRmListPodcasts(keyboard, addrmID.ChatID)

			msgparams := messages.NewEditMessageTextParams()
			msgparams.WithBody(&models.EditMessageTextBody{
				ChatID:      addrmID.ChatID,
				MessageID:   *addrmID.RmMsgID,
				Text:        pointer.ToString("Klaku por disaboni kanalo"),
				ReplyMarkup: keyboard,
			})
			EMSOK, err := Tgapi.Messages.EditMessageText(msgparams)
			if err != nil {
				logQuery.Error("Tgapi.Messages.EditMessageText error: %s", err.Error())
			}
			if EMSOK != nil && *EMSOK.Payload.Ok != true {
				logQuery.Error("EMSOK error: %s", EMSOK.Error())
			}
		}()
	}

	if addrmID.AddMsgID != nil {
		go func() {
			keyboard := &models.InlineKeyboardMarkup{
				InlineKeyboard: [][]*models.InlineKeyboardButton{},
			}

			tgApiHelperTypes.AddInlineKeyboardAddListPodcasts(keyboard, addrmID.ChatID)

			msgparams := messages.NewEditMessageTextParams()
			msgparams.WithBody(&models.EditMessageTextBody{
				ChatID:      addrmID.ChatID,
				MessageID:   *addrmID.AddMsgID,
				Text:        pointer.ToString("Klaku por aboni kanalon"),
				ReplyMarkup: keyboard,
			})
			EMSOK, err := Tgapi.Messages.EditMessageText(msgparams)
			if err != nil {
				logQuery.Error("Tgapi.Messages.EditMessageText error: %s", err.Error())
			}
			if EMSOK != nil && *EMSOK.Payload.Ok != true {
				logQuery.Error("EMSOK error: %s", EMSOK.Error())
			}
		}()
	}
}

func UpdateKB(c *tgbot.Context) error {
	answer := callbacks.NewAnswerCallbackQueryParams()
	answer.SetCallbackQueryID(c.Update.CallbackQuery.ID)
	answer.WithText("klavo renovigxas...")

	ACQOK, err := Tgapi.Callbacks.AnswerCallbackQuery(answer)
	if err != nil {
		logQuery.Error("Tgapi.Callbacks.AnswerCallbackQuery error: %s", err.Error())
	}
	if ACQOK != nil && *ACQOK.Payload.Ok != true {
		logQuery.Error("ACQOK error: %s", ACQOK.Error())
	}

	go updateAddAndRmPodkastMSG(c.Update.CallbackQuery.Message.Chat.ID)
	return nil
}
