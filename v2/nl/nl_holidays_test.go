// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package nl

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
		{Nieuwjaar, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{Nieuwjaar, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{Nieuwjaar, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{Nieuwjaar, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{Nieuwjaar, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{Nieuwjaar, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{Nieuwjaar, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{Nieuwjaar, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{GoedeVrijdag, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{GoedeVrijdag, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{GoedeVrijdag, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{GoedeVrijdag, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{GoedeVrijdag, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{GoedeVrijdag, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{GoedeVrijdag, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{GoedeVrijdag, 2022, d(2022, 4, 15), d(2022, 4, 15)},

		{EerstePaasdag, 2015, d(2015, 4, 5), d(2015, 4, 5)},
		{EerstePaasdag, 2016, d(2016, 3, 27), d(2016, 3, 27)},
		{EerstePaasdag, 2017, d(2017, 4, 16), d(2017, 4, 16)},
		{EerstePaasdag, 2018, d(2018, 4, 1), d(2018, 4, 1)},
		{EerstePaasdag, 2019, d(2019, 4, 21), d(2019, 4, 21)},
		{EerstePaasdag, 2020, d(2020, 4, 12), d(2020, 4, 12)},
		{EerstePaasdag, 2021, d(2021, 4, 4), d(2021, 4, 4)},
		{EerstePaasdag, 2022, d(2022, 4, 17), d(2022, 4, 17)},
		{EerstePaasdag, 2023, d(2023, 4, 9), d(2023, 4, 9)},

		{TweedePaasdag, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{TweedePaasdag, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{TweedePaasdag, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{TweedePaasdag, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{TweedePaasdag, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{TweedePaasdag, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{TweedePaasdag, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{TweedePaasdag, 2022, d(2022, 4, 18), d(2022, 4, 18)},

		{Koningsdag, 2015, d(2015, 4, 27), d(2015, 4, 27)},
		{Koningsdag, 2016, d(2016, 4, 27), d(2016, 4, 27)},
		{Koningsdag, 2017, d(2017, 4, 27), d(2017, 4, 27)},
		{Koningsdag, 2018, d(2018, 4, 27), d(2018, 4, 27)},
		{Koningsdag, 2019, d(2019, 4, 27), d(2019, 4, 27)},
		{Koningsdag, 2020, d(2020, 4, 27), d(2020, 4, 27)},
		{Koningsdag, 2021, d(2021, 4, 27), d(2021, 4, 27)},
		{Koningsdag, 2022, d(2022, 4, 27), d(2022, 4, 27)},
		{Koningsdag, 2025, d(2025, 4, 27), d(2025, 4, 26)},

		{BevrijdingsDag, 2015, d(2015, 5, 5), d(2015, 5, 5)},
		{BevrijdingsDag, 2016, d(2016, 5, 5), d(2016, 5, 5)},
		{BevrijdingsDag, 2017, d(2017, 5, 5), d(2017, 5, 5)},
		{BevrijdingsDag, 2018, d(2018, 5, 5), d(2018, 5, 5)},
		{BevrijdingsDag, 2019, d(2019, 5, 5), d(2019, 5, 5)},
		{BevrijdingsDag, 2020, d(2020, 5, 5), d(2020, 5, 5)},
		{BevrijdingsDag, 2021, d(2021, 5, 5), d(2021, 5, 5)},
		{BevrijdingsDag, 2022, d(2022, 5, 5), d(2022, 5, 5)},

		{Hemelvaart, 2015, d(2015, 5, 14), d(2015, 5, 14)},
		{Hemelvaart, 2016, d(2016, 5, 5), d(2016, 5, 5)},
		{Hemelvaart, 2017, d(2017, 5, 25), d(2017, 5, 25)},
		{Hemelvaart, 2018, d(2018, 5, 10), d(2018, 5, 10)},
		{Hemelvaart, 2019, d(2019, 5, 30), d(2019, 5, 30)},
		{Hemelvaart, 2020, d(2020, 5, 21), d(2020, 5, 21)},
		{Hemelvaart, 2021, d(2021, 5, 13), d(2021, 5, 13)},
		{Hemelvaart, 2022, d(2022, 5, 26), d(2022, 5, 26)},

		{EerstePinksterDag, 2015, d(2015, 5, 24), d(2015, 5, 24)},
		{EerstePinksterDag, 2016, d(2016, 5, 15), d(2016, 5, 15)},
		{EerstePinksterDag, 2017, d(2017, 6, 4), d(2017, 6, 4)},
		{EerstePinksterDag, 2018, d(2018, 5, 20), d(2018, 5, 20)},
		{EerstePinksterDag, 2019, d(2019, 6, 9), d(2019, 6, 9)},
		{EerstePinksterDag, 2020, d(2020, 5, 31), d(2020, 5, 31)},
		{EerstePinksterDag, 2021, d(2021, 5, 23), d(2021, 5, 23)},
		{EerstePinksterDag, 2022, d(2022, 6, 5), d(2022, 6, 5)},
		{EerstePinksterDag, 2023, d(2023, 5, 28), d(2023, 5, 28)},

		{TweedePinksterDag, 2015, d(2015, 5, 25), d(2015, 5, 25)},
		{TweedePinksterDag, 2016, d(2016, 5, 16), d(2016, 5, 16)},
		{TweedePinksterDag, 2017, d(2017, 6, 5), d(2017, 6, 5)},
		{TweedePinksterDag, 2018, d(2018, 5, 21), d(2018, 5, 21)},
		{TweedePinksterDag, 2019, d(2019, 6, 10), d(2019, 6, 10)},
		{TweedePinksterDag, 2020, d(2020, 6, 1), d(2020, 6, 1)},
		{TweedePinksterDag, 2021, d(2021, 5, 24), d(2021, 5, 24)},
		{TweedePinksterDag, 2022, d(2022, 6, 6), d(2022, 6, 6)},

		{EersteKerstdag, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{EersteKerstdag, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{EersteKerstdag, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{EersteKerstdag, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{EersteKerstdag, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{EersteKerstdag, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{EersteKerstdag, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{EersteKerstdag, 2022, d(2022, 12, 25), d(2022, 12, 25)},

		{TweedeKerstdag, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{TweedeKerstdag, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{TweedeKerstdag, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{TweedeKerstdag, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{TweedeKerstdag, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{TweedeKerstdag, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{TweedeKerstdag, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{TweedeKerstdag, 2022, d(2022, 12, 26), d(2022, 12, 26)},
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
