package cal

import (
	"testing"
	"time"
)

func TestUkraineHolidays(t *testing.T) {
	c := NewCalendar()
	c.Observed = ObservedExact
	AddUkraineHolidays(c)

	tests := []testStruct{
		{time.Date(2019, 1, 1, 12, 0, 0, 0, time.UTC), true, "New Year in 2019"},
		{time.Date(2011, 1, 1, 12, 0, 0, 0, time.UTC), true, "New Year in 2011"},
		{time.Date(2011, 1, 3, 12, 0, 0, 0, time.UTC), true, "New Year at weekend in 2011"},
		{time.Date(2019, 1, 7, 12, 0, 0, 0, time.UTC), true, "Orthodox Christmas in 2019"},
		{time.Date(2019, 3, 8, 12, 0, 0, 0, time.UTC), true, "Women day in 2019"},
		{time.Date(2014, 3, 10, 12, 0, 0, 0, time.UTC), true, "Women day at weekend in 2014"},
		{time.Date(2019, 5, 1, 12, 0, 0, 0, time.UTC), true, "Labour day 1 in 2019"},
		{time.Date(2010, 5, 1, 12, 0, 0, 0, time.UTC), true, "Labour day 1 in 2010"},
		{time.Date(2010, 5, 2, 12, 0, 0, 0, time.UTC), true, "Labour day 2 in 2010"},
		{time.Date(2010, 5, 3, 12, 0, 0, 0, time.UTC), true, "Labour day 1 at weekend in 2010"},
		{time.Date(2011, 5, 1, 12, 0, 0, 0, time.UTC), true, "Labour day 1 in 2011"},
		{time.Date(2011, 5, 2, 12, 0, 0, 0, time.UTC), true, "Labour day 2 in 2011"},
		{time.Date(2019, 5, 2, 12, 0, 0, 0, time.UTC), false, "Labour day 2 false test in 2019"},
		{time.Date(2017, 5, 2, 12, 0, 0, 0, time.UTC), true, "Labour day 2 in 2017"},
		{time.Date(2019, 5, 9, 12, 0, 0, 0, time.UTC), true, "Victory day in 2019"},
		{time.Date(2019, 6, 28, 12, 0, 0, 0, time.UTC), true, "Constitution day in 2019"},
		{time.Date(2019, 8, 24, 12, 0, 0, 0, time.UTC), true, "Independence day in 2019"},
		{time.Date(2019, 8, 26, 12, 0, 0, 0, time.UTC), true, "Independence day at weekend in 2019"},
		{time.Date(2019, 10, 14, 12, 0, 0, 0, time.UTC), true, "Defender Of Ukraine day in 2019"},
		{time.Date(2010, 10, 14, 12, 0, 0, 0, time.UTC), false, "Defender of Ukraine day false test in 2010"},
		{time.Date(2019, 12, 25, 12, 0, 0, 0, time.UTC), true, "Catholic Christmas day in 2019"},
		{time.Date(2017, 12, 25, 12, 0, 0, 0, time.UTC), false, "Catholic Christmas day false test in 2017"},
		{time.Date(2018, 4, 8, 12, 0, 0, 0, time.UTC), true, "Orthodox Easter day 1 in 2018"},
		{time.Date(2018, 4, 9, 12, 0, 0, 0, time.UTC), true, "Orthodox Easter day 2 in 2018"},
		{time.Date(2014, 4, 20, 12, 0, 0, 0, time.UTC), true, "Orthodox Easter day 1 in 2014"},
		{time.Date(2014, 4, 21, 12, 0, 0, 0, time.UTC), true, "Orthodox Easter day 2 in 2014"},
		{time.Date(2019, 4, 28, 12, 0, 0, 0, time.UTC), true, "Orthodox Easter day 1 in 2019"},
		{time.Date(2019, 4, 29, 12, 0, 0, 0, time.UTC), true, "Orthodox Easter day 2 in 2019"},
		{time.Date(2018, 5, 27, 12, 0, 0, 0, time.UTC), true, "Orthodox Pentecost day 1 in 2018"},
		{time.Date(2018, 5, 28, 12, 0, 0, 0, time.UTC), true, "Orthodox Pentecost day 2 in 2018"},
		{time.Date(2019, 6, 16, 12, 0, 0, 0, time.UTC), true, "Orthodox Pentecost day 1 in 2019"},
		{time.Date(2019, 6, 17, 12, 0, 0, 0, time.UTC), true, "Orthodox Pentecost day 2 in 2019"},
	}

	for _, test := range tests {
		got := c.IsHoliday(test.t)
		if got != test.want {
			t.Errorf("got: %t for %s; want: %t (%s)", got, test.name, test.want, test.t)
		}
	}
}
