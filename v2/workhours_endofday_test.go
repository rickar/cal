package cal

import (
	"testing"
	"time"
)

func TestWorkHoursEndOfDay(t *testing.T) {
	monday := time.Date(2026, time.October, 26, 12, 0, 0, 0, time.UTC)
	for _, tc := range []struct {
		name             string
		start, end, want time.Duration
	}{
		{"full day", 0, 24 * time.Hour, 24 * time.Hour},
		{"afternoon", 9 * time.Hour, 24 * time.Hour, 15 * time.Hour},
		{"partial hour", 9*time.Hour + 30*time.Minute, 24 * time.Hour, 14*time.Hour + 30*time.Minute},
		{"standard", 9 * time.Hour, 17 * time.Hour, 8 * time.Hour},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := NewBusinessCalendar()
			c.SetWorkHours(tc.start, tc.end)
			if got := c.WorkHours(monday); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
			if got := c.WorkHours(monday.AddDate(0, 0, -1)); got != 0 {
				t.Errorf("nonworking day: got %v", got)
			}
		})
	}
}
