// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package pl

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

		{ThreeKings, 2015, d(2015, 1, 6), d(2015, 1, 6)},
		{ThreeKings, 2016, d(2016, 1, 6), d(2016, 1, 6)},
		{ThreeKings, 2017, d(2017, 1, 6), d(2017, 1, 6)},
		{ThreeKings, 2018, d(2018, 1, 6), d(2018, 1, 6)},
		{ThreeKings, 2019, d(2019, 1, 6), d(2019, 1, 6)},
		{ThreeKings, 2020, d(2020, 1, 6), d(2020, 1, 6)},
		{ThreeKings, 2021, d(2021, 1, 6), d(2021, 1, 6)},
		{ThreeKings, 2022, d(2022, 1, 6), d(2022, 1, 6)},

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

		{ConstitutionDay, 2015, d(2015, 5, 3), d(2015, 5, 3)},
		{ConstitutionDay, 2016, d(2016, 5, 3), d(2016, 5, 3)},
		{ConstitutionDay, 2017, d(2017, 5, 3), d(2017, 5, 3)},
		{ConstitutionDay, 2018, d(2018, 5, 3), d(2018, 5, 3)},
		{ConstitutionDay, 2019, d(2019, 5, 3), d(2019, 5, 3)},
		{ConstitutionDay, 2020, d(2020, 5, 3), d(2020, 5, 3)},
		{ConstitutionDay, 2021, d(2021, 5, 3), d(2021, 5, 3)},
		{ConstitutionDay, 2022, d(2022, 5, 3), d(2022, 5, 3)},

		{CorpusChristi, 2015, d(2015, 6, 4), d(2015, 6, 4)},
		{CorpusChristi, 2016, d(2016, 5, 26), d(2016, 5, 26)},
		{CorpusChristi, 2017, d(2017, 6, 15), d(2017, 6, 15)},
		{CorpusChristi, 2018, d(2018, 5, 31), d(2018, 5, 31)},
		{CorpusChristi, 2019, d(2019, 6, 20), d(2019, 6, 20)},
		{CorpusChristi, 2020, d(2020, 6, 11), d(2020, 6, 11)},
		{CorpusChristi, 2021, d(2021, 6, 3), d(2021, 6, 3)},
		{CorpusChristi, 2022, d(2022, 6, 16), d(2022, 6, 16)},

		{AssumptionBlessedVirginMary, 2015, d(2015, 8, 15), d(2015, 8, 15)},
		{AssumptionBlessedVirginMary, 2016, d(2016, 8, 15), d(2016, 8, 15)},
		{AssumptionBlessedVirginMary, 2017, d(2017, 8, 15), d(2017, 8, 15)},
		{AssumptionBlessedVirginMary, 2018, d(2018, 8, 15), d(2018, 8, 15)},
		{AssumptionBlessedVirginMary, 2019, d(2019, 8, 15), d(2019, 8, 15)},
		{AssumptionBlessedVirginMary, 2020, d(2020, 8, 15), d(2020, 8, 15)},
		{AssumptionBlessedVirginMary, 2021, d(2021, 8, 15), d(2021, 8, 15)},
		{AssumptionBlessedVirginMary, 2022, d(2022, 8, 15), d(2022, 8, 15)},

		{AllSaints, 2015, d(2015, 11, 1), d(2015, 11, 1)},
		{AllSaints, 2016, d(2016, 11, 1), d(2016, 11, 1)},
		{AllSaints, 2017, d(2017, 11, 1), d(2017, 11, 1)},
		{AllSaints, 2018, d(2018, 11, 1), d(2018, 11, 1)},
		{AllSaints, 2019, d(2019, 11, 1), d(2019, 11, 1)},
		{AllSaints, 2020, d(2020, 11, 1), d(2020, 11, 1)},
		{AllSaints, 2021, d(2021, 11, 1), d(2021, 11, 1)},
		{AllSaints, 2022, d(2022, 11, 1), d(2022, 11, 1)},

		{NationalIndependenceDay, 2015, d(2015, 11, 11), d(2015, 11, 11)},
		{NationalIndependenceDay, 2016, d(2016, 11, 11), d(2016, 11, 11)},
		{NationalIndependenceDay, 2017, d(2017, 11, 11), d(2017, 11, 11)},
		{NationalIndependenceDay, 2018, d(2018, 11, 11), d(2018, 11, 11)},
		{NationalIndependenceDay, 2019, d(2019, 11, 11), d(2019, 11, 11)},
		{NationalIndependenceDay, 2020, d(2020, 11, 11), d(2020, 11, 11)},
		{NationalIndependenceDay, 2021, d(2021, 11, 11), d(2021, 11, 11)},
		{NationalIndependenceDay, 2022, d(2022, 11, 11), d(2022, 11, 11)},

		{ChristmasDayOne, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{ChristmasDayOne, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{ChristmasDayOne, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{ChristmasDayOne, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{ChristmasDayOne, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{ChristmasDayOne, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{ChristmasDayOne, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{ChristmasDayOne, 2022, d(2022, 12, 25), d(2022, 12, 25)},

		{ChristmasDayTwo, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{ChristmasDayTwo, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{ChristmasDayTwo, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{ChristmasDayTwo, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{ChristmasDayTwo, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{ChristmasDayTwo, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{ChristmasDayTwo, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{ChristmasDayTwo, 2022, d(2022, 12, 26), d(2022, 12, 26)},
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
