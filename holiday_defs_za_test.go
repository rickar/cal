package cal

import (
	"reflect"
	"testing"
	"time"
)

func TestSouthAfricanHolidays(t *testing.T) {
	c := NewCalendar()
	c.Observed = ObservedMondayTuesday
	AddSouthAfricanHolidays(c)
	loc, err := time.LoadLocation("Africa/Johannesburg")
	if err != nil {
		t.Fatal(err)
	}

	isHolidayTests := []testStruct{
		{time.Date(2020, time.January, 1, 0, 0, 0, 0, loc), true, "New year"},
		{time.Date(2020, time.March, 21, 0, 0, 0, 0, loc), true, "Human Rights"},
		{time.Date(2020, time.April, 10, 0, 0, 0, 0, loc), true, "Good Friday"},
		{time.Date(2020, time.April, 13, 0, 0, 0, 0, loc), true, "Easter Monday"},
		{time.Date(2020, time.April, 27, 0, 0, 0, 0, loc), true, "Freedom"},
		{time.Date(2020, time.May, 1, 0, 0, 0, 0, loc), true, "Workers"},
		{time.Date(2020, time.June, 16, 0, 0, 0, 0, loc), true, "Youth"},
		{time.Date(2020, time.August, 9, 0, 0, 0, 0, loc), true, "National Women's"},
		{time.Date(2020, time.August, 10, 0, 0, 0, 0, loc), false, "National Women's Observed"},
		{time.Date(2020, time.September, 24, 0, 0, 0, 0, loc), true, "Heritage"},
		{time.Date(2020, time.December, 16, 0, 0, 0, 0, loc), true, "Day of Reconciliation"},
		{time.Date(2020, time.December, 25, 0, 0, 0, 0, loc), true, "Christmas"},
		{time.Date(2020, time.December, 26, 0, 0, 0, 0, loc), true, "Day of Goodwill"},
	}

	t.Run("isHoliday", func(t *testing.T) {
		for _, tc := range isHolidayTests {
			t.Run(tc.name, func(t *testing.T) {
				if got := c.IsHoliday(tc.t); !reflect.DeepEqual(got, tc.want) {
					t.Errorf("Calendar.AddSkipNonWorkdays() = %v, want %v", got, tc.want)
				}
			})
		}
	})

	isObservedTests := []testStruct{
		{time.Date(2020, time.January, 1, 0, 0, 0, 0, loc), true, "New year"},
		{time.Date(2020, time.March, 21, 0, 0, 0, 0, loc), true, "Human Rights"},
		{time.Date(2020, time.April, 10, 0, 0, 0, 0, loc), true, "Good Friday"},
		{time.Date(2020, time.April, 13, 0, 0, 0, 0, loc), true, "Easter Monday"},
		{time.Date(2020, time.April, 27, 0, 0, 0, 0, loc), true, "Freedom"},
		{time.Date(2020, time.May, 1, 0, 0, 0, 0, loc), true, "Workers"},
		{time.Date(2020, time.June, 16, 0, 0, 0, 0, loc), true, "Youth"},
		{time.Date(2020, time.August, 10, 0, 0, 0, 0, loc), true, "National Women's Observed"},
		{time.Date(2020, time.September, 24, 0, 0, 0, 0, loc), true, "Heritage"},
		{time.Date(2020, time.December, 16, 0, 0, 0, 0, loc), true, "Day of Reconciliation"},
		{time.Date(2020, time.December, 25, 0, 0, 0, 0, loc), true, "Christmas"},
		{time.Date(2020, time.December, 26, 0, 0, 0, 0, loc), true, "Day of Goodwill"},
	}

	t.Run("isObserved", func(t *testing.T) {
		for _, tc := range isObservedTests {
			t.Run(tc.name, func(t *testing.T) {
				if got := !c.IsWorkday(tc.t); !reflect.DeepEqual(got, tc.want) {
					t.Errorf("Calendar.AddSkipNonWorkdays() = %v, want %v", got, tc.want)
				}
			})
		}
	})

}
