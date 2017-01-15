package botdb

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var cachedPodcasts map[string]bool

func init() {
	cachedPodcasts = make(map[string]bool)
}

type PodcastChannelType struct {
	//gorm.Model
	ChannelName string `sql:"unique"`
	//Podcasts    []PodcastType       //`gorm:"many2many:podkastsChannels_podkasts;"`
}

func (PodcastChannelType) TableName() string {
	return "podcastChannels"
}

func (c *PodcastChannelType) BeforeSave() error {
	if !cachedPodcasts[c.ChannelName] {
		cachedPodcasts[c.ChannelName] = true
		return nil
	}
	return ErrIsExist
}

type PodcastType struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt *time.Time `sql:"index"`
	//gorm.Model
	Date        time.Time `sql:"type:DATETIME"`
	RawDate     string          /// yyyy-MM-dd
	ChannelName string    `sql:"not null;index"` //`gorm:"FOREIGN KEY(ChannelName) REFERENCES podcastChannels(channelname),"`
	//Channel               PodcastChannelType   `gorm:"ForeignKey:ChannelName"` //`gorm:"ForeignKey:channelname"`  //`sql:"NOT NULL",gorm:""`
	Description        string `sql:"not null"`
	Href               string `sql:"unique;index"`
	TgMsgIdDescription string `sql:"not null"`
	TgMsgIdFile        string `sql:"not null"`
	TgFileId           string `sql:"not null"`
	TgFromIdChat       string `sql:"not null"`
}

func (PodcastType) TableName() string {
	return "podkasts"
}

func (p *PodcastType) BeforeSave() error {

	if p.ChannelName == "" || p.Description == "" || p.Href == "" { // || p.TgMsgIdDescription == "" || p.TgMsgIdFile == "" || p.TgFileId == "" || p.TgFromIdChat == "" {
		return ErrFieldsEmpty
	}
	//var count = 0
	//connDB.Model(&PodcastType{}).Where("href = ?", p.Href).Count(&count)
	//if count > 0 {
	//	return ErrIsExist
	//}

	//if !cachedPodcasts[p.ChannelName] {
	connDB.Create(&PodcastChannelType{ChannelName: p.ChannelName})
	//	cachedPodcasts[p.ChannelName] = true
	//}
	return nil
}

func (p *PodcastType) IsUnique() bool {
	var count = 0
	connDB.Model(&PodcastType{}).Where("href = ?", p.Href).Count(&count)
	if count > 0 {
		return false
	}
	return true
}

func (p *PodcastType) Save() error {
	return connDB.Create(&p).Error
}

func (p *PodcastType) Name() string {
	return fmt.Sprintf("%s-%s", p.ChannelName, p.RawDate)
}

func (p *PodcastType) CanSendTotelegram() CanSendToTelegram {
	cs := CanSendToTelegram{true, true}

	resp, err := http.Head(p.Href)
	if err != nil {
		log.Fatalln(err)
		cs.Status = false
		return cs
	}
	log.Printf("status:%s ; ContentLength:%s \n%#v \n\n\n", resp.Status, resp.ContentLength, resp.Header)

	if resp.Status != "200 OK" {
		log.Printf("http:%s", p.Href)
		cs.Status = false
		return cs
	}
	if resp.ContentLength > 20000000 {
		log.Printf("len:%d", resp.ContentLength)
		cs.LengthIsOk = false
		return cs
	}
	return cs
}

// reader to download file
func (p *PodcastType) InputFile() (*InputFile, error) {
	log.Printf("InputFile\n")

	IF := &InputFile{FileName:p.Name()}

	resp, err := http.Get(p.Href)
	if err != nil {
		log.Fatalln(err)
		return IF, err
	}
	IF.IOReader = resp.Body
	//resp.Body.Close()
	IF.OK = true
	//defer resp.Body.Close()
	//log.Printf("status:%s ; ContentLength:%s \n%#v \n\n\n", resp.Status, resp.ContentLength, resp.Header)

	//if resp.Status != "200 OK" {
	//	IF.OK = false
	//	log.Printf("\nstatus:%s ;http:%s", resp.Status, p.Href)
	//}
	log.Printf("\nstatus:%s ;http:%s", resp.Status, p.Href)

	//log.Printf("\n\n\n")

	return IF, nil
}

// represent InputFile interface in bot-api/telegram
type InputFile struct {
	FileName string
	IOReader io.Reader
	OK       bool
}

func (i *InputFile) Name() string {
	return i.FileName
}

func (i *InputFile) Reader() io.Reader {
	return i.IOReader
}

type CanSendToTelegram struct {
	Status     bool
	LengthIsOk bool
}