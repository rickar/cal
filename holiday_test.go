// (c) 2017 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"testing"
	"time"
)

func TestCalculateEaster(t *testing.T) {
	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(2016, 3, 27, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2017, 4, 16, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2018, 4, 1, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2019, 4, 21, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2020, 4, 12, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2021, 4, 4, 0, 0, 0, 0, time.UTC), true},
	}

	for _, test := range tests {
		easter := calculateEaster(test.t.Year(), test.t.Location())
		got := (test.t == easter)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%s)", got, test.want, test.t)
		}
	}
}

func TestCalculateGoodFriday(t *testing.T) {
	c := NewCalendar()
	c.AddHoliday(ECB_GoodFriday)

	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(2016, 3, 25, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2017, 4, 14, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2018, 3, 30, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2019, 4, 19, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2020, 4, 10, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2021, 4, 2, 0, 0, 0, 0, time.UTC), true},
	}

	for _, test := range tests {
		got := c.IsHoliday(test.t)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%s)", got, test.want, test.t)
		}
	}
}

func TestCalculateEasterMonday(t *testing.T) {
	c := NewCalendar()
	c.AddHoliday(ECB_EasterMonday)

	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(2016, 3, 28, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2017, 4, 17, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2018, 4, 2, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2019, 4, 22, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2020, 4, 13, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2021, 4, 5, 0, 0, 0, 0, time.UTC), true},
	}

	for _, test := range tests {
		got := c.IsHoliday(test.t)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%s)", got, test.want, test.t)
		}
	}
}

func TestDutchHolidays(t *testing.T) {
	c := NewCalendar()
	c.Observed = ObservedExact
	AddDutchHolidays(c)

	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(2014, 4, 27, 12, 0, 0, 0, time.UTC), false}, // Koningsdag (27th on a Sunday)
		{time.Date(2014, 4, 26, 12, 0, 0, 0, time.UTC), true},  // Koningsdag (26th in 2014)
		{time.Date(2017, 4, 27, 12, 0, 0, 0, time.UTC), true},  // Koningsdag (27th in 2017)
		{time.Date(2017, 5, 5, 12, 0, 0, 0, time.UTC), true},   // Bevrijdingsdag
	}

	for _, test := range tests {
		got := c.IsHoliday(test.t)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%s)", got, test.want, test.t)
		}
	}
}

func TestAddBritishHolidays(t *testing.T) {
	// The following is all of the national holidays in GB for the year 2017
	type date struct {
		day   int
		month time.Month
	}
	holidays := map[string]date{
		"new_year": {
			day:   2,
			month: time.January,
		},
		"good_friday": {
			day:   14,
			month: time.April,
		},
		"easter_monday": {
			day:   17,
			month: time.April,
		},
		"early_may": {
			day:   1,
			month: time.May,
		},
		"spring": {
			day:   29,
			month: time.May,
		},
		"summer": {
			day:   28,
			month: time.August,
		},
		"christmas": {
			day:   25,
			month: time.December,
		},
		"boxing": {
			day:   26,
			month: time.December,
		},
	}

	for name, holiday := range holidays {
		t.Run(name, func(t *testing.T) {
			c := NewCalendar()
			AddBritishHolidays(c)
			i := time.Date(2017, holiday.month, holiday.day, 0, 0, 0, 0, time.UTC)

			if !c.IsHoliday(i) {
				t.Errorf("Expected %q to be a holiday but wasn't", i)
			}
			if c.IsWorkday(i) {
				t.Errorf("Did not expect %q to be a holiday", i)
			}
		})
	}
}
