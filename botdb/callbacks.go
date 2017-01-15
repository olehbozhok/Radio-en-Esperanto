package botdb

import "fmt"

type CallbackGetpodcasts struct {
	Name string
	Func func() ([]PodcastType, error)
}

var callbacksArray []CallbackGetpodcasts

func RegisterCallback(callback CallbackGetpodcasts) {
	callbacksArray = append(callbacksArray, callback)
}

func GetNewPodcasts() (podcasts []PodcastType) {
	for _, f := range callbacksArray {
		lastPodcasts, err := f.Func()
		if err != nil {
			fmt.Printf("error call callback name:%s; err:%s\n", f.Name, err.Error())
		} else {
			for _, p := range lastPodcasts {
				if p.IsUnique() {
					podcasts = append(podcasts, p)
				}
			}
			fmt.Println("oK ", len(lastPodcasts))
		}
	}
	return podcasts
}
