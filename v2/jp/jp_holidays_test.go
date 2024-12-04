// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package jp

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

		{ComingOfAgeDay, 2015, d(2015, 1, 12), d(2015, 1, 12)},
		{ComingOfAgeDay, 2016, d(2016, 1, 11), d(2016, 1, 11)},
		{ComingOfAgeDay, 2017, d(2017, 1, 9), d(2017, 1, 9)},
		{ComingOfAgeDay, 2018, d(2018, 1, 8), d(2018, 1, 8)},
		{ComingOfAgeDay, 2019, d(2019, 1, 14), d(2019, 1, 14)},
		{ComingOfAgeDay, 2020, d(2020, 1, 13), d(2020, 1, 13)},
		{ComingOfAgeDay, 2021, d(2021, 1, 11), d(2021, 1, 11)},
		{ComingOfAgeDay, 2022, d(2022, 1, 10), d(2022, 1, 10)},

		{NationalFoundationDay, 2015, d(2015, 2, 11), d(2015, 2, 11)},
		{NationalFoundationDay, 2016, d(2016, 2, 11), d(2016, 2, 11)},
		{NationalFoundationDay, 2017, d(2017, 2, 11), d(2017, 2, 11)},
		{NationalFoundationDay, 2018, d(2018, 2, 11), d(2018, 2, 12)},
		{NationalFoundationDay, 2019, d(2019, 2, 11), d(2019, 2, 11)},
		{NationalFoundationDay, 2020, d(2020, 2, 11), d(2020, 2, 11)},
		{NationalFoundationDay, 2021, d(2021, 2, 11), d(2021, 2, 11)},
		{NationalFoundationDay, 2022, d(2022, 2, 11), d(2022, 2, 11)},

		{TheEmperorsBirthday, 2015, d(2015, 12, 23), d(2015, 12, 23)},
		{TheEmperorsBirthday, 2016, d(2016, 12, 23), d(2016, 12, 23)},
		{TheEmperorsBirthday, 2017, d(2017, 12, 23), d(2017, 12, 23)},
		{TheEmperorsBirthday, 2018, d(2018, 12, 23), d(2018, 12, 24)},
		{TheEmperorsBirthday, 2019, time.Time{}, time.Time{}},
		{TheEmperorsBirthday, 2020, d(2020, 2, 23), d(2020, 2, 24)},
		{TheEmperorsBirthday, 2021, d(2021, 2, 23), d(2021, 2, 23)},
		{TheEmperorsBirthday, 2022, d(2022, 2, 23), d(2022, 2, 23)},

		{VernalEquinoxDay, 1851, d(1851, 3, 21), d(1851, 3, 21)},
		{VernalEquinoxDay, 1899, d(1899, 3, 21), d(1899, 3, 21)},
		{VernalEquinoxDay, 1900, d(1900, 3, 21), d(1900, 3, 21)},
		{VernalEquinoxDay, 1979, d(1979, 3, 21), d(1979, 3, 21)},
		{VernalEquinoxDay, 1980, d(1980, 3, 20), d(1980, 3, 20)},
		{VernalEquinoxDay, 2015, d(2015, 3, 21), d(2015, 3, 21)},
		{VernalEquinoxDay, 2016, d(2016, 3, 20), d(2016, 3, 21)},
		{VernalEquinoxDay, 2017, d(2017, 3, 20), d(2017, 3, 20)},
		{VernalEquinoxDay, 2018, d(2018, 3, 21), d(2018, 3, 21)},
		{VernalEquinoxDay, 2019, d(2019, 3, 21), d(2019, 3, 21)},
		{VernalEquinoxDay, 2020, d(2020, 3, 20), d(2020, 3, 20)},
		{VernalEquinoxDay, 2021, d(2021, 3, 20), d(2021, 3, 20)},
		{VernalEquinoxDay, 2022, d(2022, 3, 21), d(2022, 3, 21)},
		{VernalEquinoxDay, 2023, d(2023, 3, 21), d(2023, 3, 21)},
		{VernalEquinoxDay, 2024, d(2024, 3, 20), d(2024, 3, 20)},
		{VernalEquinoxDay, 2025, d(2025, 3, 20), d(2025, 3, 20)},
		{VernalEquinoxDay, 2026, d(2026, 3, 20), d(2026, 3, 20)},
		{VernalEquinoxDay, 2027, d(2027, 3, 21), d(2027, 3, 22)},
		{VernalEquinoxDay, 2028, d(2028, 3, 20), d(2028, 3, 20)},
		{VernalEquinoxDay, 2029, d(2029, 3, 20), d(2029, 3, 20)},
		{VernalEquinoxDay, 2030, d(2030, 3, 20), d(2030, 3, 20)},
		{VernalEquinoxDay, 2099, d(2099, 3, 20), d(2099, 3, 20)},
		{VernalEquinoxDay, 2100, d(2100, 3, 20), d(2100, 3, 20)},
		{VernalEquinoxDay, 2150, d(2150, 3, 21), d(2150, 3, 21)},

		{ShowaDay, 2015, d(2015, 4, 29), d(2015, 4, 29)},
		{ShowaDay, 2016, d(2016, 4, 29), d(2016, 4, 29)},
		{ShowaDay, 2017, d(2017, 4, 29), d(2017, 4, 29)},
		{ShowaDay, 2018, d(2018, 4, 29), d(2018, 4, 30)},
		{ShowaDay, 2019, d(2019, 4, 29), d(2019, 4, 29)},
		{ShowaDay, 2020, d(2020, 4, 29), d(2020, 4, 29)},
		{ShowaDay, 2021, d(2021, 4, 29), d(2021, 4, 29)},
		{ShowaDay, 2022, d(2022, 4, 29), d(2022, 4, 29)},

		{ConstitutionMemorialDay, 2015, d(2015, 5, 3), d(2015, 5, 6)},
		{ConstitutionMemorialDay, 2016, d(2016, 5, 3), d(2016, 5, 3)},
		{ConstitutionMemorialDay, 2017, d(2017, 5, 3), d(2017, 5, 3)},
		{ConstitutionMemorialDay, 2018, d(2018, 5, 3), d(2018, 5, 3)},
		{ConstitutionMemorialDay, 2019, d(2019, 5, 3), d(2019, 5, 3)},
		{ConstitutionMemorialDay, 2020, d(2020, 5, 3), d(2020, 5, 6)},
		{ConstitutionMemorialDay, 2021, d(2021, 5, 3), d(2021, 5, 3)},
		{ConstitutionMemorialDay, 2022, d(2022, 5, 3), d(2022, 5, 3)},

		{GreeneryDay, 2014, d(2014, 5, 4), d(2014, 5, 6)},
		{GreeneryDay, 2015, d(2015, 5, 4), d(2015, 5, 4)},
		{GreeneryDay, 2016, d(2016, 5, 4), d(2016, 5, 4)},
		{GreeneryDay, 2017, d(2017, 5, 4), d(2017, 5, 4)},
		{GreeneryDay, 2018, d(2018, 5, 4), d(2018, 5, 4)},
		{GreeneryDay, 2019, d(2019, 5, 4), d(2019, 5, 4)},
		{GreeneryDay, 2020, d(2020, 5, 4), d(2020, 5, 4)},
		{GreeneryDay, 2021, d(2021, 5, 4), d(2021, 5, 4)},
		{GreeneryDay, 2022, d(2022, 5, 4), d(2022, 5, 4)},

		{ChildrensDay, 2015, d(2015, 5, 5), d(2015, 5, 5)},
		{ChildrensDay, 2016, d(2016, 5, 5), d(2016, 5, 5)},
		{ChildrensDay, 2017, d(2017, 5, 5), d(2017, 5, 5)},
		{ChildrensDay, 2018, d(2018, 5, 5), d(2018, 5, 5)},
		{ChildrensDay, 2019, d(2019, 5, 5), d(2019, 5, 6)},
		{ChildrensDay, 2020, d(2020, 5, 5), d(2020, 5, 5)},
		{ChildrensDay, 2021, d(2021, 5, 5), d(2021, 5, 5)},
		{ChildrensDay, 2022, d(2022, 5, 5), d(2022, 5, 5)},

		{MarineDay, 2015, d(2015, 7, 20), d(2015, 7, 20)},
		{MarineDay, 2016, d(2016, 7, 18), d(2016, 7, 18)},
		{MarineDay, 2017, d(2017, 7, 17), d(2017, 7, 17)},
		{MarineDay, 2018, d(2018, 7, 16), d(2018, 7, 16)},
		{MarineDay, 2019, d(2019, 7, 15), d(2019, 7, 15)},
		{MarineDay, 2020, d(2020, 7, 23), d(2020, 7, 23)},
		{MarineDay, 2021, d(2021, 7, 22), d(2021, 7, 22)},
		{MarineDay, 2022, d(2022, 7, 18), d(2022, 7, 18)},

		{MountainDay, 2016, d(2016, 8, 11), d(2016, 8, 11)},
		{MountainDay, 2017, d(2017, 8, 11), d(2017, 8, 11)},
		{MountainDay, 2018, d(2018, 8, 11), d(2018, 8, 11)},
		{MountainDay, 2019, d(2019, 8, 11), d(2019, 8, 12)},
		{MountainDay, 2020, d(2020, 8, 11), d(2020, 8, 11)},
		{MountainDay, 2021, d(2021, 8, 11), d(2021, 8, 11)},
		{MountainDay, 2022, d(2022, 8, 11), d(2022, 8, 11)},

		{RespectForTheAgedDay, 2015, d(2015, 9, 21), d(2015, 9, 21)},
		{RespectForTheAgedDay, 2016, d(2016, 9, 19), d(2016, 9, 19)},
		{RespectForTheAgedDay, 2017, d(2017, 9, 18), d(2017, 9, 18)},
		{RespectForTheAgedDay, 2018, d(2018, 9, 17), d(2018, 9, 17)},
		{RespectForTheAgedDay, 2019, d(2019, 9, 16), d(2019, 9, 16)},
		{RespectForTheAgedDay, 2020, d(2020, 9, 21), d(2020, 9, 21)},
		{RespectForTheAgedDay, 2021, d(2021, 9, 20), d(2021, 9, 20)},
		{RespectForTheAgedDay, 2022, d(2022, 9, 19), d(2022, 9, 19)},

		{AutumnalEquinoxDay, 1851, d(1851, 9, 24), d(1851, 9, 24)},
		{AutumnalEquinoxDay, 1899, d(1899, 9, 23), d(1899, 9, 23)},
		{AutumnalEquinoxDay, 1900, d(1900, 9, 23), d(1900, 9, 24)},
		{AutumnalEquinoxDay, 1979, d(1979, 9, 24), d(1979, 9, 24)},
		{AutumnalEquinoxDay, 1980, d(1980, 9, 23), d(1980, 9, 23)},
		{AutumnalEquinoxDay, 2015, d(2015, 9, 23), d(2015, 9, 23)},
		{AutumnalEquinoxDay, 2016, d(2016, 9, 22), d(2016, 9, 22)},
		{AutumnalEquinoxDay, 2017, d(2017, 9, 23), d(2017, 9, 23)},
		{AutumnalEquinoxDay, 2018, d(2018, 9, 23), d(2018, 9, 24)},
		{AutumnalEquinoxDay, 2019, d(2019, 9, 23), d(2019, 9, 23)},
		{AutumnalEquinoxDay, 2020, d(2020, 9, 22), d(2020, 9, 22)},
		{AutumnalEquinoxDay, 2021, d(2021, 9, 23), d(2021, 9, 23)},
		{AutumnalEquinoxDay, 2022, d(2022, 9, 23), d(2022, 9, 23)},
		{AutumnalEquinoxDay, 2023, d(2023, 9, 23), d(2023, 9, 23)},
		{AutumnalEquinoxDay, 2024, d(2024, 9, 22), d(2024, 9, 23)},
		{AutumnalEquinoxDay, 2025, d(2025, 9, 23), d(2025, 9, 23)},
		{AutumnalEquinoxDay, 2026, d(2026, 9, 23), d(2026, 9, 23)},
		{AutumnalEquinoxDay, 2027, d(2027, 9, 23), d(2027, 9, 23)},
		{AutumnalEquinoxDay, 2028, d(2028, 9, 22), d(2028, 9, 22)},
		{AutumnalEquinoxDay, 2029, d(2029, 9, 23), d(2029, 9, 24)},
		{AutumnalEquinoxDay, 2030, d(2030, 9, 23), d(2030, 9, 23)},
		{AutumnalEquinoxDay, 2099, d(2099, 9, 23), d(2099, 9, 23)},
		{AutumnalEquinoxDay, 2100, d(2100, 9, 23), d(2100, 9, 23)},
		{AutumnalEquinoxDay, 2150, d(2150, 9, 23), d(2150, 9, 23)},

		{SportsDay, 2015, d(2015, 10, 12), d(2015, 10, 12)},
		{SportsDay, 2016, d(2016, 10, 10), d(2016, 10, 10)},
		{SportsDay, 2017, d(2017, 10, 9), d(2017, 10, 9)},
		{SportsDay, 2018, d(2018, 10, 8), d(2018, 10, 8)},
		{SportsDay, 2019, d(2019, 10, 14), d(2019, 10, 14)},
		{SportsDay, 2020, d(2020, 7, 24), d(2020, 7, 24)},
		{SportsDay, 2021, d(2021, 7, 23), d(2021, 7, 23)},
		{SportsDay, 2022, d(2022, 10, 10), d(2022, 10, 10)},

		{CultureDay, 2015, d(2015, 11, 3), d(2015, 11, 3)},
		{CultureDay, 2016, d(2016, 11, 3), d(2016, 11, 3)},
		{CultureDay, 2017, d(2017, 11, 3), d(2017, 11, 3)},
		{CultureDay, 2018, d(2018, 11, 3), d(2018, 11, 3)},
		{CultureDay, 2019, d(2019, 11, 3), d(2019, 11, 4)},
		{CultureDay, 2020, d(2020, 11, 3), d(2020, 11, 3)},
		{CultureDay, 2021, d(2021, 11, 3), d(2021, 11, 3)},
		{CultureDay, 2022, d(2022, 11, 3), d(2022, 11, 3)},

		{LaborThanksgivingDay, 2014, d(2014, 11, 23), d(2014, 11, 24)},
		{LaborThanksgivingDay, 2015, d(2015, 11, 23), d(2015, 11, 23)},
		{LaborThanksgivingDay, 2016, d(2016, 11, 23), d(2016, 11, 23)},
		{LaborThanksgivingDay, 2017, d(2017, 11, 23), d(2017, 11, 23)},
		{LaborThanksgivingDay, 2018, d(2018, 11, 23), d(2018, 11, 23)},
		{LaborThanksgivingDay, 2019, d(2019, 11, 23), d(2019, 11, 23)},
		{LaborThanksgivingDay, 2020, d(2020, 11, 23), d(2020, 11, 23)},
		{LaborThanksgivingDay, 2021, d(2021, 11, 23), d(2021, 11, 23)},
		{LaborThanksgivingDay, 2022, d(2022, 11, 23), d(2022, 11, 23)},

		{NationalHolidayBetweenRespectForTheAgedDayAndAutumnalEquinoxDay, 2009, d(2009, 9, 22), d(2009, 9, 22)},
		{NationalHolidayBetweenRespectForTheAgedDayAndAutumnalEquinoxDay, 2015, d(2015, 9, 22), d(2015, 9, 22)},
		{NationalHolidayBetweenRespectForTheAgedDayAndAutumnalEquinoxDay, 2026, d(2026, 9, 22), d(2026, 9, 22)},
		{NationalHolidayBetweenRespectForTheAgedDayAndAutumnalEquinoxDay, 2022, time.Time{}, time.Time{}},
		{NationalHolidayBetweenRespectForTheAgedDayAndAutumnalEquinoxDay, 2032, d(2032, 9, 21), d(2032, 9, 21)},

		{NationalHolidayBetweenShowaDayAndNewEmperorEnthronementDay, 2019, d(2019, 4, 30), d(2019, 4, 30)},

		{TheNewEmperorEnthronementDay, 2019, d(2019, 5, 1), d(2019, 5, 1)},

		{NationalHolidayBetweenTheNewEmperorEnthronementDayAndConstitutionMemorialDay, 2019, d(2019, 5, 2), d(2019, 5, 2)},

		{TheNewEmperorEnthronementCeremony, 2019, d(2019, 10, 22), d(2019, 10, 22)},
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
