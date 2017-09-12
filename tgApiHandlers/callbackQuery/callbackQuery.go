package callbackQuery

import (
	"github.com/AlekSi/pointer"
	"github.com/azer/logger"
	"github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/client/callbacks"
	"github.com/olebedev/go-tgbot/client/messages"
	"github.com/olebedev/go-tgbot/models"

	dbModels "github.com/Oleg-MBO/Radio-en-Esperanto/db/models"
	"github.com/Oleg-MBO/Radio-en-Esperanto/tgApiHandlers/helperTypes"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

var Tgapi *tgbot.Router
var logQuery = logger.New("tgApiCallbackQuery")

//"^add_.+$"
func AddPodkastId(c *tgbot.Context) error {
	// TODO: make less data to log
	//logQuery.Info("got callbackQuery data:%s", c.Update.CallbackQuery.Data)
	//logQuery.Info("got callbackQuery ID:%s", c.Update.CallbackQuery.ID)
	//logQuery.Info("got callbackQuery c.Update.CallbackQuery.Message.Chat.ID:%s", c.Update.CallbackQuery.Message.Chat.ID)
	////logQuery.Info("got callbackQuery Message:%s", c.Update.CallbackQuery.Message.Text)
	//logQuery.Info("got callbackQuery Message ID:%d", c.Update.CallbackQuery.Message.MessageID)
	//logQuery.Info("got callbackQuery Message .From.ID:%d", c.Update.CallbackQuery.Message.From.ID)

	answer := callbacks.NewAnswerCallbackQueryParams()
	answer.SetCallbackQueryID(c.Update.CallbackQuery.ID)
	//answer.WithText(c.Update.CallbackQuery.Data[4:])
	//answer.WithShowAlert(true)
	answer.WithText("OK")

	//Save to db the subscribed channel
	pd := dbModels.SubscribersPodcast{
		Chat:        c.Update.CallbackQuery.Message.Chat.ID,
		ChannelName: c.Update.CallbackQuery.Data[4:],
	}
	//pass err if exist
	_ = pd.InsertG()


	ACQOK, err := Tgapi.Callbacks.AnswerCallbackQuery(answer)
	if err != nil {
		logQuery.Error("Tgapi.Callbacks.AnswerCallbackQuery error: %s", err.Error())
		return nil
	}
	if ACQOK != nil && ACQOK.Payload.Ok != true {
		logQuery.Error("ACQOK error: %s", ACQOK.Error())
		return nil
	}


	go updateDataInAddAndRmPodkastMSG(c.Update.CallbackQuery.Message.Chat.ID)
	return nil
}

//"^rm_.+$"
func RmPodkastId(c *tgbot.Context) error {
	//logQuery.Info("got callbackQuery data:%s", c.Update.CallbackQuery.Data)
	//logQuery.Info("got callbackQuery ID:%s", c.Update.CallbackQuery.ID)
	//logQuery.Info("got callbackQuery c.Update.CallbackQuery.Message.Chat.ID:%s", c.Update.CallbackQuery.Message.Chat.ID)
	////logQuery.Info("got callbackQuery Message:%s", c.Update.CallbackQuery.Message.Text)
	//logQuery.Info("got callbackQuery Message ID:%d", c.Update.CallbackQuery.Message.MessageID)
	//logQuery.Info("got callbackQuery Message .From.ID:%d", c.Update.CallbackQuery.Message.From.ID)

	answer := callbacks.NewAnswerCallbackQueryParams()
	answer.SetCallbackQueryID(c.Update.CallbackQuery.ID)
	answer.WithText("OK")

	// remove from db the subscribed channel
	(&dbModels.SubscribersPodcast{
		Chat:        c.Update.CallbackQuery.Message.Chat.ID,
		ChannelName: c.Update.CallbackQuery.Data[3:],
	}).DeleteGP()

	ACQOK, err := Tgapi.Callbacks.AnswerCallbackQuery(answer)
	if err != nil {
		logQuery.Error("Tgapi.Callbacks.AnswerCallbackQuery error: %s", err.Error())
		return nil
	}
	if ACQOK != nil && ACQOK.Payload.Ok != true {
		logQuery.Error("ACQOK error: %s", ACQOK.Error())
		return nil
	}

	go updateDataInAddAndRmPodkastMSG(c.Update.CallbackQuery.Message.Chat.ID)
	return nil
}

func updateDataInAddAndRmPodkastMSG(chaitd int64) {
	//create if not exist
	//(&dbModels.ChatAddRMMessagePodcast{ChatID: chaitd}).UpsertGP([]string{}, dbModels.ChatAddRMMessagePodcastColumns.ChatID)
	//get from db
	podcastMSG, err := dbModels.ChatAddRMMessagePodcastsG(qm.Where("chat_id = ?", chaitd)).One()
	if err != nil {
		logQuery.Info(err.Error())
		return
	}
	//addrmID := podcast.AddMSGID
	if podcastMSG.RMMSGID.Valid {
		// update message with rm podcasts
		go func() {
			keyboard := &models.InlineKeyboardMarkup{
				InlineKeyboard: [][]*models.InlineKeyboardButton{},
			}

			tgApiHelperTypes.AddInlineKeyboardRmListPodcasts(keyboard, podcastMSG.ChatID)

			msgparams := messages.NewEditMessageTextParams()
			msgparams.WithBody(&models.EditMessageTextBody{
				ChatID:      podcastMSG.ChatID,
				MessageID:   podcastMSG.RMMSGID.Int64,
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

	if podcastMSG.AddMSGID.Valid {
		// update message with add podcasts
		go func() {
			keyboard := &models.InlineKeyboardMarkup{
				InlineKeyboard: [][]*models.InlineKeyboardButton{},
			}

			tgApiHelperTypes.AddInlineKeyboardAddListPodcasts(keyboard, podcastMSG.ChatID)

			msgparams := messages.NewEditMessageTextParams()
			msgparams.WithBody(&models.EditMessageTextBody{
				ChatID:      podcastMSG.ChatID,
				MessageID:   podcastMSG.AddMSGID.Int64,
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
	if ACQOK != nil && ACQOK.Payload.Ok != true {
		logQuery.Error("ACQOK error: %s", ACQOK.Error())
	}

	go updateDataInAddAndRmPodkastMSG(c.Update.CallbackQuery.Message.Chat.ID)
	return nil
}
