package repository

import (
	"crypto/md5"

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

// UpdateChat is used for register chats
func (rchat *mongoChatRepository) UpdateChat(chat *radiobot.Chat) error {
	return rchat.Collection.UpdateId(chat.ID, chat)
}

// FindChat is used to find chat by id
func (rchat *mongoChatRepository) FindChat(id string) (*radiobot.Chat, error) {
	chat := new(radiobot.Chat)
	query := rchat.Collection.Find(bson.M{"_id": id})
	err := query.One(chat)
	return chat, err
}

// SubscribeChat is used to subscribe chat on channel
func (rchat *mongoChatRepository) SubscribeChat(chat *radiobot.Chat, channel *radiobot.Channel) error {
	chat, err := rchat.FindChat(chat.ID)
	if err != nil {
		return err
	}
	for _, ch := range chat.SubscribedChannelsID {
		if ch == channel.ID {
			return ErrChatAlreadySubskribed
		}
	}
	chat.SubscribedChannelsID = append(chat.SubscribedChannelsID, channel.ID)
	return rchat.UpdateChat(chat)
}

// UnsubscribeChat is used to unsubscribe chat on channel
func (rchat *mongoChatRepository) UnsubscribeChat(chat *radiobot.Chat, channel *radiobot.Channel) error {
	chat, err := rchat.FindChat(chat.ID)
	if err != nil {
		return err
	}
	for i, ch := range chat.SubscribedChannelsID {
		if ch == channel.ID {
			chat.SubscribedChannelsID = append(chat.SubscribedChannelsID[:i], chat.SubscribedChannelsID[i+1:]...)
		}
	}
	return rchat.UpdateChat(chat)
}

// GetAllChats is used to get all chats with count and offset
func (rchat *mongoChatRepository) GetAllChats(count, offset int) ([]*radiobot.Chat, error) {
	chats := make([]*radiobot.Chat, count)
	err := rchat.Collection.Find(bson.M{}).Skip(offset).Limit(count).All(&chats)
	return chats, err
}

// GetAllChatsSubscribedOn is used to fetch all chats which subscribed on channel
func (rchat *mongoChatRepository) GetAllChatsSubscribedOn(ch *radiobot.Channel, count, offset int) ([]*radiobot.Chat, error) {
	chats := make([]*radiobot.Chat, count)
	err := rchat.Collection.Find(bson.M{"subscribed_channels_id": bson.M{"$in": [][md5.Size]byte{ch.ID}}}).Skip(offset).Limit(count).All(&chats)
	return chats, err
}

// GetAllChatsIDSubscribedOn is used to fetch all chats ID which subscribed on channel
func (rchat *mongoChatRepository) GetAllChatsIDSubscribedOn(ch *radiobot.Channel, count, offset int) ([]string, error) {
	// unefficient way
	// TODO: make more efficient
	chats, err := rchat.GetAllChatsSubscribedOn(ch, count, offset)
	if err != nil {
		return nil, err
	}
	idArr := make([]string, 0, len(chats))
	for _, chat := range chats {
		idArr = append(idArr, chat.ID)
	}
	return idArr, nil
}
