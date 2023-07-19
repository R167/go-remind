package remind

import (
	"errors"
	"time"

	dps "github.com/markusmobius/go-dateparser"
)

var ErrNoDuration = errors.New("No duration found")

func ParseDuration(dur string) (time.Duration, error) {
	start := time.Unix(0, 0).UTC()
	config := dps.Configuration{
		Languages:           []string{"en"},
		PreferredDateSource: dps.Future,
		CurrentTime:         start,
	}
	parser := dps.Parser{
		ParserTypes: []dps.ParserType{dps.RelativeTime},
	}

	_, t, err := parser.Search(&config, dur)
	if err != nil {
		return time.Duration(0), err
	}
	if len(t) == 0 {
		return time.Duration(0), ErrNoDuration
	}
	return t[len(t)-1].Date.Time.Sub(start), nil
}
