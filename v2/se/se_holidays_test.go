// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package se

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
		{Nyarsdagen, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{Nyarsdagen, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{Nyarsdagen, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{Nyarsdagen, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{Nyarsdagen, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{Nyarsdagen, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{Nyarsdagen, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{Nyarsdagen, 2022, d(2022, 1, 1), d(2022, 1, 1)},
		{Nyarsdagen, 2023, d(2023, 1, 1), d(2023, 1, 1)},

		{TrettondedagJul, 2015, d(2015, 1, 6), d(2015, 1, 6)},
		{TrettondedagJul, 2016, d(2016, 1, 6), d(2016, 1, 6)},
		{TrettondedagJul, 2017, d(2017, 1, 6), d(2017, 1, 6)},
		{TrettondedagJul, 2018, d(2018, 1, 6), d(2018, 1, 6)},
		{TrettondedagJul, 2019, d(2019, 1, 6), d(2019, 1, 6)},
		{TrettondedagJul, 2020, d(2020, 1, 6), d(2020, 1, 6)},
		{TrettondedagJul, 2021, d(2021, 1, 6), d(2021, 1, 6)},
		{TrettondedagJul, 2022, d(2022, 1, 6), d(2022, 1, 6)},
		{TrettondedagJul, 2023, d(2023, 1, 6), d(2023, 1, 6)},

		{Langfredagen, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{Langfredagen, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{Langfredagen, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{Langfredagen, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{Langfredagen, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{Langfredagen, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{Langfredagen, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{Langfredagen, 2022, d(2022, 4, 15), d(2022, 4, 15)},
		{Langfredagen, 2023, d(2023, 4, 7), d(2023, 4, 7)},

		{Paskdagen, 2015, d(2015, 4, 5), d(2015, 4, 5)},
		{Paskdagen, 2016, d(2016, 3, 27), d(2016, 3, 27)},
		{Paskdagen, 2017, d(2017, 4, 16), d(2017, 4, 16)},
		{Paskdagen, 2018, d(2018, 4, 1), d(2018, 4, 1)},
		{Paskdagen, 2019, d(2019, 4, 21), d(2019, 4, 21)},
		{Paskdagen, 2020, d(2020, 4, 12), d(2020, 4, 12)},
		{Paskdagen, 2021, d(2021, 4, 4), d(2021, 4, 4)},
		{Paskdagen, 2022, d(2022, 4, 17), d(2022, 4, 17)},
		{Paskdagen, 2023, d(2023, 4, 9), d(2023, 4, 9)},

		{AnnandagPask, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{AnnandagPask, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{AnnandagPask, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{AnnandagPask, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{AnnandagPask, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{AnnandagPask, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{AnnandagPask, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{AnnandagPask, 2022, d(2022, 4, 18), d(2022, 4, 18)},
		{AnnandagPask, 2023, d(2023, 4, 10), d(2023, 4, 10)},

		{ForstaMaj, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{ForstaMaj, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{ForstaMaj, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{ForstaMaj, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{ForstaMaj, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{ForstaMaj, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{ForstaMaj, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{ForstaMaj, 2022, d(2022, 5, 1), d(2022, 5, 1)},
		{ForstaMaj, 2023, d(2023, 5, 1), d(2023, 5, 1)},

		{KristiHimmelsfardsdag, 2015, d(2015, 5, 14), d(2015, 5, 14)},
		{KristiHimmelsfardsdag, 2016, d(2016, 5, 5), d(2016, 5, 5)},
		{KristiHimmelsfardsdag, 2017, d(2017, 5, 25), d(2017, 5, 25)},
		{KristiHimmelsfardsdag, 2018, d(2018, 5, 10), d(2018, 5, 10)},
		{KristiHimmelsfardsdag, 2019, d(2019, 5, 30), d(2019, 5, 30)},
		{KristiHimmelsfardsdag, 2020, d(2020, 5, 21), d(2020, 5, 21)},
		{KristiHimmelsfardsdag, 2021, d(2021, 5, 13), d(2021, 5, 13)},
		{KristiHimmelsfardsdag, 2022, d(2022, 5, 26), d(2022, 5, 26)},
		{KristiHimmelsfardsdag, 2023, d(2023, 5, 18), d(2023, 5, 18)},

		{Pingstdagen, 2015, d(2015, 5, 24), d(2015, 5, 24)},
		{Pingstdagen, 2016, d(2016, 5, 15), d(2016, 5, 15)},
		{Pingstdagen, 2017, d(2017, 6, 4), d(2017, 6, 4)},
		{Pingstdagen, 2018, d(2018, 5, 20), d(2018, 5, 20)},
		{Pingstdagen, 2019, d(2019, 6, 9), d(2019, 6, 9)},
		{Pingstdagen, 2020, d(2020, 5, 31), d(2020, 5, 31)},
		{Pingstdagen, 2021, d(2021, 5, 23), d(2021, 5, 23)},
		{Pingstdagen, 2022, d(2022, 6, 5), d(2022, 6, 5)},
		{Pingstdagen, 2023, d(2023, 5, 28), d(2023, 5, 28)},

		{Nationaldagen, 2015, d(2015, 6, 6), d(2015, 6, 6)},
		{Nationaldagen, 2016, d(2016, 6, 6), d(2016, 6, 6)},
		{Nationaldagen, 2017, d(2017, 6, 6), d(2017, 6, 6)},
		{Nationaldagen, 2018, d(2018, 6, 6), d(2018, 6, 6)},
		{Nationaldagen, 2019, d(2019, 6, 6), d(2019, 6, 6)},
		{Nationaldagen, 2020, d(2020, 6, 6), d(2020, 6, 6)},
		{Nationaldagen, 2021, d(2021, 6, 6), d(2021, 6, 6)},
		{Nationaldagen, 2022, d(2022, 6, 6), d(2022, 6, 6)},
		{Nationaldagen, 2023, d(2023, 6, 6), d(2023, 6, 6)},

		{Midsommarafton, 2015, d(2015, 6, 19), d(2015, 6, 19)},
		{Midsommarafton, 2016, d(2016, 6, 24), d(2016, 6, 24)},
		{Midsommarafton, 2017, d(2017, 6, 23), d(2017, 6, 23)},
		{Midsommarafton, 2018, d(2018, 6, 22), d(2018, 6, 22)},
		{Midsommarafton, 2019, d(2019, 6, 21), d(2019, 6, 21)},
		{Midsommarafton, 2020, d(2020, 6, 19), d(2020, 6, 19)},
		{Midsommarafton, 2021, d(2021, 6, 25), d(2021, 6, 25)},
		{Midsommarafton, 2022, d(2022, 6, 24), d(2022, 6, 24)},
		{Midsommarafton, 2023, d(2023, 6, 23), d(2023, 6, 23)},

		{Midsommardagen, 2015, d(2015, 6, 20), d(2015, 6, 20)},
		{Midsommardagen, 2016, d(2016, 6, 25), d(2016, 6, 25)},
		{Midsommardagen, 2017, d(2017, 6, 24), d(2017, 6, 24)},
		{Midsommardagen, 2018, d(2018, 6, 23), d(2018, 6, 23)},
		{Midsommardagen, 2019, d(2019, 6, 22), d(2019, 6, 22)},
		{Midsommardagen, 2020, d(2020, 6, 20), d(2020, 6, 20)},
		{Midsommardagen, 2021, d(2021, 6, 26), d(2021, 6, 26)},
		{Midsommardagen, 2022, d(2022, 6, 25), d(2022, 6, 25)},
		{Midsommardagen, 2023, d(2023, 6, 24), d(2023, 6, 24)},

		{AllaHelgonsDag, 2015, d(2015, 10, 31), d(2015, 10, 31)},
		{AllaHelgonsDag, 2016, d(2016, 11, 5), d(2016, 11, 5)},
		{AllaHelgonsDag, 2017, d(2017, 11, 4), d(2017, 11, 4)},
		{AllaHelgonsDag, 2018, d(2018, 11, 3), d(2018, 11, 3)},
		{AllaHelgonsDag, 2019, d(2019, 11, 2), d(2019, 11, 2)},
		{AllaHelgonsDag, 2020, d(2020, 10, 31), d(2020, 10, 31)},
		{AllaHelgonsDag, 2021, d(2021, 11, 6), d(2021, 11, 6)},
		{AllaHelgonsDag, 2022, d(2022, 11, 5), d(2022, 11, 5)},
		{AllaHelgonsDag, 2023, d(2023, 11, 4), d(2023, 11, 4)},

		{Julafton, 2015, d(2015, 12, 24), d(2015, 12, 24)},
		{Julafton, 2016, d(2016, 12, 24), d(2016, 12, 24)},
		{Julafton, 2017, d(2017, 12, 24), d(2017, 12, 24)},
		{Julafton, 2018, d(2018, 12, 24), d(2018, 12, 24)},
		{Julafton, 2019, d(2019, 12, 24), d(2019, 12, 24)},
		{Julafton, 2020, d(2020, 12, 24), d(2020, 12, 24)},
		{Julafton, 2021, d(2021, 12, 24), d(2021, 12, 24)},
		{Julafton, 2022, d(2022, 12, 24), d(2022, 12, 24)},
		{Julafton, 2023, d(2023, 12, 24), d(2023, 12, 24)},

		{Juldagen, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Juldagen, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Juldagen, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Juldagen, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Juldagen, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Juldagen, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Juldagen, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Juldagen, 2022, d(2022, 12, 25), d(2022, 12, 25)},
		{Juldagen, 2023, d(2023, 12, 25), d(2023, 12, 25)},

		{AnnandagJul, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{AnnandagJul, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{AnnandagJul, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{AnnandagJul, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{AnnandagJul, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{AnnandagJul, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{AnnandagJul, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{AnnandagJul, 2022, d(2022, 12, 26), d(2022, 12, 26)},
		{AnnandagJul, 2023, d(2023, 12, 26), d(2023, 12, 26)},

		{Nyarsafton, 2015, d(2015, 12, 31), d(2015, 12, 31)},
		{Nyarsafton, 2016, d(2016, 12, 31), d(2016, 12, 31)},
		{Nyarsafton, 2017, d(2017, 12, 31), d(2017, 12, 31)},
		{Nyarsafton, 2018, d(2018, 12, 31), d(2018, 12, 31)},
		{Nyarsafton, 2019, d(2019, 12, 31), d(2019, 12, 31)},
		{Nyarsafton, 2020, d(2020, 12, 31), d(2020, 12, 31)},
		{Nyarsafton, 2021, d(2021, 12, 31), d(2021, 12, 31)},
		{Nyarsafton, 2022, d(2022, 12, 31), d(2022, 12, 31)},
		{Nyarsafton, 2023, d(2023, 12, 31), d(2023, 12, 31)},
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
