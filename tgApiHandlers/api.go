package tgApiHandlers

import (
	"time"

	"github.com/azer/logger"
	"github.com/olebedev/go-tgbot"

	"github.com/Oleg-MBO/Radio-en-Esperanto/db"
	"github.com/Oleg-MBO/Radio-en-Esperanto/parser"
	"github.com/Oleg-MBO/Radio-en-Esperanto/tgApiHandlers/callbackQuery"
	"github.com/Oleg-MBO/Radio-en-Esperanto/tgApiHandlers/message"
)

var logTg = logger.New("tgApiHelper")

var Tgapi *tgbot.Router
var ChannelId string

const MaxLengthFile = 50000000 //limit file in tg api

var TimeDelayPodcastParse = 10 * time.Minute

func InitApi(token string, channelId string) *tgbot.Router {
	Tgapi = tgbot.New(&tgbot.Options{
		Context: nil,
		Token:   token,
	})
	Message.Tgapi = Tgapi
	callbackQuery.Tgapi = Tgapi

	ChannelId = channelId

	//Bind handlers
	Tgapi.Message("^/start", Message.StartMessage)
	Tgapi.Message("^/ek", Message.StartMessage)
	Tgapi.Message("^/help", Message.StartMessage)
	Tgapi.Message("^/helpo", Message.StartMessage)
	Tgapi.Message("^/add", Message.AddListPodkasts)
	Tgapi.Message("^/aboni", Message.AddListPodkasts)
	Tgapi.Message("^/rm", Message.RmListPodkasts)
	Tgapi.Message("^/malaboni", Message.RmListPodkasts)
	Tgapi.Message("^/list", Message.ListAllPodkasts)
	Tgapi.Message("^/listo", Message.ListAllPodkasts)
	//Tgapi.Message("^.*", Message.AllText)

	Tgapi.CallbackQuery(`^add_.+$`, callbackQuery.AddPodkastId)
	Tgapi.CallbackQuery("^rm_.+$", callbackQuery.RmPodkastId)
	Tgapi.CallbackQuery("^0$", callbackQuery.UpdateKB)

	return Tgapi
}

func PodkastsProcessing() {
	defer func() {
		if r := recover(); r != nil {
			logTg.Error("recovered PodkastsProcessing: %v and sleep 1 minute\n", r)
			time.Sleep(1 * time.Minute)
			go PodkastsProcessing()
		}
	}()

	for {
		logTg.Info("start parse podcasts")
		listPodcasts := parser.GetNewPodcasts()
		logTg.Info("got %d podcasts", len(listPodcasts))
		for _, p := range listPodcasts {
			logTg.Info("parsing")
			if botdb.IsPodcastExistByHref(p.Href) {
				logTg.Info("podcast is exist with url:%s ", p.Href)
				continue
			}
			logTg.Info("try send parsed data ")
			if cs := SendPodcastToChannelAndSave(p); cs {
				logTg.Info("SendPodcastToChannelAndSave OK")
				ForwardToAllChatsPodcast(p)
			}
			time.Sleep(time.Second * 3)
		}
		logTg.Info("finish parse podcasts, sleep " + TimeDelayPodcastParse.String())
		time.Sleep(TimeDelayPodcastParse)
	}
}
