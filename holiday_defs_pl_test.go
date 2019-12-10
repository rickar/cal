package cal

import (
	"testing"
	"time"
)

func TestPolandHolidays(t *testing.T) {
	c := NewCalendar()
	c.Observed = ObservedExact
	AddPolandHolidays(c)

	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(2018, 1, 1, 12, 0, 0, 0, time.UTC), true},    // New Year
		{time.Date(2018, 1, 6, 12, 0, 0, 0, time.UTC), true},    // Three Kings
		{time.Date(2018, 4, 1, 12, 0, 0, 0, time.UTC), true},    // Easter Sunday
		{time.Date(2018, 4, 2, 12, 0, 0, 0, time.UTC), true},    // Easter Monday
		{time.Date(2018, 5, 1, 12, 0, 0, 0, time.UTC), true},    // Labour Day
		{time.Date(2018, 5, 3, 12, 0, 0, 0, time.UTC), true},    // National Day
		{time.Date(2018, 5, 20, 12, 0, 0, 0, time.UTC), true},   // Pentecost Day
		{time.Date(2018, 5, 31, 12, 0, 0, 0, time.UTC), true},   // Corpus Christi
		{time.Date(2018, 8, 15, 12, 0, 0, 0, time.UTC), true},   // Assumption Blessed Virgin Mary
		{time.Date(2018, 11, 1, 12, 0, 0, 0, time.UTC), true},   // All Saints
		{time.Date(2018, 11, 11, 12, 0, 0, 0, time.UTC), true},  // National Independence Day
		{time.Date(2018, 12, 25, 12, 0, 0, 0, time.UTC), true},  // Christmas Day One
		{time.Date(2018, 12, 26, 12, 0, 0, 0, time.UTC), true},  // Christmas Day Two
		{time.Date(2018, 02, 14, 12, 0, 0, 0, time.UTC), false}, // false test
	}

	for _, test := range tests {
		got := c.IsHoliday(test.t)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%s)", got, test.want, test.t)
		}
	}
}
