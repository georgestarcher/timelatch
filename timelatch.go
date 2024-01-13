package timelatch

import (
	"time"
)

const TimeLatchDurationDefault time.Duration = 10 * time.Minute

// TimeLatch object.
type TimeLatch struct {
	Timestamp     time.Time
	LatchDuration time.Duration
}

// Set the default latchDuration of 10 minutes.
func (latch *TimeLatch) SetDefault() {
	latch.LatchDuration = TimeLatchDurationDefault
}

// Checks if Now() is earlier than the last latch time plus the LatchDuration.
func (latch *TimeLatch) IsLatched() bool {

	// If Timestamp has never latched return false
	if latch.Timestamp.IsZero() {
		return false
	}

	// If Now is Before latched Timestamp plus LatchDuration return true
	return time.Now().Before(latch.Timestamp.Add(latch.LatchDuration))
}
