package tgApiHandlers

import (
	"fmt"
	"strings"

	dbModels "github.com/Oleg-MBO/Radio-en-Esperanto/db/models"
)

func PodcastTextToChannel(p dbModels.Podkast, withFileUrl bool) *string {
	var message string
	message = fmt.Sprintf("%s\n#%s\n%s",
		p.RawDate,
		strings.Replace(p.ChannelName, " ", "_", -1),
		p.Description,
	)
	if withFileUrl {
		message = fmt.Sprintf("%s\n\n%s", message, p.Href)
	}
	return &message
}

func PodcastAudioCaption(p dbModels.Podkast) string {
	return fmt.Sprintf("%s %s.mp3", p.ChannelName, p.RawDate)
}
