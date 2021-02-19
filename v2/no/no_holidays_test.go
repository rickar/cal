// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package no

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
		{FoersteNyttaarsdag, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{FoersteNyttaarsdag, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{FoersteNyttaarsdag, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{FoersteNyttaarsdag, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{FoersteNyttaarsdag, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{FoersteNyttaarsdag, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{FoersteNyttaarsdag, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{FoersteNyttaarsdag, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{Skjaertorsdag, 2015, d(2015, 4, 2), d(2015, 4, 2)},
		{Skjaertorsdag, 2016, d(2016, 3, 24), d(2016, 3, 24)},
		{Skjaertorsdag, 2017, d(2017, 4, 13), d(2017, 4, 13)},
		{Skjaertorsdag, 2018, d(2018, 3, 29), d(2018, 3, 29)},
		{Skjaertorsdag, 2019, d(2019, 4, 18), d(2019, 4, 18)},
		{Skjaertorsdag, 2020, d(2020, 4, 9), d(2020, 4, 9)},
		{Skjaertorsdag, 2021, d(2021, 4, 1), d(2021, 4, 1)},
		{Skjaertorsdag, 2022, d(2022, 4, 14), d(2022, 4, 14)},

		{Langfredag, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{Langfredag, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{Langfredag, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{Langfredag, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{Langfredag, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{Langfredag, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{Langfredag, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{Langfredag, 2022, d(2022, 4, 15), d(2022, 4, 15)},

		{AndrePaaskedag, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{AndrePaaskedag, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{AndrePaaskedag, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{AndrePaaskedag, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{AndrePaaskedag, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{AndrePaaskedag, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{AndrePaaskedag, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{AndrePaaskedag, 2022, d(2022, 4, 18), d(2022, 4, 18)},

		{Arbeiderenesdag, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{Arbeiderenesdag, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{Arbeiderenesdag, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{Arbeiderenesdag, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{Arbeiderenesdag, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{Arbeiderenesdag, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{Arbeiderenesdag, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{Arbeiderenesdag, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{Grunnlovsdag, 2015, d(2015, 5, 17), d(2015, 5, 17)},
		{Grunnlovsdag, 2016, d(2016, 5, 17), d(2016, 5, 17)},
		{Grunnlovsdag, 2017, d(2017, 5, 17), d(2017, 5, 17)},
		{Grunnlovsdag, 2018, d(2018, 5, 17), d(2018, 5, 17)},
		{Grunnlovsdag, 2019, d(2019, 5, 17), d(2019, 5, 17)},
		{Grunnlovsdag, 2020, d(2020, 5, 17), d(2020, 5, 17)},
		{Grunnlovsdag, 2021, d(2021, 5, 17), d(2021, 5, 17)},
		{Grunnlovsdag, 2022, d(2022, 5, 17), d(2022, 5, 17)},

		{Kristihimmelfartsdag, 2015, d(2015, 5, 14), d(2015, 5, 14)},
		{Kristihimmelfartsdag, 2016, d(2016, 5, 5), d(2016, 5, 5)},
		{Kristihimmelfartsdag, 2017, d(2017, 5, 25), d(2017, 5, 25)},
		{Kristihimmelfartsdag, 2018, d(2018, 5, 10), d(2018, 5, 10)},
		{Kristihimmelfartsdag, 2019, d(2019, 5, 30), d(2019, 5, 30)},
		{Kristihimmelfartsdag, 2020, d(2020, 5, 21), d(2020, 5, 21)},
		{Kristihimmelfartsdag, 2021, d(2021, 5, 13), d(2021, 5, 13)},
		{Kristihimmelfartsdag, 2022, d(2022, 5, 26), d(2022, 5, 26)},

		{AndrePinsedag, 2015, d(2015, 5, 25), d(2015, 5, 25)},
		{AndrePinsedag, 2016, d(2016, 5, 16), d(2016, 5, 16)},
		{AndrePinsedag, 2017, d(2017, 6, 5), d(2017, 6, 5)},
		{AndrePinsedag, 2018, d(2018, 5, 21), d(2018, 5, 21)},
		{AndrePinsedag, 2019, d(2019, 6, 10), d(2019, 6, 10)},
		{AndrePinsedag, 2020, d(2020, 6, 1), d(2020, 6, 1)},
		{AndrePinsedag, 2021, d(2021, 5, 24), d(2021, 5, 24)},
		{AndrePinsedag, 2022, d(2022, 6, 6), d(2022, 6, 6)},

		{FoersteJuledag, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{FoersteJuledag, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{FoersteJuledag, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{FoersteJuledag, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{FoersteJuledag, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{FoersteJuledag, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{FoersteJuledag, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{FoersteJuledag, 2022, d(2022, 12, 25), d(2022, 12, 25)},

		{AndreJuledag, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{AndreJuledag, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{AndreJuledag, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{AndreJuledag, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{AndreJuledag, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{AndreJuledag, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{AndreJuledag, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{AndreJuledag, 2022, d(2022, 12, 26), d(2022, 12, 26)},
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
