package remind

import (
	"time"

	dps "github.com/markusmobius/go-dateparser"
)

type Reminder struct {
	At      time.Time
	Every   time.Duration
	Message string
	Target  Target

	// Full text of the initial reminder
	text string
}

func Parse(s string) (*Reminder, error) {
	return &Reminder{text: s}, nil
}

func New() Parser {
	return &dateparse{
		parser: dps.Parser{},
		config: dps.Configuration{Languages: []string{"en"}},
	}
}
