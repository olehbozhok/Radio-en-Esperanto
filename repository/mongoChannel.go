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
func (ch *mongoChannelRepository) RegisterChannel(radioCh *radiobot.Channel) (uuid.UUID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return uuid.UUID{}, err
	}
	radioCh.ID = id

	err = ch.Collection.Insert(radioCh)
	return id, err
}

// FindChannelByID is used to find channel by id
func (ch *mongoChannelRepository) FindChannelByID(uuid.UUID) (*radiobot.Channel, error) {
	panic("implement me")
}

// FindChannelByName is used to find channel by name
func (ch *mongoChannelRepository) FindChannelByName(string) (*radiobot.Channel, error) {
	panic("implement me")
}

// GetChannels is used get channels
func (ch *mongoChannelRepository) GetChannels(count, offset int, notIn []*radiobot.Channel) ([]*radiobot.Channel, error) {
	panic("implement me")
}
