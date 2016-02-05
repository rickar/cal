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

func TestMonthStart(t *testing.T) {
	tests := []struct {
		t    time.Time
		want time.Time
	}{
		{time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)},
		{time.Date(2016, 2, 15, 12, 0, 0, 0, time.UTC), time.Date(2016, 2, 1, 12, 0, 0, 0, time.UTC)},
		{time.Date(2016, 3, 31, 23, 59, 59, 999999999, time.UTC), time.Date(2016, 3, 1, 23, 59, 59, 999999999, time.UTC)},
	}

	for _, test := range tests {
		got := MonthStart(test.t)
		if got != test.want {
			t.Errorf("got: %s; want: %s (%s)", got, test.want, test.t)
		}
	}
}

func TestMonthEnd(t *testing.T) {
	tests := []struct {
		t    time.Time
		want time.Time
	}{
		{time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2016, 1, 31, 0, 0, 0, 0, time.UTC)},
		{time.Date(2016, 2, 15, 12, 0, 0, 0, time.UTC), time.Date(2016, 2, 29, 12, 0, 0, 0, time.UTC)},
		{time.Date(2016, 3, 31, 23, 59, 59, 999999999, time.UTC), time.Date(2016, 3, 31, 23, 59, 59, 999999999, time.UTC)},
	}

	for _, test := range tests {
		got := MonthEnd(test.t)
		if got != test.want {
			t.Errorf("got: %s; want: %s (%s)", got, test.want, test.t)
		}
	}
}

func TestJulianDayNumber(t *testing.T) {
	tests := []struct {
		t    time.Time
		want int
	}{
		{time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), 2440587},
		{time.Date(2000, 1, 1, 15, 0, 0, 0, time.UTC), 2451545},
		{time.Date(2016, 2, 29, 12, 0, 0, 0, time.UTC), 2457448},
		{time.Date(2016, 3, 31, 23, 59, 59, 999999999, time.UTC), 2457479},
	}

	for _, test := range tests {
		got := JulianDayNumber(test.t)
		if got != test.want {
			t.Errorf("got: %d; want: %d (%s)", got, test.want, test.t)
		}
	}
}

func TestJulianDate(t *testing.T) {
	tests := []struct {
		t    time.Time
		want float32
	}{
		{time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), 2440587.5},
		{time.Date(2000, 1, 1, 15, 0, 0, 0, time.UTC), 2451545.125},
		{time.Date(2016, 2, 29, 12, 0, 0, 0, time.UTC), 2457448.0},
		{time.Date(2016, 3, 31, 23, 59, 59, 999999999, time.UTC), 2457479.5},
	}

	for _, test := range tests {
		got := JulianDate(test.t)
		if got != test.want {
			t.Errorf("got: %f; want: %f (%s)", got, test.want, test.t)
		}
	}
}

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

func TestHoliday(t *testing.T) {
	c := NewCalendar()
	c.AddHoliday(US_Memorial)
	c.AddHoliday(US_Independence)
	c.AddHoliday(US_Columbus)
	c.AddHoliday(Holiday{Offset: 100})

	tz, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Errorf("unable to load time zone: %v", err)
	}

	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(2014, 5, 26, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2014, 5, 26, 23, 59, 59, 0, time.UTC), true},
		{time.Date(2014, 5, 26, 00, 0, 0, 0, time.UTC), true},
		{time.Date(2014, 5, 26, 23, 59, 59, 0, tz), true},
		{time.Date(2014, 5, 26, 00, 0, 0, 0, tz), true},
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
	c.AddHoliday(US_NewYear)
	c.AddHoliday(US_Independence)
	c.AddHoliday(US_Christmas)

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
	c.AddHoliday(US_NewYear)
	c.AddHoliday(US_Independence)
	c.AddHoliday(US_Christmas)

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
	c.AddHoliday(US_NewYear)
	c.AddHoliday(US_Independence)
	c.AddHoliday(US_Christmas)

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
	c.AddHoliday(US_NewYear)
	c.AddHoliday(US_MLK)

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
	c.AddHoliday(US_NewYear)
	c.AddHoliday(US_MLK)

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
	c.AddHoliday(US_NewYear)
	c.AddHoliday(US_MLK)

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
