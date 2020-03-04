package cal

import (
	"testing"
	"time"
)

func TestCzechHolidays(t *testing.T) {
	c := NewCalendar()
	c.Observed = ObservedExact
	AddSlovakHolidays(c)
	loc, err := time.LoadLocation("Europe/Bratislava")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(2020, time.January, 1, 0, 0, 0, 0, loc), true},    // New Year
		{time.Date(2020, time.January, 6, 0, 0, 0, 0, loc), true},    // Epiphany
		{time.Date(2020, time.April, 10, 0, 0, 0, 0, loc), true},     // Good Friday
		{time.Date(2020, time.April, 13, 0, 0, 0, 0, loc), true},     // Easter Monday
		{time.Date(2020, time.May, 1, 0, 0, 0, 0, loc), true},        // Labor Day
		{time.Date(2020, time.May, 8, 0, 0, 0, 0, loc), true},        // Liberation Day
		{time.Date(2020, time.July, 5, 0, 0, 0, 0, loc), true},       // Saints Cyril and Methodius Day
		{time.Date(2020, time.August, 29, 0, 0, 0, 0, loc), true},    // SNP
		{time.Date(2020, time.September, 1, 0, 0, 0, 0, loc), true},  // Slovak Consitution day
		{time.Date(2020, time.September, 15, 0, 0, 0, 0, loc), true}, // Our Lady of Sorrows
		{time.Date(2020, time.November, 1, 0, 0, 0, 0, loc), true},   // All Saints' Day
		{time.Date(2020, time.November, 17, 0, 0, 0, 0, loc), true},  // Freedom and Democracy Day
		{time.Date(2020, time.December, 24, 0, 0, 0, 0, loc), true},  // Christmas Eve
		{time.Date(2020, time.December, 25, 0, 0, 0, 0, loc), true},  // Christmas Day
		{time.Date(2020, time.December, 26, 0, 0, 0, 0, loc), true},  // Saint Stephen's Day

		{time.Date(2020, 02, 14, 0, 0, 0, 0, loc), false}, // false test
	}

	for _, test := range tests {
		got := c.IsHoliday(test.t)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%s)", got, test.want, test.t)
		}
	}
}
