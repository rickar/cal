package fi

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
		{Uudenvuodenpaiva, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{Uudenvuodenpaiva, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{Uudenvuodenpaiva, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{Uudenvuodenpaiva, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{Uudenvuodenpaiva, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{Uudenvuodenpaiva, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{Uudenvuodenpaiva, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{Uudenvuodenpaiva, 2022, d(2022, 1, 1), d(2022, 1, 1)},
		{Uudenvuodenpaiva, 2023, d(2023, 1, 1), d(2023, 1, 1)},

		{Loppiainen, 2015, d(2015, 1, 6), d(2015, 1, 6)},
		{Loppiainen, 2016, d(2016, 1, 6), d(2016, 1, 6)},
		{Loppiainen, 2017, d(2017, 1, 6), d(2017, 1, 6)},
		{Loppiainen, 2018, d(2018, 1, 6), d(2018, 1, 6)},
		{Loppiainen, 2019, d(2019, 1, 6), d(2019, 1, 6)},
		{Loppiainen, 2020, d(2020, 1, 6), d(2020, 1, 6)},
		{Loppiainen, 2021, d(2021, 1, 6), d(2021, 1, 6)},
		{Loppiainen, 2022, d(2022, 1, 6), d(2022, 1, 6)},
		{Loppiainen, 2023, d(2023, 1, 6), d(2023, 1, 6)},

		{Pitkaperjantai, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{Pitkaperjantai, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{Pitkaperjantai, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{Pitkaperjantai, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{Pitkaperjantai, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{Pitkaperjantai, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{Pitkaperjantai, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{Pitkaperjantai, 2022, d(2022, 4, 15), d(2022, 4, 15)},
		{Pitkaperjantai, 2023, d(2023, 4, 7), d(2023, 4, 7)},

		{Paasiaispaiva, 2015, d(2015, 4, 5), d(2015, 4, 5)},
		{Paasiaispaiva, 2016, d(2016, 3, 27), d(2016, 3, 27)},
		{Paasiaispaiva, 2017, d(2017, 4, 16), d(2017, 4, 16)},
		{Paasiaispaiva, 2018, d(2018, 4, 1), d(2018, 4, 1)},
		{Paasiaispaiva, 2019, d(2019, 4, 21), d(2019, 4, 21)},
		{Paasiaispaiva, 2020, d(2020, 4, 12), d(2020, 4, 12)},
		{Paasiaispaiva, 2021, d(2021, 4, 4), d(2021, 4, 4)},
		{Paasiaispaiva, 2022, d(2022, 4, 17), d(2022, 4, 17)},
		{Paasiaispaiva, 2023, d(2023, 4, 9), d(2023, 4, 9)},

		{ToinenPaasiaispaiva, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{ToinenPaasiaispaiva, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{ToinenPaasiaispaiva, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{ToinenPaasiaispaiva, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{ToinenPaasiaispaiva, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{ToinenPaasiaispaiva, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{ToinenPaasiaispaiva, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{ToinenPaasiaispaiva, 2022, d(2022, 4, 18), d(2022, 4, 18)},
		{ToinenPaasiaispaiva, 2023, d(2023, 4, 10), d(2023, 4, 10)},

		{Vappu, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{Vappu, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{Vappu, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{Vappu, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{Vappu, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{Vappu, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{Vappu, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{Vappu, 2022, d(2022, 5, 1), d(2022, 5, 1)},
		{Vappu, 2023, d(2023, 5, 1), d(2023, 5, 1)},

		{Helatorstai, 2015, d(2015, 5, 14), d(2015, 5, 14)},
		{Helatorstai, 2016, d(2016, 5, 5), d(2016, 5, 5)},
		{Helatorstai, 2017, d(2017, 5, 25), d(2017, 5, 25)},
		{Helatorstai, 2018, d(2018, 5, 10), d(2018, 5, 10)},
		{Helatorstai, 2019, d(2019, 5, 30), d(2019, 5, 30)},
		{Helatorstai, 2020, d(2020, 5, 21), d(2020, 5, 21)},
		{Helatorstai, 2021, d(2021, 5, 13), d(2021, 5, 13)},
		{Helatorstai, 2022, d(2022, 5, 26), d(2022, 5, 26)},
		{Helatorstai, 2023, d(2023, 5, 18), d(2023, 5, 18)},

		{Helluntaipaiva, 2015, d(2015, 5, 24), d(2015, 5, 24)},
		{Helluntaipaiva, 2016, d(2016, 5, 15), d(2016, 5, 15)},
		{Helluntaipaiva, 2017, d(2017, 6, 4), d(2017, 6, 4)},
		{Helluntaipaiva, 2018, d(2018, 5, 20), d(2018, 5, 20)},
		{Helluntaipaiva, 2019, d(2019, 6, 9), d(2019, 6, 9)},
		{Helluntaipaiva, 2020, d(2020, 5, 31), d(2020, 5, 31)},
		{Helluntaipaiva, 2021, d(2021, 5, 23), d(2021, 5, 23)},
		{Helluntaipaiva, 2022, d(2022, 6, 5), d(2022, 6, 5)},
		{Helluntaipaiva, 2023, d(2023, 5, 28), d(2023, 5, 28)},

		{Juhannusaatto, 2015, d(2015, 6, 19), d(2015, 6, 19)},
		{Juhannusaatto, 2016, d(2016, 6, 24), d(2016, 6, 24)},
		{Juhannusaatto, 2017, d(2017, 6, 23), d(2017, 6, 23)},
		{Juhannusaatto, 2018, d(2018, 6, 22), d(2018, 6, 22)},
		{Juhannusaatto, 2019, d(2019, 6, 21), d(2019, 6, 21)},
		{Juhannusaatto, 2020, d(2020, 6, 19), d(2020, 6, 19)},
		{Juhannusaatto, 2021, d(2021, 6, 25), d(2021, 6, 25)},
		{Juhannusaatto, 2022, d(2022, 6, 24), d(2022, 6, 24)},
		{Juhannusaatto, 2023, d(2023, 6, 23), d(2023, 6, 23)},

		{Juhannuspaiva, 2015, d(2015, 6, 20), d(2015, 6, 20)},
		{Juhannuspaiva, 2016, d(2016, 6, 25), d(2016, 6, 25)},
		{Juhannuspaiva, 2017, d(2017, 6, 24), d(2017, 6, 24)},
		{Juhannuspaiva, 2018, d(2018, 6, 23), d(2018, 6, 23)},
		{Juhannuspaiva, 2019, d(2019, 6, 22), d(2019, 6, 22)},
		{Juhannuspaiva, 2020, d(2020, 6, 20), d(2020, 6, 20)},
		{Juhannuspaiva, 2021, d(2021, 6, 26), d(2021, 6, 26)},
		{Juhannuspaiva, 2022, d(2022, 6, 25), d(2022, 6, 25)},
		{Juhannuspaiva, 2023, d(2023, 6, 24), d(2023, 6, 24)},

		{Pyhainpaiva, 2015, d(2015, 10, 31), d(2015, 10, 31)},
		{Pyhainpaiva, 2016, d(2016, 11, 5), d(2016, 11, 5)},
		{Pyhainpaiva, 2017, d(2017, 11, 4), d(2017, 11, 4)},
		{Pyhainpaiva, 2018, d(2018, 11, 3), d(2018, 11, 3)},
		{Pyhainpaiva, 2019, d(2019, 11, 2), d(2019, 11, 2)},
		{Pyhainpaiva, 2020, d(2020, 10, 31), d(2020, 10, 31)},
		{Pyhainpaiva, 2021, d(2021, 11, 6), d(2021, 11, 6)},
		{Pyhainpaiva, 2022, d(2022, 11, 5), d(2022, 11, 5)},
		{Pyhainpaiva, 2023, d(2023, 11, 4), d(2023, 11, 4)},

		{Itsenaisyyspaiva, 2015, d(2015, 12, 6), d(2015, 12, 6)},
		{Itsenaisyyspaiva, 2016, d(2016, 12, 6), d(2016, 12, 6)},
		{Itsenaisyyspaiva, 2017, d(2017, 12, 6), d(2017, 12, 6)},
		{Itsenaisyyspaiva, 2018, d(2018, 12, 6), d(2018, 12, 6)},
		{Itsenaisyyspaiva, 2019, d(2019, 12, 6), d(2019, 12, 6)},
		{Itsenaisyyspaiva, 2020, d(2020, 12, 6), d(2020, 12, 6)},
		{Itsenaisyyspaiva, 2021, d(2021, 12, 6), d(2021, 12, 6)},
		{Itsenaisyyspaiva, 2022, d(2022, 12, 6), d(2022, 12, 6)},
		{Itsenaisyyspaiva, 2023, d(2023, 12, 6), d(2023, 12, 6)},

		{Jouluaatto, 2015, d(2015, 12, 24), d(2015, 12, 24)},
		{Jouluaatto, 2016, d(2016, 12, 24), d(2016, 12, 24)},
		{Jouluaatto, 2017, d(2017, 12, 24), d(2017, 12, 24)},
		{Jouluaatto, 2018, d(2018, 12, 24), d(2018, 12, 24)},
		{Jouluaatto, 2019, d(2019, 12, 24), d(2019, 12, 24)},
		{Jouluaatto, 2020, d(2020, 12, 24), d(2020, 12, 24)},
		{Jouluaatto, 2021, d(2021, 12, 24), d(2021, 12, 24)},
		{Jouluaatto, 2022, d(2022, 12, 24), d(2022, 12, 24)},
		{Jouluaatto, 2023, d(2023, 12, 24), d(2023, 12, 24)},

		{Joulupaiva, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Joulupaiva, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Joulupaiva, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Joulupaiva, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Joulupaiva, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Joulupaiva, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Joulupaiva, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Joulupaiva, 2022, d(2022, 12, 25), d(2022, 12, 25)},
		{Joulupaiva, 2023, d(2023, 12, 25), d(2023, 12, 25)},

		{Tapaninpaiva, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{Tapaninpaiva, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{Tapaninpaiva, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{Tapaninpaiva, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{Tapaninpaiva, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{Tapaninpaiva, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{Tapaninpaiva, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{Tapaninpaiva, 2022, d(2022, 12, 26), d(2022, 12, 26)},
		{Tapaninpaiva, 2023, d(2023, 12, 26), d(2023, 12, 26)},
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
