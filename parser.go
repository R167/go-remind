package remind

import (
	"time"

	dps "github.com/markusmobius/go-dateparser"
	"github.com/tj/go-naturaldate"
)

type Parser interface {
	Parse(s string) (*Reminder, error)
}

type dateparse struct {
	parser dps.Parser
	config dps.Configuration
}

func (dp *dateparse) Parse(s string) (*Reminder, error) {
	return &Reminder{text: s}, nil
}

type naturaldateParser struct {
	time time.Time
}

func (nd *naturaldateParser) Parse(s string) (*Reminder, error) {
	_, err := naturaldate.Parse(s, nd.time, naturaldate.WithDirection(naturaldate.Future))
	if err != nil {
		return nil, err
	}
	return &Reminder{text: s}, nil
}
