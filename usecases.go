package radiobot

import (
	"crypto/md5"

	telebot "gopkg.in/tucnak/telebot.v2"
)

// Usecase отвечают за стандартизацию взаимодействия с подсистемой.
// Все взаимодействия с объектами должны проходить исключително через них.
// Создание и предобразование объектов происходит только в данных функциях.
// Прямые обращения в базу данных в обход usecase'ов запрещены
type Usecase interface {
	FindChannelByID([md5.Size]byte) (*Channel, error)
	FindChannelByName(string) (*Channel, error)

	// Зберегти новий канал
	RegisterORFindChannel(*Channel) error

	GetChannels(count, offset int) ([]*Channel, error)

	// save podcast and channel that return parser (not sended to tg channel yet)
	SaveOnlyNewPodcastAndChannel(PodcastAndChannel) (isPodcastNew bool, err error)

	FindUnsendedPodcasts(count, offset int) ([]Podcast, error)
	// send podcast to tg channel and update podcast in db
	SendToTgChannelAndUpdatePodcast(*Podcast) error

	// send message to one user
	SendTgMessage(to telebot.Recipient, what interface{}, options ...interface{}) (*telebot.Message, error)
	// EditTgMessage is magic, it lets you change already sent message.
	EditTgMessage(message telebot.Editable, what interface{}, options ...interface{}) (*telebot.Message, error)
	// ForwardTgMessage behaves just like Send() but of all options it only supports Silent (see Bots API).
	ForwardTgMessage(to telebot.Recipient, what *telebot.Message, options ...interface{}) (*telebot.Message, error)
	// Respond is used to send callback responce from inline keyboard
	Respond(callback *telebot.Callback, responseOptional ...*telebot.CallbackResponse) error

	// SendPodcastToSubscribers send podcast to all chats which is subscribed on podcast channel
	SendPodcastToSubscribers(Podcast) error

	FindOrRegisterChat(*Chat) error
	SubscribeChat(*Chat, *Channel) error
	UnsubscribeChat(*Chat, *Channel) error

	GetAllChats(count, offset int) ([]*Chat, error)
	GetAllChatsIDSubscribedOn(ch *Channel, count, offset int) ([]string, error)

	// // add Recipient interface
	// SendMessageTo(*Chat, interface{})

	HandleTg(endpoint interface{}, handler interface{})
}
