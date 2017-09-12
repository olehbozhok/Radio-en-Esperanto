package parser

import (
	"fmt"
	"github.com/Oleg-MBO/Radio-en-Esperanto/db/models"
	"github.com/labstack/gommon/log"
)

type CallbackGetpodcasts struct {
	Name string
	Func func() ([]models.Podkast, error)
}

var callbacksArray []CallbackGetpodcasts

func RegisterParser(callback CallbackGetpodcasts) {
	callbacksArray = append(callbacksArray, callback)
}

func GetNewPodcasts() (podcasts []models.Podkast) {
	for _, f := range callbacksArray {
		//log.Infof("handle parser %s\n", f.Name)
		lastPodcasts, err := f.Func()
		log.Infof("handle parser %s, len list:%d\n", f.Name, len(lastPodcasts))
		if err != nil {
			fmt.Printf("error call callback name:%s; err:%s\n", f.Name, err.Error())
		} else {
			for i := len(lastPodcasts) - 1; i >= 0; i-- {
				p := lastPodcasts[i]
				//if !botdb.IsPodcastExistByHref(p.Href) {
				podcasts = append(podcasts, p)
				//}
			}
			fmt.Println("OK ", len(lastPodcasts))
		}
	}
	return podcasts
}
