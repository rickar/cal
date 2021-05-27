// (c) Matías Blasi. Licensed under the BSD license (see LICENSE).

package ar

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
		{NewYear, 2023, d(2023, 1, 1), d(2023, 1, 1)},

		{CarnivalDay1, 2015, d(2015, 2, 16), d(2015, 2, 16)},
		{CarnivalDay1, 2016, d(2016, 2, 8), d(2016, 2, 8)},
		{CarnivalDay1, 2017, d(2017, 2, 27), d(2017, 2, 27)},
		{CarnivalDay1, 2018, d(2018, 2, 12), d(2018, 2, 12)},
		{CarnivalDay1, 2019, d(2019, 3, 4), d(2019, 3, 4)},
		{CarnivalDay1, 2020, d(2020, 2, 24), d(2020, 2, 24)},
		{CarnivalDay1, 2021, d(2021, 2, 15), d(2021, 2, 15)},
		{CarnivalDay1, 2022, d(2022, 2, 28), d(2022, 2, 28)},
		{CarnivalDay1, 2023, d(2023, 2, 20), d(2023, 2, 20)},

		{CarnivalDay2, 2015, d(2015, 2, 17), d(2015, 2, 17)},
		{CarnivalDay2, 2016, d(2016, 2, 9), d(2016, 2, 9)},
		{CarnivalDay2, 2017, d(2017, 2, 28), d(2017, 2, 28)},
		{CarnivalDay2, 2018, d(2018, 2, 13), d(2018, 2, 13)},
		{CarnivalDay2, 2019, d(2019, 3, 5), d(2019, 3, 5)},
		{CarnivalDay2, 2020, d(2020, 2, 25), d(2020, 2, 25)},
		{CarnivalDay2, 2021, d(2021, 2, 16), d(2021, 2, 16)},
		{CarnivalDay2, 2022, d(2022, 3, 1), d(2022, 3, 1)},
		{CarnivalDay2, 2023, d(2023, 2, 21), d(2023, 2, 21)},

		{TruethDay, 2015, d(2015, 3, 24), d(2015, 3, 24)},
		{TruethDay, 2016, d(2016, 3, 24), d(2016, 3, 24)},
		{TruethDay, 2017, d(2017, 3, 24), d(2017, 3, 24)},
		{TruethDay, 2018, d(2018, 3, 24), d(2018, 3, 24)},
		{TruethDay, 2019, d(2019, 3, 24), d(2019, 3, 24)},
		{TruethDay, 2020, d(2020, 3, 24), d(2020, 3, 24)},
		{TruethDay, 2021, d(2021, 3, 24), d(2021, 3, 24)},
		{TruethDay, 2022, d(2022, 3, 24), d(2022, 3, 24)},
		{TruethDay, 2023, d(2023, 3, 24), d(2023, 3, 24)},

		{MalvinasVeterans, 2015, d(2015, 4, 2), d(2015, 4, 2)},
		{MalvinasVeterans, 2016, d(2016, 4, 2), d(2016, 4, 2)},
		{MalvinasVeterans, 2017, d(2017, 4, 2), d(2017, 4, 2)},
		{MalvinasVeterans, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{MalvinasVeterans, 2019, d(2019, 4, 2), d(2019, 4, 2)},
		{MalvinasVeterans, 2020, d(2020, 4, 2), d(2020, 4, 2)},
		{MalvinasVeterans, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{MalvinasVeterans, 2022, d(2022, 4, 2), d(2022, 4, 2)},
		{MalvinasVeterans, 2023, d(2023, 4, 2), d(2023, 4, 2)},

		{EasternsDay, 2015, d(2015, 4, 2), d(2015, 4, 2)},
		{EasternsDay, 2016, d(2016, 3, 24), d(2016, 3, 24)},
		{EasternsDay, 2017, d(2017, 4, 13), d(2017, 4, 13)},
		{EasternsDay, 2018, d(2018, 3, 29), d(2018, 3, 29)},
		{EasternsDay, 2019, d(2019, 4, 18), d(2019, 4, 18)},
		{EasternsDay, 2020, d(2020, 4, 9), d(2020, 4, 9)},
		{EasternsDay, 2021, d(2021, 4, 1), d(2021, 4, 1)},
		{EasternsDay, 2022, d(2022, 4, 14), d(2022, 4, 14)},
		{EasternsDay, 2023, d(2023, 4, 6), d(2023, 4, 6)},

		{LaborDay, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{LaborDay, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{LaborDay, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{LaborDay, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{LaborDay, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{LaborDay, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{LaborDay, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{LaborDay, 2022, d(2022, 5, 1), d(2022, 5, 1)},
		{LaborDay, 2023, d(2023, 5, 1), d(2023, 5, 1)},

		{RevolutionDay, 2015, d(2015, 5, 25), d(2015, 5, 25)},
		{RevolutionDay, 2016, d(2016, 5, 25), d(2016, 5, 25)},
		{RevolutionDay, 2017, d(2017, 5, 25), d(2017, 5, 25)},
		{RevolutionDay, 2018, d(2018, 5, 25), d(2018, 5, 25)},
		{RevolutionDay, 2019, d(2019, 5, 25), d(2019, 5, 25)},
		{RevolutionDay, 2020, d(2020, 5, 25), d(2020, 5, 25)},
		{RevolutionDay, 2021, d(2021, 5, 25), d(2021, 5, 25)},
		{RevolutionDay, 2022, d(2022, 5, 25), d(2022, 5, 25)},
		{RevolutionDay, 2023, d(2023, 5, 25), d(2023, 5, 25)},

		{GuemesDay, 2015, d(2015, 6, 17), d(2015, 6, 17)},
		{GuemesDay, 2016, d(2016, 6, 17), d(2016, 6, 17)},
		{GuemesDay, 2017, d(2017, 6, 17), d(2017, 6, 16)},
		{GuemesDay, 2018, d(2018, 6, 17), d(2018, 6, 18)},
		{GuemesDay, 2019, d(2019, 6, 17), d(2019, 6, 17)},
		{GuemesDay, 2020, d(2020, 6, 17), d(2020, 6, 17)},
		{GuemesDay, 2021, d(2021, 6, 17), d(2021, 6, 17)},
		{GuemesDay, 2022, d(2022, 6, 17), d(2022, 6, 17)},
		{GuemesDay, 2023, d(2023, 6, 17), d(2023, 6, 16)},

		{BelgranoDay, 2015, d(2015, 6, 20), d(2015, 6, 20)},
		{BelgranoDay, 2016, d(2016, 6, 20), d(2016, 6, 20)},
		{BelgranoDay, 2017, d(2017, 6, 20), d(2017, 6, 20)},
		{BelgranoDay, 2018, d(2018, 6, 20), d(2018, 6, 20)},
		{BelgranoDay, 2019, d(2019, 6, 20), d(2019, 6, 20)},
		{BelgranoDay, 2020, d(2020, 6, 20), d(2020, 6, 20)},
		{BelgranoDay, 2021, d(2021, 6, 20), d(2021, 6, 20)},
		{BelgranoDay, 2022, d(2022, 6, 20), d(2022, 6, 20)},
		{BelgranoDay, 2023, d(2023, 6, 20), d(2023, 6, 20)},

		{IndependenceDay, 2015, d(2015, 7, 9), d(2015, 7, 9)},
		{IndependenceDay, 2016, d(2016, 7, 9), d(2016, 7, 9)},
		{IndependenceDay, 2017, d(2017, 7, 9), d(2017, 7, 9)},
		{IndependenceDay, 2018, d(2018, 7, 9), d(2018, 7, 9)},
		{IndependenceDay, 2019, d(2019, 7, 9), d(2019, 7, 9)},
		{IndependenceDay, 2020, d(2020, 7, 9), d(2020, 7, 9)},
		{IndependenceDay, 2021, d(2021, 7, 9), d(2021, 7, 9)},
		{IndependenceDay, 2022, d(2022, 7, 9), d(2022, 7, 9)},
		{IndependenceDay, 2023, d(2023, 7, 9), d(2023, 7, 9)},

		{SanMartinDay, 2015, d(2015, 8, 17), d(2015, 8, 17)},
		{SanMartinDay, 2016, d(2016, 8, 17), d(2016, 8, 17)},
		{SanMartinDay, 2017, d(2017, 8, 17), d(2017, 8, 17)},
		{SanMartinDay, 2018, d(2018, 8, 17), d(2018, 8, 17)},
		{SanMartinDay, 2019, d(2019, 8, 17), d(2019, 8, 16)},
		{SanMartinDay, 2020, d(2020, 8, 17), d(2020, 8, 17)},
		{SanMartinDay, 2021, d(2021, 8, 17), d(2021, 8, 17)},
		{SanMartinDay, 2022, d(2022, 8, 17), d(2022, 8, 17)},
		{SanMartinDay, 2023, d(2023, 8, 17), d(2023, 8, 17)},

		{DiversityDay, 2015, d(2015, 10, 12), d(2015, 10, 12)},
		{DiversityDay, 2016, d(2016, 10, 12), d(2016, 10, 12)},
		{DiversityDay, 2017, d(2017, 10, 12), d(2017, 10, 12)},
		{DiversityDay, 2018, d(2018, 10, 12), d(2018, 10, 12)},
		{DiversityDay, 2019, d(2019, 10, 12), d(2019, 10, 11)},
		{DiversityDay, 2020, d(2020, 10, 12), d(2020, 10, 12)},
		{DiversityDay, 2021, d(2021, 10, 12), d(2021, 10, 12)},
		{DiversityDay, 2022, d(2022, 10, 12), d(2022, 10, 12)},
		{DiversityDay, 2023, d(2023, 10, 12), d(2023, 10, 12)},

		{SovereigntyDay, 2015, d(2015, 11, 20), d(2015, 11, 20)},
		{SovereigntyDay, 2016, d(2016, 11, 20), d(2016, 11, 21)},
		{SovereigntyDay, 2017, d(2017, 11, 20), d(2017, 11, 20)},
		{SovereigntyDay, 2018, d(2018, 11, 20), d(2018, 11, 20)},
		{SovereigntyDay, 2019, d(2019, 11, 20), d(2019, 11, 20)},
		{SovereigntyDay, 2020, d(2020, 11, 20), d(2020, 11, 20)},
		{SovereigntyDay, 2021, d(2021, 11, 20), d(2021, 11, 19)},
		{SovereigntyDay, 2022, d(2022, 11, 20), d(2022, 11, 21)},
		{SovereigntyDay, 2023, d(2023, 11, 20), d(2023, 11, 20)},

		{VirgenDay, 2015, d(2015, 12, 8), d(2015, 12, 8)},
		{VirgenDay, 2016, d(2016, 12, 8), d(2016, 12, 8)},
		{VirgenDay, 2017, d(2017, 12, 8), d(2017, 12, 8)},
		{VirgenDay, 2018, d(2018, 12, 8), d(2018, 12, 8)},
		{VirgenDay, 2019, d(2019, 12, 8), d(2019, 12, 8)},
		{VirgenDay, 2020, d(2020, 12, 8), d(2020, 12, 8)},
		{VirgenDay, 2021, d(2021, 12, 8), d(2021, 12, 8)},
		{VirgenDay, 2022, d(2022, 12, 8), d(2022, 12, 8)},
		{VirgenDay, 2023, d(2023, 12, 8), d(2023, 12, 8)},

		{ChristmasDay, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{ChristmasDay, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{ChristmasDay, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{ChristmasDay, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{ChristmasDay, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{ChristmasDay, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{ChristmasDay, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{ChristmasDay, 2022, d(2022, 12, 25), d(2022, 12, 25)},
		{ChristmasDay, 2023, d(2023, 12, 25), d(2023, 12, 25)},
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
