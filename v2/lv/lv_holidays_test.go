// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package lv

import (
	"fmt"
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
		{NewYear, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{NewYear, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{NewYear, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{NewYear, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{NewYear, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{NewYear, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{GoodFriday, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{GoodFriday, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{GoodFriday, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{GoodFriday, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{GoodFriday, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{GoodFriday, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{GoodFriday, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{GoodFriday, 2022, d(2022, 4, 15), d(2022, 4, 15)},
		{GoodFriday, 2023, d(2023, 4, 7), d(2023, 4, 7)},

		{Easter, 2015, d(2015, 4, 5), d(2015, 4, 5)},
		{Easter, 2016, d(2016, 3, 27), d(2016, 3, 27)},
		{Easter, 2017, d(2017, 4, 16), d(2017, 4, 16)},
		{Easter, 2018, d(2018, 4, 1), d(2018, 4, 1)},
		{Easter, 2019, d(2019, 4, 21), d(2019, 4, 21)},
		{Easter, 2020, d(2020, 4, 12), d(2020, 4, 12)},
		{Easter, 2021, d(2021, 4, 4), d(2021, 4, 4)},
		{Easter, 2022, d(2022, 4, 17), d(2022, 4, 17)},
		{Easter, 2023, d(2023, 4, 9), d(2023, 4, 9)},

		{EasterMonday, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{EasterMonday, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{EasterMonday, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{EasterMonday, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{EasterMonday, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{EasterMonday, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{EasterMonday, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{EasterMonday, 2022, d(2022, 4, 18), d(2022, 4, 18)},
		{EasterMonday, 2023, d(2023, 4, 10), d(2023, 4, 10)},

		{LabourDay, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{LabourDay, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{LabourDay, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{LabourDay, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{LabourDay, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{LabourDay, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{LabourDay, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{LabourDay, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{StateRestorationDay, 2015, d(2015, 5, 4), d(2015, 5, 4)},
		{StateRestorationDay, 2016, d(2016, 5, 4), d(2016, 5, 4)},
		{StateRestorationDay, 2017, d(2017, 5, 4), d(2017, 5, 4)},
		{StateRestorationDay, 2018, d(2018, 5, 4), d(2018, 5, 4)},
		{StateRestorationDay, 2019, d(2019, 5, 4), d(2019, 5, 6)},
		{StateRestorationDay, 2020, d(2020, 5, 4), d(2020, 5, 4)},
		{StateRestorationDay, 2021, d(2021, 5, 4), d(2021, 5, 4)},
		{StateRestorationDay, 2022, d(2022, 5, 4), d(2022, 5, 4)},

		{MidsummerEve, 2015, d(2015, 6, 23), d(2015, 6, 23)},
		{MidsummerEve, 2016, d(2016, 6, 23), d(2016, 6, 23)},
		{MidsummerEve, 2017, d(2017, 6, 23), d(2017, 6, 23)},
		{MidsummerEve, 2018, d(2018, 6, 23), d(2018, 6, 23)},
		{MidsummerEve, 2019, d(2019, 6, 23), d(2019, 6, 23)},
		{MidsummerEve, 2020, d(2020, 6, 23), d(2020, 6, 23)},
		{MidsummerEve, 2021, d(2021, 6, 23), d(2021, 6, 23)},
		{MidsummerEve, 2022, d(2022, 6, 23), d(2022, 6, 23)},

		{MidsummeDay, 2015, d(2015, 6, 24), d(2015, 6, 24)},
		{MidsummeDay, 2016, d(2016, 6, 24), d(2016, 6, 24)},
		{MidsummeDay, 2017, d(2017, 6, 24), d(2017, 6, 24)},
		{MidsummeDay, 2018, d(2018, 6, 24), d(2018, 6, 24)},
		{MidsummeDay, 2019, d(2019, 6, 24), d(2019, 6, 24)},
		{MidsummeDay, 2020, d(2020, 6, 24), d(2020, 6, 24)},
		{MidsummeDay, 2021, d(2021, 6, 24), d(2021, 6, 24)},
		{MidsummeDay, 2022, d(2022, 6, 24), d(2022, 6, 24)},

		{StateProclamationDay, 2015, d(2015, 11, 18), d(2015, 11, 18)},
		{StateProclamationDay, 2016, d(2016, 11, 18), d(2016, 11, 18)},
		{StateProclamationDay, 2017, d(2017, 11, 18), d(2017, 11, 20)},
		{StateProclamationDay, 2018, d(2018, 11, 18), d(2018, 11, 19)},
		{StateProclamationDay, 2019, d(2019, 11, 18), d(2019, 11, 18)},
		{StateProclamationDay, 2020, d(2020, 11, 18), d(2020, 11, 18)},
		{StateProclamationDay, 2021, d(2021, 11, 18), d(2021, 11, 18)},
		{StateProclamationDay, 2022, d(2022, 11, 18), d(2022, 11, 18)},

		{ChristmasEve, 2015, d(2015, 12, 24), d(2015, 12, 24)},
		{ChristmasEve, 2016, d(2016, 12, 24), d(2016, 12, 24)},
		{ChristmasEve, 2017, d(2017, 12, 24), d(2017, 12, 24)},
		{ChristmasEve, 2018, d(2018, 12, 24), d(2018, 12, 24)},
		{ChristmasEve, 2019, d(2019, 12, 24), d(2019, 12, 24)},
		{ChristmasEve, 2020, d(2020, 12, 24), d(2020, 12, 24)},
		{ChristmasEve, 2021, d(2021, 12, 24), d(2021, 12, 24)},
		{ChristmasEve, 2022, d(2022, 12, 24), d(2022, 12, 24)},

		{ChristmasDay, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{ChristmasDay, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{ChristmasDay, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{ChristmasDay, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{ChristmasDay, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{ChristmasDay, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{ChristmasDay, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{ChristmasDay, 2022, d(2022, 12, 25), d(2022, 12, 25)},
		{ChristmasDay, 2023, d(2023, 12, 25), d(2023, 12, 25)},

		{ChristmasDay2, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{ChristmasDay2, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{ChristmasDay2, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{ChristmasDay2, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{ChristmasDay2, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{ChristmasDay2, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{ChristmasDay2, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{ChristmasDay2, 2022, d(2022, 12, 26), d(2022, 12, 26)},
		{ChristmasDay2, 2023, d(2023, 12, 26), d(2023, 12, 26)},

		{NewYearEve, 2015, d(2015, 12, 31), d(2015, 12, 31)},
		{NewYearEve, 2016, d(2016, 12, 31), d(2016, 12, 31)},
		{NewYearEve, 2017, d(2017, 12, 31), d(2017, 12, 31)},
		{NewYearEve, 2018, d(2018, 12, 31), d(2018, 12, 31)},
		{NewYearEve, 2019, d(2019, 12, 31), d(2019, 12, 31)},
		{NewYearEve, 2020, d(2020, 12, 31), d(2020, 12, 31)},
		{NewYearEve, 2021, d(2021, 12, 31), d(2021, 12, 31)},
		{NewYearEve, 2022, d(2022, 12, 31), d(2022, 12, 31)},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s-%d", test.h.Name, test.y), func(t *testing.T) {
			gotAct, gotObs := test.h.Calc(test.y)
			if !gotAct.Equal(test.wantAct) {
				t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
			}
			if !gotObs.Equal(test.wantObs) {
				t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
			}
		})
	}
}
