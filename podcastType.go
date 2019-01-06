package radiobot

import (
	"crypto/md5"
	"time"

	telebot "gopkg.in/tucnak/telebot.v2"

	"github.com/google/uuid"
)

// Podcast represent podcast of the channel
type Podcast struct {
	ID [md5.Size]byte `json:"id" bson:"_id"`

	FileURL   string    `json:"file_url" bson:"file_url"`
	ChannelID uuid.UUID `json:"channel_id" bson:"channel_id"`
	Comment   string    `json:"comment" bson:"comment"`
	CreatedOn time.Time `json:"created_on" bson:"created_on"`
	ParsedOn  time.Time `json:"parsed_on" bson:"parsed_on"`

	// telegram messages data to forward messages from channel to another chat
	Recipient    Recipient `json:"recipient" bson:"recipient"`
	CommentMsgID int       `json:"comment_msg_id" bson:"comment_msg_id"`
	FileMsgID    int       `json:"file_msg_id" bson:"file_msg_id"`
	FIleTgID     string    `json:"file_tg_id" bson:"file_tg_id"`
}

// CalcID calc and set ID of podcast based on the FileURL
func (p *Podcast) CalcID() {
	p.ID = md5.Sum([]byte(p.FileURL))
}

// SetParsedOnNow set field ParsedOn now
func (p *Podcast) SetParsedOnNow() {
	p.ParsedOn = time.Now()
}

// IsSended check if podcas is sended to channel
func (p *Podcast) IsSended() bool {
	return p.Recipient.Recipient() != "" && (p.CommentMsgID != 0 || p.FileMsgID != 0)
}

// SetRecipient set Recipient from telebot.Recipient interface
func (p *Podcast) SetRecipient(r telebot.Recipient) {
	p.Recipient = Recipient(r.Recipient())
}

// Recipient implement telebot.Recipient interface
type Recipient string

// Recipient return string of recipient
func (r *Recipient) Recipient() string {
	return string(*r)
}
