package botdb

import (
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/Oleg-MBO/Radio-en-Esperanto/db/models"
)

func IsPodcastExistByHref(href string) bool {
	var exists bool
	sql := "select exists(select 1 from `podkasts` where `href`=? limit 1)"

	row := boil.GetDB().QueryRow(sql, href)

	err := row.Scan(&exists)
	if err != nil {
		panic(err)
	}

	return exists
}

//TODO: get list unique channels make paginate
func GetAllUniqueChannels() models.SubscribersPodcastSlice {
	sql := "SELECT DISTINCT  channel_name from podkasts"
	//row := boil.GetDB().QueryRow(sql)
	//
	//err := row.Scan(&channels)
	//if err != nil {
	//	panic(err)
	//}
	//return channels
	return models.SubscribersPodcastsG(qm.SQL(sql)).AllP()
}

func GetSubscribesChannelsChat(chatID int64) (subskribesChannels []string) {
	l := models.SubscribersPodcastsG(
		qm.Select(models.SubscribersPodcastColumns.ChannelName),
		qm.Where("chat = ?", chatID),
		qm.OrderBy(models.SubscribersPodcastColumns.ChannelName),
	).AllP()
	for _, q := range l {
		subskribesChannels = append(subskribesChannels, q.ChannelName)
	}
	return subskribesChannels
}

//TODO: get all rows with one query
func GetUnsubscribesChannelsChat(chatID int64) (unsubskribesChannels []string) {
	listUniqueChannels := GetAllUniqueChannels()
	mapChannels := make(map[string]struct{})
	for _, l := range listUniqueChannels {
		mapChannels[l.ChannelName] = struct{}{}
	}

	for _, ch := range GetSubscribesChannelsChat(chatID) {
		delete(mapChannels, ch)
	}

	for key := range mapChannels {
		unsubskribesChannels = append(unsubskribesChannels, key)
	}

	return unsubskribesChannels
}
