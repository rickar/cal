// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package mx

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

		{ConstitutionDay, 2015, d(2015, 2, 5), d(2015, 2, 5)},
		{ConstitutionDay, 2016, d(2016, 2, 5), d(2016, 2, 5)},
		{ConstitutionDay, 2017, d(2017, 2, 5), d(2017, 2, 6)},
		{ConstitutionDay, 2018, d(2018, 2, 5), d(2018, 2, 5)},
		{ConstitutionDay, 2019, d(2019, 2, 5), d(2019, 2, 5)},
		{ConstitutionDay, 2020, d(2020, 2, 5), d(2020, 2, 5)},
		{ConstitutionDay, 2021, d(2021, 2, 5), d(2021, 2, 5)},
		{ConstitutionDay, 2022, d(2022, 2, 5), d(2022, 2, 4)},

		{BenitoJuarezDay, 2015, d(2015, 3, 21), d(2015, 3, 20)},
		{BenitoJuarezDay, 2016, d(2016, 3, 21), d(2016, 3, 21)},
		{BenitoJuarezDay, 2017, d(2017, 3, 21), d(2017, 3, 21)},
		{BenitoJuarezDay, 2018, d(2018, 3, 21), d(2018, 3, 21)},
		{BenitoJuarezDay, 2019, d(2019, 3, 21), d(2019, 3, 21)},
		{BenitoJuarezDay, 2020, d(2020, 3, 21), d(2020, 3, 20)},
		{BenitoJuarezDay, 2021, d(2021, 3, 21), d(2021, 3, 22)},
		{BenitoJuarezDay, 2022, d(2022, 3, 21), d(2022, 3, 21)},

		{LabourDay, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{LabourDay, 2016, d(2016, 5, 1), d(2016, 5, 2)},
		{LabourDay, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{LabourDay, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{LabourDay, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{LabourDay, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{LabourDay, 2021, d(2021, 5, 1), d(2021, 4, 30)},
		{LabourDay, 2022, d(2022, 5, 1), d(2022, 5, 2)},

		{IndependenceDay, 2015, d(2015, 9, 16), d(2015, 9, 16)},
		{IndependenceDay, 2016, d(2016, 9, 16), d(2016, 9, 16)},
		{IndependenceDay, 2017, d(2017, 9, 16), d(2017, 9, 15)},
		{IndependenceDay, 2018, d(2018, 9, 16), d(2018, 9, 17)},
		{IndependenceDay, 2019, d(2019, 9, 16), d(2019, 9, 16)},
		{IndependenceDay, 2020, d(2020, 9, 16), d(2020, 9, 16)},
		{IndependenceDay, 2021, d(2021, 9, 16), d(2021, 9, 16)},
		{IndependenceDay, 2022, d(2022, 9, 16), d(2022, 9, 16)},

		{RevolutionDay, 2015, d(2015, 11, 20), d(2015, 11, 20)},
		{RevolutionDay, 2016, d(2016, 11, 20), d(2016, 11, 21)},
		{RevolutionDay, 2017, d(2017, 11, 20), d(2017, 11, 20)},
		{RevolutionDay, 2018, d(2018, 11, 20), d(2018, 11, 20)},
		{RevolutionDay, 2019, d(2019, 11, 20), d(2019, 11, 20)},
		{RevolutionDay, 2020, d(2020, 11, 20), d(2020, 11, 20)},
		{RevolutionDay, 2021, d(2021, 11, 20), d(2021, 11, 19)},
		{RevolutionDay, 2022, d(2022, 11, 20), d(2022, 11, 21)},

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
