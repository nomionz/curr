package ecb

import (
	"log/slog"
	"time"
)

// DefaultTicker is just a simple ticker that ticks at 14:10:00 CET every day.
// more info https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html
func DefaultTicker() *time.Timer {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		slog.Error("failed to load location", "err", err)
		return nil
	}

	t := time.Now().In(loc)
	desiredTime := time.Date(t.Year(), t.Month(), t.Day(), 14, 10, 0, 0, loc)
	durationUntilTick := desiredTime.Sub(t)

	if durationUntilTick < 0 {
		durationUntilTick += 24 * time.Hour
	}

	return time.NewTimer(durationUntilTick)
}
