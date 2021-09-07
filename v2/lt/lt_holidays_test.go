// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package lt

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
		{NewYear, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{NewYear, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{NewYear, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{NewYear, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{NewYear, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{NewYear, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{StateRestorationDay, 2015, d(2015, 2, 16), d(2015, 2, 16)},
		{StateRestorationDay, 2016, d(2016, 2, 16), d(2016, 2, 16)},
		{StateRestorationDay, 2017, d(2017, 2, 16), d(2017, 2, 16)},
		{StateRestorationDay, 2018, d(2018, 2, 16), d(2018, 2, 16)},
		{StateRestorationDay, 2019, d(2019, 2, 16), d(2019, 2, 16)},
		{StateRestorationDay, 2020, d(2020, 2, 16), d(2020, 2, 16)},
		{StateRestorationDay, 2021, d(2021, 2, 16), d(2021, 2, 16)},
		{StateRestorationDay, 2022, d(2022, 2, 16), d(2022, 2, 16)},

		{IndependenceDay, 2015, d(2015, 3, 11), d(2015, 3, 11)},
		{IndependenceDay, 2016, d(2016, 3, 11), d(2016, 3, 11)},
		{IndependenceDay, 2017, d(2017, 3, 11), d(2017, 3, 11)},
		{IndependenceDay, 2018, d(2018, 3, 11), d(2018, 3, 11)},
		{IndependenceDay, 2019, d(2019, 3, 11), d(2019, 3, 11)},
		{IndependenceDay, 2020, d(2020, 3, 11), d(2020, 3, 11)},
		{IndependenceDay, 2021, d(2021, 3, 11), d(2021, 3, 11)},
		{IndependenceDay, 2022, d(2022, 3, 11), d(2022, 3, 11)},

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

		{SaintJohnsEve, 2015, d(2015, 6, 24), d(2015, 6, 24)},
		{SaintJohnsEve, 2016, d(2016, 6, 24), d(2016, 6, 24)},
		{SaintJohnsEve, 2017, d(2017, 6, 24), d(2017, 6, 24)},
		{SaintJohnsEve, 2018, d(2018, 6, 24), d(2018, 6, 24)},
		{SaintJohnsEve, 2019, d(2019, 6, 24), d(2019, 6, 24)},
		{SaintJohnsEve, 2020, d(2020, 6, 24), d(2020, 6, 24)},
		{SaintJohnsEve, 2021, d(2021, 6, 24), d(2021, 6, 24)},
		{SaintJohnsEve, 2022, d(2022, 6, 24), d(2022, 6, 24)},

		{StatehoodDay, 2015, d(2015, 7, 6), d(2015, 7, 6)},
		{StatehoodDay, 2016, d(2016, 7, 6), d(2016, 7, 6)},
		{StatehoodDay, 2017, d(2017, 7, 6), d(2017, 7, 6)},
		{StatehoodDay, 2018, d(2018, 7, 6), d(2018, 7, 6)},
		{StatehoodDay, 2019, d(2019, 7, 6), d(2019, 7, 6)},
		{StatehoodDay, 2020, d(2020, 7, 6), d(2020, 7, 6)},
		{StatehoodDay, 2021, d(2021, 7, 6), d(2021, 7, 6)},
		{StatehoodDay, 2022, d(2022, 7, 6), d(2022, 7, 6)},

		{AssumptionDay, 2015, d(2015, 8, 15), d(2015, 8, 15)},
		{AssumptionDay, 2016, d(2016, 8, 15), d(2016, 8, 15)},
		{AssumptionDay, 2017, d(2017, 8, 15), d(2017, 8, 15)},
		{AssumptionDay, 2018, d(2018, 8, 15), d(2018, 8, 15)},
		{AssumptionDay, 2019, d(2019, 8, 15), d(2019, 8, 15)},
		{AssumptionDay, 2020, d(2020, 8, 15), d(2020, 8, 15)},
		{AssumptionDay, 2021, d(2021, 8, 15), d(2021, 8, 15)},
		{AssumptionDay, 2022, d(2022, 8, 15), d(2022, 8, 15)},

		{AllSaintsDay, 2015, d(2015, 11, 1), d(2015, 11, 1)},
		{AllSaintsDay, 2016, d(2016, 11, 1), d(2016, 11, 1)},
		{AllSaintsDay, 2017, d(2017, 11, 1), d(2017, 11, 1)},
		{AllSaintsDay, 2018, d(2018, 11, 1), d(2018, 11, 1)},
		{AllSaintsDay, 2019, d(2019, 11, 1), d(2019, 11, 1)},
		{AllSaintsDay, 2020, d(2020, 11, 1), d(2020, 11, 1)},
		{AllSaintsDay, 2021, d(2021, 11, 1), d(2021, 11, 1)},
		{AllSaintsDay, 2022, d(2022, 11, 1), d(2022, 11, 1)},

		{AllSoulsDay, 2015, d(2015, 11, 2), d(2015, 11, 2)},
		{AllSoulsDay, 2016, d(2016, 11, 2), d(2016, 11, 2)},
		{AllSoulsDay, 2017, d(2017, 11, 2), d(2017, 11, 2)},
		{AllSoulsDay, 2018, d(2018, 11, 2), d(2018, 11, 2)},
		{AllSoulsDay, 2019, d(2019, 11, 2), d(2019, 11, 2)},
		{AllSoulsDay, 2020, d(2020, 11, 2), d(2020, 11, 2)},
		{AllSoulsDay, 2021, d(2021, 11, 2), d(2021, 11, 2)},
		{AllSoulsDay, 2022, d(2022, 11, 2), d(2022, 11, 2)},

		{ChristmasEve, 2015, d(2015, 12, 24), d(2015, 12, 24)},
		{ChristmasEve, 2016, d(2016, 12, 24), d(2016, 12, 24)},
		{ChristmasEve, 2017, d(2017, 12, 24), d(2017, 12, 24)},
		{ChristmasEve, 2018, d(2018, 12, 24), d(2018, 12, 24)},
		{ChristmasEve, 2019, d(2019, 12, 24), d(2019, 12, 24)},
		{ChristmasEve, 2020, d(2020, 12, 24), d(2020, 12, 24)},
		{ChristmasEve, 2021, d(2021, 12, 24), d(2021, 12, 24)},
		{ChristmasEve, 2022, d(2022, 12, 24), d(2022, 12, 24)},

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
