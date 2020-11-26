// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package ie

import (
	"testing"
	"time"

	"github.com/rickar/cal/v2"
)

func d(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, cal.DefaultLoc)
}

func TestHolidays(t *testing.T) {
	tests := []struct {
		h    *cal.Holiday
		y    int
		date time.Time
	}{
		{NewYear, 2020, d(2020, 1, 1)},
		{NewYear, 2021, d(2021, 1, 1)},
		{NewYear, 2022, d(2022, 1, 1)},

		{SaintPatrickDay, 2020, d(2020, 3, 17)},
		{SaintPatrickDay, 2021, d(2021, 3, 17)},
		{SaintPatrickDay, 2022, d(2022, 3, 17)},

		{EasterMonday, 2020, d(2020, 4, 13)},
		{EasterMonday, 2021, d(2021, 4, 5)},
		{EasterMonday, 2022, d(2022, 4, 18)},

		{FirstMondayMay, 2020, d(2020, 5, 4)},
		{FirstMondayMay, 2021, d(2021, 5, 3)},
		{FirstMondayMay, 2022, d(2022, 5, 2)},
		{FirstMondayMay, 2023, d(2023, 5, 1)},

		{FirstMondayJune, 2020, d(2020, 6, 1)},
		{FirstMondayJune, 2021, d(2021, 6, 7)},
		{FirstMondayJune, 2022, d(2022, 6, 6)},
		{FirstMondayJune, 2023, d(2023, 6, 5)},

		{FirstMondayAugust, 2020, d(2020, 8, 3)},
		{FirstMondayAugust, 2021, d(2021, 8, 2)},
		{FirstMondayAugust, 2022, d(2022, 8, 1)},
		{FirstMondayAugust, 2023, d(2023, 8, 7)},

		{LastMondayInOctober, 2020, d(2020, 10, 26)},
		{LastMondayInOctober, 2021, d(2021, 10, 25)},
		{LastMondayInOctober, 2022, d(2022, 10, 31)},
		{LastMondayInOctober, 2023, d(2023, 10, 30)},

		{ChristmasDay, 2020, d(2020, 12, 25)},
		{ChristmasDay, 2021, d(2021, 12, 25)},
		{ChristmasDay, 2022, d(2022, 12, 25)},

		{SaintStephenDay, 2020, d(2020, 12, 26)},
		{SaintStephenDay, 2021, d(2021, 12, 26)},
		{SaintStephenDay, 2022, d(2022, 12, 26)},
	}

	for _, test := range tests {
		gotAct, _ := test.h.Calc(test.y)
		if !gotAct.Equal(test.date) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.date.String())
		}
	}
}
