package radiobot

import (
	"github.com/google/uuid"
)

// Usecase отвечают за стандартизацию взаимодействия с подсистемой.
// Все взаимодействия с объектами должны проходить исключително через них.
// Создание и предобразование объектов происходит только в данных функциях.
// Прямые обращения в базу данных в обход usecase'ов запрещены
type Usecase interface {
	RegisterChannel(*Channel) (uuid.UUID, error)
	GetChannelByID(uuid.UUID) *Channel
	GetChannelByName(string) *Channel

	RegisterChat(*Chat) error
	FindChat(id int64) (*Chat, error)
	SubscribeChat(*Chat, *Channel) error
	UnsubscribeChat(*Chat, *Channel) error

	GetAllChats(count, offset int) []*Chat
	GetAllChatsSubscribedOn(ch *Channel, count, offset int) []*Chat

	SendMessageTo(*Chat, interface{})
}
