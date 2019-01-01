package radiobot

import "github.com/google/uuid"

// Repository отвечают за стандартизацию взаимодействия с хранилищем (БД + кэш + файловое хранилище + ...) подсистемы.
// Все IO операции с объектами должны проходить исключително через данный интерфейс
type Repository interface {
	ChannelRepository
	ChatRepository
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
}
