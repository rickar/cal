package cal

import (
	"testing"
	"time"
)

func TestCzechHolidays(t *testing.T) {
	c := NewCalendar()
	c.Observed = ObservedExact
	AddCzechHolidays(c)
	loc, err := time.LoadLocation("Europe/Prague")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(2020, 1, 1, 0, 0, 0, 0, loc), true},   // New Year
		{time.Date(2020, 4, 10, 0, 0, 0, 0, loc), true},  // Good Friday
		{time.Date(2020, 4, 13, 0, 0, 0, 0, loc), true},  // Easter Monday
		{time.Date(2020, 5, 1, 0, 0, 0, 0, loc), true},   // Labor Day
		{time.Date(2020, 5, 8, 0, 0, 0, 0, loc), true},   // Liberation Day
		{time.Date(2020, 7, 5, 0, 0, 0, 0, loc), true},   // Saints Cyril and Methodius Day
		{time.Date(2020, 7, 6, 0, 0, 0, 0, loc), true},   // Jan Hus Day
		{time.Date(2020, 9, 28, 0, 0, 0, 0, loc), true},  // Saint Wenceslas' Day
		{time.Date(2020, 10, 28, 0, 0, 0, 0, loc), true}, // Foundation of the Independent Czechoslovak Statey
		{time.Date(2020, 11, 17, 0, 0, 0, 0, loc), true}, // Struggle for Freedom and Democracy Day
		{time.Date(2020, 12, 24, 0, 0, 0, 0, loc), true}, // Christmas Eve
		{time.Date(2020, 12, 25, 0, 0, 0, 0, loc), true}, // Christmas Day
		{time.Date(2020, 12, 26, 0, 0, 0, 0, loc), true}, // Saint Stephen's Day

		{time.Date(2020, 02, 14, 0, 0, 0, 0, loc), false}, // false test
	}

	for _, test := range tests {
		got := c.IsHoliday(test.t)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%s)", got, test.want, test.t)
		}
	}
}
