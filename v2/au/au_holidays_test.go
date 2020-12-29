// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package au

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

		{AustraliaDay, 2015, d(2015, 1, 26), d(2015, 1, 26)},
		{AustraliaDay, 2016, d(2016, 1, 26), d(2016, 1, 26)},
		{AustraliaDay, 2017, d(2017, 1, 26), d(2017, 1, 26)},
		{AustraliaDay, 2018, d(2018, 1, 26), d(2018, 1, 26)},
		{AustraliaDay, 2019, d(2019, 1, 26), d(2019, 1, 28)},
		{AustraliaDay, 2020, d(2020, 1, 26), d(2020, 1, 27)},
		{AustraliaDay, 2021, d(2021, 1, 26), d(2021, 1, 26)},
		{AustraliaDay, 2022, d(2022, 1, 26), d(2022, 1, 26)},

		{GoodFriday, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{GoodFriday, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{GoodFriday, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{GoodFriday, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{GoodFriday, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{GoodFriday, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{GoodFriday, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{GoodFriday, 2022, d(2022, 4, 15), d(2022, 4, 15)},

		{EasterSaturday, 2015, d(2015, 4, 4), d(2015, 4, 4)},
		{EasterSaturday, 2016, d(2016, 3, 26), d(2016, 3, 26)},
		{EasterSaturday, 2017, d(2017, 4, 15), d(2017, 4, 15)},
		{EasterSaturday, 2018, d(2018, 3, 31), d(2018, 3, 31)},
		{EasterSaturday, 2019, d(2019, 4, 20), d(2019, 4, 20)},
		{EasterSaturday, 2020, d(2020, 4, 11), d(2020, 4, 11)},
		{EasterSaturday, 2021, d(2021, 4, 3), d(2021, 4, 3)},
		{EasterSaturday, 2022, d(2022, 4, 16), d(2022, 4, 16)},

		{EasterSunday, 2015, d(2015, 4, 5), d(2015, 4, 5)},
		{EasterSunday, 2016, d(2016, 3, 27), d(2016, 3, 27)},
		{EasterSunday, 2017, d(2017, 4, 16), d(2017, 4, 16)},
		{EasterSunday, 2018, d(2018, 4, 1), d(2018, 4, 1)},
		{EasterSunday, 2019, d(2019, 4, 21), d(2019, 4, 21)},
		{EasterSunday, 2020, d(2020, 4, 12), d(2020, 4, 12)},
		{EasterSunday, 2021, d(2021, 4, 4), d(2021, 4, 4)},
		{EasterSunday, 2022, d(2022, 4, 17), d(2022, 4, 17)},

		{EasterMonday, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{EasterMonday, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{EasterMonday, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{EasterMonday, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{EasterMonday, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{EasterMonday, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{EasterMonday, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{EasterMonday, 2022, d(2022, 4, 18), d(2022, 4, 18)},

		{LabourDayWa, 2015, d(2015, 3, 2), d(2015, 3, 2)},
		{LabourDayWa, 2016, d(2016, 3, 7), d(2016, 3, 7)},
		{LabourDayWa, 2017, d(2017, 3, 6), d(2017, 3, 6)},
		{LabourDayWa, 2018, d(2018, 3, 5), d(2018, 3, 5)},
		{LabourDayWa, 2019, d(2019, 3, 4), d(2019, 3, 4)},
		{LabourDayWa, 2020, d(2020, 3, 2), d(2020, 3, 2)},
		{LabourDayWa, 2021, d(2021, 3, 1), d(2021, 3, 1)},
		{LabourDayWa, 2022, d(2022, 3, 7), d(2022, 3, 7)},

		{LabourDayVic, 2015, d(2015, 3, 9), d(2015, 3, 9)},
		{LabourDayVic, 2016, d(2016, 3, 14), d(2016, 3, 14)},
		{LabourDayVic, 2017, d(2017, 3, 13), d(2017, 3, 13)},
		{LabourDayVic, 2018, d(2018, 3, 12), d(2018, 3, 12)},
		{LabourDayVic, 2019, d(2019, 3, 11), d(2019, 3, 11)},
		{LabourDayVic, 2020, d(2020, 3, 9), d(2020, 3, 9)},
		{LabourDayVic, 2021, d(2021, 3, 8), d(2021, 3, 8)},
		{LabourDayVic, 2022, d(2022, 3, 14), d(2022, 3, 14)},

		{LabourDayTas, 2015, d(2015, 3, 9), d(2015, 3, 9)},
		{LabourDayTas, 2016, d(2016, 3, 14), d(2016, 3, 14)},
		{LabourDayTas, 2017, d(2017, 3, 13), d(2017, 3, 13)},
		{LabourDayTas, 2018, d(2018, 3, 12), d(2018, 3, 12)},
		{LabourDayTas, 2019, d(2019, 3, 11), d(2019, 3, 11)},
		{LabourDayTas, 2020, d(2020, 3, 9), d(2020, 3, 9)},
		{LabourDayTas, 2021, d(2021, 3, 8), d(2021, 3, 8)},
		{LabourDayTas, 2022, d(2022, 3, 14), d(2022, 3, 14)},

		{CanberraDay, 2015, d(2015, 3, 9), d(2015, 3, 9)},
		{CanberraDay, 2016, d(2016, 3, 14), d(2016, 3, 14)},
		{CanberraDay, 2017, d(2017, 3, 13), d(2017, 3, 13)},
		{CanberraDay, 2018, d(2018, 3, 12), d(2018, 3, 12)},
		{CanberraDay, 2019, d(2019, 3, 11), d(2019, 3, 11)},
		{CanberraDay, 2020, d(2020, 3, 9), d(2020, 3, 9)},
		{CanberraDay, 2021, d(2021, 3, 8), d(2021, 3, 8)},
		{CanberraDay, 2022, d(2022, 3, 14), d(2022, 3, 14)},

		{MarchPublicHoliday, 2015, d(2015, 3, 9), d(2015, 3, 9)},
		{MarchPublicHoliday, 2016, d(2016, 3, 14), d(2016, 3, 14)},
		{MarchPublicHoliday, 2017, d(2017, 3, 13), d(2017, 3, 13)},
		{MarchPublicHoliday, 2018, d(2018, 3, 12), d(2018, 3, 12)},
		{MarchPublicHoliday, 2019, d(2019, 3, 11), d(2019, 3, 11)},
		{MarchPublicHoliday, 2020, d(2020, 3, 9), d(2020, 3, 9)},
		{MarchPublicHoliday, 2021, d(2021, 3, 8), d(2021, 3, 8)},
		{MarchPublicHoliday, 2022, d(2022, 3, 14), d(2022, 3, 14)},

		{AnzacDay, 2015, d(2015, 4, 25), d(2015, 4, 25)},
		{AnzacDay, 2016, d(2016, 4, 25), d(2016, 4, 25)},
		{AnzacDay, 2017, d(2017, 4, 25), d(2017, 4, 25)},
		{AnzacDay, 2018, d(2018, 4, 25), d(2018, 4, 25)},
		{AnzacDay, 2019, d(2019, 4, 25), d(2019, 4, 25)},
		{AnzacDay, 2020, d(2020, 4, 25), d(2020, 4, 25)},
		{AnzacDay, 2021, d(2021, 4, 25), d(2021, 4, 25)},
		{AnzacDay, 2022, d(2022, 4, 25), d(2022, 4, 25)},

		{LabourDayNtQld, 2015, d(2015, 5, 4), d(2015, 5, 4)},
		{LabourDayNtQld, 2016, d(2016, 5, 2), d(2016, 5, 2)},
		{LabourDayNtQld, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{LabourDayNtQld, 2018, d(2018, 5, 7), d(2018, 5, 7)},
		{LabourDayNtQld, 2019, d(2019, 5, 6), d(2019, 5, 6)},
		{LabourDayNtQld, 2020, d(2020, 5, 4), d(2020, 5, 4)},
		{LabourDayNtQld, 2021, d(2021, 5, 3), d(2021, 5, 3)},
		{LabourDayNtQld, 2022, d(2022, 5, 2), d(2022, 5, 2)},

		{ReconciliationDay, 2015, time.Time{}, time.Time{}},
		{ReconciliationDay, 2016, time.Time{}, time.Time{}},
		{ReconciliationDay, 2017, time.Time{}, time.Time{}},
		{ReconciliationDay, 2018, d(2018, 5, 28), d(2018, 5, 28)},
		{ReconciliationDay, 2019, d(2019, 5, 27), d(2019, 5, 27)},
		{ReconciliationDay, 2020, d(2020, 6, 1), d(2020, 6, 1)},
		{ReconciliationDay, 2021, d(2021, 5, 31), d(2021, 5, 31)},
		{ReconciliationDay, 2022, d(2022, 5, 30), d(2022, 5, 30)},

		{WesternAustraliaDay, 2015, d(2015, 6, 1), d(2015, 6, 1)},
		{WesternAustraliaDay, 2016, d(2016, 6, 6), d(2016, 6, 6)},
		{WesternAustraliaDay, 2017, d(2017, 6, 5), d(2017, 6, 5)},
		{WesternAustraliaDay, 2018, d(2018, 6, 4), d(2018, 6, 4)},
		{WesternAustraliaDay, 2019, d(2019, 6, 3), d(2019, 6, 3)},
		{WesternAustraliaDay, 2020, d(2020, 6, 1), d(2020, 6, 1)},
		{WesternAustraliaDay, 2021, d(2021, 6, 7), d(2021, 6, 7)},
		{WesternAustraliaDay, 2022, d(2022, 6, 6), d(2022, 6, 6)},

		{QueensBirthday, 2015, d(2015, 6, 8), d(2015, 6, 8)},
		{QueensBirthday, 2016, d(2016, 6, 13), d(2016, 6, 13)},
		{QueensBirthday, 2017, d(2017, 6, 12), d(2017, 6, 12)},
		{QueensBirthday, 2018, d(2018, 6, 11), d(2018, 6, 11)},
		{QueensBirthday, 2019, d(2019, 6, 10), d(2019, 6, 10)},
		{QueensBirthday, 2020, d(2020, 6, 8), d(2020, 6, 8)},
		{QueensBirthday, 2021, d(2021, 6, 14), d(2021, 6, 14)},
		{QueensBirthday, 2022, d(2022, 6, 13), d(2022, 6, 13)},

		{PicnicDay, 2015, d(2015, 8, 3), d(2015, 8, 3)},
		{PicnicDay, 2016, d(2016, 8, 1), d(2016, 8, 1)},
		{PicnicDay, 2017, d(2017, 8, 7), d(2017, 8, 7)},
		{PicnicDay, 2018, d(2018, 8, 6), d(2018, 8, 6)},
		{PicnicDay, 2019, d(2019, 8, 5), d(2019, 8, 5)},
		{PicnicDay, 2020, d(2020, 8, 3), d(2020, 8, 3)},
		{PicnicDay, 2021, d(2021, 8, 2), d(2021, 8, 2)},
		{PicnicDay, 2022, d(2022, 8, 1), d(2022, 8, 1)},

		{QueensBirthdayWa, 2015, d(2015, 9, 28), d(2015, 9, 28)},
		{QueensBirthdayWa, 2016, d(2016, 9, 26), d(2016, 9, 26)},
		{QueensBirthdayWa, 2017, d(2017, 9, 25), d(2017, 9, 25)},
		{QueensBirthdayWa, 2018, d(2018, 9, 24), d(2018, 9, 24)},
		{QueensBirthdayWa, 2019, d(2019, 9, 30), d(2019, 9, 30)},
		{QueensBirthdayWa, 2020, d(2020, 9, 28), d(2020, 9, 28)},
		{QueensBirthdayWa, 2021, d(2021, 9, 27), d(2021, 9, 27)},
		{QueensBirthdayWa, 2022, d(2022, 9, 26), d(2022, 9, 26)},

		{FridayBeforeAflFinal, 2015, d(2015, 9, 25), d(2015, 9, 25)},
		{FridayBeforeAflFinal, 2016, d(2016, 9, 30), d(2016, 9, 30)},
		{FridayBeforeAflFinal, 2017, d(2017, 9, 29), d(2017, 9, 29)},
		{FridayBeforeAflFinal, 2018, d(2018, 9, 28), d(2018, 9, 28)},
		{FridayBeforeAflFinal, 2019, d(2019, 9, 27), d(2019, 9, 27)},
		{FridayBeforeAflFinal, 2020, d(2020, 10, 23), d(2020, 10, 23)},
		{FridayBeforeAflFinal, 2021, d(2021, 9, 24), d(2021, 9, 24)},
		{FridayBeforeAflFinal, 2022, d(2022, 9, 30), d(2022, 9, 30)},

		{QueensBirthdayQld, 2015, d(2015, 10, 5), d(2015, 10, 5)},
		{QueensBirthdayQld, 2016, d(2016, 10, 3), d(2016, 10, 3)},
		{QueensBirthdayQld, 2017, d(2017, 10, 2), d(2017, 10, 2)},
		{QueensBirthdayQld, 2018, d(2018, 10, 1), d(2018, 10, 1)},
		{QueensBirthdayQld, 2019, d(2019, 10, 7), d(2019, 10, 7)},
		{QueensBirthdayQld, 2020, d(2020, 10, 5), d(2020, 10, 5)},
		{QueensBirthdayQld, 2021, d(2021, 10, 4), d(2021, 10, 4)},
		{QueensBirthdayQld, 2022, d(2022, 10, 3), d(2022, 10, 3)},

		{LabourDayActNswSa, 2015, d(2015, 10, 5), d(2015, 10, 5)},
		{LabourDayActNswSa, 2016, d(2016, 10, 3), d(2016, 10, 3)},
		{LabourDayActNswSa, 2017, d(2017, 10, 2), d(2017, 10, 2)},
		{LabourDayActNswSa, 2018, d(2018, 10, 1), d(2018, 10, 1)},
		{LabourDayActNswSa, 2019, d(2019, 10, 7), d(2019, 10, 7)},
		{LabourDayActNswSa, 2020, d(2020, 10, 5), d(2020, 10, 5)},
		{LabourDayActNswSa, 2021, d(2021, 10, 4), d(2021, 10, 4)},
		{LabourDayActNswSa, 2022, d(2022, 10, 3), d(2022, 10, 3)},

		{MelbourneCup, 2015, d(2015, 11, 3), d(2015, 11, 3)},
		{MelbourneCup, 2016, d(2016, 11, 1), d(2016, 11, 1)},
		{MelbourneCup, 2017, d(2017, 11, 7), d(2017, 11, 7)},
		{MelbourneCup, 2018, d(2018, 11, 6), d(2018, 11, 6)},
		{MelbourneCup, 2019, d(2019, 11, 5), d(2019, 11, 5)},
		{MelbourneCup, 2020, d(2020, 11, 3), d(2020, 11, 3)},
		{MelbourneCup, 2021, d(2021, 11, 2), d(2021, 11, 2)},
		{MelbourneCup, 2022, d(2022, 11, 1), d(2022, 11, 1)},

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

		{ProclamationDay, 2015, d(2015, 12, 26), d(2015, 12, 28)},
		{ProclamationDay, 2016, d(2016, 12, 26), d(2016, 12, 27)},
		{ProclamationDay, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{ProclamationDay, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{ProclamationDay, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{ProclamationDay, 2020, d(2020, 12, 26), d(2020, 12, 28)},
		{ProclamationDay, 2021, d(2021, 12, 26), d(2021, 12, 28)},
		{ProclamationDay, 2022, d(2022, 12, 26), d(2022, 12, 27)},
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
