package botdb

var cachedPodcasts map[string]bool

func init() {
	cachedPodcasts = make(map[string]bool)
}

type PodcastChannelType struct {
	ChannelName string `sql:"primary_key"`
}

func (PodcastChannelType) TableName() string {
	return "podcastChannels"
}

func (PodcastChannelType) CreateTable() error {
	return connDB.Exec(`
	CREATE TABLE IF NOT EXISTS podcastChannels
(
channel_name varchar(30) ,
PRIMARY KEY (channel_name)
) CHARACTER SET=utf8mb4
	`).Error

}

func (c *PodcastChannelType) BeforeSave() error {
	if !cachedPodcasts[c.ChannelName] {
		cachedPodcasts[c.ChannelName] = true
		return nil
	}
	return ErrIsExist
}

func (c PodcastChannelType) IsExist() bool {
	var count = 0
	connDB.Model(&PodcastChannelType{}).Where("channel_name = ?", c.ChannelName).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

type ChatPodcasts struct {
	Chat    int64
	Channel string
}

func (ChatPodcasts) TableName() string {
	return "PodcastToChat"
}

func (ChatPodcasts) CreateTable() error {
	return connDB.Exec(`
	CREATE TABLE IF NOT EXISTS PodcastToChat
(
chat BIGINT NOT NULL ,
channel varchar(30) NOT NULL,
 PRIMARY KEY(chat, channel),
CONSTRAINT constr_PodcastToChat UNIQUE (chat, channel),
INDEX (channel)
) CHARACTER SET=utf8mb4
	`).Error

}

func (c *ChatPodcasts) BeforeSave() error {
	//var podcast PodcastChannelType
	//connDB.Where("channel_name = ?", c.Channel).First(&podcast)
	//if podcast.ChannelName == "" {
	//	return ErrNotCorrectId
	//}
	return nil
}

func (cp ChatPodcasts) Save() error {
	return connDB.Create(&cp).Error
}

func (cp ChatPodcasts) Delete() error {
	return connDB.Where("chat = ? AND channel = ?", cp.Chat, cp.Channel).Delete(&cp).Error
}

type ChatAddRmMessagePodcasts struct {
	ChatID   int64  `sql:"not null;primary_key"`
	AddMsgID *int64 //`sql:"DEFAULT:null"`
	RmMsgID  *int64 //`sql:"DEFAULT:null"`
}

func (ChatAddRmMessagePodcasts) TableName() string {
	return "chat_add_rm_message_podcasts"
}

func (ChatAddRmMessagePodcasts) CreateTable() error {
	return connDB.Exec(`
	CREATE TABLE IF NOT EXISTS chat_add_rm_message_podcasts
	(
	chat_id bigint AUTO_INCREMENT NOT NULL,
	add_msg_id bigint,
	rm_msg_id bigint ,
	PRIMARY KEY (chat_id)

	) CHARACTER SET=utf8mb4
	`).Error

}

func (cadd ChatAddRmMessagePodcasts) UpdateAddMsgID() error {
	q := ChatAddRmMessagePodcasts{}
	if connDB.First(&q, cadd.ChatID).RecordNotFound() {
		return connDB.Save(&cadd).Error
	}
	return connDB.Model(&cadd).Omit("rm_msg_id").Update(map[string]interface{}{"add_msg_id": cadd.AddMsgID, "chat_id": cadd.ChatID}).Error

	//if connDB.Model(&cadd).Omit("rm_msg_id").Update(map[string]interface{}{"add_msg_id": cadd.AddMsgID, "chat_id": cadd.ChatID}) {
	//	logDB.Error("UpdateRmMsgID RecordNotFound == true")
	//	return connDB.Save(&cadd).Error
	//}
	//return nil
}

func (cadd ChatAddRmMessagePodcasts) UpdateRmMsgID() error {
	q := ChatAddRmMessagePodcasts{}
	if connDB.First(&q, cadd.ChatID).RecordNotFound() {
		return connDB.Save(&cadd).Error
	}
	return connDB.Model(&cadd).Omit("add_msg_id").Update(map[string]interface{}{"rm_msg_id": cadd.RmMsgID, "chat_id": cadd.ChatID}).Error

	//
	//if  {
	//	logDB.Error("UpdateRmMsgID RecordNotFound == true")
	//	return connDB.Save(&cadd).Error
	//}
	//return nil
}

func (cadd ChatAddRmMessagePodcasts) Find() ChatAddRmMessagePodcasts {
	connDB.Where("chat_id = ?", cadd.ChatID).First(&cadd)
	return cadd
}
