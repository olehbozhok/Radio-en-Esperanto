package radiobot

import (
	"strconv"

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
	ID int64 `json:"id" bson:"_id"`

	// See telebot.ChatType and consts.
	Type telebot.ChatType `json:"type"`

	// Won't be there for ChatPrivate.
	Title string `json:"title"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`

	SubscribedChannels []*Channel
}

// Recipient returns chat ID (see Recipient interface).
func (c *Chat) Recipient() string {
	if c.Type == telebot.ChatChannel {
		return "@" + c.Username
	}
	return strconv.FormatInt(c.ID, 10)
}
