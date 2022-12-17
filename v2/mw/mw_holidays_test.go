// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package mw

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

		{ChilembweDay, 2015, d(2015, 1, 15), d(2015, 1, 15)},
		{ChilembweDay, 2016, d(2016, 1, 15), d(2016, 1, 15)},
		{ChilembweDay, 2017, d(2017, 1, 15), d(2017, 1, 16)},
		{ChilembweDay, 2018, d(2018, 1, 15), d(2018, 1, 15)},
		{ChilembweDay, 2019, d(2019, 1, 15), d(2019, 1, 15)},
		{ChilembweDay, 2020, d(2020, 1, 15), d(2020, 1, 15)},
		{ChilembweDay, 2021, d(2021, 1, 15), d(2021, 1, 15)},
		{ChilembweDay, 2022, d(2022, 1, 15), d(2022, 1, 17)},

		{MartyrsDay, 2015, d(2015, 3, 3), d(2015, 3, 3)},
		{MartyrsDay, 2016, d(2016, 3, 3), d(2016, 3, 3)},
		{MartyrsDay, 2017, d(2017, 3, 3), d(2017, 3, 3)},
		{MartyrsDay, 2018, d(2018, 3, 3), d(2018, 3, 5)},
		{MartyrsDay, 2019, d(2019, 3, 3), d(2019, 3, 4)},
		{MartyrsDay, 2020, d(2020, 3, 3), d(2020, 3, 3)},
		{MartyrsDay, 2021, d(2021, 3, 3), d(2021, 3, 3)},
		{MartyrsDay, 2022, d(2022, 3, 3), d(2022, 3, 3)},

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

		{LabourDay, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{LabourDay, 2016, d(2016, 5, 1), d(2016, 5, 2)},
		{LabourDay, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{LabourDay, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{LabourDay, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{LabourDay, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{LabourDay, 2021, d(2021, 5, 1), d(2021, 5, 3)},
		{LabourDay, 2022, d(2022, 5, 1), d(2022, 5, 2)},

		{KamuzuDay, 2015, d(2015, 5, 14), d(2015, 5, 14)},
		{KamuzuDay, 2016, d(2016, 5, 14), d(2016, 5, 16)},
		{KamuzuDay, 2017, d(2017, 5, 14), d(2017, 5, 15)},
		{KamuzuDay, 2018, d(2018, 5, 14), d(2018, 5, 14)},
		{KamuzuDay, 2019, d(2019, 5, 14), d(2019, 5, 14)},
		{KamuzuDay, 2020, d(2020, 5, 14), d(2020, 5, 14)},
		{KamuzuDay, 2021, d(2021, 5, 14), d(2021, 5, 14)},
		{KamuzuDay, 2022, d(2022, 5, 14), d(2022, 5, 16)},

		{MothersDay, 2015, d(2015, 10, 15), d(2015, 10, 15)},
		{MothersDay, 2016, d(2016, 10, 15), d(2016, 10, 17)},
		{MothersDay, 2017, d(2017, 10, 15), d(2017, 10, 16)},
		{MothersDay, 2018, d(2018, 10, 15), d(2018, 10, 15)},
		{MothersDay, 2019, d(2019, 10, 15), d(2019, 10, 15)},
		{MothersDay, 2020, d(2020, 10, 15), d(2020, 10, 15)},
		{MothersDay, 2021, d(2021, 10, 15), d(2021, 10, 15)},
		{MothersDay, 2022, d(2022, 10, 15), d(2022, 10, 17)},

		{IndependenceDay, 2015, d(2015, 7, 6), d(2015, 7, 6)},
		{IndependenceDay, 2016, d(2016, 7, 6), d(2016, 7, 6)},
		{IndependenceDay, 2017, d(2017, 7, 6), d(2017, 7, 6)},
		{IndependenceDay, 2018, d(2018, 7, 6), d(2018, 7, 6)},
		{IndependenceDay, 2019, d(2019, 7, 6), d(2019, 7, 8)},
		{IndependenceDay, 2020, d(2020, 7, 6), d(2020, 7, 6)},
		{IndependenceDay, 2021, d(2021, 7, 6), d(2021, 7, 6)},
		{IndependenceDay, 2022, d(2022, 7, 6), d(2022, 7, 6)},

		{BoxingDay, 2021, d(2021, 12, 26), d(2021, 12, 27)},
		{BoxingDay, 2022, d(2022, 12, 26), d(2022, 12, 26)},
		{BoxingDay, 2023, d(2023, 12, 26), d(2023, 12, 26)},
		{BoxingDay, 2024, d(2024, 12, 26), d(2024, 12, 26)},

		{ChristmasDay, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{ChristmasDay, 2016, d(2016, 12, 25), d(2016, 12, 26)},
		{ChristmasDay, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{ChristmasDay, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{ChristmasDay, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{ChristmasDay, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{ChristmasDay, 2021, d(2021, 12, 25), d(2021, 12, 27)},
		{ChristmasDay, 2022, d(2022, 12, 25), d(2022, 12, 26)},
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
