package cal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAustralianHolidays(t *testing.T) {
	c := NewCalendar()
	c.Observed = ObservedExact
	AddAustralianHolidays(c)
	loc, err := time.LoadLocation("Australia/Sydney")
	assert.NoError(t, err)

	currentTime := time.Now()
	year := currentTime.Year()

	tests := []testStruct{
		{time.Date(year, time.January, 1, 0, 0, 0, 0, loc), true, "New year"},
		{time.Date(2017, time.January, 2, 0, 0, 0, 0, loc), true, "New year Sunday"},
		{time.Date(year, time.January, 2, 0, 0, 0, 0, loc), false, "New year Sunday invalid"},
		{time.Date(year, time.December, 25, 0, 0, 0, 0, loc), true, "Christmas"},
		{time.Date(2010, time.December, 27, 0, 0, 0, 0, loc), true, "Christmas on Saturday"},
		{time.Date(2011, time.December, 27, 0, 0, 0, 0, loc), true, "Christmas on Sunday"},
		{time.Date(2018, time.December, 26, 0, 0, 0, 0, loc), true, "Boxing Day"},
		{time.Date(2019, time.April, 19, 0, 0, 0, 0, loc), true, "Good Friday"},
		{time.Date(2019, time.April, 22, 0, 0, 0, 0, loc), true, "Easter Monday"},
		{time.Date(2019, time.April, 25, 0, 0, 0, 0, loc), true, "Anzac Day"},
		{time.Date(2004, time.April, 25, 0, 0, 0, 0, loc), true, "Anzac Day"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := c.IsHoliday(test.t)
			assert.Equal(t, test.want, got)
		})
	}
}
