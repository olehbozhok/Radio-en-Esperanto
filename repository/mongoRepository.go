package repository

import (
	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
	"github.com/globalsign/mgo"
)

type mongoRepository struct {
	mongoPodcastRepository
	mongoChannelRepository
	mongoChatRepository
}

// NewMongoRepository mongo repository
func NewMongoRepository(podcastCollection, channelCollection, chatCollection *mgo.Collection) radiobot.Repository {
	return &mongoRepository{
		mongoChannelRepository: mongoChannelRepository{channelCollection},
		mongoChatRepository:    mongoChatRepository{chatCollection},
		mongoPodcastRepository: mongoPodcastRepository{podcastCollection},
	}
}
