// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package cz

import (
	"testing"
	"time"

	"github.com/Tamh/cal/v2"
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

		{EasterMonday, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{EasterMonday, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{EasterMonday, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{EasterMonday, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{EasterMonday, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{EasterMonday, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{EasterMonday, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{EasterMonday, 2022, d(2022, 4, 18), d(2022, 4, 18)},

		{LabourDay, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{LabourDay, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{LabourDay, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{LabourDay, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{LabourDay, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{LabourDay, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{LabourDay, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{LabourDay, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{LiberationDay, 2015, d(2015, 5, 8), d(2015, 5, 8)},
		{LiberationDay, 2016, d(2016, 5, 8), d(2016, 5, 8)},
		{LiberationDay, 2017, d(2017, 5, 8), d(2017, 5, 8)},
		{LiberationDay, 2018, d(2018, 5, 8), d(2018, 5, 8)},
		{LiberationDay, 2019, d(2019, 5, 8), d(2019, 5, 8)},
		{LiberationDay, 2020, d(2020, 5, 8), d(2020, 5, 8)},
		{LiberationDay, 2021, d(2021, 5, 8), d(2021, 5, 8)},
		{LiberationDay, 2022, d(2022, 5, 8), d(2022, 5, 8)},

		{SaintsCyrilMethodius, 2015, d(2015, 7, 5), d(2015, 7, 5)},
		{SaintsCyrilMethodius, 2016, d(2016, 7, 5), d(2016, 7, 5)},
		{SaintsCyrilMethodius, 2017, d(2017, 7, 5), d(2017, 7, 5)},
		{SaintsCyrilMethodius, 2018, d(2018, 7, 5), d(2018, 7, 5)},
		{SaintsCyrilMethodius, 2019, d(2019, 7, 5), d(2019, 7, 5)},
		{SaintsCyrilMethodius, 2020, d(2020, 7, 5), d(2020, 7, 5)},
		{SaintsCyrilMethodius, 2021, d(2021, 7, 5), d(2021, 7, 5)},
		{SaintsCyrilMethodius, 2022, d(2022, 7, 5), d(2022, 7, 5)},

		{JanHusDay, 2015, d(2015, 7, 6), d(2015, 7, 6)},
		{JanHusDay, 2016, d(2016, 7, 6), d(2016, 7, 6)},
		{JanHusDay, 2017, d(2017, 7, 6), d(2017, 7, 6)},
		{JanHusDay, 2018, d(2018, 7, 6), d(2018, 7, 6)},
		{JanHusDay, 2019, d(2019, 7, 6), d(2019, 7, 6)},
		{JanHusDay, 2020, d(2020, 7, 6), d(2020, 7, 6)},
		{JanHusDay, 2021, d(2021, 7, 6), d(2021, 7, 6)},
		{JanHusDay, 2022, d(2022, 7, 6), d(2022, 7, 6)},

		{SaintWenceslasDay, 2015, d(2015, 9, 28), d(2015, 9, 28)},
		{SaintWenceslasDay, 2016, d(2016, 9, 28), d(2016, 9, 28)},
		{SaintWenceslasDay, 2017, d(2017, 9, 28), d(2017, 9, 28)},
		{SaintWenceslasDay, 2018, d(2018, 9, 28), d(2018, 9, 28)},
		{SaintWenceslasDay, 2019, d(2019, 9, 28), d(2019, 9, 28)},
		{SaintWenceslasDay, 2020, d(2020, 9, 28), d(2020, 9, 28)},
		{SaintWenceslasDay, 2021, d(2021, 9, 28), d(2021, 9, 28)},
		{SaintWenceslasDay, 2022, d(2022, 9, 28), d(2022, 9, 28)},

		{IndependenceDay, 2015, d(2015, 10, 28), d(2015, 10, 28)},
		{IndependenceDay, 2016, d(2016, 10, 28), d(2016, 10, 28)},
		{IndependenceDay, 2017, d(2017, 10, 28), d(2017, 10, 28)},
		{IndependenceDay, 2018, d(2018, 10, 28), d(2018, 10, 28)},
		{IndependenceDay, 2019, d(2019, 10, 28), d(2019, 10, 28)},
		{IndependenceDay, 2020, d(2020, 10, 28), d(2020, 10, 28)},
		{IndependenceDay, 2021, d(2021, 10, 28), d(2021, 10, 28)},
		{IndependenceDay, 2022, d(2022, 10, 28), d(2022, 10, 28)},

		{FreedomDay, 2015, d(2015, 11, 17), d(2015, 11, 17)},
		{FreedomDay, 2016, d(2016, 11, 17), d(2016, 11, 17)},
		{FreedomDay, 2017, d(2017, 11, 17), d(2017, 11, 17)},
		{FreedomDay, 2018, d(2018, 11, 17), d(2018, 11, 17)},
		{FreedomDay, 2019, d(2019, 11, 17), d(2019, 11, 17)},
		{FreedomDay, 2020, d(2020, 11, 17), d(2020, 11, 17)},
		{FreedomDay, 2021, d(2021, 11, 17), d(2021, 11, 17)},
		{FreedomDay, 2022, d(2022, 11, 17), d(2022, 11, 17)},

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

		{SaintStephensDay, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{SaintStephensDay, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{SaintStephensDay, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{SaintStephensDay, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{SaintStephensDay, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{SaintStephensDay, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{SaintStephensDay, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{SaintStephensDay, 2022, d(2022, 12, 26), d(2022, 12, 26)},
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
