// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package gb

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
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{NewYear, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{NewYear, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{NewYear, 2017, d(2017, 1, 1), d(2017, 1, 2)},
		{NewYear, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{NewYear, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{NewYear, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{NewYear, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{NewYear, 2022, d(2022, 1, 1), d(2022, 1, 3)},

		{GoodFriday, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{GoodFriday, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{GoodFriday, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{GoodFriday, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{GoodFriday, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{GoodFriday, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{GoodFriday, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{GoodFriday, 2022, d(2022, 4, 15), d(2022, 4, 15)},

		{EasterMonday, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{EasterMonday, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{EasterMonday, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{EasterMonday, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{EasterMonday, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{EasterMonday, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{EasterMonday, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{EasterMonday, 2022, d(2022, 4, 18), d(2022, 4, 18)},

		{EarlyMay, 2015, d(2015, 5, 4), d(2015, 5, 4)},
		{EarlyMay, 2016, d(2016, 5, 2), d(2016, 5, 2)},
		{EarlyMay, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{EarlyMay, 2018, d(2018, 5, 7), d(2018, 5, 7)},
		{EarlyMay, 2019, d(2019, 5, 6), d(2019, 5, 6)},
		{EarlyMay, 2020, time.Time{}, time.Time{}},
		{EarlyMay, 2021, d(2021, 5, 3), d(2021, 5, 3)},
		{EarlyMay, 2022, d(2022, 5, 2), d(2022, 5, 2)},

		{VEDay, 2015, time.Time{}, time.Time{}},
		{VEDay, 2016, time.Time{}, time.Time{}},
		{VEDay, 2017, time.Time{}, time.Time{}},
		{VEDay, 2018, time.Time{}, time.Time{}},
		{VEDay, 2019, time.Time{}, time.Time{}},
		{VEDay, 2020, d(2020, 5, 8), d(2020, 5, 8)},
		{VEDay, 2021, time.Time{}, time.Time{}},
		{VEDay, 2022, time.Time{}, time.Time{}},

		{SpringHoliday, 2015, d(2015, 5, 25), d(2015, 5, 25)},
		{SpringHoliday, 2016, d(2016, 5, 30), d(2016, 5, 30)},
		{SpringHoliday, 2017, d(2017, 5, 29), d(2017, 5, 29)},
		{SpringHoliday, 2018, d(2018, 5, 28), d(2018, 5, 28)},
		{SpringHoliday, 2019, d(2019, 5, 27), d(2019, 5, 27)},
		{SpringHoliday, 2020, d(2020, 5, 25), d(2020, 5, 25)},
		{SpringHoliday, 2021, d(2021, 5, 31), d(2021, 5, 31)},
		{SpringHoliday, 2022, time.Time{}, time.Time{}},
		{SpringHoliday, 2023, d(2023, 5, 29), d(2023, 5, 29)},

		{SpringHoliday2022, 2021, time.Time{}, time.Time{}},
		{SpringHoliday2022, 2022, d(2022, 6, 2), d(2022, 6, 2)},
		{SpringHoliday2022, 2023, time.Time{}, time.Time{}},

		{PlatinumJubilee, 2021, time.Time{}, time.Time{}},
		{PlatinumJubilee, 2022, d(2022, 6, 3), d(2022, 6, 3)},
		{PlatinumJubilee, 2023, time.Time{}, time.Time{}},

		{SummerHolidayScotland, 2015, d(2015, 8, 3), d(2015, 8, 3)},
		{SummerHolidayScotland, 2016, d(2016, 8, 1), d(2016, 8, 1)},
		{SummerHolidayScotland, 2017, d(2017, 8, 7), d(2017, 8, 7)},
		{SummerHolidayScotland, 2018, d(2018, 8, 6), d(2018, 8, 6)},
		{SummerHolidayScotland, 2019, d(2019, 8, 5), d(2019, 8, 5)},
		{SummerHolidayScotland, 2020, d(2020, 8, 3), d(2020, 8, 3)},
		{SummerHolidayScotland, 2021, d(2021, 8, 2), d(2021, 8, 2)},
		{SummerHolidayScotland, 2022, d(2022, 8, 1), d(2022, 8, 1)},

		{SummerHoliday, 2015, d(2015, 8, 31), d(2015, 8, 31)},
		{SummerHoliday, 2016, d(2016, 8, 29), d(2016, 8, 29)},
		{SummerHoliday, 2017, d(2017, 8, 28), d(2017, 8, 28)},
		{SummerHoliday, 2018, d(2018, 8, 27), d(2018, 8, 27)},
		{SummerHoliday, 2019, d(2019, 8, 26), d(2019, 8, 26)},
		{SummerHoliday, 2020, d(2020, 8, 31), d(2020, 8, 31)},
		{SummerHoliday, 2021, d(2021, 8, 30), d(2021, 8, 30)},
		{SummerHoliday, 2022, d(2022, 8, 29), d(2022, 8, 29)},

		{ChristmasDay, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{ChristmasDay, 2016, d(2016, 12, 25), d(2016, 12, 26)},
		{ChristmasDay, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{ChristmasDay, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{ChristmasDay, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{ChristmasDay, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{ChristmasDay, 2021, d(2021, 12, 25), d(2021, 12, 27)},
		{ChristmasDay, 2022, d(2022, 12, 25), d(2022, 12, 26)},

		{BoxingDay, 2015, d(2015, 12, 26), d(2015, 12, 28)},
		{BoxingDay, 2016, d(2016, 12, 26), d(2016, 12, 27)},
		{BoxingDay, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{BoxingDay, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{BoxingDay, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{BoxingDay, 2020, d(2020, 12, 26), d(2020, 12, 28)},
		{BoxingDay, 2021, d(2021, 12, 26), d(2021, 12, 28)},
		{BoxingDay, 2022, d(2022, 12, 26), d(2022, 12, 27)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}
