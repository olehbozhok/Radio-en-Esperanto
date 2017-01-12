package botdb

import (
	"time"

	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type PodcastChannelType struct {
	//gorm.Model
	ChannelName string `sql:"unique"`
	//Podcasts    []PodcastType       //`gorm:"many2many:podkastsChannels_podkasts;"`
}

func (PodcastChannelType) TableName() string {
	return "podcastChannels"
}

type PodcastType struct {
	gorm.Model
	Date        time.Time `sql:"type:DATETIME"`
	ChannelName string    `sql:"not null"` //`gorm:"FOREIGN KEY(ChannelName) REFERENCES podcastChannels(channelname),"`
	//Channel               PodcastChannelType   `gorm:"ForeignKey:ChannelName"` //`gorm:"ForeignKey:channelname"`  //`sql:"NOT NULL",gorm:""`
	Description           string `sql:"not null"`
	Href                  string `sql:"unique"`
	Tg_msg_id_description string
	Tg_msg_id_file        string
	Tg_file_id            string
	Tg_from_chat_id       string
}

func (PodcastType) TableName() string {
	return "podkasts"
}

func (p *PodcastType) BeforeSave() (err error) {

	if p.ChannelName == "" || p.Description == "" || p.Href == "" {
		err = errors.New("one of fills is empty")
		return
	}
	var count = 0
	connDB.Select("nil").Model(PodcastType{}).Where("href = ?", p.Href).Count(&count)
	fmt.Println("count: ", count)
	if count > 0 {
		err = errors.New("podkast is exist")
		return
	}

	//connDB.Create(&PodcastChannelType{ChannelName:p.ChannelName})
	return
}
