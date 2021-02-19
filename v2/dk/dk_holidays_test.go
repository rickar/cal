// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package dk

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
		{Nytaarsdag, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{Nytaarsdag, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{Nytaarsdag, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{Nytaarsdag, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{Nytaarsdag, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{Nytaarsdag, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{Nytaarsdag, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{Nytaarsdag, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{Skaertorsdag, 2015, d(2015, 4, 2), d(2015, 4, 2)},
		{Skaertorsdag, 2016, d(2016, 3, 24), d(2016, 3, 24)},
		{Skaertorsdag, 2017, d(2017, 4, 13), d(2017, 4, 13)},
		{Skaertorsdag, 2018, d(2018, 3, 29), d(2018, 3, 29)},
		{Skaertorsdag, 2019, d(2019, 4, 18), d(2019, 4, 18)},
		{Skaertorsdag, 2020, d(2020, 4, 9), d(2020, 4, 9)},
		{Skaertorsdag, 2021, d(2021, 4, 1), d(2021, 4, 1)},
		{Skaertorsdag, 2022, d(2022, 4, 14), d(2022, 4, 14)},

		{Langfredag, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{Langfredag, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{Langfredag, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{Langfredag, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{Langfredag, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{Langfredag, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{Langfredag, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{Langfredag, 2022, d(2022, 4, 15), d(2022, 4, 15)},

		{AndenPaaskedag, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{AndenPaaskedag, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{AndenPaaskedag, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{AndenPaaskedag, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{AndenPaaskedag, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{AndenPaaskedag, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{AndenPaaskedag, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{AndenPaaskedag, 2022, d(2022, 4, 18), d(2022, 4, 18)},

		{StoreBededag, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{StoreBededag, 2016, d(2016, 4, 22), d(2016, 4, 22)},
		{StoreBededag, 2017, d(2017, 5, 12), d(2017, 5, 12)},
		{StoreBededag, 2018, d(2018, 4, 27), d(2018, 4, 27)},
		{StoreBededag, 2019, d(2019, 5, 17), d(2019, 5, 17)},
		{StoreBededag, 2020, d(2020, 5, 8), d(2020, 5, 8)},
		{StoreBededag, 2021, d(2021, 4, 30), d(2021, 4, 30)},
		{StoreBededag, 2022, d(2022, 5, 13), d(2022, 5, 13)},

		{KristiHimmelfartsdag, 2015, d(2015, 5, 14), d(2015, 5, 14)},
		{KristiHimmelfartsdag, 2016, d(2016, 5, 5), d(2016, 5, 5)},
		{KristiHimmelfartsdag, 2017, d(2017, 5, 25), d(2017, 5, 25)},
		{KristiHimmelfartsdag, 2018, d(2018, 5, 10), d(2018, 5, 10)},
		{KristiHimmelfartsdag, 2019, d(2019, 5, 30), d(2019, 5, 30)},
		{KristiHimmelfartsdag, 2020, d(2020, 5, 21), d(2020, 5, 21)},
		{KristiHimmelfartsdag, 2021, d(2021, 5, 13), d(2021, 5, 13)},
		{KristiHimmelfartsdag, 2022, d(2022, 5, 26), d(2022, 5, 26)},

		{AndenPinsedag, 2015, d(2015, 5, 25), d(2015, 5, 25)},
		{AndenPinsedag, 2016, d(2016, 5, 16), d(2016, 5, 16)},
		{AndenPinsedag, 2017, d(2017, 6, 5), d(2017, 6, 5)},
		{AndenPinsedag, 2018, d(2018, 5, 21), d(2018, 5, 21)},
		{AndenPinsedag, 2019, d(2019, 6, 10), d(2019, 6, 10)},
		{AndenPinsedag, 2020, d(2020, 6, 1), d(2020, 6, 1)},
		{AndenPinsedag, 2021, d(2021, 5, 24), d(2021, 5, 24)},
		{AndenPinsedag, 2022, d(2022, 6, 6), d(2022, 6, 6)},

		{Grundlovsdag, 2015, d(2015, 6, 5), d(2015, 6, 5)},
		{Grundlovsdag, 2016, d(2016, 6, 5), d(2016, 6, 5)},
		{Grundlovsdag, 2017, d(2017, 6, 5), d(2017, 6, 5)},
		{Grundlovsdag, 2018, d(2018, 6, 5), d(2018, 6, 5)},
		{Grundlovsdag, 2019, d(2019, 6, 5), d(2019, 6, 5)},
		{Grundlovsdag, 2020, d(2020, 6, 5), d(2020, 6, 5)},
		{Grundlovsdag, 2021, d(2021, 6, 5), d(2021, 6, 5)},
		{Grundlovsdag, 2022, d(2022, 6, 5), d(2022, 6, 5)},

		{Juledag, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Juledag, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Juledag, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Juledag, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Juledag, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Juledag, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Juledag, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Juledag, 2022, d(2022, 12, 25), d(2022, 12, 25)},

		{AndenJuledag, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{AndenJuledag, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{AndenJuledag, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{AndenJuledag, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{AndenJuledag, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{AndenJuledag, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{AndenJuledag, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{AndenJuledag, 2022, d(2022, 12, 26), d(2022, 12, 26)},
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
