package botdb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"

	//_ "github.com/mattn/go-sqlite3"
	"github.com/azer/logger"
)

var DEBUG bool = true
var connDB *gorm.DB

var logDB = logger.New("botDB")

func Initdb(args ...interface{}) *gorm.DB {
	//err := os.Remove("db.sqlite3")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//db, err := gorm.Open("sqlite3", args...)
	db, err := gorm.Open("mysql", args...)
	if err != nil {
		logDB.Error("Error connect to db")
		panic(err)
	}
	db.LogMode(DEBUG)

	connDB = db

	err = connDB.DB().Ping()
	if err != nil {
		panic(err)
	}
	connDB.Exec("set names utf8mb4")
	//db.AutoMigrate(  ChatAddRmMessagePodcasts{})
	PodcastType{}.CreateTable()
	ChatPodcasts{}.CreateTable()
	PodcastChannelType{}.CreateTable()
	ChatAddRmMessagePodcasts{}.CreateTable()
	return db
}


