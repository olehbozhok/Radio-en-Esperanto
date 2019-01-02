package usecases

import (
	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
	"github.com/globalsign/mgo"
	"github.com/google/uuid"
)

type usecases struct {
	repo radiobot.Repository
}

// NewUsecases create usecases
func NewUsecases(repo radiobot.Repository) radiobot.Usecase {
	return &usecases{repo: repo}
}

func (u *usecases) FindChannelByID(id uuid.UUID) (*radiobot.Channel, error) {
	return u.repo.FindChannelByID(id)
}

func (u *usecases) FindChannelByName(name string) (*radiobot.Channel, error) {
	return u.repo.FindChannelByName(name)
}

func (u *usecases) RegisterORFindChannel(ch *radiobot.Channel) (uuid.UUID, error) {
	chOld, err := u.repo.FindChannelByName(ch.Name)
	if err != nil && err != mgo.ErrNotFound {
		return uuid.UUID{}, err
	}
	if err == mgo.ErrNotFound {
		u.repo.RegisterChannel(ch)
	}
	ch = chOld
	return ch.ID, err
}
