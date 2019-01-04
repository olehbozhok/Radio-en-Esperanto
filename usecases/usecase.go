package usecases

import (
	"bytes"
	"fmt"
	"html"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
	"github.com/globalsign/mgo"
	"github.com/google/uuid"
	telebot "gopkg.in/tucnak/telebot.v2"
)

const maxFileLengthTg = 20000

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

func checkDataPodcastHEAD(url string) (isMP3 bool, contentlength int64, err error) {
	var res *http.Response
	res, err = http.Head(url)
	if err != nil {
		return
	}
	contentlength = res.ContentLength
	isMP3 = strings.Contains(res.Header.Get("Content-Type"), "mpeg")
	return
}

// // send podcast to tg channel and update podcast in db
// SendAndUpdatePodcast(*Podcast) error
func (u *usecases) SendAndUpdatePodcast(p *radiobot.Podcast) error {

	podcastChannel, err := u.repo.FindChannelByID(p.ChannelID)
	if err != nil {
		return errors.Wrap(err, "error find channel id")
	}
	// TODO Replace p.ParsedOn.String() to better print
	// markdownFile := fmt.Sprintf(`[Dosiero](%s)`, p.FileURL)
	markdownFile := fmt.Sprintf(`<a href="%s">Dosiero</a>`, p.FileURL)
	buffer := bytes.Buffer{}
	buffer.WriteString("#")
	buffer.WriteString(strings.Replace(strings.Replace(podcastChannel.Name, " ", "_", -1), ".", "_", -1))
	buffer.WriteString(" ")
	buffer.WriteString(p.CreatedOn.Format("2006-01-02"))
	buffer.WriteString("\n")
	buffer.WriteString(markdownFile)
	buffer.WriteString("\n\n")
	if p.Comment != "" {
		buffer.WriteString(html.EscapeString(p.Comment))
		buffer.WriteString("\n\n")

		buffer.WriteString(markdownFile)

	}
	fmt.Println(p.Comment)

	descriptionMsg, err := u.tgBot.Send(u.tgChannel, buffer.String(), &telebot.SendOptions{
		DisableWebPagePreview: false,
		ParseMode:             telebot.ModeHTML,
	})
	if err != nil {
		return err
	}
	p.CommentMsgID = descriptionMsg.ID

	// var fileSentable telebot.Sendable

	// if isMP3 {
	// 	fileSentable = &telebot.Audio{
	// 		File:    telebot.FromURL(p.FileURL),
	// 		Caption: "#" + title,
	// 		// Title:   title + ".mp3",
	// 	}
	// } else {
	// 	fileSentable = &telebot.Document{
	// 		File:    telebot.FromURL(p.FileURL),
	// 		Caption: "#" + title,
	// 		// FileName:   title + ".mp3",
	// 	}
	// }

	// if contentlength <= maxFileLengthTg {
	// 	message, err := u.tgBot.Send(u.tgChannel, fileSentable, &telebot.SendOptions{
	// 		DisableNotification: true,
	// 	})
	// 	if err != nil {
	// 		return err
	// 	}
	// 	p.SetRecipient(u.tgChannel)
	// 	p.FileMsgID = message.ID

	// 	if message.Audio != nil {
	// 		p.FIleTgID = message.Audio.FileID
	// 	} else if message.Document != nil {
	// 		p.FIleTgID = message.Audio.FileID
	// 	}

	// }
	// TODO: ENABLE UPDATE podcast in db
	// return u.repo.UpdatePodcast(*p)
	return nil
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
