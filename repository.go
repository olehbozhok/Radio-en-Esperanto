package radiobot

import "github.com/google/uuid"

// Repository отвечают за стандартизацию взаимодействия с хранилищем (БД + кэш + файловое хранилище + ...) подсистемы.
// Все IO операции с объектами должны проходить исключително через данный интерфейс
type Repository interface {
	RegisterChannel(*Channel) (uuid.UUID, error)
	GetChannelByID(uuid.UUID) *Channel
	GetChannelByName(string) *Channel

	RegisterChat(*Chat) error
	FindChat(id int64) (*Chat, error)
	SubscribeChat(*Chat, *Channel) error
	UnsubscribeChat(*Chat, *Channel) error

	GetAllChats(count, offset int) []*Chat
	GetAllChatsSubscribedOn(ch *Channel, count, offset int) []*Chat
}
