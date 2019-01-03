package parser

import (
	"io"
	"net/http"

	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
)

type parserPodcasts struct {
	ParserName string
	DataURL    string
	ParseFunc  func(body io.Reader) ([]radiobot.PodcastAndChannel, error)
}

func (p *parserPodcasts) Name() string {
	return p.ParserName
}

func (p *parserPodcasts) fetchDataURL() (io.ReadCloser, error) {
	resp, err := http.Get(p.DataURL)
	if err != nil {
		return nil, err
	}
	return resp.Body, err
}

func (p *parserPodcasts) Parse() ([]radiobot.PodcastAndChannel, error) {
	rc, err := p.fetchDataURL()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return p.ParseFunc(rc)
}
