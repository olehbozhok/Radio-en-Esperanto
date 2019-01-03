package parser

import (
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"

	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
	"github.com/PuerkitoBio/goquery"
)

// NewEsperantoRadio parser to parse podcasts from esperanto-radio.com
func NewEsperantoRadio() Parser {
	ParserName := "esperanto-radio.com"
	var RegexpDataAndChannel = regexp.MustCompile(`(\d\d\d\d-\d\d-\d\d)\s(.*)`)
	return &parserPodcasts{
		ParserName: ParserName,
		DataURL:    "http://esperanto-radio.com/m",
		ParseFunc: func(body io.Reader) ([]radiobot.PodcastAndChannel, error) {
			doc, err := goquery.NewDocumentFromReader(body)
			if err != nil {
				return nil, err
			}
			podcastsList := make([]radiobot.PodcastAndChannel, 0)
			doc.Find("div > a").EachWithBreak(func(i int, s *goquery.Selection) bool {
				podcast := new(radiobot.Podcast)
				podcast.ParsedOn = time.Now()
				channel := new(radiobot.Channel)
				channel.Parser = ParserName

				url, exist := s.Attr("href")
				if !exist {
					err = fmt.Errorf("err parse parser %s, could not find href attribute", ParserName)
				}
				podcast.FileURL = url
				strongElem := s.Find("strong")
				tmp := RegexpDataAndChannel.FindStringSubmatch(strongElem.Text())
				if len(tmp) == 3 {
					timeText := tmp[1]
					var t time.Time
					t, err = time.Parse("2006-01-02", timeText)
					if err != nil {
						return false
					}
					podcast.CreatedOn = t
					channel.Name = strings.TrimSpace(tmp[2])

				} else {
					err = fmt.Errorf("RegexpDataAndChannel.FindStringSubmatch(strongElem.Text()) != 3")
					return false
				}

				channel.CalcID()
				podcast.ChannelID = channel.ID
				podcast.CalcID()
				podcastsList = append(podcastsList, radiobot.PodcastAndChannel{
					Channel: channel,
					Podcast: podcast,
				})
				return true
			})
			return podcastsList, err
		},
	}
}
