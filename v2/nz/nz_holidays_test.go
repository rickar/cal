// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package nz

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
		{NewYear, 2017, d(2017, 1, 1), d(2017, 1, 2)},
		{NewYear, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{NewYear, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{NewYear, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{NewYear, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{NewYear, 2022, d(2022, 1, 1), d(2022, 1, 3)},

		{DayAfterNewYear, 2015, d(2015, 1, 2), d(2015, 1, 2)},
		{DayAfterNewYear, 2016, d(2016, 1, 2), d(2016, 1, 4)},
		{DayAfterNewYear, 2017, d(2017, 1, 2), d(2017, 1, 3)},
		{DayAfterNewYear, 2018, d(2018, 1, 2), d(2018, 1, 2)},
		{DayAfterNewYear, 2019, d(2019, 1, 2), d(2019, 1, 2)},
		{DayAfterNewYear, 2020, d(2020, 1, 2), d(2020, 1, 2)},
		{DayAfterNewYear, 2021, d(2021, 1, 2), d(2021, 1, 4)},
		{DayAfterNewYear, 2022, d(2022, 1, 2), d(2022, 1, 4)},

		{WaitangiDay, 2015, d(2015, 2, 6), d(2015, 2, 6)},
		{WaitangiDay, 2016, d(2016, 2, 6), d(2016, 2, 8)},
		{WaitangiDay, 2017, d(2017, 2, 6), d(2017, 2, 6)},
		{WaitangiDay, 2018, d(2018, 2, 6), d(2018, 2, 6)},
		{WaitangiDay, 2019, d(2019, 2, 6), d(2019, 2, 6)},
		{WaitangiDay, 2020, d(2020, 2, 6), d(2020, 2, 6)},
		{WaitangiDay, 2021, d(2021, 2, 6), d(2021, 2, 8)},
		{WaitangiDay, 2022, d(2022, 2, 6), d(2022, 2, 7)},

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

		{AnzacDay, 2015, d(2015, 4, 25), d(2015, 4, 27)},
		{AnzacDay, 2016, d(2016, 4, 25), d(2016, 4, 25)},
		{AnzacDay, 2017, d(2017, 4, 25), d(2017, 4, 25)},
		{AnzacDay, 2018, d(2018, 4, 25), d(2018, 4, 25)},
		{AnzacDay, 2019, d(2019, 4, 25), d(2019, 4, 25)},
		{AnzacDay, 2020, d(2020, 4, 25), d(2020, 4, 27)},
		{AnzacDay, 2021, d(2021, 4, 25), d(2021, 4, 26)},
		{AnzacDay, 2022, d(2022, 4, 25), d(2022, 4, 25)},

		{QueensBirthday, 2015, d(2015, 6, 1), d(2015, 6, 1)},
		{QueensBirthday, 2016, d(2016, 6, 6), d(2016, 6, 6)},
		{QueensBirthday, 2017, d(2017, 6, 5), d(2017, 6, 5)},
		{QueensBirthday, 2018, d(2018, 6, 4), d(2018, 6, 4)},
		{QueensBirthday, 2019, d(2019, 6, 3), d(2019, 6, 3)},
		{QueensBirthday, 2020, d(2020, 6, 1), d(2020, 6, 1)},
		{QueensBirthday, 2021, d(2021, 6, 7), d(2021, 6, 7)},
		{QueensBirthday, 2022, d(2022, 6, 6), d(2022, 6, 6)},

		{LabourDay, 2015, d(2015, 10, 26), d(2015, 10, 26)},
		{LabourDay, 2016, d(2016, 10, 24), d(2016, 10, 24)},
		{LabourDay, 2017, d(2017, 10, 23), d(2017, 10, 23)},
		{LabourDay, 2018, d(2018, 10, 22), d(2018, 10, 22)},
		{LabourDay, 2019, d(2019, 10, 28), d(2019, 10, 28)},
		{LabourDay, 2020, d(2020, 10, 26), d(2020, 10, 26)},
		{LabourDay, 2021, d(2021, 10, 25), d(2021, 10, 25)},
		{LabourDay, 2022, d(2022, 10, 24), d(2022, 10, 24)},

		{ChristmasDay, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{ChristmasDay, 2016, d(2016, 12, 25), d(2016, 12, 26)},
		{ChristmasDay, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{ChristmasDay, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{ChristmasDay, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{ChristmasDay, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{ChristmasDay, 2021, d(2021, 12, 25), d(2021, 12, 27)},
		{ChristmasDay, 2022, d(2022, 12, 25), d(2022, 12, 26)},

		{BoxingDay, 2015, d(2015, 12, 26), d(2015, 12, 28)},
		{BoxingDay, 2016, d(2016, 12, 26), d(2016, 12, 27)},
		{BoxingDay, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{BoxingDay, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{BoxingDay, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{BoxingDay, 2020, d(2020, 12, 26), d(2020, 12, 28)},
		{BoxingDay, 2021, d(2021, 12, 26), d(2021, 12, 28)},
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
