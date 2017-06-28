package main

import (
	"github.com/Oleg-MBO/Radio-en-Esperanto/botdb"
	"github.com/Oleg-MBO/Radio-en-Esperanto/plagins/parserERadio"
	_ "github.com/Oleg-MBO/Radio-en-Esperanto/plagins/parserERadio"
	"github.com/Oleg-MBO/Radio-en-Esperanto/tgApiHelper"

	"golang.org/x/net/context"

	"github.com/azer/logger"
	"os"

	//"github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/models"
	//"gopkg.in/natefinch/lumberjack.v2"
)

var logMain = logger.New("main")

func main() {

	//logger.SetOutput(&lumberjack.Logger{
	//	//Filename:   "/home/pi/tgBotRadioEnEsperanto/log/logfile.log",
	//	Filename:   "log/logfile.log",
	//	MaxSize:    1, // megabytes
	//	MaxBackups: 20,
	//	MaxAge:     28, //days
	//})

	logMain.Info("START PROGRAM")
	logger.SetOutput(os.Stdout)

	db := botdb.Initdb(os.Getenv("dbString"))


	defer db.Close()
	//db.LogMode(false)



	token := os.Getenv("tgBotToken")


	channelId := os.Getenv("channelId")
	r := tgApiHelper.InitApi(token, channelId)

	botdb.RegisterParser(botdb.CallbackGetpodcasts{"ParserERadio", parserERadio.GetLastPodcasts})
	go tgApiHelper.PodkastsProcessing()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//setup global middleware
	//r.Use(tgbot.Recover)

	//Bind handler

	if err := r.Poll(ctx, []models.AllowedUpdate{
		models.AllowedUpdateMessage,
		models.AllowedUpdateCallbackQuery,
	}); err != nil {
		logMain.Error("Poll Error: %s", err.Error())
	}
}
