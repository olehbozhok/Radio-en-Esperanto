package botdb

import "fmt"

type CallbackGetpodcasts struct {
	Name string
	Func func() ([]PodcastType, error)
}

var callbacksArray []CallbackGetpodcasts

func RegisterParser(callback CallbackGetpodcasts) {
	callbacksArray = append(callbacksArray, callback)
}

func GetNewPodcasts() (podcasts []PodcastType) {
	for _, f := range callbacksArray {
		lastPodcasts, err := f.Func()
		if err != nil {
			fmt.Printf("error call callback name:%s; err:%s\n", f.Name, err.Error())
		} else {
			for i := len(lastPodcasts) - 1; i >= 0; i-- {
				p := lastPodcasts[i]
				if p.IsUnique() {
					podcasts = append(podcasts, p)
				}
			}
			fmt.Println("OK ", len(lastPodcasts))
		}
	}
	return podcasts
}
