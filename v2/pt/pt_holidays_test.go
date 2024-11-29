// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package pt

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
		{AnoNovo, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{AnoNovo, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{AnoNovo, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{AnoNovo, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{AnoNovo, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{AnoNovo, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{AnoNovo, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{AnoNovo, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{SextaFeiraSanta, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{SextaFeiraSanta, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{SextaFeiraSanta, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{SextaFeiraSanta, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{SextaFeiraSanta, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{SextaFeiraSanta, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{SextaFeiraSanta, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{SextaFeiraSanta, 2022, d(2022, 4, 15), d(2022, 4, 15)},

		{DomingoPascoa, 2015, d(2015, 4, 5), d(2015, 4, 5)},
		{DomingoPascoa, 2016, d(2016, 3, 27), d(2016, 3, 27)},
		{DomingoPascoa, 2017, d(2017, 4, 16), d(2017, 4, 16)},
		{DomingoPascoa, 2018, d(2018, 4, 1), d(2018, 4, 1)},
		{DomingoPascoa, 2019, d(2019, 4, 21), d(2019, 4, 21)},
		{DomingoPascoa, 2020, d(2020, 4, 12), d(2020, 4, 12)},
		{DomingoPascoa, 2021, d(2021, 4, 4), d(2021, 4, 4)},
		{DomingoPascoa, 2022, d(2022, 4, 17), d(2022, 4, 17)},

		{DiaDaLiberdade, 2015, d(2015, 4, 25), d(2015, 4, 25)},
		{DiaDaLiberdade, 2016, d(2016, 4, 25), d(2016, 4, 25)},
		{DiaDaLiberdade, 2017, d(2017, 4, 25), d(2017, 4, 25)},
		{DiaDaLiberdade, 2018, d(2018, 4, 25), d(2018, 4, 25)},
		{DiaDaLiberdade, 2019, d(2019, 4, 25), d(2019, 4, 25)},

		{CorpoDeDeus, 2015, d(2015, 6, 4), d(2015, 6, 4)},
		{CorpoDeDeus, 2016, d(2016, 5, 26), d(2016, 5, 26)},
		{CorpoDeDeus, 2017, d(2017, 6, 15), d(2017, 6, 15)},
		{CorpoDeDeus, 2018, d(2018, 5, 31), d(2018, 5, 31)},
		{CorpoDeDeus, 2019, d(2019, 6, 20), d(2019, 6, 20)},

		{Natal, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Natal, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Natal, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Natal, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Natal, 2019, d(2019, 12, 25), d(2019, 12, 25)},
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
