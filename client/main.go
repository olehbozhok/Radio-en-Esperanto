package main

import (
	"fmt"
	"log"
	"os"
	"time"

	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	tb "gopkg.in/tucnak/telebot.v2"

	tgbothandlers "github.com/Oleg-MBO/Radio-en-Esperanto/delivery/tgBotHandlers"
	"github.com/Oleg-MBO/Radio-en-Esperanto/parser"
	"github.com/Oleg-MBO/Radio-en-Esperanto/repository"
	"github.com/Oleg-MBO/Radio-en-Esperanto/usecases"

	"github.com/globalsign/mgo"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getEnvVar(name string) string {
	val := os.Getenv(name)
	if val == "" {
		checkErr(fmt.Errorf("error not specified env var %s", name))
	}
	return val
}

func main() {
	mondoConnStr := getEnvVar("MONGOCONNSTR")
	dbName := getEnvVar("MONGODBNAME")
	podcastCollectionName := getEnvVar("PODCASTCOLLECTIONMANE")
	channelsCollectionName := getEnvVar("CHANNELCOLLECTIONMANE")
	chatsCollectionName := getEnvVar("CHATCOLLECTIONMANE")

	tgChannelID := getEnvVar("TGCHANNELID")
	tgBotAPI := getEnvVar("TGBOTAPI")

	logFile := getEnvVar("LOGFILE")

	log.SetOutput(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    20, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		// Compress:   true, // disabled by default
		LocalTime: true,
	})

	session, err := mgo.Dial(mondoConnStr)
	checkErr(err)
	db := session.DB(dbName)
	repo := repository.NewMongoRepository(db.C(podcastCollectionName), db.C(channelsCollectionName), db.C(chatsCollectionName))

	bot, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL. If field is empty it equals to "https://api.telegram.org"
		// URL: "http://195.129.111.17:8012",
		Poller: &tb.LongPoller{Timeout: 20 * time.Second},
		Token:  tgBotAPI,
		Reporter: func(err error) {
			log.Printf("telegram recovered err: %+v", err)
		},
	})
	checkErr(err)

	usecase := usecases.NewUsecases(repo, tgChannelID, bot)

	tgbothandlers.RegisterHandlers(usecase)

	runPodcastParce := getPodcastParceAndSendFunc(usecase)

	var preparedPodcastParce func()
	preparedPodcastParce = func() {
		runPodcastParce()
		time.AfterFunc(time.Minute*30, preparedPodcastParce)
	}
	go preparedPodcastParce()

	bot.Start()

}

func getPodcastParceAndSendFunc(usecase radiobot.Usecase) func() {
	podcastParser := parser.NewEsperantoRadio()
	return func() {
		podcastsAndChannels, err := podcastParser.Parse()
		if err != nil {
			log.Println("err podcastParser.Parse: ", err)
			return
		}
		for _, pAndCh := range podcastsAndChannels {
			_, err = usecase.SaveOnlyNewPodcastAndChannel(pAndCh)
			if err != nil {
				log.Println("err usecase.SaveOnlyNewPodcastAndChannel: ", err)
				continue
			}
		}

		sendedPodcasts := 0
		for {
			podcasts, err := usecase.FindUnsendedPodcasts(100, 0)
			checkErr(err)
			if len(podcasts) == 0 {
				break
			}
			for _, p := range podcasts {

				err = usecase.SendToTgChannelAndUpdatePodcast(&p)
				if err != nil {
					log.Println("err usecase.SendAndUpdatePodcast: ", err)
				}
				sendedPodcasts++
				time.Sleep(3000 * time.Millisecond)
			}
		}
		log.Println("sended podcasts ", sendedPodcasts)
	}
}
