// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package si

import (
	"testing"
	"time"

	"github.com/rickar/cal/v2"
)

func d(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, cal.DefaultLoc)
}

func TestHolidays(t *testing.T) {
	t.Parallel()

	tests := []struct {
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{NovoLeto, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{NovoLeto, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{NovoLeto, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{NovoLeto, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{NovoLeto, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{NovoLeto, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{NovoLeto, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{NovoLeto, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{NovoLeto2, 2015, d(2015, 1, 2), d(2015, 1, 2)},
		{NovoLeto2, 2016, d(2016, 1, 2), d(2016, 1, 2)},
		{NovoLeto2, 2017, d(2017, 1, 2), d(2017, 1, 2)},
		{NovoLeto2, 2018, d(2018, 1, 2), d(2018, 1, 2)},
		{NovoLeto2, 2019, d(2019, 1, 2), d(2019, 1, 2)},
		{NovoLeto2, 2020, d(2020, 1, 2), d(2020, 1, 2)},
		{NovoLeto2, 2021, d(2021, 1, 2), d(2021, 1, 2)},
		{NovoLeto2, 2022, d(2022, 1, 2), d(2022, 1, 2)},

		{PresernovDan, 2015, d(2015, 2, 8), d(2015, 2, 8)},
		{PresernovDan, 2016, d(2016, 2, 8), d(2016, 2, 8)},
		{PresernovDan, 2017, d(2017, 2, 8), d(2017, 2, 8)},
		{PresernovDan, 2018, d(2018, 2, 8), d(2018, 2, 8)},
		{PresernovDan, 2019, d(2019, 2, 8), d(2019, 2, 8)},
		{PresernovDan, 2020, d(2020, 2, 8), d(2020, 2, 8)},
		{PresernovDan, 2021, d(2021, 2, 8), d(2021, 2, 8)},
		{PresernovDan, 2022, d(2022, 2, 8), d(2022, 2, 8)},

		{DanUporaProtiOkupatorju, 2015, d(2015, 4, 27), d(2015, 4, 27)},
		{DanUporaProtiOkupatorju, 2016, d(2016, 4, 27), d(2016, 4, 27)},
		{DanUporaProtiOkupatorju, 2017, d(2017, 4, 27), d(2017, 4, 27)},
		{DanUporaProtiOkupatorju, 2018, d(2018, 4, 27), d(2018, 4, 27)},
		{DanUporaProtiOkupatorju, 2019, d(2019, 4, 27), d(2019, 4, 27)},
		{DanUporaProtiOkupatorju, 2020, d(2020, 4, 27), d(2020, 4, 27)},
		{DanUporaProtiOkupatorju, 2021, d(2021, 4, 27), d(2021, 4, 27)},
		{DanUporaProtiOkupatorju, 2022, d(2022, 4, 27), d(2022, 4, 27)},

		{VelikaNoc, 2015, d(2015, 4, 5), d(2015, 4, 5)},
		{VelikaNoc, 2016, d(2016, 3, 27), d(2016, 3, 27)},
		{VelikaNoc, 2017, d(2017, 4, 16), d(2017, 4, 16)},
		{VelikaNoc, 2018, d(2018, 4, 1), d(2018, 4, 1)},
		{VelikaNoc, 2019, d(2019, 4, 21), d(2019, 4, 21)},
		{VelikaNoc, 2020, d(2020, 4, 12), d(2020, 4, 12)},
		{VelikaNoc, 2021, d(2021, 4, 4), d(2021, 4, 4)},
		{VelikaNoc, 2022, d(2022, 4, 17), d(2022, 4, 17)},

		{VelikonocniPonedeljek, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{VelikonocniPonedeljek, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{VelikonocniPonedeljek, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{VelikonocniPonedeljek, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{VelikonocniPonedeljek, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{VelikonocniPonedeljek, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{VelikonocniPonedeljek, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{VelikonocniPonedeljek, 2022, d(2022, 4, 18), d(2022, 4, 18)},

		{BinkostnaNedelja, 2015, d(2015, 5, 24), d(2015, 5, 24)},
		{BinkostnaNedelja, 2016, d(2016, 5, 15), d(2016, 5, 15)},
		{BinkostnaNedelja, 2017, d(2017, 6, 4), d(2017, 6, 4)},
		{BinkostnaNedelja, 2018, d(2018, 5, 20), d(2018, 5, 20)},
		{BinkostnaNedelja, 2019, d(2019, 6, 9), d(2019, 6, 9)},
		{BinkostnaNedelja, 2020, d(2020, 5, 31), d(2020, 5, 31)},
		{BinkostnaNedelja, 2021, d(2021, 5, 23), d(2021, 5, 23)},
		{BinkostnaNedelja, 2022, d(2022, 6, 5), d(2022, 6, 5)},

		{PrviMaj, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{PrviMaj, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{PrviMaj, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{PrviMaj, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{PrviMaj, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{PrviMaj, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{PrviMaj, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{PrviMaj, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{DrugiMaj, 2015, d(2015, 5, 2), d(2015, 5, 2)},
		{DrugiMaj, 2016, d(2016, 5, 2), d(2016, 5, 2)},
		{DrugiMaj, 2017, d(2017, 5, 2), d(2017, 5, 2)},
		{DrugiMaj, 2018, d(2018, 5, 2), d(2018, 5, 2)},
		{DrugiMaj, 2019, d(2019, 5, 2), d(2019, 5, 2)},
		{DrugiMaj, 2020, d(2020, 5, 2), d(2020, 5, 2)},
		{DrugiMaj, 2021, d(2021, 5, 2), d(2021, 5, 2)},
		{DrugiMaj, 2022, d(2022, 5, 2), d(2022, 5, 2)},

		{DanDrzavnosti, 2015, d(2015, 6, 25), d(2015, 6, 25)},
		{DanDrzavnosti, 2016, d(2016, 6, 25), d(2016, 6, 25)},
		{DanDrzavnosti, 2017, d(2017, 6, 25), d(2017, 6, 25)},
		{DanDrzavnosti, 2018, d(2018, 6, 25), d(2018, 6, 25)},
		{DanDrzavnosti, 2019, d(2019, 6, 25), d(2019, 6, 25)},
		{DanDrzavnosti, 2020, d(2020, 6, 25), d(2020, 6, 25)},
		{DanDrzavnosti, 2021, d(2021, 6, 25), d(2021, 6, 25)},
		{DanDrzavnosti, 2022, d(2022, 6, 25), d(2022, 6, 25)},

		{MarijinoVnebovzetje, 2015, d(2015, 8, 15), d(2015, 8, 15)},
		{MarijinoVnebovzetje, 2016, d(2016, 8, 15), d(2016, 8, 15)},
		{MarijinoVnebovzetje, 2017, d(2017, 8, 15), d(2017, 8, 15)},
		{MarijinoVnebovzetje, 2018, d(2018, 8, 15), d(2018, 8, 15)},
		{MarijinoVnebovzetje, 2019, d(2019, 8, 15), d(2019, 8, 15)},
		{MarijinoVnebovzetje, 2020, d(2020, 8, 15), d(2020, 8, 15)},
		{MarijinoVnebovzetje, 2021, d(2021, 8, 15), d(2021, 8, 15)},
		{MarijinoVnebovzetje, 2022, d(2022, 8, 15), d(2022, 8, 15)},

		{DanReformacije, 2015, d(2015, 10, 31), d(2015, 10, 31)},
		{DanReformacije, 2016, d(2016, 10, 31), d(2016, 10, 31)},
		{DanReformacije, 2017, d(2017, 10, 31), d(2017, 10, 31)},
		{DanReformacije, 2018, d(2018, 10, 31), d(2018, 10, 31)},
		{DanReformacije, 2019, d(2019, 10, 31), d(2019, 10, 31)},
		{DanReformacije, 2020, d(2020, 10, 31), d(2020, 10, 31)},
		{DanReformacije, 2021, d(2021, 10, 31), d(2021, 10, 31)},
		{DanReformacije, 2022, d(2022, 10, 31), d(2022, 10, 31)},

		{DanSpominaNaMrtve, 2015, d(2015, 11, 1), d(2015, 11, 1)},
		{DanSpominaNaMrtve, 2016, d(2016, 11, 1), d(2016, 11, 1)},
		{DanSpominaNaMrtve, 2017, d(2017, 11, 1), d(2017, 11, 1)},
		{DanSpominaNaMrtve, 2018, d(2018, 11, 1), d(2018, 11, 1)},
		{DanSpominaNaMrtve, 2019, d(2019, 11, 1), d(2019, 11, 1)},
		{DanSpominaNaMrtve, 2020, d(2020, 11, 1), d(2020, 11, 1)},
		{DanSpominaNaMrtve, 2021, d(2021, 11, 1), d(2021, 11, 1)},
		{DanSpominaNaMrtve, 2022, d(2022, 11, 1), d(2022, 11, 1)},

		{Bozic, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Bozic, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Bozic, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Bozic, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Bozic, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Bozic, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Bozic, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Bozic, 2022, d(2022, 12, 25), d(2022, 12, 25)},

		{DanSamostojnostiInEnotnosti, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{DanSamostojnostiInEnotnosti, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{DanSamostojnostiInEnotnosti, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{DanSamostojnostiInEnotnosti, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{DanSamostojnostiInEnotnosti, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{DanSamostojnostiInEnotnosti, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{DanSamostojnostiInEnotnosti, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{DanSamostojnostiInEnotnosti, 2022, d(2022, 12, 26), d(2022, 12, 26)},
	}

	for _, test := range tests {
		test := test
		gotAct, gotObs := test.h.Calc(test.y)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}
