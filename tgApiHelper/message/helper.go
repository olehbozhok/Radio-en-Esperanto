package Message

import (
	"github.com/olebedev/go-tgbot/client/messages"
	"github.com/olebedev/go-tgbot/models"
)

func Sendreply(m *models.Message, text string) (*messages.SendMessageOK, error) {
	return Tgapi.Messages.SendMessage(
		messages.NewSendMessageParams().WithBody(&models.SendMessageBody{
			Text:             &text,
			ChatID:           m.Chat.ID,
			ReplyToMessageID: m.MessageID,
		}),
	)
}

//func ForwardMessage()  {
//
//}
