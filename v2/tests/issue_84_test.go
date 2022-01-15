package tests

import (
	"testing"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/us"
)

// Holidays that were pushed into the next or previous year due to observance
// rules were not properly handled.
func TestNewYear(t *testing.T) {
	c := cal.NewBusinessCalendar()
	c.AddHoliday(us.NewYear)

	tests := []struct {
		date    time.Time
		wantAct bool
		wantObs bool
	}{
		{time.Date(2021, 1, 1, 12, 30, 0, 0, time.UTC), true, true},
		{time.Date(2021, 12, 31, 12, 30, 0, 0, time.UTC), false, true},
		{time.Date(2022, 1, 1, 12, 30, 0, 0, time.UTC), true, false},
	}

	for i, test := range tests {
		gotAct, gotObs, _ := c.IsHoliday(test.date)
		if gotAct != test.wantAct || gotObs != test.wantObs {
			t.Errorf("[%d] gotAct: %t, wantAct: %t; gotObs: %t, wantObs: %t", i, gotAct, test.wantAct, gotObs,
				test.wantObs)
		}
	}
}
