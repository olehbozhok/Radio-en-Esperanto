package tgApiHelper

import (
	"github.com/AlekSi/pointer"
	"github.com/olebedev/go-tgbot/client/attachments"
	"github.com/olebedev/go-tgbot/models"
	"time"

	"github.com/Oleg-MBO/Radio-en-Esperanto/botdb"
	"github.com/Oleg-MBO/Radio-en-Esperanto/tgApiHelper/helperTypes"
	"github.com/olebedev/go-tgbot/client/messages"
	"net/http"
	"strconv"
)

func CanSendToTelegram(p *botdb.PodcastType) tgApiHelperTypes.CanSendToTelegramType {
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
func ReadCloser(p botdb.PodcastType) (*tgApiHelperTypes.InputFile, error) {
	resp, err := http.Get(p.Href)
	if err != nil {
		return &tgApiHelperTypes.InputFile{}, err
	}
	IF := &tgApiHelperTypes.InputFile{
		ReadCloser: resp.Body,
		FileName:   p.Name(),
	}

	IF.OK = true
	logTg.Info("status:%s ;http:%s", resp.Status, p.Href)
	return IF, nil
}

func getFileAndSendToChannel(p botdb.PodcastType) (*attachments.SendAudioOK, error) {
	Audio, err := ReadCloser(p)
	if err != nil {
		logTg.Info("ReadCloser:%s", err.Error())
		return nil, err
	}

	AudioParams := attachments.NewSendAudioParamsWithTimeout(10 * time.Minute)

	AudioParams.WithAudio(Audio)
	AudioParams.WithChatID(ChannelId)
	AudioParams.WithPerformer(pointer.ToString(p.ChannelName))
	AudioParams.WithCaption(pointer.ToString(p.Caption()))

	return Tgapi.Attachments.SendAudio(AudioParams)
}

func sendFileToChannelViaLink(p botdb.PodcastType) (*attachments.SendAudioLinkOK, error) {
	AudioLinkParams := attachments.NewSendAudioLinkParams()
	AudioLinkParams.SetBody(&models.SendAudioLinkBody{
		ChatID:    ChannelId,
		Audio:     &p.Href,
		Performer: p.ChannelName,
		Title:     p.Name(),
		Caption:   p.Caption(),
	})

	return Tgapi.Attachments.SendAudioLink(AudioLinkParams)
}

//func (p *PodcastType) SendFileToChannel() *models.Message {

// send Description to channel and get && sen msg id to  p.TgMsgIdFile
func SendFileToChannel(p *botdb.PodcastType) bool {
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
			p.TgMsgFileId = &SAOK.Payload.Result.MessageID
			p.TgFromIdChat = &SAOK.Payload.Result.Chat.ID
			return true
		}
	} else if SAOK != nil && SAOK.Payload.Ok {
		logTg.Info("sendFileToChannelViaLink: OK")
		p.TgMsgFileId = &SAOK.Payload.Result.MessageID
		p.TgFromIdChat = &SAOK.Payload.Result.Chat.ID
		return true
	}

	return false
}

// send Description to channel and get && sen msg id to  p.TgMsgIdDescription
func SendDescriptionToChannel(p *botdb.PodcastType, withFileUrl bool) bool {
	logTg.Info("send description")

	message := messages.NewSendMessageParams()
	message.WithBody(&models.SendMessageBody{
		ChatID: ChannelId,
		Text:   p.TextToChannel(withFileUrl),
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
	p.TgMsgIdDescription = &SMOK.Payload.Result.MessageID
	p.TgFromIdChat = &SMOK.Payload.Result.Chat.ID
	return true
}

func EditDescriptionToChannel(p *botdb.PodcastType) bool {
	logTg.Info("send description")

	msgparams := messages.NewEditMessageTextParams()
	msgparams.WithBody(&models.EditMessageTextBody{
		ChatID:    *p.TgFromIdChat,
		MessageID: *p.TgMsgIdDescription,
		Text:      p.TextToChannel(true),
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

func SendPodcastToChannelAndSave(p *botdb.PodcastType) bool {
	if cs := CanSendToTelegram(p); cs.Status == true {
		if cs.LengthIsOk {
			logTg.Info("LengthIsOk: OK")
			// send description
			if !SendDescriptionToChannel(p, false) {
				logTg.Info("SendDescriptionToChannel: false")
				return false
			}
			if SendFileToChannel(p) {
				logTg.Info("SendFileToChannel OK")
				if err := p.Save(); err != nil {
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
			if SendDescriptionToChannel(p, true) {
				if err := p.Save(); err != nil {
					logTg.Error("error Save SendDescriptionToChannel : %s", p.Save().Error())
				}
				return true
			}
		}
	}
	return false
}

func ForwardToAllChatsPodcast(podcast botdb.PodcastType) {
	if podcast.TgMsgIdDescription != nil {
		for _, cp1 := range podcast.Subscrbers() {
			go func(cp botdb.ChatPodcasts) {

				logTg.Info(" try to send to chat: &d", cp.Chat)
				fmp := messages.NewForwardMessageParams()
				fmp.WithChatID(strconv.FormatInt(cp.Chat, 10))
				fmp.WithFromChatID(strconv.FormatInt(*podcast.TgFromIdChat, 10))
				fmp.WithMessageID(*podcast.TgMsgIdDescription)
				Tgapi.Messages.ForwardMessage(fmp)

				if podcast.TgMsgFileId != nil {
					fmp := messages.NewForwardMessageParams()
					fmp.WithChatID(strconv.FormatInt(cp.Chat, 10))
					fmp.WithFromChatID(strconv.FormatInt(*podcast.TgFromIdChat, 10))
					fmp.WithMessageID(*podcast.TgMsgFileId)
					fmp.WithDisableNotification(pointer.ToBool(true))
					Tgapi.Messages.ForwardMessage(fmp)
				}
			}(cp1)

		}
	}
}
