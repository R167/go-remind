package remind_test

import (
	"testing"
	"time"

	"github.com/R167/go-remind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func hours(h time.Duration) time.Duration {
	return h * time.Hour
}

func minutes(m time.Duration) time.Duration {
	return m * time.Minute
}

func tt(y, m, d int, parts ...int) time.Time {
	var h, min, s int
	if len(parts) > 0 {
		h = parts[0]
	}
	if len(parts) > 1 {
		min = parts[1]
	}
	if len(parts) > 2 {
		s = parts[2]
	}
	return time.Date(y, time.Month(m), d, h, min, s, 0, time.UTC)
}

func TestParsingReminders(t *testing.T) {
	start := tt(2023, 1, 1)
	after := func(dur time.Duration) time.Time {
		return start.Add(dur)
	}

	reminders := []struct {
		input    string
		reminder *remind.Reminder
	}{
		{"update the sitrep every 2 hours", &remind.Reminder{
			At:      after(hours(2)),
			Every:   hours(2),
			Message: "update the sitrep",
			Target:  remind.Me,
		}},
		{"update the sitrep every 2 hours starting at 9am", &remind.Reminder{
			At:      tt(2023, 1, 1, 9),
			Every:   hours(2),
			Message: "update the sitrep",
			Target:  remind.Me,
		}},
		// @bob to clear the deploy lock tomorrow morning
		{"@bob to clear the deploy lock tomorrow at 9am", &remind.Reminder{
			At:      tt(2023, 1, 2, 9),
			Message: "clear the deploy lock",
			Target:  remind.User("bob"),
		}},
		// in 5 minutes, make sure I clear page in eim in case things get worse
		{"in 5 minutes, make sure I clear page in eim in case things get worse", &remind.Reminder{
			At:      after(minutes(5)),
			Message: "make sure I clear page in eim in case things get worse",
			Target:  remind.Me,
		}},
		// NOTE: The `Target` is not parsed correctly because it is deep in the sentence.
		{"In 12 minutes, make sure @george does something important", &remind.Reminder{
			At:      after(minutes(12)),
			Message: "make sure @george does something important",
			Target:  remind.Me,
		}},
		// Every 30 minutes, post a reminder to the #standup channel
		{"Every 30 minutes, post a reminder to the #standup channel", &remind.Reminder{
			At:      after(minutes(30)),
			Every:   minutes(30),
			Message: "post a reminder to the #standup channel",
			Target:  remind.Me,
		}},
		// in an hour and a half, remind me to check the status of the deploy every 19 minutes
		{"in an hour and a half, remind me to check the status of the deploy every 19 minutes", &remind.Reminder{
			At:      after(hours(1) + minutes(30)),
			Every:   minutes(19),
			Message: "remind me to check the status of the deploy",
			Target:  remind.Me,
		}},
		// me on tueday at 17:00 to clear the checks with tiffany and greg
		{"me on tuesday at 17:00 to clear the checks with tiffany and greg", &remind.Reminder{
			At:      tt(2023, 1, 3, 17),
			Message: "clear the checks with tiffany and greg",
			Target:  remind.Me,
		}},
		// @winston to post an update to customers every friday at 4pm
		{"@winston to post an update to customers every friday at 4pm", &remind.Reminder{
			At:      tt(2023, 1, 6, 16),
			Every:   hours(24 * 7),
			Message: "post an update to customers",
			Target:  remind.User("winston"),
		}},
	}

	failures := []string{
		// Tomorrow morning is imprecise, so it returns and error instead
		"@bob to clear the deploy lock tomorrow morning",
		// post an update on wednesday at 3pm and every tuesday after that
		// NOTE: too complex to parse
		"post an update on wednesday at 3pm and every tuesday after that",
	}

	for _, tc := range reminders {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			r, err := remind.Parse(tc.input)
			require.NoError(t, err)
			assert.Equal(t, tc.reminder, *r)
		})
	}

	for _, tc := range failures {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			v, err := remind.Parse(tc)
			require.Error(t, err, v)
		})
	}
}
