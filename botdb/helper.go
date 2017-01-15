package botdb

var CountPodcastsOnPage = 2

func GetChannelList() (channels []PodcastChannelType) {
	connDB.Find(&channels)
	return channels
}

func GetPodkastsFromChannelDB(channel string, page int) (podcasts []PodcastType) {
	connDB.Where("channel_name = ?", channel).Limit(CountPodcastsOnPage).Offset(CountPodcastsOnPage * page).Find(&podcasts)
	return podcasts
}
