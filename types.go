package radiobot

import (
	"github.com/google/uuid"
	telebot "gopkg.in/tucnak/telebot.v2"
)

// Channel represent radio channel
type Channel struct {
	ID      uuid.UUID `json:"id" bson:"_id"`
	Name    string
	Comment string
}

// Chat telegram info
type Chat struct {
	telebot.Chat
	ID int64 `json:"id" bson:"_id"`

	SubscribedChannels []*Channel
}
