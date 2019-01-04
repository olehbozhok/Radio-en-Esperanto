package radiobot

import (
	"github.com/google/uuid"
	telebot "gopkg.in/tucnak/telebot.v2"
)

// Usecase отвечают за стандартизацию взаимодействия с подсистемой.
// Все взаимодействия с объектами должны проходить исключително через них.
// Создание и предобразование объектов происходит только в данных функциях.
// Прямые обращения в базу данных в обход usecase'ов запрещены
type Usecase interface {
	FindChannelByID(uuid.UUID) (*Channel, error)
	FindChannelByName(string) (*Channel, error)

	// Зберегти новий канал
	RegisterORFindChannel(*Channel) error

	// save podcast and channel that return parser (not sended to tg channel yet)
	SaveOnlyNewPodcastAndChannel(PodcastAndChannel) (isPodcastNew bool, err error)

	FindUnsendedPodcasts(count, offset int) ([]Podcast, error)
	// send podcast to tg channel and update podcast in db
	SendAndUpdatePodcast(*Podcast) error

	// send message to one user
	Send(to telebot.Recipient, what interface{}, options ...interface{}) (*telebot.Message, error)

	// TODO: Implement in future
	// SendPodcastToSubscribers(Podcast) error

	FindOrRegisterChat(*Chat) error
	SubscribeChat(*Chat, *Channel) error
	UnsubscribeChat(*Chat, *Channel) error

	GetAllChats(count, offset int) ([]*Chat, error)
	GetAllChatsIDSubscribedOn(ch *Channel, count, offset int) ([]string, error)

	// // add Recipient interface
	// SendMessageTo(*Chat, interface{})

	HandleTg(endpoint interface{}, handler interface{})
}
