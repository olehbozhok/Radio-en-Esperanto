package parserERadio

import (
	"golang.org/x/net/html"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/Oleg-MBO/Radio-en-Esperanto/botdb"
)

// parser http://esperanto-radio.com/m
const urlpodkasts = `http://esperanto-radio.com/m`

type PodkastType struct {
	RawDate     time.Time
	ChannelName string
	Description string
	Href        string
}

var RegexpDataAndChannel = regexp.MustCompile(`(\d\d\d\d-\d\d-\d\d)\s(.*)`)

func GetLastPodcasts() (podcasts []botdb.PodcastType, err error) {
	resp, err := http.Get(urlpodkasts)

	if err != nil {
		return podcasts, err
	}

	return parseLastPodcasts(resp.Body)
}

func parseLastPodcasts(body io.ReadCloser) (podcasts []botdb.PodcastType, err error) {

	defer body.Close() // close Body when the function returns

	z := html.NewTokenizer(body)

	var (
		inDiv  bool
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
				inDiv = true
			case inDiv && t.Data == "a":
				podcast = botdb.PodcastType{}
				// parse podcast href

				setHrefPodcastFromToken(t, &podcast)
				status = 1
			case inDiv && t.Data == "strong" && status == 1:
				status = 2
			}
		case tt == html.EndTagToken:
			t := z.Token()
			if inDiv && t.Data == "div" {
				inDiv = false
			}
		case tt == html.TextToken:
			if inDiv {
				t := z.Token()
				t.Data = strings.TrimSpace(t.Data)
				switch status {
				case 2:
					//parse time and channel
					tmp := RegexpDataAndChannel.FindStringSubmatch(t.Data)
					if len(tmp) == 3 {
						podcast.RawDate = tmp[1]
						//t, err := time.Parse("2006-01-02", tmp[1])
						//if err != nil {
						//	podcast.Date = t
						//}
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
