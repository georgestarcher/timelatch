package timelatch

import (
	"testing"
	"time"
)

// Test the now vs latch time cases
func TestIsLatched(t *testing.T) {

	var testLatch TimeLatch
	testLatch.SetDefault()
	t.Logf("Default Latch: %v", testLatch)

	cases := []struct {
		name      string
		latchTime time.Time
		duration  time.Duration
		expected  bool
	}{
		{
			name:      "Default",
			latchTime: time.Time{},
			duration:  10 * time.Minute,
			expected:  false,
		},
		{
			name:      "TimeLatched",
			latchTime: time.Now().Add(-time.Minute),
			duration:  10 * time.Minute,
			expected:  true,
		},
		{
			name:      "Expired",
			latchTime: time.Now().Add(-20 * time.Minute),
			duration:  5 * time.Minute,
			expected:  false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if !c.latchTime.IsZero() {
				testLatch.Timestamp = c.latchTime
			}
			testLatch.LatchDuration = c.duration
			got := testLatch.IsLatched()

			if got != c.expected {
				t.Errorf("Expected %v, got %v", c.expected, got)
			}
		})
	}
}
