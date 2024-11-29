// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package mt

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
		{LEwwelTasSena, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{LEwwelTasSena, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{LEwwelTasSena, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{NawfragjuSanPawl, 2015, d(2015, 2, 10), d(2015, 2, 10)},
		{NawfragjuSanPawl, 2020, d(2020, 2, 10), d(2020, 2, 10)},
		{NawfragjuSanPawl, 2022, d(2022, 2, 10), d(2022, 2, 10)},

		{SanGuzepp, 2015, d(2015, 3, 19), d(2015, 3, 19)},
		{SanGuzepp, 2020, d(2020, 3, 19), d(2020, 3, 19)},
		{SanGuzepp, 2022, d(2022, 3, 19), d(2022, 3, 19)},

		{IlGimghaLKbira, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{IlGimghaLKbira, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{IlGimghaLKbira, 2022, d(2022, 4, 15), d(2022, 4, 15)},

		{JumIlĦelsien, 2015, d(2015, 3, 31), d(2015, 3, 31)},
		{JumIlĦelsien, 2020, d(2020, 3, 31), d(2020, 3, 31)},
		{JumIlĦelsien, 2022, d(2022, 3, 31), d(2022, 3, 31)},

		{JumIlĦaddiem, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{JumIlĦaddiem, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{JumIlĦaddiem, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{SetteGiugno, 2015, d(2015, 6, 7), d(2015, 6, 7)},
		{SetteGiugno, 2020, d(2020, 6, 7), d(2020, 6, 7)},
		{SetteGiugno, 2022, d(2022, 6, 7), d(2022, 6, 7)},

		{LImnarja, 2015, d(2015, 6, 29), d(2015, 6, 29)},
		{LImnarja, 2020, d(2020, 6, 29), d(2020, 6, 29)},
		{LImnarja, 2022, d(2022, 6, 29), d(2022, 6, 29)},

		{SantaMarija, 2015, d(2015, 8, 15), d(2015, 8, 15)},
		{SantaMarija, 2020, d(2020, 8, 15), d(2020, 8, 15)},
		{SantaMarija, 2022, d(2022, 8, 15), d(2022, 8, 15)},

		{JumIlVitorja, 2015, d(2015, 9, 8), d(2015, 9, 8)},
		{JumIlVitorja, 2020, d(2020, 9, 8), d(2020, 9, 8)},
		{JumIlVitorja, 2022, d(2022, 9, 8), d(2022, 9, 8)},

		{JumLIndipendenza, 2015, d(2015, 9, 21), d(2015, 9, 21)},
		{JumLIndipendenza, 2020, d(2020, 9, 21), d(2020, 9, 21)},
		{JumLIndipendenza, 2022, d(2022, 9, 21), d(2022, 9, 21)},

		{IlKuncizzjoni, 2015, d(2015, 12, 8), d(2015, 12, 8)},
		{IlKuncizzjoni, 2020, d(2020, 12, 8), d(2020, 12, 8)},
		{IlKuncizzjoni, 2022, d(2022, 12, 8), d(2022, 12, 8)},

		{JumIrRepubblika, 2015, d(2015, 12, 13), d(2015, 12, 13)},
		{JumIrRepubblika, 2020, d(2020, 12, 13), d(2020, 12, 13)},
		{JumIrRepubblika, 2022, d(2022, 12, 13), d(2022, 12, 13)},

		{IlMilied, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{IlMilied, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{IlMilied, 2022, d(2022, 12, 25), d(2022, 12, 25)},
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
