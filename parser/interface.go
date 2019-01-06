package parser

import (
	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
)

// Parser interface for parser podcasts
type Parser interface {
	Parse() ([]radiobot.PodcastAndChannel, error)
	Name() string
}
