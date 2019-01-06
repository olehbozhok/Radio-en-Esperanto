package repository

import (
	"crypto/md5"

	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type mongoChannelRepository struct {
	Collection *mgo.Collection
}

// NewMongoChannelRepository mongo repository for channels
func NewMongoChannelRepository(collection *mgo.Collection) radiobot.ChannelRepository {
	return &mongoChannelRepository{Collection: collection}
}

// RegisterChannel is used to register channel
func (ch *mongoChannelRepository) RegisterChannel(radioCh *radiobot.Channel) error {
	radioCh.CalcID()
	return ch.Collection.Insert(radioCh)
}

// FindChannelByID is used to find channel by id
func (ch *mongoChannelRepository) FindChannelByID(id [md5.Size]byte) (*radiobot.Channel, error) {
	radioCh := new(radiobot.Channel)
	query := ch.Collection.Find(bson.M{"_id": id})
	err := query.One(radioCh)
	return radioCh, err
}

// FindChannelByName is used to find channel by name
func (ch *mongoChannelRepository) FindChannelByName(name string) (*radiobot.Channel, error) {
	radioCh := new(radiobot.Channel)
	query := ch.Collection.Find(bson.M{"name": name})
	err := query.One(radioCh)
	return radioCh, err
}

// GetChannels is used get channels
func (ch *mongoChannelRepository) GetChannels(count, offset int) ([]*radiobot.Channel, error) {
	radioChannels := make([]*radiobot.Channel, 0, count)
	err := ch.Collection.Find(bson.M{}).Skip(offset).Limit(count).All(&radioChannels)
	return radioChannels, err
}
