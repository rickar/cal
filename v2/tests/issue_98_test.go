package tests

import (
	"testing"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/us"
)

// New Year calculation was not checking that the observed year was correct so
// 12/31/2022 was treated as a holiday because 12/31/2021 was.
func TestNewYear2022(t *testing.T) {
	c := cal.NewBusinessCalendar()
	c.AddHoliday(us.NewYear)

	tests := []struct {
		date    time.Time
		wantAct bool
		wantObs bool
	}{
		{time.Date(2022, 12, 31, 12, 30, 0, 0, time.UTC), false, false},
		{time.Date(2023, 1, 1, 12, 30, 0, 0, time.UTC), true, false},
		{time.Date(2023, 1, 2, 12, 30, 0, 0, time.UTC), false, true},
	}

	for i, test := range tests {
		gotAct, gotObs, _ := c.IsHoliday(test.date)
		if gotAct != test.wantAct || gotObs != test.wantObs {
			t.Errorf("[%d] gotAct: %t, wantAct: %t; gotObs: %t, wantObs: %t", i, gotAct, test.wantAct, gotObs,
				test.wantObs)
		}
	}
}
