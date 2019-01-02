package radiobot

import (
	"github.com/google/uuid"
)

// Repository отвечают за стандартизацию взаимодействия с хранилищем (БД + кэш + файловое хранилище + ...) подсистемы.
// Все IO операции с объектами должны проходить исключително через данный интерфейс
type Repository interface {
	PodcastRepository
	ChannelRepository
	ChatRepository
}

// PodcastRepository represent repository for podcasts
type PodcastRepository interface {
	AddPocast(Podcast) error
	IsNewPodcast(Podcast) (bool, error)

	// TODO: implement in future
	// GetUnsendedPodcasts(count, offset int) ([]Podcast, error)
	// FindAllPocastParsedFromTo(from, to time.Time, count, offset int) ([]Podcast, error)
	// FindPocastParsedFromToByChannelID(from, to time.Time, count, offset int, channelID uuid.UUID) ([]Podcast, error)
}

// ChannelRepository represent repository for channels
type ChannelRepository interface {
	RegisterChannel(*Channel) (uuid.UUID, error)
	FindChannelByID(uuid.UUID) (*Channel, error)
	FindChannelByName(string) (*Channel, error)
	GetChannels(count, offset int) ([]*Channel, error)
}

// ChatRepository represent repository for chats
type ChatRepository interface {
	RegisterChat(*Chat) error
	FindChat(id int64) (*Chat, error)
	SubscribeChat(*Chat, *Channel) error
	UnsubscribeChat(*Chat, *Channel) error

	GetAllChats(count, offset int) ([]*Chat, error)
	// GetAllChatsSubscribedOn(ch *Channel, count, offset int) ([]*Chat, error)
	GetAllChatsIDSubscribedOn(ch *Channel, count, offset int) ([]int64, error)

	// TODO
	// SetChatLastSendNow() error
}
