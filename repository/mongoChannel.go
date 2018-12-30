package repository

import (
	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
	"github.com/globalsign/mgo"
	"github.com/google/uuid"
)

type mongoChannelRepository struct {
	Collection *mgo.Collection
}

// NewMongoChannelRepository mongo repository for channels
func NewMongoChannelRepository(collection *mgo.Collection) radiobot.ChannelRepository {
	return &mongoChannelRepository{Collection: collection}
}

// RegisterChannel is used to register channel
func (ch *mongoChannelRepository) RegisterChannel(*radiobot.Channel) (uuid.UUID, error) {
	panic("implement me")
}

// FindChannelByID is used to find channel by id
func (ch *mongoChannelRepository) FindChannelByID(uuid.UUID) (*radiobot.Channel, error) {
	panic("implement me")
}

// FindChannelByName is used to find channel by name
func (ch *mongoChannelRepository) FindChannelByName(string) (*radiobot.Channel, error) {
	panic("implement me")
}
