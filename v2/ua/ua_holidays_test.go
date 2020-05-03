// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package ua

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

		{OrthodoxChristmas, 2015, d(2015, 1, 7), d(2015, 1, 7)},
		{OrthodoxChristmas, 2016, d(2016, 1, 7), d(2016, 1, 7)},
		{OrthodoxChristmas, 2017, d(2017, 1, 7), d(2017, 1, 9)},
		{OrthodoxChristmas, 2018, d(2018, 1, 7), d(2018, 1, 8)},
		{OrthodoxChristmas, 2019, d(2019, 1, 7), d(2019, 1, 7)},
		{OrthodoxChristmas, 2020, d(2020, 1, 7), d(2020, 1, 7)},
		{OrthodoxChristmas, 2021, d(2021, 1, 7), d(2021, 1, 7)},
		{OrthodoxChristmas, 2022, d(2022, 1, 7), d(2022, 1, 7)},

		{WomensDay, 2015, d(2015, 3, 8), d(2015, 3, 9)},
		{WomensDay, 2016, d(2016, 3, 8), d(2016, 3, 8)},
		{WomensDay, 2017, d(2017, 3, 8), d(2017, 3, 8)},
		{WomensDay, 2018, d(2018, 3, 8), d(2018, 3, 8)},
		{WomensDay, 2019, d(2019, 3, 8), d(2019, 3, 8)},
		{WomensDay, 2020, d(2020, 3, 8), d(2020, 3, 9)},
		{WomensDay, 2021, d(2021, 3, 8), d(2021, 3, 8)},
		{WomensDay, 2022, d(2022, 3, 8), d(2022, 3, 8)},

		{OrthodoxEasterMonday, 2015, d(2015, 4, 13), d(2015, 4, 13)},
		{OrthodoxEasterMonday, 2016, d(2016, 5, 2), d(2016, 5, 2)},
		{OrthodoxEasterMonday, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{OrthodoxEasterMonday, 2018, d(2018, 4, 9), d(2018, 4, 9)},
		{OrthodoxEasterMonday, 2019, d(2019, 4, 29), d(2019, 4, 29)},
		{OrthodoxEasterMonday, 2020, d(2020, 4, 20), d(2020, 4, 20)},
		{OrthodoxEasterMonday, 2021, d(2021, 5, 3), d(2021, 5, 3)},
		{OrthodoxEasterMonday, 2022, d(2022, 4, 25), d(2022, 4, 25)},

		{LabourDay, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{LabourDay, 2016, d(2016, 5, 1), d(2016, 5, 2)},
		{LabourDay, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{LabourDay, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{LabourDay, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{LabourDay, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{LabourDay, 2021, d(2021, 5, 1), d(2021, 5, 3)},
		{LabourDay, 2022, d(2022, 5, 1), d(2022, 5, 2)},

		{LabourDay2, 2015, d(2015, 5, 2), d(2015, 5, 4)},
		{LabourDay2, 2016, d(2016, 5, 2), d(2016, 5, 2)},
		{LabourDay2, 2017, d(2017, 5, 2), d(2017, 5, 2)},
		{LabourDay2, 2018, time.Time{}, time.Time{}},
		{LabourDay2, 2019, time.Time{}, time.Time{}},
		{LabourDay2, 2020, time.Time{}, time.Time{}},
		{LabourDay2, 2021, time.Time{}, time.Time{}},
		{LabourDay2, 2022, time.Time{}, time.Time{}},

		{VictoryDay, 2015, d(2015, 5, 9), d(2015, 5, 11)},
		{VictoryDay, 2016, d(2016, 5, 9), d(2016, 5, 9)},
		{VictoryDay, 2017, d(2017, 5, 9), d(2017, 5, 9)},
		{VictoryDay, 2018, d(2018, 5, 9), d(2018, 5, 9)},
		{VictoryDay, 2019, d(2019, 5, 9), d(2019, 5, 9)},
		{VictoryDay, 2020, d(2020, 5, 9), d(2020, 5, 11)},
		{VictoryDay, 2021, d(2021, 5, 9), d(2021, 5, 10)},
		{VictoryDay, 2022, d(2022, 5, 9), d(2022, 5, 9)},

		{OrthodoxPentecostMonday, 2015, d(2015, 5, 31), d(2015, 5, 31)},
		{OrthodoxPentecostMonday, 2016, d(2016, 6, 19), d(2016, 6, 19)},
		{OrthodoxPentecostMonday, 2017, d(2017, 6, 4), d(2017, 6, 4)},
		{OrthodoxPentecostMonday, 2018, d(2018, 5, 27), d(2018, 5, 27)},
		{OrthodoxPentecostMonday, 2019, d(2019, 6, 16), d(2019, 6, 16)},
		{OrthodoxPentecostMonday, 2020, d(2020, 6, 7), d(2020, 6, 7)},
		{OrthodoxPentecostMonday, 2021, d(2021, 6, 20), d(2021, 6, 20)},
		{OrthodoxPentecostMonday, 2022, d(2022, 6, 12), d(2022, 6, 12)},

		{ConstitutionDay, 2015, d(2015, 6, 28), d(2015, 6, 29)},
		{ConstitutionDay, 2016, d(2016, 6, 28), d(2016, 6, 28)},
		{ConstitutionDay, 2017, d(2017, 6, 28), d(2017, 6, 28)},
		{ConstitutionDay, 2018, d(2018, 6, 28), d(2018, 6, 28)},
		{ConstitutionDay, 2019, d(2019, 6, 28), d(2019, 6, 28)},
		{ConstitutionDay, 2020, d(2020, 6, 28), d(2020, 6, 29)},
		{ConstitutionDay, 2021, d(2021, 6, 28), d(2021, 6, 28)},
		{ConstitutionDay, 2022, d(2022, 6, 28), d(2022, 6, 28)},

		{IndependenceDay, 2015, d(2015, 8, 24), d(2015, 8, 24)},
		{IndependenceDay, 2016, d(2016, 8, 24), d(2016, 8, 24)},
		{IndependenceDay, 2017, d(2017, 8, 24), d(2017, 8, 24)},
		{IndependenceDay, 2018, d(2018, 8, 24), d(2018, 8, 24)},
		{IndependenceDay, 2019, d(2019, 8, 24), d(2019, 8, 26)},
		{IndependenceDay, 2020, d(2020, 8, 24), d(2020, 8, 24)},
		{IndependenceDay, 2021, d(2021, 8, 24), d(2021, 8, 24)},
		{IndependenceDay, 2022, d(2022, 8, 24), d(2022, 8, 24)},

		{DefenderOfUkraineDay, 2015, d(2015, 10, 14), d(2015, 10, 14)},
		{DefenderOfUkraineDay, 2016, d(2016, 10, 14), d(2016, 10, 14)},
		{DefenderOfUkraineDay, 2017, d(2017, 10, 14), d(2017, 10, 16)},
		{DefenderOfUkraineDay, 2018, d(2018, 10, 14), d(2018, 10, 15)},
		{DefenderOfUkraineDay, 2019, d(2019, 10, 14), d(2019, 10, 14)},
		{DefenderOfUkraineDay, 2020, d(2020, 10, 14), d(2020, 10, 14)},
		{DefenderOfUkraineDay, 2021, d(2021, 10, 14), d(2021, 10, 14)},
		{DefenderOfUkraineDay, 2022, d(2022, 10, 14), d(2022, 10, 14)},

		{CatholicChristmas, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{CatholicChristmas, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{CatholicChristmas, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{CatholicChristmas, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{CatholicChristmas, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{CatholicChristmas, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{CatholicChristmas, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{CatholicChristmas, 2022, d(2022, 12, 25), d(2022, 12, 25)},
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
