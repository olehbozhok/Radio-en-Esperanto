package repository

import (
	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type mongoChatRepository struct {
	Collection *mgo.Collection
}

// NewMongoChatRepository mongo repository for chats
func NewMongoChatRepository(collection *mgo.Collection) radiobot.ChatRepository {
	return &mongoChatRepository{Collection: collection}
}

// RegisterChat is used for register chats
func (rchat *mongoChatRepository) RegisterChat(chat *radiobot.Chat) error {
	return rchat.Collection.Insert(chat)
}

// FindChat is used to find chat by id
func (rchat *mongoChatRepository) FindChat(id int64) (*radiobot.Chat, error) {
	chat := new(radiobot.Chat)
	query := rchat.Collection.Find(bson.M{"_id": id})
	err := query.One(chat)
	return chat, err
}

// SubscribeChat is used to subscribe chat on channel
func (rchat *mongoChatRepository) SubscribeChat(chat *radiobot.Chat, channel *radiobot.Channel) error {
	panic("implement me")
}

// UnsubscribeChat is used to unsubscribe chat on channel
func (rchat *mongoChatRepository) UnsubscribeChat(*radiobot.Chat, *radiobot.Channel) error {
	panic("implement me")
}

// GetAllChats is used to get all chats with count and offset
func (rchat *mongoChatRepository) GetAllChats(count, offset int) ([]*radiobot.Chat, error) {
	panic("implement me")
}

// GetAllChatsSubscribedOn is used to fetch all chats which subscribed on channel
func (rchat *mongoChatRepository) GetAllChatsSubscribedOn(ch *radiobot.Channel, count, offset int) ([]*radiobot.Chat, error) {
	panic("implement me")
}

// GetAllChatsIDSubscribedOn is used to fetch all chats ID which subscribed on channel
func (rchat *mongoChatRepository) GetAllChatsIDSubscribedOn(ch *radiobot.Channel, count, offset int) ([]int64, error) {
	panic("implement me")
}
