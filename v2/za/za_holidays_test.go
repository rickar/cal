// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package za

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
		{NewYear, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{HumanRightsDay, 2015, d(2015, 3, 21), d(2015, 3, 21)},
		{HumanRightsDay, 2016, d(2016, 3, 21), d(2016, 3, 21)},
		{HumanRightsDay, 2017, d(2017, 3, 21), d(2017, 3, 21)},
		{HumanRightsDay, 2018, d(2018, 3, 21), d(2018, 3, 21)},
		{HumanRightsDay, 2019, d(2019, 3, 21), d(2019, 3, 21)},
		{HumanRightsDay, 2020, d(2020, 3, 21), d(2020, 3, 21)},
		{HumanRightsDay, 2021, d(2021, 3, 21), d(2021, 3, 22)},
		{HumanRightsDay, 2022, d(2022, 3, 21), d(2022, 3, 21)},

		{GoodFriday, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{GoodFriday, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{GoodFriday, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{GoodFriday, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{GoodFriday, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{GoodFriday, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{GoodFriday, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{GoodFriday, 2022, d(2022, 4, 15), d(2022, 4, 15)},

		{FamilyDay, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{FamilyDay, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{FamilyDay, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{FamilyDay, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{FamilyDay, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{FamilyDay, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{FamilyDay, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{FamilyDay, 2022, d(2022, 4, 18), d(2022, 4, 18)},

		{FreedomDay, 2015, d(2015, 4, 27), d(2015, 4, 27)},
		{FreedomDay, 2016, d(2016, 4, 27), d(2016, 4, 27)},
		{FreedomDay, 2017, d(2017, 4, 27), d(2017, 4, 27)},
		{FreedomDay, 2018, d(2018, 4, 27), d(2018, 4, 27)},
		{FreedomDay, 2019, d(2019, 4, 27), d(2019, 4, 27)},
		{FreedomDay, 2020, d(2020, 4, 27), d(2020, 4, 27)},
		{FreedomDay, 2021, d(2021, 4, 27), d(2021, 4, 27)},
		{FreedomDay, 2022, d(2022, 4, 27), d(2022, 4, 27)},
		{FreedomDay, 2025, d(2025, 4, 27), d(2025, 4, 28)},

		{WorkersDay, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{WorkersDay, 2016, d(2016, 5, 1), d(2016, 5, 2)},
		{WorkersDay, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{WorkersDay, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{WorkersDay, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{WorkersDay, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{WorkersDay, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{WorkersDay, 2022, d(2022, 5, 1), d(2022, 5, 2)},

		{YouthDay, 2015, d(2015, 6, 16), d(2015, 6, 16)},
		{YouthDay, 2016, d(2016, 6, 16), d(2016, 6, 16)},
		{YouthDay, 2017, d(2017, 6, 16), d(2017, 6, 16)},
		{YouthDay, 2018, d(2018, 6, 16), d(2018, 6, 16)},
		{YouthDay, 2019, d(2019, 6, 16), d(2019, 6, 17)},
		{YouthDay, 2020, d(2020, 6, 16), d(2020, 6, 16)},
		{YouthDay, 2021, d(2021, 6, 16), d(2021, 6, 16)},
		{YouthDay, 2022, d(2022, 6, 16), d(2022, 6, 16)},

		{WomensDay, 2015, d(2015, 8, 9), d(2015, 8, 10)},
		{WomensDay, 2016, d(2016, 8, 9), d(2016, 8, 9)},
		{WomensDay, 2017, d(2017, 8, 9), d(2017, 8, 9)},
		{WomensDay, 2018, d(2018, 8, 9), d(2018, 8, 9)},
		{WomensDay, 2019, d(2019, 8, 9), d(2019, 8, 9)},
		{WomensDay, 2020, d(2020, 8, 9), d(2020, 8, 10)},
		{WomensDay, 2021, d(2021, 8, 9), d(2021, 8, 9)},
		{WomensDay, 2022, d(2022, 8, 9), d(2022, 8, 9)},

		{HeritageDay, 2015, d(2015, 9, 24), d(2015, 9, 24)},
		{HeritageDay, 2016, d(2016, 9, 24), d(2016, 9, 24)},
		{HeritageDay, 2017, d(2017, 9, 24), d(2017, 9, 25)},
		{HeritageDay, 2018, d(2018, 9, 24), d(2018, 9, 24)},
		{HeritageDay, 2019, d(2019, 9, 24), d(2019, 9, 24)},
		{HeritageDay, 2020, d(2020, 9, 24), d(2020, 9, 24)},
		{HeritageDay, 2021, d(2021, 9, 24), d(2021, 9, 24)},
		{HeritageDay, 2022, d(2022, 9, 24), d(2022, 9, 24)},

		{ReconciliationDay, 2015, d(2015, 12, 16), d(2015, 12, 16)},
		{ReconciliationDay, 2016, d(2016, 12, 16), d(2016, 12, 16)},
		{ReconciliationDay, 2017, d(2017, 12, 16), d(2017, 12, 16)},
		{ReconciliationDay, 2018, d(2018, 12, 16), d(2018, 12, 17)},
		{ReconciliationDay, 2019, d(2019, 12, 16), d(2019, 12, 16)},
		{ReconciliationDay, 2020, d(2020, 12, 16), d(2020, 12, 16)},
		{ReconciliationDay, 2021, d(2021, 12, 16), d(2021, 12, 16)},
		{ReconciliationDay, 2022, d(2022, 12, 16), d(2022, 12, 16)},

		{ChristmasDay, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{ChristmasDay, 2016, d(2016, 12, 25), d(2016, 12, 26)},
		{ChristmasDay, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{ChristmasDay, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{ChristmasDay, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{ChristmasDay, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{ChristmasDay, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{ChristmasDay, 2022, d(2022, 12, 25), d(2022, 12, 26)},

		{GoodwillDay, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{GoodwillDay, 2016, d(2016, 12, 26), d(2016, 12, 27)},
		{GoodwillDay, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{GoodwillDay, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{GoodwillDay, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{GoodwillDay, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{GoodwillDay, 2021, d(2021, 12, 26), d(2021, 12, 27)},
		{GoodwillDay, 2022, d(2022, 12, 26), d(2022, 12, 27)},
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
