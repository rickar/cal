// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"reflect"
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

func TestWorkday(t *testing.T) {
	c := NewCalendar()
	c.SetWorkday(time.Monday, false)
	c.SetWorkday(time.Saturday, true)
	c.AddHoliday(Holiday{Month: time.June, Day: 12})
	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(2017, 6, 4, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2017, 6, 5, 12, 0, 0, 0, time.UTC), false},
		{time.Date(2017, 6, 6, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2017, 6, 7, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2017, 6, 8, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2017, 6, 9, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2017, 6, 10, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2017, 6, 12, 12, 0, 0, 0, time.UTC), false},
	}

	for _, test := range tests {
		got := c.IsWorkday(test.t)
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

func TestWorkdaysRemainCustom(t *testing.T) {
	c := NewCalendar()
	c.SetWorkday(time.Monday, false)
	c.SetWorkday(time.Tuesday, false)
	c.SetWorkday(time.Wednesday, false)
	c.SetWorkday(time.Thursday, false)
	c.SetWorkday(time.Friday, false)
	c.SetWorkday(time.Saturday, true)
	days := c.WorkdaysRemain(time.Date(2017, 6, 1, 12, 0, 0, 0, time.UTC))
	if days != 4 {
		t.Errorf("got: %d; want %d", days, 4)
	}

	c.SetWorkday(time.Friday, true)
	days = c.WorkdaysRemain(time.Date(2017, 6, 1, 12, 0, 0, 0, time.UTC))
	if days != 9 {
		t.Errorf("got: %d; want %d", days, 9)
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

func TestCountWorkdays(t *testing.T) {
	c := NewCalendar()
	c.Observed = ObservedExact
	c.AddHoliday(US_NewYear)

	/*
	      Dezember 2015
	   Mo Di Mi Do Fr Sa So
	   14 15 16 17 18 19 20
	   21 22 23 24 25 26 27
	   28 29 30 31

	       Januar 2016
	   Mo Di Mi Do Fr Sa So
	                1  2  3
	    4  5  6  7  8  9 10
	*/

	yearend := time.Date(2015, 12, 31, 12, 0, 0, 0, time.UTC)
	newyear := time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC)
	fifth := time.Date(2016, 1, 5, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		t    time.Time
		u    time.Time
		want int64
	}{
		{
			time.Date(2015, 12, 17, 12, 0, 0, 0, time.UTC),
			time.Date(2015, 12, 20, 12, 0, 0, 0, time.UTC),
			2,
		},
		{newyear, newyear, 0},
		{yearend, newyear, 1},
		{yearend, fifth, 3},
	}

	for _, test := range tests {
		got := c.CountWorkdays(test.t, test.u)
		if got != test.want {
			t.Errorf("got: %v; want: %v (%s-%s)", got, test.want, test.t, test.u)
		}
		got = c.CountWorkdays(test.u, test.t)
		if got != -test.want {
			t.Errorf("got: %v; want: %v (%s-%s)", got, -test.want, test.u, test.t)
		}
	}
}

func TestAddWorkDuration(t *testing.T) {
	c := NewCalendar()

	tests := []struct {
		name  string
		start time.Time
		d     time.Duration
		want  time.Time
	}{
		{
			"Sat0Hour",
			time.Date(2017, 7, 29, 12, 0, 0, 0, time.UTC),
			0,
			time.Date(2017, 7, 31, 12, 0, 0, 0, time.UTC),
		},
		{
			"Sun0Hour",
			time.Date(2017, 7, 30, 12, 0, 0, 0, time.UTC),
			0,
			time.Date(2017, 7, 31, 12, 0, 0, 0, time.UTC),
		},
		{
			"Sun24Hour",
			time.Date(2017, 7, 30, 12, 0, 0, 0, time.UTC),
			24 * time.Hour,
			time.Date(2017, 8, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			"Sun26Hour",
			time.Date(2017, 7, 30, 12, 0, 0, 0, time.UTC),
			26 * time.Hour,
			time.Date(2017, 8, 1, 14, 0, 0, 0, time.UTC),
		},
		{
			"Fri0Hour",
			time.Date(2017, 7, 28, 12, 0, 0, 0, time.UTC),
			0,
			time.Date(2017, 7, 28, 12, 0, 0, 0, time.UTC),
		},
		{
			"Fri14Hour",
			time.Date(2017, 7, 28, 12, 0, 0, 0, time.UTC),
			14 * time.Hour,
			time.Date(2017, 7, 31, 2, 0, 0, 0, time.UTC),
		},
		{
			"Fri24Hour",
			time.Date(2017, 7, 28, 12, 0, 0, 0, time.UTC),
			24 * time.Hour,
			time.Date(2017, 7, 31, 12, 0, 0, 0, time.UTC),
		},
		{
			"Tue72Hour",
			time.Date(2017, 7, 25, 12, 0, 0, 0, time.UTC),
			72 * time.Hour,
			time.Date(2017, 7, 28, 12, 0, 0, 0, time.UTC),
		},
		{
			"Tue5Day",
			time.Date(2017, 7, 25, 12, 0, 0, 0, time.UTC),
			5 * 24 * time.Hour,
			time.Date(2017, 8, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			"Tue7Day",
			time.Date(2017, 7, 25, 12, 0, 0, 0, time.UTC),
			7 * 24 * time.Hour,
			time.Date(2017, 8, 3, 12, 0, 0, 0, time.UTC),
		},
		{
			"Tue10Day",
			time.Date(2017, 7, 25, 12, 0, 0, 0, time.UTC),
			10 * 24 * time.Hour,
			time.Date(2017, 8, 8, 12, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := c.AddWorkDuration(tt.start, tt.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calendar.AddWorkDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
