// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package br

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

		{Tiradentes, 2015, d(2015, 4, 21), d(2015, 4, 21)},
		{Tiradentes, 2016, d(2016, 4, 21), d(2016, 4, 21)},
		{Tiradentes, 2017, d(2017, 4, 21), d(2017, 4, 21)},
		{Tiradentes, 2018, d(2018, 4, 21), d(2018, 4, 21)},
		{Tiradentes, 2019, d(2019, 4, 21), d(2019, 4, 21)},
		{Tiradentes, 2020, d(2020, 4, 21), d(2020, 4, 21)},
		{Tiradentes, 2021, d(2021, 4, 21), d(2021, 4, 21)},
		{Tiradentes, 2022, d(2022, 4, 21), d(2022, 4, 21)},

		{Trabalhador, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{Trabalhador, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{Trabalhador, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{Trabalhador, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{Trabalhador, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{Trabalhador, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{Trabalhador, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{Trabalhador, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{Independencia, 2015, d(2015, 9, 7), d(2015, 9, 7)},
		{Independencia, 2016, d(2016, 9, 7), d(2016, 9, 7)},
		{Independencia, 2017, d(2017, 9, 7), d(2017, 9, 7)},
		{Independencia, 2018, d(2018, 9, 7), d(2018, 9, 7)},
		{Independencia, 2019, d(2019, 9, 7), d(2019, 9, 7)},
		{Independencia, 2020, d(2020, 9, 7), d(2020, 9, 7)},
		{Independencia, 2021, d(2021, 9, 7), d(2021, 9, 7)},
		{Independencia, 2022, d(2022, 9, 7), d(2022, 9, 7)},

		{NossaSenhoraAparecida, 2015, d(2015, 10, 12), d(2015, 10, 12)},
		{NossaSenhoraAparecida, 2016, d(2016, 10, 12), d(2016, 10, 12)},
		{NossaSenhoraAparecida, 2017, d(2017, 10, 12), d(2017, 10, 12)},
		{NossaSenhoraAparecida, 2018, d(2018, 10, 12), d(2018, 10, 12)},
		{NossaSenhoraAparecida, 2019, d(2019, 10, 12), d(2019, 10, 12)},
		{NossaSenhoraAparecida, 2020, d(2020, 10, 12), d(2020, 10, 12)},
		{NossaSenhoraAparecida, 2021, d(2021, 10, 12), d(2021, 10, 12)},
		{NossaSenhoraAparecida, 2022, d(2022, 10, 12), d(2022, 10, 12)},

		{Finados, 2015, d(2015, 11, 2), d(2015, 11, 2)},
		{Finados, 2016, d(2016, 11, 2), d(2016, 11, 2)},
		{Finados, 2017, d(2017, 11, 2), d(2017, 11, 2)},
		{Finados, 2018, d(2018, 11, 2), d(2018, 11, 2)},
		{Finados, 2019, d(2019, 11, 2), d(2019, 11, 2)},
		{Finados, 2020, d(2020, 11, 2), d(2020, 11, 2)},
		{Finados, 2021, d(2021, 11, 2), d(2021, 11, 2)},
		{Finados, 2022, d(2022, 11, 2), d(2022, 11, 2)},

		{Republica, 2015, d(2015, 11, 15), d(2015, 11, 15)},
		{Republica, 2016, d(2016, 11, 15), d(2016, 11, 15)},
		{Republica, 2017, d(2017, 11, 15), d(2017, 11, 15)},
		{Republica, 2018, d(2018, 11, 15), d(2018, 11, 15)},
		{Republica, 2019, d(2019, 11, 15), d(2019, 11, 15)},
		{Republica, 2020, d(2020, 11, 15), d(2020, 11, 15)},
		{Republica, 2021, d(2021, 11, 15), d(2021, 11, 15)},
		{Republica, 2022, d(2022, 11, 15), d(2022, 11, 15)},

		{CorpusChristi, 2015, d(2015, 6, 4), d(2015, 6, 4)},
		{CorpusChristi, 2016, d(2016, 5, 26), d(2016, 5, 26)},
		{CorpusChristi, 2017, d(2017, 6, 15), d(2017, 6, 15)},
		{CorpusChristi, 2018, d(2018, 5, 31), d(2018, 5, 31)},
		{CorpusChristi, 2019, d(2019, 6, 20), d(2019, 6, 20)},
		{CorpusChristi, 2020, d(2020, 6, 11), d(2020, 6, 11)},
		{CorpusChristi, 2021, d(2021, 6, 3), d(2021, 6, 3)},
		{CorpusChristi, 2022, d(2022, 6, 16), d(2022, 6, 16)},

		{SextaFeiraSanta, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{SextaFeiraSanta, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{SextaFeiraSanta, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{SextaFeiraSanta, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{SextaFeiraSanta, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{SextaFeiraSanta, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{SextaFeiraSanta, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{SextaFeiraSanta, 2022, d(2022, 4, 15), d(2022, 4, 15)},

		{Carnaval, 2015, d(2015, 2, 17), d(2015, 2, 17)},
		{Carnaval, 2016, d(2016, 2, 9), d(2016, 2, 9)},
		{Carnaval, 2017, d(2017, 2, 28), d(2017, 2, 28)},
		{Carnaval, 2018, d(2018, 2, 13), d(2018, 2, 13)},
		{Carnaval, 2019, d(2019, 3, 5), d(2019, 3, 5)},
		{Carnaval, 2020, d(2020, 2, 25), d(2020, 2, 25)},
		{Carnaval, 2021, d(2021, 2, 16), d(2021, 2, 16)},
		{Carnaval, 2022, d(2022, 3, 1), d(2022, 3, 1)},
		{Carnaval, 2023, d(2023, 2, 21), d(2023, 2, 21)},

		{Natal, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Natal, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Natal, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Natal, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Natal, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Natal, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Natal, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Natal, 2022, d(2022, 12, 25), d(2022, 12, 25)},
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
