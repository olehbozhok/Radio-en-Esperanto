package repository

import (
	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
	"github.com/globalsign/mgo"
)

type mongoRepository struct {
	mongoChatRepository
	mongoChannelRepository
}

// NewMongoRepository mongo repository
func NewMongoRepository(chatCollection, channelCollection *mgo.Collection) radiobot.Repository {
	return &mongoRepository{
		mongoChannelRepository: mongoChannelRepository{channelCollection},
		mongoChatRepository:    mongoChatRepository{chatCollection},
	}
}
