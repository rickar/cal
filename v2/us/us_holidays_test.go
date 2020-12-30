// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package us

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
		{NewYear, 2022, d(2022, 1, 1), d(2021, 12, 31)},

		{MlkDay, 2015, d(2015, 1, 19), d(2015, 1, 19)},
		{MlkDay, 2016, d(2016, 1, 18), d(2016, 1, 18)},
		{MlkDay, 2017, d(2017, 1, 16), d(2017, 1, 16)},
		{MlkDay, 2018, d(2018, 1, 15), d(2018, 1, 15)},
		{MlkDay, 2019, d(2019, 1, 21), d(2019, 1, 21)},
		{MlkDay, 2020, d(2020, 1, 20), d(2020, 1, 20)},
		{MlkDay, 2021, d(2021, 1, 18), d(2021, 1, 18)},
		{MlkDay, 2022, d(2022, 1, 17), d(2022, 1, 17)},

		{PresidentsDay, 2015, d(2015, 2, 16), d(2015, 2, 16)},
		{PresidentsDay, 2016, d(2016, 2, 15), d(2016, 2, 15)},
		{PresidentsDay, 2017, d(2017, 2, 20), d(2017, 2, 20)},
		{PresidentsDay, 2018, d(2018, 2, 19), d(2018, 2, 19)},
		{PresidentsDay, 2019, d(2019, 2, 18), d(2019, 2, 18)},
		{PresidentsDay, 2020, d(2020, 2, 17), d(2020, 2, 17)},
		{PresidentsDay, 2021, d(2021, 2, 15), d(2021, 2, 15)},
		{PresidentsDay, 2022, d(2022, 2, 21), d(2022, 2, 21)},

		{MemorialDay, 2015, d(2015, 5, 25), d(2015, 5, 25)},
		{MemorialDay, 2016, d(2016, 5, 30), d(2016, 5, 30)},
		{MemorialDay, 2017, d(2017, 5, 29), d(2017, 5, 29)},
		{MemorialDay, 2018, d(2018, 5, 28), d(2018, 5, 28)},
		{MemorialDay, 2019, d(2019, 5, 27), d(2019, 5, 27)},
		{MemorialDay, 2020, d(2020, 5, 25), d(2020, 5, 25)},
		{MemorialDay, 2021, d(2021, 5, 31), d(2021, 5, 31)},
		{MemorialDay, 2022, d(2022, 5, 30), d(2022, 5, 30)},

		{IndependenceDay, 2015, d(2015, 7, 4), d(2015, 7, 3)},
		{IndependenceDay, 2016, d(2016, 7, 4), d(2016, 7, 4)},
		{IndependenceDay, 2017, d(2017, 7, 4), d(2017, 7, 4)},
		{IndependenceDay, 2018, d(2018, 7, 4), d(2018, 7, 4)},
		{IndependenceDay, 2019, d(2019, 7, 4), d(2019, 7, 4)},
		{IndependenceDay, 2020, d(2020, 7, 4), d(2020, 7, 3)},
		{IndependenceDay, 2021, d(2021, 7, 4), d(2021, 7, 5)},
		{IndependenceDay, 2022, d(2022, 7, 4), d(2022, 7, 4)},

		{LaborDay, 2015, d(2015, 9, 7), d(2015, 9, 7)},
		{LaborDay, 2016, d(2016, 9, 5), d(2016, 9, 5)},
		{LaborDay, 2017, d(2017, 9, 4), d(2017, 9, 4)},
		{LaborDay, 2018, d(2018, 9, 3), d(2018, 9, 3)},
		{LaborDay, 2019, d(2019, 9, 2), d(2019, 9, 2)},
		{LaborDay, 2020, d(2020, 9, 7), d(2020, 9, 7)},
		{LaborDay, 2021, d(2021, 9, 6), d(2021, 9, 6)},
		{LaborDay, 2022, d(2022, 9, 5), d(2022, 9, 5)},

		{ColumbusDay, 2015, d(2015, 10, 12), d(2015, 10, 12)},
		{ColumbusDay, 2016, d(2016, 10, 10), d(2016, 10, 10)},
		{ColumbusDay, 2017, d(2017, 10, 9), d(2017, 10, 9)},
		{ColumbusDay, 2018, d(2018, 10, 8), d(2018, 10, 8)},
		{ColumbusDay, 2019, d(2019, 10, 14), d(2019, 10, 14)},
		{ColumbusDay, 2020, d(2020, 10, 12), d(2020, 10, 12)},
		{ColumbusDay, 2021, d(2021, 10, 11), d(2021, 10, 11)},
		{ColumbusDay, 2022, d(2022, 10, 10), d(2022, 10, 10)},

		{VeteransDay, 2015, d(2015, 11, 11), d(2015, 11, 11)},
		{VeteransDay, 2016, d(2016, 11, 11), d(2016, 11, 11)},
		{VeteransDay, 2017, d(2017, 11, 11), d(2017, 11, 10)},
		{VeteransDay, 2018, d(2018, 11, 11), d(2018, 11, 12)},
		{VeteransDay, 2019, d(2019, 11, 11), d(2019, 11, 11)},
		{VeteransDay, 2020, d(2020, 11, 11), d(2020, 11, 11)},
		{VeteransDay, 2021, d(2021, 11, 11), d(2021, 11, 11)},
		{VeteransDay, 2022, d(2022, 11, 11), d(2022, 11, 11)},

		{ThanksgivingDay, 2015, d(2015, 11, 26), d(2015, 11, 26)},
		{ThanksgivingDay, 2016, d(2016, 11, 24), d(2016, 11, 24)},
		{ThanksgivingDay, 2017, d(2017, 11, 23), d(2017, 11, 23)},
		{ThanksgivingDay, 2018, d(2018, 11, 22), d(2018, 11, 22)},
		{ThanksgivingDay, 2019, d(2019, 11, 28), d(2019, 11, 28)},
		{ThanksgivingDay, 2020, d(2020, 11, 26), d(2020, 11, 26)},
		{ThanksgivingDay, 2021, d(2021, 11, 25), d(2021, 11, 25)},
		{ThanksgivingDay, 2022, d(2022, 11, 24), d(2022, 11, 24)},

		{DayAfterThanksgivingDay, 2021, d(2021, 11, 26), d(2021, 11, 26)},
		{DayAfterThanksgivingDay, 2022, d(2022, 11, 25), d(2022, 11, 25)},
		{DayAfterThanksgivingDay, 2023, d(2023, 11, 24), d(2023, 11, 24)},
		{DayAfterThanksgivingDay, 2024, d(2024, 11, 29), d(2024, 11, 29)},

		{ChristmasDay, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{ChristmasDay, 2016, d(2016, 12, 25), d(2016, 12, 26)},
		{ChristmasDay, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{ChristmasDay, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{ChristmasDay, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{ChristmasDay, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{ChristmasDay, 2021, d(2021, 12, 25), d(2021, 12, 24)},
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
