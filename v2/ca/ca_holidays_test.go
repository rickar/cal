// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package ca

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

		{VictoriaDay, 2015, d(2015, 5, 18), d(2015, 5, 18)},
		{VictoriaDay, 2016, d(2016, 5, 23), d(2016, 5, 23)},
		{VictoriaDay, 2017, d(2017, 5, 22), d(2017, 5, 22)},
		{VictoriaDay, 2018, d(2018, 5, 21), d(2018, 5, 21)},
		{VictoriaDay, 2019, d(2019, 5, 20), d(2019, 5, 20)},
		{VictoriaDay, 2020, d(2020, 5, 18), d(2020, 5, 18)},
		{VictoriaDay, 2021, d(2021, 5, 24), d(2021, 5, 24)},
		{VictoriaDay, 2022, d(2022, 5, 23), d(2022, 5, 23)},

		{CanadaDay, 2015, d(2015, 7, 1), d(2015, 7, 1)},
		{CanadaDay, 2016, d(2016, 7, 1), d(2016, 7, 1)},
		{CanadaDay, 2017, d(2017, 7, 1), d(2017, 7, 1)},
		{CanadaDay, 2018, d(2018, 7, 1), d(2018, 7, 2)},
		{CanadaDay, 2019, d(2019, 7, 1), d(2019, 7, 1)},
		{CanadaDay, 2020, d(2020, 7, 1), d(2020, 7, 1)},
		{CanadaDay, 2021, d(2021, 7, 1), d(2021, 7, 1)},
		{CanadaDay, 2022, d(2022, 7, 1), d(2022, 7, 1)},

		{CivicDay, 2015, d(2015, 8, 3), d(2015, 8, 3)},
		{CivicDay, 2016, d(2016, 8, 1), d(2016, 8, 1)},
		{CivicDay, 2017, d(2017, 8, 7), d(2017, 8, 7)},
		{CivicDay, 2018, d(2018, 8, 6), d(2018, 8, 6)},
		{CivicDay, 2019, d(2019, 8, 5), d(2019, 8, 5)},
		{CivicDay, 2020, d(2020, 8, 3), d(2020, 8, 3)},
		{CivicDay, 2021, d(2021, 8, 2), d(2021, 8, 2)},
		{CivicDay, 2022, d(2022, 8, 1), d(2022, 8, 1)},

		{LabourDay, 2015, d(2015, 9, 7), d(2015, 9, 7)},
		{LabourDay, 2016, d(2016, 9, 5), d(2016, 9, 5)},
		{LabourDay, 2017, d(2017, 9, 4), d(2017, 9, 4)},
		{LabourDay, 2018, d(2018, 9, 3), d(2018, 9, 3)},
		{LabourDay, 2019, d(2019, 9, 2), d(2019, 9, 2)},
		{LabourDay, 2020, d(2020, 9, 7), d(2020, 9, 7)},
		{LabourDay, 2021, d(2021, 9, 6), d(2021, 9, 6)},
		{LabourDay, 2022, d(2022, 9, 5), d(2022, 9, 5)},

		{ThanksgivingDay, 2015, d(2015, 10, 12), d(2015, 10, 12)},
		{ThanksgivingDay, 2016, d(2016, 10, 10), d(2016, 10, 10)},
		{ThanksgivingDay, 2017, d(2017, 10, 9), d(2017, 10, 9)},
		{ThanksgivingDay, 2018, d(2018, 10, 8), d(2018, 10, 8)},
		{ThanksgivingDay, 2019, d(2019, 10, 14), d(2019, 10, 14)},
		{ThanksgivingDay, 2020, d(2020, 10, 12), d(2020, 10, 12)},
		{ThanksgivingDay, 2021, d(2021, 10, 11), d(2021, 10, 11)},
		{ThanksgivingDay, 2022, d(2022, 10, 10), d(2022, 10, 10)},

		{RemembranceDay, 2015, d(2015, 11, 11), d(2015, 11, 11)},
		{RemembranceDay, 2016, d(2016, 11, 11), d(2016, 11, 11)},
		{RemembranceDay, 2017, d(2017, 11, 11), d(2017, 11, 11)},
		{RemembranceDay, 2018, d(2018, 11, 11), d(2018, 11, 12)},
		{RemembranceDay, 2019, d(2019, 11, 11), d(2019, 11, 11)},
		{RemembranceDay, 2020, d(2020, 11, 11), d(2020, 11, 11)},
		{RemembranceDay, 2021, d(2021, 11, 11), d(2021, 11, 11)},
		{RemembranceDay, 2022, d(2022, 11, 11), d(2022, 11, 11)},

		{ChristmasDay, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{ChristmasDay, 2016, d(2016, 12, 25), d(2016, 12, 26)},
		{ChristmasDay, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{ChristmasDay, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{ChristmasDay, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{ChristmasDay, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{ChristmasDay, 2021, d(2021, 12, 25), d(2021, 12, 24)},
		{ChristmasDay, 2022, d(2022, 12, 25), d(2022, 12, 26)},

		{BoxingDay, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{BoxingDay, 2016, d(2016, 12, 26), d(2016, 12, 27)},
		{BoxingDay, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{BoxingDay, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{BoxingDay, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{BoxingDay, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{BoxingDay, 2021, d(2021, 12, 26), d(2021, 12, 26)},
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
