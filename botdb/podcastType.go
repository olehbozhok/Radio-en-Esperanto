package botdb

import (
	"fmt"
	"strings"
	"time"

)

type PodcastType struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	//UpdatedAt time.Time
	//Date               time.Time
	RawDate            string  /// yyyy-MM-dd
	ChannelName        string  `sql:"not null;index"` //`gorm:"FOREIGN KEY(ChannelName) REFERENCES podcastChannels(channelname),"`
	Description        string  `sql:"blob"`//;gorm:"blob"`
	Href               string  `sql:"unique;index"`
	TgMsgIdDescription *int64  `sql:"not null"`
	TgMsgFileId        *int64
	TgFromIdChat       *int64   `sql:"not null"`
}

func (PodcastType) TableName() string {
	return "podkasts"
}

func (PodcastType) CreateTable() error {
	return connDB.Exec(`
	CREATE TABLE IF NOT EXISTS podkasts
	(
	id int unsigned AUTO_INCREMENT,
	created_at timestamp    NULL,
	updated_at timestamp    NULL,
	raw_date varchar(15)    NOT NULL,
	channel_name varchar(30) NOT NULL,
	description TEXT        NOT NULL,
	href varchar(190) UNIQUE,
	tg_msg_id_description bigint NOT NULL,
	tg_msg_file_id bigint,
	tg_from_id_chat bigint NOT NULL ,
	PRIMARY KEY (id),
	INDEX (href)
	) CHARACTER SET=utf8mb4
	`).Error

}

func (p *PodcastType) BeforeSave() error {

	if p.ChannelName == "" || p.Description == "" || p.Href == "" {
		// || p.TgMsgIdDescription == "" || p.TgMsgIdFile == "" || p.TgFileId == "" || p.TgFromIdChat == "" {
		return ErrFieldsEmpty
	}

	if !cachedPodcasts[p.ChannelName] {
		connDB.Create(&PodcastChannelType{ChannelName: p.ChannelName})
	}
	logDB.Info("!!!!!!!!!!!! SAVE PODKSAST")
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

func (p *PodcastType) Delete() error {
	//return connDB.Model(&p).Where("href = ?", p.Href).Error
	return connDB.Exec("DELETE FROM podkasts WHERE href=?", p.Href).Error
	//return connDB.Delete(&p).Error
}

func (p *PodcastType) Name() string {
	return fmt.Sprintf("%s %s.mp3", p.ChannelName, p.RawDate)
}

func (p *PodcastType) Caption() string {
	return fmt.Sprintf("#%s %s", strings.Replace(p.ChannelName, " ", "_", -1), p.RawDate)
}

func (p *PodcastType) TextToChannel(withFileUrl bool) *string {
	var message string
	message = fmt.Sprintf("%s\n#%s\n%s",
		p.RawDate,
		strings.Replace(p.ChannelName, " ", "_", -1),
		p.Description,
	)
	if withFileUrl {
		message = fmt.Sprintf("%s\n\n%s", message, p.Href)
	}
	return &message
}

func (p *PodcastType) Subscrbers() (cpodcasts []ChatPodcasts) {
	connDB.Where("channel = ?", p.ChannelName).Find(&cpodcasts)
	return cpodcasts
}
