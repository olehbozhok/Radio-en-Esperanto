package radiobot

import (
	"crypto/md5"
	"strconv"
	"time"

	"github.com/google/uuid"
	telebot "gopkg.in/tucnak/telebot.v2"
)

// Podcast represent podcast of the channel
type Podcast struct {
	ID [md5.Size]byte `json:"id" bson:"_id"`

	FileURL   string    `json:"file_url" bson:"file_url"`
	ChannelID uuid.UUID `json:"channel_id" bson:"channel_id"`
	Comment   string    `json:"comment" bson:"comment"`
	CreatedOn time.Time `json:"created_on" bson:"created_on"`
	ParsedOn  time.Time `json:"parsed_on" bson:"parsed_on"`
}

// CalcID calc and set ID of podcast based on the FileURL
func (p *Podcast) CalcID() {
	p.ID = md5.Sum([]byte(p.FileURL))
}

// SetParsedOnNow set field ParsedOn now
func (p *Podcast) SetParsedOnNow() {
	p.ParsedOn = time.Now()
}

// Channel represent radio channel
type Channel struct {
	ID      uuid.UUID `json:"id" bson:"_id"`
	Name    string    `json:"name" bson:"name"`
	Comment string    `json:"comment" bson:"comment"`
}

// Chat telegram info
type Chat struct {
	ID int64 `json:"id" bson:"_id"`

	// See telebot.ChatType and consts.
	Type telebot.ChatType `json:"type" bson:"type"`

	// Won't be there for ChatPrivate.
	Title string `json:"title" bson:"title"`

	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Username  string `json:"username" bson:"username"`

	SubscribedChannelsID []uuid.UUID `json:"subscribed_channels_id" bson:"subscribed_channels_id"`
}

// Recipient returns chat ID (see Recipient interface).
func (c *Chat) Recipient() string {
	if c.Type == telebot.ChatChannel {
		return "@" + c.Username
	}
	return strconv.FormatInt(c.ID, 10)
}
