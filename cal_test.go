// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"testing"
	"time"
)

func TestWeekend(t *testing.T) {
	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(2014, 6, 1, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2014, 6, 2, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2014, 6, 10, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2014, 6, 18, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2014, 6, 26, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2014, 6, 27, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2014, 6, 14, 12, 0, 0, 0, time.UTC), true},
	}

	for _, test := range tests {
		got := IsWeekend(test.t)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%s)", got, test.want, test.t)
		}
	}
}

func TestWeekdayN(t *testing.T) {
	tests := []struct {
		t    time.Time
		d    time.Weekday
		n    int
		want bool
	}{
		{time.Date(2014, 6, 1, 12, 0, 0, 0, time.UTC), time.Sunday, 1, true},
		{time.Date(2014, 6, 2, 12, 0, 0, 0, time.UTC), time.Monday, 1, true},
		{time.Date(2014, 6, 10, 12, 0, 0, 0, time.UTC), time.Tuesday, 2, true},
		{time.Date(2014, 6, 18, 12, 0, 0, 0, time.UTC), time.Wednesday, 3, true},
		{time.Date(2014, 6, 26, 12, 0, 0, 0, time.UTC), time.Thursday, 4, true},
		{time.Date(2014, 6, 27, 12, 0, 0, 0, time.UTC), time.Friday, 4, true},
		{time.Date(2014, 6, 14, 12, 0, 0, 0, time.UTC), time.Saturday, 2, true},
		{time.Date(2014, 6, 29, 12, 0, 0, 0, time.UTC), time.Sunday, 5, true},
		{time.Date(2014, 6, 16, 12, 0, 0, 0, time.UTC), time.Wednesday, 3, false},
		{time.Date(2014, 6, 16, 12, 0, 0, 0, time.UTC), time.Monday, 2, false},
		{time.Date(2014, 6, 16, 12, 0, 0, 0, time.UTC), time.Monday, 4, false},
		{time.Date(2014, 6, 30, 12, 0, 0, 0, time.UTC), time.Monday, -1, true},
		{time.Date(2014, 6, 17, 12, 0, 0, 0, time.UTC), time.Tuesday, -2, true},
		{time.Date(2014, 6, 17, 12, 0, 0, 0, time.UTC), time.Tuesday, -6, false},
	}

	for _, test := range tests {
		got := IsWeekdayN(test.t, test.d, test.n)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%s, %d, %d)", got, test.want, test.t, test.d, test.n)
		}
	}
}

func TestHoliday(t *testing.T) {
	c := NewCalendar()
	c.AddHoliday(USMemorial)
	c.AddHoliday(USIndependence)
	c.AddHoliday(USColumbus)
	c.AddHoliday(Holiday{Offset: 100})

	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(2014, 5, 26, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2014, 7, 4, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2014, 10, 13, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2014, 5, 25, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2014, 5, 27, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2014, 1, 1, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2014, 4, 10, 12, 0, 0, 0, time.UTC), true},
	}

	for _, test := range tests {
		got := c.IsHoliday(test.t)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%s)", got, test.want, test.t)
		}
	}
}

func TestWorkdayNearest(t *testing.T) {
	c := NewCalendar()
	c.AddHoliday(USNewYear)
	c.AddHoliday(USIndependence)
	c.AddHoliday(USChristmas)

	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(2012, 1, 1, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2012, 1, 2, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2012, 1, 3, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2015, 7, 3, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2015, 7, 4, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2015, 7, 5, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2015, 7, 6, 12, 0, 0, 0, time.UTC), true},
	}

	for _, test := range tests {
		got := c.IsWorkday(test.t)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%s)", got, test.want, test.t)
		}
	}
}

func TestWorkdayExact(t *testing.T) {
	c := NewCalendar()
	c.Observed = ObservedExact
	c.AddHoliday(USNewYear)
	c.AddHoliday(USIndependence)
	c.AddHoliday(USChristmas)

	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(2011, 12, 30, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2012, 1, 1, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2012, 1, 2, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2015, 7, 3, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2015, 7, 4, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2015, 7, 5, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2015, 7, 6, 12, 0, 0, 0, time.UTC), true},
	}

	for _, test := range tests {
		got := c.IsWorkday(test.t)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%s)", got, test.want, test.t)
		}
	}
}

func TestWorkdayMonday(t *testing.T) {
	c := NewCalendar()
	c.Observed = ObservedMonday
	c.AddHoliday(USNewYear)
	c.AddHoliday(USIndependence)
	c.AddHoliday(USChristmas)

	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(2011, 12, 30, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2012, 1, 1, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2012, 1, 2, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2012, 1, 3, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2015, 7, 3, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2015, 7, 4, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2015, 7, 5, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2015, 7, 6, 12, 0, 0, 0, time.UTC), false},
	}

	for _, test := range tests {
		got := c.IsWorkday(test.t)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%s)", got, test.want, test.t)
		}
	}
}

func TestWorkdays(t *testing.T) {
	c := NewCalendar()
	c.AddHoliday(USNewYear)
	c.AddHoliday(USMLK)

	tests := []struct {
		y    int
		m    time.Month
		want int
	}{
		{2016, time.January, 19},
		{2011, time.January, 20},
		{2014, time.January, 21},
		{2016, time.February, 21},
		{2016, time.March, 23},
	}

	for _, test := range tests {
		got := c.Workdays(test.y, test.m)
		if got != test.want {
			t.Errorf("got: %d; want: %d (%d %d)", got, test.want, test.y, test.m)
		}
	}
}

func TestWorkdaysRemain(t *testing.T) {
	c := NewCalendar()
	c.AddHoliday(USNewYear)
	c.AddHoliday(USMLK)

	tests := []struct {
		t    time.Time
		want int
	}{
		{time.Date(2016, 1, 12, 12, 0, 0, 0, time.UTC), 12},
		{time.Date(2016, 1, 19, 12, 0, 0, 0, time.UTC), 8},
		{time.Date(2014, 2, 26, 12, 0, 0, 0, time.UTC), 2},
		{time.Date(2014, 3, 31, 12, 0, 0, 0, time.UTC), 0},
		{time.Date(2014, 3, 30, 12, 0, 0, 0, time.UTC), 1},
		{time.Date(2014, 3, 27, 12, 0, 0, 0, time.UTC), 2},
	}

	for _, test := range tests {
		got := c.WorkdaysRemain(test.t)
		if got != test.want {
			t.Errorf("got: %d; want: %d (%s)", got, test.want, test.t)
		}
	}
}

func TestWorkdayN(t *testing.T) {
	c := NewCalendar()
	c.AddHoliday(USNewYear)
	c.AddHoliday(USMLK)

	tests := []struct {
		y    int
		m    time.Month
		n    int
		want int
	}{
		{2016, time.January, 14, 22},
		{2016, time.January, -5, 25},
		{2016, time.January, -14, 11},
		{2016, time.February, 21, 29},
		{2016, time.February, 22, 0},
		{2016, time.February, -1, 29},
		{2016, time.February, 0, 0},
	}

	for _, test := range tests {
		got := c.WorkdayN(test.y, test.m, test.n)
		if got != test.want {
			t.Errorf("got: %d; want: %d (%d %d %d)", got, test.want, test.y, test.m, test.n)
		}
	}
}
