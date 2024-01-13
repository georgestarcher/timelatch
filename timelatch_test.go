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

func TestRunIntegrated(t *testing.T) {

	var testLatch TimeLatch
	testLatch.SetDefault()
	testLatch.LatchDuration = 1 * time.Second
	t.Logf("Test Latch: %v\n", testLatch)

	if !testLatch.IsLatched() {
		t.Log("Latching")
		testLatch.Timestamp = time.Now()
	} else {
		t.Log("Latched")
	}

	got := testLatch.IsLatched()
	expected := true

	if got != expected {
		t.Errorf("IsLatched Active: Expected %v, got %v\n", expected, got)
	}

	time.Sleep(2 * time.Second)
	t.Logf("Now: %v", time.Now())

	got = testLatch.IsLatched()
	expected = false

	if got != expected {
		t.Errorf("IsLatched Expired: Expected %v, got %v", expected, got)
	}

}
