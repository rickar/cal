package cal

import (
	"testing"
	"time"
)

func TestCanadianHolidays(t *testing.T) {
	c := NewCalendar()
	AddCanadianHolidays(c)

	tests := []testStruct{
		{time.Date(2019, time.January, 1, 12, 0, 0, 0, time.UTC), true, "New Years"},
		{time.Date(2019, time.February, 18, 12, 0, 0, 0, time.UTC), true, "Family Day"},
		{time.Date(2019, time.April, 19, 12, 0, 0, 0, time.UTC), true, "Good Friday"},
		{time.Date(2019, time.May, 20, 12, 0, 0, 0, time.UTC), true, "Victoria Day"},
		{time.Date(2019, time.July, 1, 12, 0, 0, 0, time.UTC), true, "Canada Day"},
		{time.Date(2019, time.August, 5, 12, 0, 0, 0, time.UTC), true, "Civic Holiday"},
		{time.Date(2019, time.September, 2, 12, 0, 0, 0, time.UTC), true, "Labour Day"},
		{time.Date(2019, time.October, 14, 12, 0, 0, 0, time.UTC), true, "Thanksgiving"},
		{time.Date(2019, time.November, 11, 12, 0, 0, 0, time.UTC), true, "Remembrance Day"},
		{time.Date(2019, time.December, 25, 12, 0, 0, 0, time.UTC), true, "Christmas"},
		{time.Date(2019, time.December, 26, 12, 0, 0, 0, time.UTC), true, "Boxing Day"},
	}

	for _, test := range tests {
		got := c.IsHoliday(test.t)
		if got != test.want {
			t.Errorf("got: %t for %s; want: %t (%s)", got, test.name, test.want, test.t)
		}
	}
}
