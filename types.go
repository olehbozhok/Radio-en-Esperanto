package radiobot

import (
	"crypto/md5"
	"strconv"

	telebot "gopkg.in/tucnak/telebot.v2"
)

// Channel represent radio channel
type Channel struct {
	ID      [md5.Size]byte `json:"id" bson:"_id"`
	Name    string         `json:"name" bson:"name"`
	Comment string         `json:"comment" bson:"comment"`
	Parser  string         `json:"parser" bson:"parser"`
}

// CalcID calc and set ID of podcast based on the FileURL
func (c *Channel) CalcID() {
	c.ID = md5.Sum([]byte(c.Parser + "-" + c.Name))
}

// PodcastAndChannel is used for parsing result
type PodcastAndChannel struct {
	Podcast *Podcast
	Channel *Channel
}

// Chat telegram info
type Chat struct {
	ID     string `json:"id" bson:"_id"`
	ChatID int64  `json:"chat_id" bson:"chat_id"`

	// See telebot.ChatType and consts.
	Type telebot.ChatType `json:"type" bson:"type"`

	// Won't be there for ChatPrivate.
	Title string `json:"title" bson:"title"`

	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Username  string `json:"username" bson:"username"`

	SubscribedChannelsID [][md5.Size]byte `json:"subscribed_channels_id" bson:"subscribed_channels_id"`
}

// Recipient returns chat ID (see Recipient interface).
func (c *Chat) Recipient() string {
	if c.Type == telebot.ChatChannel {
		return "@" + c.Username
	}
	return strconv.FormatInt(c.ChatID, 10)
}

// SetID calc and set ID from Recipient
func (c *Chat) SetID() {
	c.ID = c.Recipient()
}

// FromTgChat fill fields from tg.Chat
func (c *Chat) FromTgChat(ch *telebot.Chat) {
	c.ChatID = ch.ID
	c.Type = ch.Type
	c.Title = ch.Title
	c.FirstName = ch.FirstName
	c.Username = ch.Username
	c.ID = c.Recipient()
}
