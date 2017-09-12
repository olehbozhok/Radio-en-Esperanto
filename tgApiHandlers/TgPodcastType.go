package tgApiHandlers

import (
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/go-tgbot/client/attachments"
	"github.com/olebedev/go-tgbot/client/messages"
	"github.com/olebedev/go-tgbot/models"
	"gopkg.in/volatiletech/null.v6"
	"net/http"
	"strconv"

	dbModels "github.com/Oleg-MBO/Radio-en-Esperanto/db/models"
	"github.com/Oleg-MBO/Radio-en-Esperanto/tgApiHandlers/helperTypes"
)

func CanSendToTelegram(p dbModels.Podkast) tgApiHelperTypes.CanSendToTelegramType {
	cs := tgApiHelperTypes.CanSendToTelegramType{true, true}

	resp, err := http.Head(p.Href)
	if err != nil {
		logTg.Info(err.Error())
		cs.Status = false
		return cs
	}
	logTg.Info("status:%s ; ContentLength:%d", resp.Status, resp.ContentLength)

	if resp.Status != "200 OK" {
		cs.Status = false
		return cs
	}
	if resp.ContentLength > MaxLengthFile {
		cs.LengthIsOk = false
	}
	return cs
}

// reader to download file
func ReadCloser(p dbModels.Podkast) (*tgApiHelperTypes.InputFile, error) {
	resp, err := http.Get(p.Href)
	if err != nil {
		return &tgApiHelperTypes.InputFile{}, err
	}
	IF := &tgApiHelperTypes.InputFile{
		ReadCloser: resp.Body,
		FileName:   p.ChannelName + " " + p.RawDate + ".mp3",
	}

	IF.OK = true
	logTg.Info("status:%s ;http:%s", resp.Status, p.Href)
	return IF, nil
}

func getFileAndSendToChannel(p dbModels.Podkast) (*attachments.SendAudioOK, error) {
	Audio, err := ReadCloser(p)
	if err != nil {
		logTg.Info("ReadCloser:%s", err.Error())
		return nil, err
	}

	AudioParams := attachments.NewSendAudioParamsWithTimeout(10 * time.Minute)

	AudioParams.WithAudio(Audio)
	AudioParams.WithChatID(ChannelId)
	AudioParams.WithPerformer(pointer.ToString(p.ChannelName))
	AudioParams.WithCaption(pointer.ToString(PodcastAudioCaption(p)))

	return Tgapi.Attachments.SendAudio(AudioParams)
}

func sendFileToChannelViaLink(p dbModels.Podkast) (*attachments.SendAudioLinkOK, error) {
	AudioLinkParams := attachments.NewSendAudioLinkParams()
	AudioLinkParams.SetBody(&models.SendAudioLinkBody{
		ChatID:    ChannelId,
		Audio:     &p.Href,
		Performer: p.ChannelName,
		Title:     p.ChannelName + " " + p.RawDate,
		Caption:   PodcastAudioCaption(p),
	})

	return Tgapi.Attachments.SendAudioLink(AudioLinkParams)
}

// send Description to channel and get && set msg id to  p.TGMSGFileID
func SendFileToChannel(p *dbModels.Podkast) bool {
	logTg.Info("SendFileToChannel")

	SAOK, err := sendFileToChannelViaLink(*p)

	if err != nil {
		logTg.Info("sendFileToChannelViaLink: ERROR  ", err.Error())
	}
	if err != nil || (SAOK != nil && !SAOK.Payload.Ok) {

		SAOK, err := getFileAndSendToChannel(*p)
		if err == nil {
			logTg.Info("ERROR send file via file SAOK: %s", SAOK.Error())
		}
		if err == nil && SAOK != nil && SAOK.Payload.Ok {
			p.TGMSGFileID = null.NewInt64(SAOK.Payload.Result.MessageID, true)
			p.TGFromIDChat = SAOK.Payload.Result.Chat.ID
			return true
		}
	} else if SAOK != nil && SAOK.Payload.Ok {
		logTg.Info("sendFileToChannelViaLink: OK")
		p.TGMSGFileID = null.NewInt64(SAOK.Payload.Result.MessageID, true)
		p.TGFromIDChat = SAOK.Payload.Result.Chat.ID
		return true
	}

	return false
}

// send Description to channel and get && set msg id to  p.TgMsgIdDescription
func SendDescriptionToChannel(p *dbModels.Podkast, withFileUrl bool) bool {
	logTg.Info("send description")

	message := messages.NewSendMessageParams()
	message.WithBody(&models.SendMessageBody{
		ChatID: ChannelId,
		Text:   PodcastTextToChannel(*p, withFileUrl),
	})
	SMOK, err := Tgapi.Messages.SendMessage(message)
	if err != nil {
		logTg.Error("SendDescriptionToChannel ERROR:%s", err.Error())
		return false
	}
	if SMOK != nil && SMOK.Payload.Ok != true {
		logTg.Info("SendDescriptionToChannel ERROR:%s", SMOK.Error())
		return false
	}

	p.TGMSGFileID = null.NewInt64(SMOK.Payload.Result.MessageID, true)
	p.TGFromIDChat = SMOK.Payload.Result.Chat.ID
	return true
}

func EditDescriptionToChannel(p dbModels.Podkast) bool {
	logTg.Info("send description")

	msgparams := messages.NewEditMessageTextParams()
	msgparams.WithBody(&models.EditMessageTextBody{
		ChatID:    p.TGFromIDChat,
		MessageID: p.TGMSGIDDescription,
		Text:      PodcastTextToChannel(p, true),
	})
	EMSOK, err := Tgapi.Messages.EditMessageText(msgparams)
	if err != nil {
		logTg.Error("Tgapi.Messages.EditMessageText error: %s", err.Error())
		return false
	}
	if EMSOK != nil && *EMSOK.Payload.Ok != true {
		logTg.Error("EMSOK error: %s", EMSOK.Error())
		return false
	}
	return true
}

func SendPodcastToChannelAndSave(p dbModels.Podkast) bool {
	if cs := CanSendToTelegram(p); cs.Status == true {
		if cs.LengthIsOk {
			logTg.Info("LengthIsOk: OK")
			// send description
			if !SendDescriptionToChannel(&p, false) {
				logTg.Info("SendDescriptionToChannel: false")
				return false
			}
			if SendFileToChannel(&p) {
				logTg.Info("SendFileToChannel OK")
				if err := p.InsertG(); err != nil {
					logTg.Error("error Save SendFileToChannel : %s", err.Error())
				}
				return true
			} else {
				//error send file
				if EditDescriptionToChannel(p) {
					return true
				}
			}
		} else {
			if SendDescriptionToChannel(&p, true) {
				if err := p.InsertG(); err != nil {
					logTg.Error("error Save SendDescriptionToChannel : %s", err.Error())
				}
				return true
			}
		}
	}
	return false
}

func ForwardToAllChatsPodcast(podcast dbModels.Podkast) {
	// TODO: add iteration
	subscribers := dbModels.SubscribersPodcastsG().AllP()

	for _, sp := range subscribers {
		go func(cp dbModels.SubscribersPodcast) {

			logTg.Info(" try to send to chat: &d", cp.Chat)
			fmp := messages.NewForwardMessageParams()
			fmp.WithChatID(strconv.FormatInt(cp.Chat, 10))
			fmp.WithFromChatID(strconv.FormatInt(podcast.TGFromIDChat, 10))
			fmp.WithMessageID(podcast.TGMSGIDDescription)
			Tgapi.Messages.ForwardMessage(fmp)

			if podcast.TGMSGFileID.Valid {
				fmp := messages.NewForwardMessageParams()
				fmp.WithChatID(strconv.FormatInt(cp.Chat, 10))
				fmp.WithFromChatID(strconv.FormatInt(podcast.TGFromIDChat, 10))
				fmp.WithMessageID(podcast.TGMSGFileID.Int64)
				fmp.WithDisableNotification(pointer.ToBool(true))
				Tgapi.Messages.ForwardMessage(fmp)
			}
		}(*sp)

	}
}
