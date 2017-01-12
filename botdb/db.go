package botdb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//_ "github.com/mattn/go-sqlite3"
)
var DEBUG   bool = false
var connDB    *gorm.DB


func Initdb(args ...interface{}) (*gorm.DB)  {
	//err := os.Remove("db.sqlite3")
	//if err != nil {
	//	fmt.Println(err)
	//}

	db, err := gorm.Open("sqlite3", args...)
	if err != nil {
		panic(err)
	}
	connDB = db

	//db.LogMode(DEBUG)

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(PodcastChannelType{}, PodcastType{})

	return db
}