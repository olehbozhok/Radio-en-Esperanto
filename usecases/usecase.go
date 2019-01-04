package usecases

import (
	"bytes"
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
	"github.com/globalsign/mgo"
	"github.com/google/uuid"
	telebot "gopkg.in/tucnak/telebot.v2"
)

// 50 MB
const maxAudioLengthTg = 50000000

type usecases struct {
	repo  radiobot.Repository
	tgBot *telebot.Bot

	tgChannel telebot.Recipient
}

// NewUsecases create usecases
func NewUsecases(repo radiobot.Repository, tgChannelID string, bot *telebot.Bot) radiobot.Usecase {
	recipient := radiobot.Recipient(tgChannelID)
	return &usecases{repo: repo,
		tgChannel: &recipient,
		tgBot:     bot,
	}
}

func (u *usecases) FindChannelByID(id uuid.UUID) (*radiobot.Channel, error) {
	return u.repo.FindChannelByID(id)
}

func (u *usecases) FindChannelByName(name string) (*radiobot.Channel, error) {
	return u.repo.FindChannelByName(name)
}

func (u *usecases) RegisterORFindChannel(ch *radiobot.Channel) error {
	chOld, err := u.repo.FindChannelByName(ch.Name)
	if err != nil && err != mgo.ErrNotFound {
		return err
	}
	if err == mgo.ErrNotFound {
		u.repo.RegisterChannel(ch)
	}
	ch = chOld
	return err
}

// SaveOnlyNewPodcast save podcast (not sended to tg channel)
// return true if podcast is new
func (u *usecases) SaveOnlyNewPodcast(p radiobot.Podcast) (bool, error) {
	isNew, err := u.repo.IsNewPodcast(p)
	if err != nil {
		return isNew, err
	}
	if isNew {
		err = u.repo.AddPocast(p)
	}
	return isNew, err
}

// SaveOnlyNewPodcast save podcast (not sended to tg channel)
// return true if podcast is new
func (u *usecases) FindUnsendedPodcasts(count, offset int) ([]radiobot.Podcast, error) {
	return u.repo.FindUnsendedPodcasts(count, offset)
}

// SaveOnlyNewPodcastAndChannel save to db podcast and channel that return parser
func (u *usecases) SaveOnlyNewPodcastAndChannel(pAndCh radiobot.PodcastAndChannel) (isPodcastNew bool, err error) {
	p := pAndCh.Podcast
	if p == nil {
		return false, fmt.Errorf("podcast must be not nil")
	}
	ch := pAndCh.Channel
	if ch == nil {
		return false, fmt.Errorf("channel must be not nil")
	}
	err = u.RegisterORFindChannel(ch)
	if err != nil {
		return false, err
	}
	return u.SaveOnlyNewPodcast(*p)

}

func getDataPodcast(url string) (isMP3 bool, contentlength int64, body io.ReadCloser, err error) {
	var res *http.Response
	res, err = http.Get(url)
	if err != nil {
		return
	}
	contentlength = res.ContentLength
	isMP3 = strings.Contains(res.Header.Get("Content-Type"), "mpeg")
	if err == nil {
		if res.StatusCode != 200 {
			res.Body.Close()
			err = fmt.Errorf("Status code not 200")
			return
		}
		body = res.Body
	}
	return
}

// // send podcast to tg channel and update podcast in db
// SendAndUpdatePodcast(*Podcast) error
func (u *usecases) SendAndUpdatePodcast(p *radiobot.Podcast) error {

	podcastChannel, err := u.repo.FindChannelByID(p.ChannelID)
	if err != nil {
		return errors.Wrap(err, "error find channel id")
	}

	// Send comment to podcast
	title := strings.Replace(strings.Replace(podcastChannel.Name, " ", "_", -1), ".", "_", -1) + " " + p.CreatedOn.Format("2006-01-02")
	htmlURLFile := fmt.Sprintf(`<a href="%s">Dosiero</a>`, p.FileURL)
	buffer := bytes.Buffer{}
	buffer.WriteString("#")
	buffer.WriteString(title)
	buffer.WriteString("\n")
	buffer.WriteString(htmlURLFile)
	buffer.WriteString("\n\n")
	if p.Comment != "" {
		buffer.WriteString(html.EscapeString(p.Comment))
		buffer.WriteString("\n\n")

		buffer.WriteString(htmlURLFile)
	}

	descriptionMsg, err := u.tgBot.Send(u.tgChannel, buffer.String(), &telebot.SendOptions{
		DisableWebPagePreview: true,
		ParseMode:             telebot.ModeHTML,
	})
	if err != nil {
		return err
	}
	p.CommentMsgID = descriptionMsg.ID

	// Send  podcast file

	isMP3, contLength, body, err := getDataPodcast(p.FileURL)
	if err == nil {
		defer body.Close()
	}
	if err == nil && contLength <= maxAudioLengthTg {

		var fileSentable telebot.Sendable

		if isMP3 {
			fileSentable = &telebot.Audio{
				File:      telebot.FromReader(body),
				Caption:   "#" + title,
				Title:     title + ".mp3",
				Performer: podcastChannel.Name,
			}
		} else {
			fileSentable = &telebot.Document{
				File:     telebot.FromReader(body),
				Caption:  "#" + title,
				FileName: title + ".mp3",
			}
		}

		message, err := u.tgBot.Send(u.tgChannel, fileSentable, &telebot.SendOptions{
			DisableNotification: true,
		})
		if err != nil {
			return err
		}
		p.FileMsgID = message.ID

		if message.Audio != nil {
			p.FIleTgID = message.Audio.FileID
		} else if message.Document != nil {
			p.FIleTgID = message.Document.FileID
		}

	}
	p.SetRecipient(u.tgChannel)
	return u.repo.UpdatePodcast(*p)
}

// Send message to one user
func (u *usecases) Send(to telebot.Recipient, what interface{}, options ...interface{}) (*telebot.Message, error) {
	return u.tgBot.Send(to, what, options...)
}

//func (u *usecases) SendPodcastToSubscribers(Podcast) error{
//
// }

// FindOrRegisterChat register chat if not exist
// if chat exist set chat data to ch from db
func (u *usecases) FindOrRegisterChat(ch *radiobot.Chat) error {
	ch.SetID()
	ch1, err := u.repo.FindChat(ch.ID)
	if err != nil && err != mgo.ErrNotFound {
		return err
	}
	if err == mgo.ErrNotFound {
		err := u.repo.RegisterChat(ch)
		if err != nil {
			return err
		}
		return nil
	}
	ch = ch1
	return nil
}

// SubscribeChat is used to subscribe chat on channel
func (u *usecases) SubscribeChat(chat *radiobot.Chat, channel *radiobot.Channel) error {
	err := u.FindOrRegisterChat(chat)
	if err != nil {
		return err
	}
	return u.repo.SubscribeChat(chat, channel)
}

// UnsubscribeChat is used to unsubscribe chat on channel
func (u *usecases) UnsubscribeChat(chat *radiobot.Chat, channel *radiobot.Channel) error {
	err := u.FindOrRegisterChat(chat)
	if err != nil {
		return err
	}
	return u.repo.UnsubscribeChat(chat, channel)
}

func (u *usecases) GetAllChats(count, offset int) ([]*radiobot.Chat, error) {
	return u.repo.GetAllChats(count, offset)
}

// GetAllChatsIDSubscribedOn is used to fetch all chats ID which subscribed on channel
func (u *usecases) GetAllChatsIDSubscribedOn(ch *radiobot.Channel, count, offset int) ([]string, error) {
	return u.repo.GetAllChatsIDSubscribedOn(ch, count, offset)
}

// HandleTg to handle telegram handlers
func (u *usecases) HandleTg(endpoint interface{}, handler interface{}) {
	u.tgBot.Handle(endpoint, handler)
}
