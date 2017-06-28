package botdb

import (
//"fmt"
//"log"
//"time"
)

const CountPodcastsOnPage = 4
const CountChannelsOnPage = 3

func GetChannelsList() (channels []PodcastChannelType) {
	connDB.Order("channel_name").Find(&channels)
	return channels
}

func GetChannelsListPage(page int) (channels []PodcastChannelType) {
	connDB.Offset(CountChannelsOnPage * page).Limit(CountChannelsOnPage).Find(&channels)
	return channels
}

func GetPodkastsFromChannelDB(channel string, page int) (podcasts []PodcastType) {
	connDB.Where("channel_name = ?", channel).Limit(CountPodcastsOnPage).Offset(CountPodcastsOnPage * page).Find(&podcasts)
	return podcasts
}

func GetSubscribesChannelsChat(chatid int64) (channels []ChatPodcasts) {
	//var subscribesChannels []ChatPodcasts
	if err := connDB.Where("chat = ?", chatid).Find(&channels).Error; err != nil {
		logDB.Error("GetUnsubscribesChannelsChat1 error:", err)
	}
	return channels
}

func GetUnsubscribesChannelsChat(chatid int64) (channels []PodcastChannelType) {
	var subscribesChannels []ChatPodcasts
	if err := connDB.Where("chat = ?", chatid).Find(&subscribesChannels).Error; err != nil {
		logDB.Error("GetUnsubscribesChannelsChat1 error:", err)
	}
	var subscribesChannelsName []string
	for _, sc := range subscribesChannels {
		subscribesChannelsName = append(subscribesChannelsName, sc.Channel)
	}

	if len(subscribesChannelsName) == 0 {
		if err := connDB.Find(&channels).Error; err != nil {
			logDB.Error("GetUnsubscribesChannelsChat2 error:", err)
		}
		return channels
	}
	if err := connDB.Where("channel_name NOT IN (?)", subscribesChannelsName).Find(&channels).Error; err != nil {
		logDB.Error("GetUnsubscribesChannelsChat2 error:", err)
	}
	return channels
}
