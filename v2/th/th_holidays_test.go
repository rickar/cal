package th

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
		{NewYear, 2017, d(2017, 1, 1), d(2017, 1, 2)},
		{NewYear, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{NewYear, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{NewYear, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{NewYear, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{NewYear, 2022, d(2022, 1, 1), d(2022, 1, 3)},
		{NewYear, 2023, d(2023, 1, 1), d(2023, 1, 2)},
		{ChakriDay, 2017, d(2017, 4, 6), d(2017, 4, 6)},
		{ChakriDay, 2018, d(2018, 4, 6), d(2018, 4, 6)},
		{ChakriDay, 2019, d(2019, 4, 6), d(2019, 4, 8)},
		{ChakriDay, 2020, d(2020, 4, 6), d(2020, 4, 6)},
		{ChakriDay, 2021, d(2021, 4, 6), d(2021, 4, 6)},
		{ChakriDay, 2022, d(2022, 4, 6), d(2022, 4, 6)},
		{ChakriDay, 2023, d(2023, 4, 6), d(2023, 4, 6)},
		{ElderlyDay, 2017, d(2017, 4, 13), d(2017, 4, 13)},
		{ElderlyDay, 2018, d(2018, 4, 13), d(2018, 4, 13)},
		{ElderlyDay, 2019, d(2019, 4, 13), d(2019, 4, 15)},
		{ElderlyDay, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{ElderlyDay, 2021, d(2021, 4, 13), d(2021, 4, 13)},
		{ElderlyDay, 2022, d(2022, 4, 13), d(2022, 4, 13)},
		{ElderlyDay, 2023, d(2023, 4, 13), d(2023, 4, 13)},
		{FamilyDay, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{FamilyDay, 2018, d(2018, 4, 14), d(2018, 4, 16)},
		{FamilyDay, 2019, d(2019, 4, 14), d(2019, 4, 16)},
		{FamilyDay, 2020, d(2020, 4, 14), d(2020, 4, 14)},
		{FamilyDay, 2021, d(2021, 4, 14), d(2021, 4, 14)},
		{FamilyDay, 2022, d(2022, 4, 14), d(2022, 4, 14)},
		{FamilyDay, 2023, d(2023, 4, 14), d(2023, 4, 14)},
		{ThaiNewYear, 2017, d(2017, 4, 15), d(2017, 4, 17)},
		{ThaiNewYear, 2018, d(2018, 4, 15), d(2018, 4, 17)},
		{ThaiNewYear, 2019, d(2019, 4, 15), d(2019, 4, 15)},
		{ThaiNewYear, 2020, d(2020, 4, 15), d(2020, 4, 15)},
		{ThaiNewYear, 2021, d(2021, 4, 15), d(2021, 4, 15)},
		{ThaiNewYear, 2022, d(2022, 4, 15), d(2022, 4, 15)},
		{ThaiNewYear, 2023, d(2023, 4, 15), d(2023, 4, 17)},
		{LaborDay, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{LaborDay, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{LaborDay, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{LaborDay, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{LaborDay, 2021, d(2021, 5, 1), d(2021, 5, 3)},
		{LaborDay, 2022, d(2022, 5, 1), d(2022, 5, 2)},
		{LaborDay, 2023, d(2023, 5, 1), d(2023, 5, 1)},
		{CoronationDay, 2019, d(2019, 5, 4), d(2019, 5, 6)},
		{CoronationDay, 2020, d(2020, 5, 4), d(2020, 5, 4)},
		{CoronationDay, 2021, d(2021, 5, 4), d(2021, 5, 4)},
		{CoronationDay, 2022, d(2022, 5, 4), d(2022, 5, 4)},
		{CoronationDay, 2023, d(2023, 5, 4), d(2023, 5, 4)},
		{QueensBirthday, 2019, d(2019, 6, 3), d(2019, 6, 3)},
		{QueensBirthday, 2020, d(2020, 6, 3), d(2020, 6, 3)},
		{QueensBirthday, 2021, d(2021, 6, 3), d(2021, 6, 3)},
		{QueensBirthday, 2022, d(2022, 6, 3), d(2022, 6, 3)},
		{QueensBirthday, 2023, d(2023, 6, 3), d(2023, 6, 5)},
		{KingsBirthday, 2017, d(2017, 7, 28), d(2017, 7, 28)},
		{KingsBirthday, 2018, d(2018, 7, 28), d(2018, 7, 30)},
		{KingsBirthday, 2019, d(2019, 7, 28), d(2019, 7, 29)},
		{KingsBirthday, 2020, d(2020, 7, 28), d(2020, 7, 28)},
		{KingsBirthday, 2021, d(2021, 7, 28), d(2021, 7, 28)},
		{KingsBirthday, 2022, d(2022, 7, 28), d(2022, 7, 28)},
		{KingsBirthday, 2023, d(2023, 7, 28), d(2023, 7, 28)},
		{MothersDay, 2017, d(2017, 8, 12), d(2017, 8, 14)},
		{MothersDay, 2018, d(2018, 8, 12), d(2018, 8, 13)},
		{MothersDay, 2019, d(2019, 8, 12), d(2019, 8, 12)},
		{MothersDay, 2020, d(2020, 8, 12), d(2020, 8, 12)},
		{MothersDay, 2021, d(2021, 8, 12), d(2021, 8, 12)},
		{MothersDay, 2022, d(2022, 8, 12), d(2022, 8, 12)},
		{MothersDay, 2023, d(2023, 8, 12), d(2023, 8, 14)},
		{KingBhumibolDay, 2017, d(2017, 10, 13), d(2017, 10, 13)},
		{KingBhumibolDay, 2018, d(2018, 10, 13), d(2018, 10, 15)},
		{KingBhumibolDay, 2019, d(2019, 10, 13), d(2019, 10, 14)},
		{KingBhumibolDay, 2020, d(2020, 10, 13), d(2020, 10, 13)},
		{KingBhumibolDay, 2021, d(2021, 10, 13), d(2021, 10, 13)},
		{KingBhumibolDay, 2022, d(2022, 10, 13), d(2022, 10, 13)},
		{KingBhumibolDay, 2023, d(2023, 10, 13), d(2023, 10, 13)},
		{KingChulalongkornDay, 2017, d(2017, 10, 23), d(2017, 10, 23)},
		{KingChulalongkornDay, 2018, d(2018, 10, 23), d(2018, 10, 23)},
		{KingChulalongkornDay, 2019, d(2019, 10, 23), d(2019, 10, 23)},
		{KingChulalongkornDay, 2020, d(2020, 10, 23), d(2020, 10, 23)},
		{KingChulalongkornDay, 2021, d(2021, 10, 23), d(2021, 10, 25)},
		{KingChulalongkornDay, 2022, d(2022, 10, 23), d(2022, 10, 24)},
		{KingChulalongkornDay, 2023, d(2023, 10, 23), d(2023, 10, 23)},
		{FathersDay, 2017, d(2017, 12, 5), d(2017, 12, 5)},
		{FathersDay, 2018, d(2018, 12, 5), d(2018, 12, 5)},
		{FathersDay, 2019, d(2019, 12, 5), d(2019, 12, 5)},
		{FathersDay, 2020, d(2020, 12, 5), d(2020, 12, 7)},
		{FathersDay, 2021, d(2021, 12, 5), d(2021, 12, 6)},
		{FathersDay, 2022, d(2022, 12, 5), d(2022, 12, 5)},
		{FathersDay, 2023, d(2023, 12, 5), d(2023, 12, 5)},
		{ConstitutionDay, 2017, d(2017, 12, 10), d(2017, 12, 11)},
		{ConstitutionDay, 2018, d(2018, 12, 10), d(2018, 12, 10)},
		{ConstitutionDay, 2019, d(2019, 12, 10), d(2019, 12, 10)},
		{ConstitutionDay, 2020, d(2020, 12, 10), d(2020, 12, 10)},
		{ConstitutionDay, 2021, d(2021, 12, 10), d(2021, 12, 10)},
		{ConstitutionDay, 2022, d(2022, 12, 10), d(2022, 12, 12)},
		{ConstitutionDay, 2023, d(2023, 12, 10), d(2023, 12, 11)},
		{NewYearEve, 2017, d(2017, 12, 31), d(2017, 12, 31)},
		{NewYearEve, 2018, d(2018, 12, 31), d(2018, 12, 31)},
		{NewYearEve, 2019, d(2019, 12, 31), d(2019, 12, 31)},
		{NewYearEve, 2020, d(2020, 12, 31), d(2020, 12, 31)},
		{NewYearEve, 2021, d(2021, 12, 31), d(2021, 12, 31)},
		{NewYearEve, 2022, d(2022, 12, 31), d(2022, 12, 31)},
		{NewYearEve, 2023, d(2023, 12, 31), d(2023, 12, 31)},
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
