// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package it

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
		{Capodanno, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{Capodanno, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{Capodanno, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{Capodanno, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{Capodanno, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{Capodanno, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{Capodanno, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{Capodanno, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{Epifania, 2015, d(2015, 1, 6), d(2015, 1, 6)},
		{Epifania, 2016, d(2016, 1, 6), d(2016, 1, 6)},
		{Epifania, 2017, d(2017, 1, 6), d(2017, 1, 6)},
		{Epifania, 2018, d(2018, 1, 6), d(2018, 1, 6)},
		{Epifania, 2019, d(2019, 1, 6), d(2019, 1, 6)},
		{Epifania, 2020, d(2020, 1, 6), d(2020, 1, 6)},
		{Epifania, 2021, d(2021, 1, 6), d(2021, 1, 6)},
		{Epifania, 2022, d(2022, 1, 6), d(2022, 1, 6)},

		{Pasquetta, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{Pasquetta, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{Pasquetta, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{Pasquetta, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{Pasquetta, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{Pasquetta, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{Pasquetta, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{Pasquetta, 2022, d(2022, 4, 18), d(2022, 4, 18)},

		{FestaDellaLiberazione, 2015, d(2015, 4, 25), d(2015, 4, 25)},
		{FestaDellaLiberazione, 2016, d(2016, 4, 25), d(2016, 4, 25)},
		{FestaDellaLiberazione, 2017, d(2017, 4, 25), d(2017, 4, 25)},
		{FestaDellaLiberazione, 2018, d(2018, 4, 25), d(2018, 4, 25)},
		{FestaDellaLiberazione, 2019, d(2019, 4, 25), d(2019, 4, 25)},
		{FestaDellaLiberazione, 2020, d(2020, 4, 25), d(2020, 4, 25)},
		{FestaDellaLiberazione, 2021, d(2021, 4, 25), d(2021, 4, 25)},
		{FestaDellaLiberazione, 2022, d(2022, 4, 25), d(2022, 4, 25)},

		{FestaDelLavoro, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{FestaDelLavoro, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{FestaDelLavoro, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{FestaDelLavoro, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{FestaDelLavoro, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{FestaDelLavoro, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{FestaDelLavoro, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{FestaDelLavoro, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{FestaDellaRepubblica, 2015, d(2015, 6, 2), d(2015, 6, 2)},
		{FestaDellaRepubblica, 2016, d(2016, 6, 2), d(2016, 6, 2)},
		{FestaDellaRepubblica, 2017, d(2017, 6, 2), d(2017, 6, 2)},
		{FestaDellaRepubblica, 2018, d(2018, 6, 2), d(2018, 6, 2)},
		{FestaDellaRepubblica, 2019, d(2019, 6, 2), d(2019, 6, 2)},
		{FestaDellaRepubblica, 2020, d(2020, 6, 2), d(2020, 6, 2)},
		{FestaDellaRepubblica, 2021, d(2021, 6, 2), d(2021, 6, 2)},
		{FestaDellaRepubblica, 2022, d(2022, 6, 2), d(2022, 6, 2)},

		{Assunzione, 2015, d(2015, 8, 15), d(2015, 8, 15)},
		{Assunzione, 2016, d(2016, 8, 15), d(2016, 8, 15)},
		{Assunzione, 2017, d(2017, 8, 15), d(2017, 8, 15)},
		{Assunzione, 2018, d(2018, 8, 15), d(2018, 8, 15)},
		{Assunzione, 2019, d(2019, 8, 15), d(2019, 8, 15)},
		{Assunzione, 2020, d(2020, 8, 15), d(2020, 8, 15)},
		{Assunzione, 2021, d(2021, 8, 15), d(2021, 8, 15)},
		{Assunzione, 2022, d(2022, 8, 15), d(2022, 8, 15)},

		{TuttiISanti, 2015, d(2015, 11, 1), d(2015, 11, 1)},
		{TuttiISanti, 2016, d(2016, 11, 1), d(2016, 11, 1)},
		{TuttiISanti, 2017, d(2017, 11, 1), d(2017, 11, 1)},
		{TuttiISanti, 2018, d(2018, 11, 1), d(2018, 11, 1)},
		{TuttiISanti, 2019, d(2019, 11, 1), d(2019, 11, 1)},
		{TuttiISanti, 2020, d(2020, 11, 1), d(2020, 11, 1)},
		{TuttiISanti, 2021, d(2021, 11, 1), d(2021, 11, 1)},
		{TuttiISanti, 2022, d(2022, 11, 1), d(2022, 11, 1)},

		{Immacolata, 2015, d(2015, 12, 8), d(2015, 12, 8)},
		{Immacolata, 2016, d(2016, 12, 8), d(2016, 12, 8)},
		{Immacolata, 2017, d(2017, 12, 8), d(2017, 12, 8)},
		{Immacolata, 2018, d(2018, 12, 8), d(2018, 12, 8)},
		{Immacolata, 2019, d(2019, 12, 8), d(2019, 12, 8)},
		{Immacolata, 2020, d(2020, 12, 8), d(2020, 12, 8)},
		{Immacolata, 2021, d(2021, 12, 8), d(2021, 12, 8)},
		{Immacolata, 2022, d(2022, 12, 8), d(2022, 12, 8)},

		{Natale, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Natale, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Natale, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Natale, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Natale, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Natale, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Natale, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Natale, 2022, d(2022, 12, 25), d(2022, 12, 25)},

		{SantoStefano, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{SantoStefano, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{SantoStefano, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{SantoStefano, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{SantoStefano, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{SantoStefano, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{SantoStefano, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{SantoStefano, 2022, d(2022, 12, 26), d(2022, 12, 26)},
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
