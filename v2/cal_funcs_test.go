// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"math"
	"testing"
	"time"
)

func TestIsWeekend(t *testing.T) {
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

func TestWeekdayNFrom(t *testing.T) {
	tests := []struct {
		t    time.Time
		day  time.Weekday
		n    int
		want time.Time
	}{
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Sunday, 1, time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Saturday, 1, time.Date(2014, time.June, 7, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Sunday, -1, time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Saturday, -1, time.Date(2014, time.May, 31, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Sunday, 2, time.Date(2014, time.June, 8, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Saturday, 2, time.Date(2014, time.June, 14, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Sunday, 6, time.Date(2014, time.July, 6, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Saturday, 6, time.Date(2014, time.July, 12, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Sunday, 10, time.Date(2014, time.August, 3, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Saturday, 10, time.Date(2014, time.August, 9, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Sunday, -2, time.Date(2014, time.May, 25, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Saturday, -2, time.Date(2014, time.May, 24, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Sunday, -6, time.Date(2014, time.April, 27, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Saturday, -6, time.Date(2014, time.April, 26, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Sunday, -10, time.Date(2014, time.March, 30, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Saturday, -10, time.Date(2014, time.March, 29, 12, 0, 0, 0, time.UTC)},
		{time.Date(2014, time.June, 1, 12, 0, 0, 0, time.UTC), time.Sunday, 0, time.Time{}},
	}

	for _, test := range tests {
		got := WeekdayNFrom(test.t, test.day, test.n)
		if !got.Equal(test.want) {
			t.Errorf("got: %s; want: %s (%s, %d, %d)", got.String(), test.want.String(), test.t.String(), test.day, test.n)
		}
	}
}

func TestWeekdayN(t *testing.T) {
	tests := []struct {
		y         int
		m         time.Month
		d         time.Weekday
		n         int
		wantDay   int
		wantMonth time.Month
	}{
		{2014, time.June, time.Sunday, 1, 1, time.June},
		{2014, time.June, time.Monday, 1, 2, time.June},
		{2014, time.June, time.Tuesday, 2, 10, time.June},
		{2014, time.June, time.Wednesday, 3, 18, time.June},
		{2014, time.June, time.Thursday, 4, 26, time.June},
		{2014, time.June, time.Friday, 4, 27, time.June},
		{2014, time.June, time.Saturday, 2, 14, time.June},
		{2014, time.June, time.Sunday, 5, 29, time.June},
		{2014, time.June, time.Monday, -1, 30, time.June},
		{2014, time.June, time.Sunday, -1, 29, time.June},
		{2014, time.June, time.Saturday, -1, 28, time.June},
		{2014, time.June, time.Friday, -1, 27, time.June},
		{2014, time.June, time.Thursday, -1, 26, time.June},
		{2014, time.June, time.Wednesday, -1, 25, time.June},
		{2014, time.June, time.Tuesday, -1, 24, time.June},
		{2014, time.June, time.Monday, -2, 23, time.June},
		{2014, time.June, time.Tuesday, -2, 17, time.June},
		{2014, time.June, time.Monday, -5, 2, time.June},
		{2014, time.June, time.Sunday, -5, 1, time.June},
		{2014, time.June, time.Sunday, 0, 1, time.January},
		{2014, time.June, time.Saturday, 0, 1, time.January},
		{2014, time.June, time.Monday, -6, 26, time.May},
	}

	for _, test := range tests {
		got := WeekdayN(test.y, test.m, test.d, test.n)
		if got.Day() != test.wantDay || got.Month() != test.wantMonth {
			t.Errorf("got: %s; want: %d-%d (%d, %d, %d, %d)", got.String(), test.wantDay, test.wantMonth,
				test.y, test.m, test.d, test.n)
		}
	}
}

func TestIsWeekdayN(t *testing.T) {
	tests := []struct {
		t    time.Time
		d    time.Weekday
		n    int
		want bool
	}{
		{time.Date(2014, 6, 1, 12, 0, 0, 0, time.UTC), time.Sunday, 1, true},
		{time.Date(2014, 6, 2, 12, 0, 0, 0, time.UTC), time.Monday, 1, true},
		{time.Date(2014, 6, 8, 12, 0, 0, 0, time.UTC), time.Sunday, 2, true},
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
		{time.Date(2014, 6, 29, 12, 0, 0, 0, time.UTC), time.Sunday, -1, true},
		{time.Date(2014, 6, 28, 12, 0, 0, 0, time.UTC), time.Saturday, -1, true},
		{time.Date(2014, 6, 27, 12, 0, 0, 0, time.UTC), time.Friday, -1, true},
		{time.Date(2014, 6, 27, 12, 0, 0, 0, time.UTC), time.Thursday, -1, false},
		{time.Date(2014, 6, 26, 12, 0, 0, 0, time.UTC), time.Thursday, -1, true},
		{time.Date(2014, 6, 25, 12, 0, 0, 0, time.UTC), time.Wednesday, -1, true},
		{time.Date(2014, 6, 24, 12, 0, 0, 0, time.UTC), time.Tuesday, -1, true},
		{time.Date(2014, 6, 23, 12, 0, 0, 0, time.UTC), time.Monday, -2, true},
		{time.Date(2014, 6, 17, 12, 0, 0, 0, time.UTC), time.Tuesday, -2, true},
		{time.Date(2014, 6, 17, 12, 0, 0, 0, time.UTC), time.Tuesday, -6, false},
		{time.Date(2014, 6, 2, 12, 0, 0, 0, time.UTC), time.Monday, -5, true},
		{time.Date(2014, 6, 1, 12, 0, 0, 0, time.UTC), time.Sunday, -5, true},
	}

	for _, test := range tests {
		got := IsWeekdayN(test.t, test.d, test.n)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%s, %d, %d)", got, test.want, test.t, test.d, test.n)
		}
	}
}

func TestDayStart(t *testing.T) {
	val := time.Date(2016, 6, 12, 12, 30, 40, 1234, time.UTC)
	got := DayStart(val)
	want := time.Date(2016, 6, 12, 0, 0, 0, 0, time.UTC)
	if got.UnixNano() != want.UnixNano() {
		t.Errorf("got: %s; want: %s", got.String(), want.String())
	}
}

func TestDayEnd(t *testing.T) {
	val := time.Date(2016, 6, 12, 12, 30, 40, 1234, time.UTC)
	got := DayEnd(val)
	want := time.Date(2016, 6, 12, 23, 59, 59, 999999999, time.UTC)
	if got.UnixNano() != want.UnixNano() {
		t.Errorf("got: %s; want: %s", got.String(), want.String())
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

func TestReplaceLocation(t *testing.T) {
	tests := []struct {
		t    time.Time
		want time.Time
	}{
		{time.Date(2000, 1, 1, 12, 30, 0, 0, time.UTC), time.Date(2000, 1, 1, 12, 30, 0, 0, DefaultLoc)},
		{time.Date(2000, 1, 1, 12, 30, 0, 0, time.Local), time.Date(2000, 1, 1, 12, 30, 0, 0, DefaultLoc)},
	}

	for _, test := range tests {
		got := ReplaceLocation(test.t, DefaultLoc)
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

func TestModifiedJulianDayNumber(t *testing.T) {
	tests := []struct {
		t    time.Time
		want int
	}{
		{time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), 44239},
		{time.Date(1980, 1, 1, 12, 0, 0, 0, time.UTC), 44239},
		{time.Date(2000, 1, 1, 15, 0, 0, 0, time.UTC), 51544},
		{time.Date(2016, 2, 29, 12, 0, 0, 0, time.UTC), 57447},
		{time.Date(2016, 3, 31, 23, 59, 59, 999999999, time.UTC), 57478},
	}

	for _, test := range tests {
		got := ModifiedJulianDayNumber(test.t)
		if got != test.want {
			t.Errorf("got: %d; want: %d (%s)", got, test.want, test.t)
		}
	}
}

func TestJulianDate(t *testing.T) {
	tests := []struct {
		t    time.Time
		want float64
	}{
		{time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), 2440587.5},
		{time.Date(2000, 1, 1, 15, 0, 0, 0, time.UTC), 2451545.125},
		{time.Date(2016, 2, 29, 12, 0, 0, 0, time.UTC), 2457448.0},
		{time.Date(2016, 3, 31, 23, 59, 59, 999999999, time.UTC), 2457479.499988},
	}

	for _, test := range tests {
		got := JulianDate(test.t)
		if math.Abs(got-test.want) > 0.000001 {
			t.Errorf("got: %f; want: %f (%s)", got, test.want, test.t)
		}
	}
}

func TestModifiedJulianDate(t *testing.T) {
	tests := []struct {
		t    time.Time
		want float64
	}{
		{time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), 40587.0},
		{time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), 44239.0},
		{time.Date(2000, 1, 1, 15, 0, 0, 0, time.UTC), 51544.625},
		{time.Date(2016, 2, 29, 12, 0, 0, 0, time.UTC), 57447.5},
		{time.Date(2016, 3, 31, 23, 59, 59, 999999999, time.UTC), 57478.999988},
	}

	for _, test := range tests {
		got := ModifiedJulianDate(test.t)
		if math.Abs(got-test.want) > 0.000001 {
			t.Errorf("got: %f; want: %f (%s)", got, test.want, test.t)
		}
	}
}

func TestMaxTime(t *testing.T) {
	tests := []struct {
		t    []time.Time
		want time.Time
	}{
		{[]time.Time{}, time.Time{}},
		{[]time.Time{
			time.Date(2000, 1, 1, 12, 30, 0, 0, time.UTC),
			time.Date(2001, 1, 1, 12, 30, 0, 0, time.UTC),
			time.Date(2002, 1, 1, 12, 30, 0, 0, time.UTC)},
			time.Date(2002, 1, 1, 12, 30, 0, 0, time.UTC)},
		{[]time.Time{
			time.Date(2000, 1, 1, 12, 30, 0, 0, time.UTC),
			time.Date(2000, 1, 1, 10, 30, 0, 0, time.UTC),
			time.Date(2000, 1, 1, 8, 30, 0, 0, time.UTC)},
			time.Date(2000, 1, 1, 12, 30, 0, 0, time.UTC)},
	}

	for _, test := range tests {
		got := MaxTime(test.t...)
		if got != test.want {
			t.Errorf("got: %s; want: %s (%s)", got, test.want, test.t)
		}
	}
}

func TestMinTime(t *testing.T) {
	tests := []struct {
		t    []time.Time
		want time.Time
	}{
		{[]time.Time{}, time.Time{}},
		{[]time.Time{
			time.Date(2000, 1, 1, 12, 30, 0, 0, time.UTC),
			time.Date(2001, 1, 1, 12, 30, 0, 0, time.UTC),
			time.Date(2002, 1, 1, 12, 30, 0, 0, time.UTC)},
			time.Date(2000, 1, 1, 12, 30, 0, 0, time.UTC)},
		{[]time.Time{
			time.Date(2000, 1, 1, 12, 30, 0, 0, time.UTC),
			time.Date(2000, 1, 1, 10, 30, 0, 0, time.UTC),
			time.Date(2000, 1, 1, 8, 30, 0, 0, time.UTC)},
			time.Date(2000, 1, 1, 8, 30, 0, 0, time.UTC)},
	}

	for _, test := range tests {
		got := MinTime(test.t...)
		if got != test.want {
			t.Errorf("got: %s; want: %s (%s)", got, test.want, test.t)
		}
	}
}
