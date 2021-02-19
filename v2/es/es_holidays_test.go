// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package es

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
		{AñoNuevo, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{AñoNuevo, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{AñoNuevo, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{AñoNuevo, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{AñoNuevo, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{AñoNuevo, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{AñoNuevo, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{AñoNuevo, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{Reyes, 2015, d(2015, 1, 6), d(2015, 1, 6)},
		{Reyes, 2016, d(2016, 1, 6), d(2016, 1, 6)},
		{Reyes, 2017, d(2017, 1, 6), d(2017, 1, 6)},
		{Reyes, 2018, d(2018, 1, 6), d(2018, 1, 6)},
		{Reyes, 2019, d(2019, 1, 6), d(2019, 1, 6)},
		{Reyes, 2020, d(2020, 1, 6), d(2020, 1, 6)},
		{Reyes, 2021, d(2021, 1, 6), d(2021, 1, 6)},
		{Reyes, 2022, d(2022, 1, 6), d(2022, 1, 6)},

		{ViernesSanto, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{ViernesSanto, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{ViernesSanto, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{ViernesSanto, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{ViernesSanto, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{ViernesSanto, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{ViernesSanto, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{ViernesSanto, 2022, d(2022, 4, 15), d(2022, 4, 15)},

		{Trabajador, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{Trabajador, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{Trabajador, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{Trabajador, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{Trabajador, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{Trabajador, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{Trabajador, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{Trabajador, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{Asunción, 2015, d(2015, 8, 15), d(2015, 8, 15)},
		{Asunción, 2016, d(2016, 8, 15), d(2016, 8, 15)},
		{Asunción, 2017, d(2017, 8, 15), d(2017, 8, 15)},
		{Asunción, 2018, d(2018, 8, 15), d(2018, 8, 15)},
		{Asunción, 2019, d(2019, 8, 15), d(2019, 8, 15)},
		{Asunción, 2020, d(2020, 8, 15), d(2020, 8, 15)},
		{Asunción, 2021, d(2021, 8, 15), d(2021, 8, 15)},
		{Asunción, 2022, d(2022, 8, 15), d(2022, 8, 15)},

		{FiestaNacionalDeEspaña, 2015, d(2015, 10, 12), d(2015, 10, 12)},
		{FiestaNacionalDeEspaña, 2016, d(2016, 10, 12), d(2016, 10, 12)},
		{FiestaNacionalDeEspaña, 2017, d(2017, 10, 12), d(2017, 10, 12)},
		{FiestaNacionalDeEspaña, 2018, d(2018, 10, 12), d(2018, 10, 12)},
		{FiestaNacionalDeEspaña, 2019, d(2019, 10, 12), d(2019, 10, 12)},
		{FiestaNacionalDeEspaña, 2020, d(2020, 10, 12), d(2020, 10, 12)},
		{FiestaNacionalDeEspaña, 2021, d(2021, 10, 12), d(2021, 10, 12)},
		{FiestaNacionalDeEspaña, 2022, d(2022, 10, 12), d(2022, 10, 12)},

		{TodosLosSantos, 2015, d(2015, 11, 1), d(2015, 11, 1)},
		{TodosLosSantos, 2016, d(2016, 11, 1), d(2016, 11, 1)},
		{TodosLosSantos, 2017, d(2017, 11, 1), d(2017, 11, 1)},
		{TodosLosSantos, 2018, d(2018, 11, 1), d(2018, 11, 1)},
		{TodosLosSantos, 2019, d(2019, 11, 1), d(2019, 11, 1)},
		{TodosLosSantos, 2020, d(2020, 11, 1), d(2020, 11, 1)},
		{TodosLosSantos, 2021, d(2021, 11, 1), d(2021, 11, 1)},
		{TodosLosSantos, 2022, d(2022, 11, 1), d(2022, 11, 1)},

		{Constitucion, 2015, d(2015, 12, 6), d(2015, 12, 6)},
		{Constitucion, 2016, d(2016, 12, 6), d(2016, 12, 6)},
		{Constitucion, 2017, d(2017, 12, 6), d(2017, 12, 6)},
		{Constitucion, 2018, d(2018, 12, 6), d(2018, 12, 6)},
		{Constitucion, 2019, d(2019, 12, 6), d(2019, 12, 6)},
		{Constitucion, 2020, d(2020, 12, 6), d(2020, 12, 6)},
		{Constitucion, 2021, d(2021, 12, 6), d(2021, 12, 6)},
		{Constitucion, 2022, d(2022, 12, 6), d(2022, 12, 6)},

		{InmaculadaConcepcion, 2015, d(2015, 12, 8), d(2015, 12, 8)},
		{InmaculadaConcepcion, 2016, d(2016, 12, 8), d(2016, 12, 8)},
		{InmaculadaConcepcion, 2017, d(2017, 12, 8), d(2017, 12, 8)},
		{InmaculadaConcepcion, 2018, d(2018, 12, 8), d(2018, 12, 8)},
		{InmaculadaConcepcion, 2019, d(2019, 12, 8), d(2019, 12, 8)},
		{InmaculadaConcepcion, 2020, d(2020, 12, 8), d(2020, 12, 8)},
		{InmaculadaConcepcion, 2021, d(2021, 12, 8), d(2021, 12, 8)},
		{InmaculadaConcepcion, 2022, d(2022, 12, 8), d(2022, 12, 8)},

		{Navidad, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Navidad, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Navidad, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Navidad, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Navidad, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Navidad, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Navidad, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Navidad, 2022, d(2022, 12, 25), d(2022, 12, 25)},
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
