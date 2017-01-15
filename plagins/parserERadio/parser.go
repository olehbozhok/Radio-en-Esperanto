package parserERadio

import (
	"github.com/Oleg-MBO/Radio-en-Esperanto-in-tg/botdb"
	"golang.org/x/net/html"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func init() {
	botdb.RegisterCallback(botdb.CallbackGetpodcasts{"ParserRadio", GetLastPodcasts})
}

// parser http://esperanto-radio.com/m
const urlpodkasts = `http://esperanto-radio.com/m`

type PodkastType struct {
	Date        time.Time
	Channel     string
	Description string
	Href        string
}

var RegexpDataAndChannel = regexp.MustCompile(`(\d\d\d\d-\d\d-\d\d)\s(.*)`)

var podcastsMap map[string]*botdb.PodcastChannelType

func init() {
	podcastsMap = make(map[string]*botdb.PodcastChannelType)
}

func getPodkastChannelIfExist(podkastname string) *botdb.PodcastChannelType {
	if val, ok := podcastsMap[podkastname]; ok {
		return val
	} else {
		return &botdb.PodcastChannelType{ChannelName: podkastname}
	}
}

func GetLastPodcasts() (podcasts []botdb.PodcastType, err error) {
	resp, err := http.Get(urlpodkasts)

	if err != nil {
		return podcasts, err
	}

	b := resp.Body
	defer b.Close() // close Body when the function returns

	z := html.NewTokenizer(b)

	var (
		indiv  bool
		status int
	)

	podcast := botdb.PodcastType{}
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			switch {
			case inCorrectDiv(t):
				indiv = true
			case indiv && t.Data == "a":
				podcast = botdb.PodcastType{}
				// parse podcast href

				setHrefPodcastFromToken(t, &podcast)
				status = 1
			case indiv && t.Data == "strong" && status == 1:
				status = 2
			}
		case tt == html.EndTagToken:
			t := z.Token()
			if indiv && t.Data == "div" {
				indiv = false
			}
		case tt == html.TextToken:
			if indiv {
				t := z.Token()
				t.Data = strings.TrimSpace(t.Data)
				switch status {
				case 2:
					//parse time and channel
					tmp := RegexpDataAndChannel.FindStringSubmatch(t.Data)
					if len(tmp) == 3 {
						podcast.RawDate = tmp[1]
						t, err := time.Parse("2006-01-02", tmp[1])
						if err != nil {
							podcast.Date = t
						}
						podcast.ChannelName = strings.TrimSpace(tmp[2])
						//fmt.Println(strings.TrimSpace(tmp[2]))
					}
					status = 3
				case 3:
					if strings.TrimSpace(t.Data) != "" {
						podcast.Description = t.Data
						// done add podcast to list
						podcasts = append(podcasts, podcast)

					}
				}
			}
		}
	}
	return podcasts, nil
}

func inCorrectDiv(t html.Token) bool {
	if t.Data == "div" {
		for _, a := range t.Attr {
			if a.Key == "id" && strings.HasPrefix(a.Val, "versio_") {
				return true
			}
		}
	}
	return false
}

func setHrefPodcastFromToken(t html.Token, p *botdb.PodcastType) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			p.Href = a.Val
		}
	}
}
