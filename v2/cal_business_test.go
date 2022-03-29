// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"testing"
	"time"
)

func d(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}

func dt(y, m, d, h, min int) time.Time {
	return time.Date(y, time.Month(m), d, h, min, 0, 0, time.UTC)
}

func TestNewBusinessCalendar(t *testing.T) {
	b := NewBusinessCalendar()
	if b.workday[time.Sunday] || !b.workday[time.Monday] || !b.workday[time.Tuesday] || !b.workday[time.Wednesday] ||
		!b.workday[time.Thursday] || !b.workday[time.Friday] || b.workday[time.Saturday] {
		t.Errorf("invalid work days: want M-F, got: %v", b.workday)
	}

	if b.workdayStart != 9*time.Hour || b.workdayEnd != 17*time.Hour {
		t.Errorf("invalid work hours; want 9am-5pm; got: %d, %d", b.workdayStart, b.workdayEnd)
	}
}

func TestSetWorkday(t *testing.T) {
	b := NewBusinessCalendar()
	b.SetWorkday(time.Saturday, true)
	b.SetWorkday(time.Wednesday, false)
	if !b.workday[time.Saturday] || b.workday[time.Wednesday] {
		t.Errorf("invalid work days: want Sat, Wed, got: %v", b.workday)
	}
}

func TestSetWorkingHours(t *testing.T) {
	b := NewBusinessCalendar()
	b.SetWorkHours(8*time.Hour+30*time.Minute, 18*time.Hour+15*time.Minute)
	if b.workdayStart != 8*time.Hour+30*time.Minute || b.workdayEnd != 18*time.Hour+15*time.Minute {
		t.Errorf("invalid work hours; want 9am-5pm; got: %d, %d", b.workdayStart, b.workdayEnd)
	}
}

func TestIsWorkday(t *testing.T) {
	cal1 := NewBusinessCalendar()
	cal2 := NewBusinessCalendar()
	cal2.WorkdayFunc = func(date time.Time) bool {
		return date.Day() >= 1 && date.Day() <= 15
	}
	hol := &Holiday{
		Type:  ObservancePublic,
		Month: time.July,
		Day:   4,
		Observed: []AltDay{
			{Day: time.Saturday, Offset: -1},
			{Day: time.Sunday, Offset: 1},
		},
		Func: CalcDayOfMonth,
	}
	cal2.AddHoliday(hol)

	tests := []struct {
		c    *BusinessCalendar
		d    time.Time
		want bool
	}{
		{cal1, d(2020, 4, 1), true},
		{cal1, d(2020, 4, 3), true},
		{cal1, d(2020, 4, 4), false},
		{cal1, d(2020, 4, 5), false},
		{cal1, d(2020, 4, 6), true},

		{cal2, d(2020, 4, 1), true},
		{cal2, d(2020, 4, 3), true},
		{cal2, d(2020, 4, 4), true},
		{cal2, d(2020, 4, 16), false},
		{cal2, d(2020, 4, 30), false},

		{cal2, d(2015, 7, 3), false},
		{cal2, d(2015, 7, 4), true},
		{cal2, d(2016, 7, 3), true},
		{cal2, d(2016, 7, 4), false},
	}

	for i, test := range tests {
		got := test.c.IsWorkday(test.d)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%d)", got, test.want, i)
		}
	}
}

func TestIsWorkTime(t *testing.T) {
	cal1 := NewBusinessCalendar()
	cal2 := NewBusinessCalendar()
	cal2.WorkdayStartFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12, 30, 0, 0, time.UTC)
	}
	cal2.WorkdayEndFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12+6, 45, 0, 0, time.UTC)
	}

	tests := []struct {
		c    *BusinessCalendar
		d    time.Time
		want bool
	}{
		{cal1, dt(2020, 4, 1, 12, 00), true},
		{cal1, dt(2020, 4, 1, 0, 00), false},
		{cal1, dt(2020, 4, 1, 18, 00), false},
		{cal1, dt(2020, 4, 5, 12, 00), false},

		{cal2, dt(2020, 4, 1, 3, 00), true},
		{cal2, dt(2020, 4, 1, 0, 00), false},
		{cal2, dt(2020, 4, 1, 7, 00), true},
		{cal2, dt(2020, 4, 1, 7, 50), false},
	}

	for i, test := range tests {
		got := test.c.IsWorkTime(test.d)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%d)", got, test.want, i)
		}
	}
}

func TestWorkdaysRemain(t *testing.T) {
	c := NewBusinessCalendar()
	hol := &Holiday{
		Type:  ObservancePublic,
		Month: time.July,
		Day:   4,
		Observed: []AltDay{
			{Day: time.Saturday, Offset: -1},
			{Day: time.Sunday, Offset: 1},
		},
		Func: CalcDayOfMonth,
	}
	c.AddHoliday(hol)

	tests := []struct {
		t    time.Time
		want int
	}{
		{d(2020, 4, 1), 21},
		{d(2020, 4, 3), 19},
		{d(2020, 4, 4), 19},
		{d(2020, 4, 5), 19},
		{d(2020, 4, 6), 18},

		{d(2015, 7, 1), 21},
		{d(2015, 7, 2), 20},
		{d(2015, 7, 3), 20},
		{d(2015, 7, 4), 20},
		{d(2015, 7, 5), 20},
		{d(2015, 7, 6), 19},
	}

	for _, test := range tests {
		got := c.WorkdaysRemain(test.t)
		if got != test.want {
			t.Errorf("got: %d; want: %d (%s)", got, test.want, test.t)
		}
	}
}

func TestWorkdaysInMonth(t *testing.T) {
	c := NewBusinessCalendar()
	hol := &Holiday{
		Type:  ObservancePublic,
		Month: time.July,
		Day:   4,
		Observed: []AltDay{
			{Day: time.Saturday, Offset: -1},
			{Day: time.Sunday, Offset: 1},
		},
		Func: CalcDayOfMonth,
	}
	c.AddHoliday(hol)

	tests := []struct {
		y    int
		m    time.Month
		want int
	}{
		{2020, time.April, 22},
		{2020, time.May, 21},
		{2020, time.June, 22},
		{2020, time.July, 22},
		{2020, time.November, 21},
	}

	for _, test := range tests {
		got := c.WorkdaysInMonth(test.y, test.m)
		if got != test.want {
			t.Errorf("got: %d; want: %d (%d-%d)", got, test.want, test.y, test.m)
		}
	}
}

func TestWorkdaysInRange(t *testing.T) {
	c := NewBusinessCalendar()
	hol := &Holiday{
		Type:  ObservancePublic,
		Month: time.July,
		Day:   4,
		Observed: []AltDay{
			{Day: time.Saturday, Offset: -1},
			{Day: time.Sunday, Offset: 1},
		},
		Func: CalcDayOfMonth,
	}
	c.AddHoliday(hol)

	tests := []struct {
		f    time.Time
		t    time.Time
		want int
	}{
		{d(2015, 4, 1), d(2015, 4, 10), 8},
		{d(2015, 4, 1), d(2015, 4, 30), 22},
		{d(2015, 4, 1), d(2015, 5, 16), 33},
		{d(2015, 4, 1), d(2015, 4, 1), 1},
		{d(2015, 4, 4), d(2015, 4, 5), 0},
		{d(2015, 7, 1), d(2015, 7, 6), 3},
	}

	for _, test := range tests {
		got := c.WorkdaysInRange(test.f, test.t)
		if got != test.want {
			t.Errorf("got: %d; want: %d (%s-%s)", got, test.want, test.f, test.t)
		}
		if !test.f.Equal(test.t) {
			got = c.WorkdaysInRange(test.t, test.f)
			if got != -test.want {
				t.Errorf("got: %d; want: %d (%s-%s)", got, -test.want, test.t, test.f)
			}
		}
	}
}

func TestHolidaysInRange(t *testing.T) {
	c := NewBusinessCalendar()
	hol := &Holiday{
		Type:  ObservancePublic,
		Month: time.July,
		Day:   4,
		Observed: []AltDay{
			{Day: time.Saturday, Offset: -1},
			{Day: time.Sunday, Offset: 1},
		},
		Func: CalcDayOfMonth,
	}
	c.AddHoliday(hol)

	tests := []struct {
		f    time.Time
		t    time.Time
		want int
	}{
		{d(2015, 4, 4), d(2015, 4, 5), 0},
		{d(2015, 7, 1), d(2015, 7, 6), 1},
	}

	for _, test := range tests {
		got := c.HolidaysInRange(test.f, test.t)
		if got != test.want {
			t.Errorf("got: %d; want: %d (%s-%s)", got, test.want, test.f, test.t)
		}
		if !test.f.Equal(test.t) {
			got = c.HolidaysInRange(test.t, test.f)
			if got != -test.want {
				t.Errorf("got: %d; want: %d (%s-%s)", got, -test.want, test.t, test.f)
			}
		}
	}
}

func TestWorkdayN(t *testing.T) {
	c := NewBusinessCalendar()
	hol := &Holiday{
		Type:  ObservancePublic,
		Month: time.July,
		Day:   4,
		Observed: []AltDay{
			{Day: time.Saturday, Offset: -1},
			{Day: time.Sunday, Offset: 1},
		},
		Func: CalcDayOfMonth,
	}
	c.AddHoliday(hol)

	tests := []struct {
		y    int
		m    time.Month
		n    int
		want int
	}{
		{2016, time.January, 14, 20},
		{2016, time.January, -5, 25},
		{2016, time.January, -14, 12},
		{2016, time.February, 21, 29},
		{2016, time.February, 22, 0},
		{2016, time.February, -1, 29},
		{2016, time.February, 0, 0},
		{2016, time.July, 4, 7},
	}

	for _, test := range tests {
		got := c.WorkdayN(test.y, test.m, test.n)
		if got != test.want {
			t.Errorf("got: %d; want: %d (%d %d %d)", got, test.want, test.y, test.m, test.n)
		}
	}
}

func TestWorkdaysFrom(t *testing.T) {
	c := NewBusinessCalendar()
	hol := &Holiday{
		Type:  ObservancePublic,
		Month: time.July,
		Day:   4,
		Observed: []AltDay{
			{Day: time.Saturday, Offset: -1},
			{Day: time.Sunday, Offset: 1},
		},
		Func: CalcDayOfMonth,
	}
	c.AddHoliday(hol)

	tests := []struct {
		d    time.Time
		o    int
		want time.Time
	}{
		{d(2016, 1, 5), 0, d(2016, 1, 5)},
		{d(2016, 1, 5), 1, d(2016, 1, 6)},
		{d(2016, 1, 5), -1, d(2016, 1, 4)},
		{d(2016, 1, 15), 1, d(2016, 1, 18)},
		{d(2016, 1, 15), -12, d(2015, 12, 30)},

		{d(2016, 7, 1), 1, d(2016, 7, 5)},
		{d(2016, 7, 4), 1, d(2016, 7, 5)},
		{d(2016, 7, 4), -1, d(2016, 7, 1)},
		{d(2016, 1, 1), 366, d(2017, 5, 30)},
		{d(2016, 1, 1), -366, d(2014, 8, 6)},
	}

	for _, test := range tests {
		got := c.WorkdaysFrom(test.d, test.o)
		if got != test.want {
			t.Errorf("got: %s; want: %s (%s %d)", got, test.want, test.d, test.o)
		}
	}
}

func TestWorkingHours(t *testing.T) {
	cal1 := NewBusinessCalendar()
	cal2 := NewBusinessCalendar()
	cal2.WorkdayStartFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12, 30, 0, 0, time.UTC)
	}
	cal2.WorkdayEndFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12+6, 45, 0, 0, time.UTC)
	}

	tests := []struct {
		c    *BusinessCalendar
		d    time.Time
		want time.Duration
	}{
		{cal1, d(2020, 4, 1), 8 * time.Hour},
		{cal1, d(2020, 4, 5), 0 * time.Hour},

		{cal2, d(2020, 4, 1), 6*time.Hour + 15*time.Minute},
		{cal2, d(2020, 4, 6), 6*time.Hour + 15*time.Minute},
	}

	for i, test := range tests {
		got := test.c.WorkHours(test.d)
		if got != test.want {
			t.Errorf("got: %d; want: %d (%d)", got/time.Hour, test.want/time.Hour, i)
		}
	}
}

func TestWorkdayStart(t *testing.T) {
	cal1 := NewBusinessCalendar()
	cal2 := NewBusinessCalendar()
	cal2.WorkdayStartFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12, 30, 0, 0, time.UTC)
	}
	cal2.WorkdayEndFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12+6, 45, 0, 0, time.UTC)
	}

	tests := []struct {
		c    *BusinessCalendar
		d    time.Time
		want time.Time
	}{
		{cal1, d(2020, 4, 1), dt(2020, 4, 1, 9, 0)},
		{cal1, d(2020, 4, 5), time.Time{}},

		{cal2, d(2020, 4, 1), dt(2020, 4, 1, 1, 30)},
		{cal2, d(2020, 4, 6), dt(2020, 4, 6, 6, 30)},
	}

	for i, test := range tests {
		got := test.c.WorkdayStart(test.d)
		if got != test.want {
			t.Errorf("got: %s; want: %s (%d)", got, test.want, i)
		}
	}
}

func TestWorkdayEnd(t *testing.T) {
	cal1 := NewBusinessCalendar()
	cal2 := NewBusinessCalendar()
	cal2.WorkdayStartFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12, 30, 0, 0, time.UTC)
	}
	cal2.WorkdayEndFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12+6, 45, 0, 0, time.UTC)
	}

	tests := []struct {
		c    *BusinessCalendar
		d    time.Time
		want time.Time
	}{
		{cal1, d(2020, 4, 1), dt(2020, 4, 1, 17, 0)},
		{cal1, d(2020, 4, 5), time.Time{}},

		{cal2, d(2020, 4, 1), dt(2020, 4, 1, 7, 45)},
		{cal2, d(2020, 4, 6), dt(2020, 4, 6, 12, 45)},
	}

	for i, test := range tests {
		got := test.c.WorkdayEnd(test.d)
		if got != test.want {
			t.Errorf("got: %s; want: %s (%d)", got, test.want, i)
		}
	}
}

func TestNextWorkdayStart(t *testing.T) {
	cal1 := NewBusinessCalendar()
	cal2 := NewBusinessCalendar()
	cal2.WorkdayStartFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12, 30, 0, 0, time.UTC)
	}
	cal2.WorkdayEndFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12+6, 45, 0, 0, time.UTC)
	}

	tests := []struct {
		c    *BusinessCalendar
		d    time.Time
		want time.Time
	}{
		{cal1, dt(2020, 4, 1, 6, 0), dt(2020, 4, 1, 9, 0)},
		{cal1, dt(2020, 4, 1, 20, 0), dt(2020, 4, 2, 9, 0)},
		{cal1, dt(2020, 4, 4, 12, 0), dt(2020, 4, 6, 9, 0)},
		{cal1, dt(2020, 4, 5, 12, 0), dt(2020, 4, 6, 9, 0)},

		{cal2, dt(2020, 4, 1, 3, 0), dt(2020, 4, 2, 2, 30)},
		{cal2, dt(2020, 4, 6, 8, 0), dt(2020, 4, 7, 7, 30)},
	}

	for i, test := range tests {
		got := test.c.NextWorkdayStart(test.d)
		if got != test.want {
			t.Errorf("got: %s; want: %s (%d)", got, test.want, i)
		}
	}
}

func TestNextWorkdayEnd(t *testing.T) {
	cal1 := NewBusinessCalendar()
	cal2 := NewBusinessCalendar()
	cal2.WorkdayStartFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12, 30, 0, 0, time.UTC)
	}
	cal2.WorkdayEndFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12+6, 45, 0, 0, time.UTC)
	}

	tests := []struct {
		c    *BusinessCalendar
		d    time.Time
		want time.Time
	}{
		{cal1, dt(2020, 4, 1, 6, 0), dt(2020, 4, 1, 17, 0)},
		{cal1, dt(2020, 4, 1, 20, 0), dt(2020, 4, 2, 17, 0)},
		{cal1, dt(2020, 4, 4, 12, 0), dt(2020, 4, 6, 17, 0)},
		{cal1, dt(2020, 4, 5, 12, 0), dt(2020, 4, 6, 17, 0)},

		{cal2, dt(2020, 4, 1, 3, 0), dt(2020, 4, 1, 7, 45)},
		{cal2, dt(2020, 4, 6, 8, 0), dt(2020, 4, 6, 12, 45)},
		{cal2, dt(2020, 4, 6, 22, 0), dt(2020, 4, 7, 13, 45)},
	}

	for i, test := range tests {
		got := test.c.NextWorkdayEnd(test.d)
		if got != test.want {
			t.Errorf("got: %s; want: %s (%d)", got, test.want, i)
		}
	}
}

func TestWorkHoursInRange(t *testing.T) {
	cal1 := NewBusinessCalendar()
	cal2 := NewBusinessCalendar()
	cal2.WorkdayStartFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12, 30, 0, 0, time.UTC)
	}
	cal2.WorkdayEndFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12+6, 45, 0, 0, time.UTC)
	}

	tests := []struct {
		c    *BusinessCalendar
		f    time.Time
		t    time.Time
		want time.Duration
	}{
		{cal1, dt(2020, 4, 1, 6, 0), dt(2020, 4, 1, 23, 0), 8 * time.Hour},
		{cal1, dt(2020, 4, 1, 18, 0), dt(2020, 4, 2, 12, 0), 3 * time.Hour},
		{cal1, dt(2020, 4, 1, 14, 0), dt(2020, 4, 1, 18, 0), 3 * time.Hour},
		{cal1, dt(2020, 4, 1, 9, 15), dt(2020, 4, 2, 12, 0), 10*time.Hour + 45*time.Minute},
		{cal1, dt(2020, 4, 4, 9, 0), dt(2020, 4, 6, 8, 0), 0 * time.Hour},

		{cal2, dt(2020, 4, 1, 1, 0), dt(2020, 4, 2, 5, 00), 8*time.Hour + 45*time.Minute},
		{cal2, dt(2020, 4, 13, 0, 0), dt(2020, 4, 18, 0, 0), 5*6*time.Hour + 5*15*time.Minute},
	}

	for i, test := range tests {
		got := test.c.WorkHoursInRange(test.f, test.t)
		if got != test.want {
			t.Errorf("got: %s; want: %s (%d)", got, test.want, i)
		}

		got = test.c.WorkHoursInRange(test.t, test.f)
		if got != test.want {
			t.Errorf("got: %s; want: %s (%d flipped)", got, test.want, i)
		}
	}
}

func TestAddWorkHours(t *testing.T) {
	cal1 := NewBusinessCalendar()
	cal2 := NewBusinessCalendar()
	cal2.WorkdayStartFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12, 30, 0, 0, time.UTC)
	}
	cal2.WorkdayEndFunc = func(date time.Time) time.Time {
		return time.Date(date.Year(), date.Month(), date.Day(), date.Day()%12+6, 45, 0, 0, time.UTC)
	}

	tests := []struct {
		c    *BusinessCalendar
		t    time.Time
		d    time.Duration
		want time.Time
	}{
		{cal1, dt(2020, 4, 1, 6, 0), 8 * time.Hour, dt(2020, 4, 1, 17, 0)},
		{cal1, dt(2020, 4, 1, 18, 0), 3 * time.Hour, dt(2020, 4, 2, 12, 0)},
		{cal1, dt(2020, 4, 1, 12, 0), 3*time.Hour + 20*time.Minute, dt(2020, 4, 1, 15, 20)},
		{cal1, dt(2020, 4, 1, 9, 15), 10*time.Hour + 45*time.Minute, dt(2020, 4, 2, 12, 0)},
		{cal1, dt(2020, 4, 4, 9, 0), 0 * time.Hour, dt(2020, 4, 4, 9, 0)},
		{cal1, dt(2020, 4, 4, 9, 0), 24 * time.Hour, dt(2020, 4, 8, 17, 0)},

		{cal2, dt(2020, 4, 1, 1, 0), 8*time.Hour + 45*time.Minute, dt(2020, 4, 2, 5, 00)},
		{cal2, dt(2020, 4, 13, 0, 0), 5*6*time.Hour + 5*15*time.Minute, dt(2020, 4, 17, 11, 45)},
	}

	for i, test := range tests {
		got := test.c.AddWorkHours(test.t, test.d)
		if got != test.want {
			t.Errorf("got: %s; want: %s (%d)", got, test.want, i)
		}
	}
}

func TestWorkDurationInRange(t *testing.T) {
	c := NewBusinessCalendar()
	hol := &Holiday{
		Type:  ObservancePublic,
		Month: time.July,
		Day:   4,
		Observed: []AltDay{
			{Day: time.Saturday, Offset: -1},
			{Day: time.Sunday, Offset: 1},
		},
		Func: CalcDayOfMonth,
	}
	c.AddHoliday(hol)
	days := func(days time.Duration) time.Duration {
		return days * 24 * time.Hour
	}
	tests := []struct {
		f    time.Time
		t    time.Time
		want time.Duration
	}{
		{d(2015, 4, 1), d(2015, 4, 10), days(7)},
		{d(2015, 4, 1), d(2015, 4, 30), days(21)},
		{d(2015, 4, 1), d(2015, 5, 16), days(33)},
		{d(2015, 4, 1), d(2015, 4, 1), days(1)},
		{d(2015, 4, 4), d(2015, 4, 5), days(0)},
		{d(2015, 7, 1), d(2015, 7, 6), days(2)},
	}

	for _, test := range tests {
		got := c.WorkDurationInRange(test.f, test.t)
		if got != test.want {
			t.Errorf("got: %s; want: %s (%s-%s)", got, test.want, test.f, test.t)
		}
		if !test.f.Equal(test.t) {
			got = c.WorkDurationInRange(test.t, test.f)
			if got != -test.want {
				t.Errorf("got: %s; want: %s (%s-%s)", got, -test.want, test.t, test.f)
			}
		}
	}
}
