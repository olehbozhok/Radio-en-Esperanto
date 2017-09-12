package main

import (
	"golang.org/x/net/context"

	"github.com/azer/logger"
	"os"

	//"github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/models"
	//"gopkg.in/natefinch/lumberjack.v2"
	//"github.com/Oleg-MBO/Radio-en-Esperanto/botdb"
	"github.com/Oleg-MBO/Radio-en-Esperanto/db"
	"github.com/Oleg-MBO/Radio-en-Esperanto/parser"
	"github.com/Oleg-MBO/Radio-en-Esperanto/parser/parserERadio"
	"github.com/Oleg-MBO/Radio-en-Esperanto/tgApiHandlers"
	"github.com/olebedev/go-tgbot"
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

	debugStr := os.Getenv("DEBUG")
	var debug = false
	if debugStr != "" {
		debug = true
	}

	db, err := botdb.InitDB(os.Getenv("dbString"), debug)
	if err != nil {
		logMain.Error(err.Error())
		os.Exit(2)
	}

	defer db.Close()

	token := os.Getenv("tgBotToken")
	channelId := os.Getenv("channelId")

	r := tgApiHandlers.InitApi(token, channelId)

	parser.RegisterParser(parser.CallbackGetpodcasts{"ParserERadio", parserERadio.GetLastPodcasts})
	go tgApiHandlers.PodkastsProcessing()

	ctx := context.Background()

	//setup global middleware
	r.Use(tgbot.Recover)

	//Bind handler
	i := 0
	for {
		i++
		if err := r.Poll(ctx, []models.AllowedUpdate{
			models.AllowedUpdateMessage,
			models.AllowedUpdateCallbackQuery,
		}); err != nil {
			logMain.Error("Poll Error: %s", err.Error())
		}
		if i > 4 {
			break
		}
	}
}
