// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package fr

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
		{NouvelAn, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{NouvelAn, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{NouvelAn, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{NouvelAn, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{NouvelAn, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{NouvelAn, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{NouvelAn, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{NouvelAn, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{LundiDePâques, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{LundiDePâques, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{LundiDePâques, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{LundiDePâques, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{LundiDePâques, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{LundiDePâques, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{LundiDePâques, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{LundiDePâques, 2022, d(2022, 4, 18), d(2022, 4, 18)},

		{FêteDuTravail, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{FêteDuTravail, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{FêteDuTravail, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{FêteDuTravail, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{FêteDuTravail, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{FêteDuTravail, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{FêteDuTravail, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{FêteDuTravail, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{FêteDeLaVictoire, 2015, d(2015, 5, 8), d(2015, 5, 8)},
		{FêteDeLaVictoire, 2016, d(2016, 5, 8), d(2016, 5, 8)},
		{FêteDeLaVictoire, 2017, d(2017, 5, 8), d(2017, 5, 8)},
		{FêteDeLaVictoire, 2018, d(2018, 5, 8), d(2018, 5, 8)},
		{FêteDeLaVictoire, 2019, d(2019, 5, 8), d(2019, 5, 8)},
		{FêteDeLaVictoire, 2020, d(2020, 5, 8), d(2020, 5, 8)},
		{FêteDeLaVictoire, 2021, d(2021, 5, 8), d(2021, 5, 8)},
		{FêteDeLaVictoire, 2022, d(2022, 5, 8), d(2022, 5, 8)},

		{Ascension, 2015, d(2015, 5, 14), d(2015, 5, 14)},
		{Ascension, 2016, d(2016, 5, 5), d(2016, 5, 5)},
		{Ascension, 2017, d(2017, 5, 25), d(2017, 5, 25)},
		{Ascension, 2018, d(2018, 5, 10), d(2018, 5, 10)},
		{Ascension, 2019, d(2019, 5, 30), d(2019, 5, 30)},
		{Ascension, 2020, d(2020, 5, 21), d(2020, 5, 21)},
		{Ascension, 2021, d(2021, 5, 13), d(2021, 5, 13)},
		{Ascension, 2022, d(2022, 5, 26), d(2022, 5, 26)},

		{LundiDePentecôte, 2015, d(2015, 5, 25), d(2015, 5, 25)},
		{LundiDePentecôte, 2016, d(2016, 5, 16), d(2016, 5, 16)},
		{LundiDePentecôte, 2017, d(2017, 6, 5), d(2017, 6, 5)},
		{LundiDePentecôte, 2018, d(2018, 5, 21), d(2018, 5, 21)},
		{LundiDePentecôte, 2019, d(2019, 6, 10), d(2019, 6, 10)},
		{LundiDePentecôte, 2020, d(2020, 6, 1), d(2020, 6, 1)},
		{LundiDePentecôte, 2021, d(2021, 5, 24), d(2021, 5, 24)},
		{LundiDePentecôte, 2022, d(2022, 6, 6), d(2022, 6, 6)},

		{FêteNationale, 2015, d(2015, 7, 14), d(2015, 7, 14)},
		{FêteNationale, 2016, d(2016, 7, 14), d(2016, 7, 14)},
		{FêteNationale, 2017, d(2017, 7, 14), d(2017, 7, 14)},
		{FêteNationale, 2018, d(2018, 7, 14), d(2018, 7, 14)},
		{FêteNationale, 2019, d(2019, 7, 14), d(2019, 7, 14)},
		{FêteNationale, 2020, d(2020, 7, 14), d(2020, 7, 14)},
		{FêteNationale, 2021, d(2021, 7, 14), d(2021, 7, 14)},
		{FêteNationale, 2022, d(2022, 7, 14), d(2022, 7, 14)},

		{Assomption, 2015, d(2015, 8, 15), d(2015, 8, 15)},
		{Assomption, 2016, d(2016, 8, 15), d(2016, 8, 15)},
		{Assomption, 2017, d(2017, 8, 15), d(2017, 8, 15)},
		{Assomption, 2018, d(2018, 8, 15), d(2018, 8, 15)},
		{Assomption, 2019, d(2019, 8, 15), d(2019, 8, 15)},
		{Assomption, 2020, d(2020, 8, 15), d(2020, 8, 15)},
		{Assomption, 2021, d(2021, 8, 15), d(2021, 8, 15)},
		{Assomption, 2022, d(2022, 8, 15), d(2022, 8, 15)},

		{Toussaint, 2015, d(2015, 11, 1), d(2015, 11, 1)},
		{Toussaint, 2016, d(2016, 11, 1), d(2016, 11, 1)},
		{Toussaint, 2017, d(2017, 11, 1), d(2017, 11, 1)},
		{Toussaint, 2018, d(2018, 11, 1), d(2018, 11, 1)},
		{Toussaint, 2019, d(2019, 11, 1), d(2019, 11, 1)},
		{Toussaint, 2020, d(2020, 11, 1), d(2020, 11, 1)},
		{Toussaint, 2021, d(2021, 11, 1), d(2021, 11, 1)},
		{Toussaint, 2022, d(2022, 11, 1), d(2022, 11, 1)},

		{Armistice1918, 2015, d(2015, 11, 11), d(2015, 11, 11)},
		{Armistice1918, 2016, d(2016, 11, 11), d(2016, 11, 11)},
		{Armistice1918, 2017, d(2017, 11, 11), d(2017, 11, 11)},
		{Armistice1918, 2018, d(2018, 11, 11), d(2018, 11, 11)},
		{Armistice1918, 2019, d(2019, 11, 11), d(2019, 11, 11)},
		{Armistice1918, 2020, d(2020, 11, 11), d(2020, 11, 11)},
		{Armistice1918, 2021, d(2021, 11, 11), d(2021, 11, 11)},
		{Armistice1918, 2022, d(2022, 11, 11), d(2022, 11, 11)},

		{Noël, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Noël, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Noël, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Noël, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Noël, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Noël, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Noël, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Noël, 2022, d(2022, 12, 25), d(2022, 12, 25)},
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
