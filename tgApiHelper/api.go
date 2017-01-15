package tgApiHelper

import (
	tg "github.com/bot-api/telegram"
)

var tgapi *tg.API

func InitApi(token string) *tg.API {
	tgapi = tg.New(token)
	return tgapi
}
