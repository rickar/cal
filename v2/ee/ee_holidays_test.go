// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package ee

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
		{Uusaasta, 2024, d(2024, 1, 1), d(2024, 1, 1)},
		{Uusaasta, 2025, d(2025, 1, 1), d(2025, 1, 1)},
		{Uusaasta, 2026, d(2026, 1, 1), d(2026, 1, 1)},

		{Iseseisvuspaev, 2024, d(2024, 2, 24), d(2024, 2, 24)},
		{Iseseisvuspaev, 2025, d(2025, 2, 24), d(2025, 2, 24)},
		{Iseseisvuspaev, 2026, d(2026, 2, 24), d(2026, 2, 24)},

		{SuurReede, 2024, d(2024, 3, 29), d(2024, 3, 29)},
		{SuurReede, 2025, d(2025, 4, 18), d(2025, 4, 18)},
		{SuurReede, 2026, d(2026, 4, 3), d(2026, 4, 3)},

		{Ulestousmispuhade, 2024, d(2024, 3, 31), d(2024, 3, 31)},
		{Ulestousmispuhade, 2025, d(2025, 4, 20), d(2025, 4, 20)},
		{Ulestousmispuhade, 2026, d(2026, 4, 5), d(2026, 4, 5)},

		{Kevadpuha, 2024, d(2024, 5, 1), d(2024, 5, 1)},
		{Kevadpuha, 2025, d(2025, 5, 1), d(2025, 5, 1)},
		{Kevadpuha, 2026, d(2026, 5, 1), d(2026, 5, 1)},

		{Nelipuha, 2024, d(2024, 5, 19), d(2024, 5, 19)},
		{Nelipuha, 2025, d(2025, 6, 8), d(2025, 6, 8)},
		{Nelipuha, 2026, d(2026, 5, 24), d(2026, 5, 24)},

		{Voidupuha, 2024, d(2024, 6, 23), d(2024, 6, 23)},
		{Voidupuha, 2025, d(2025, 6, 23), d(2025, 6, 23)},
		{Voidupuha, 2026, d(2026, 6, 23), d(2026, 6, 23)},

		{Jaanipaev, 2024, d(2024, 6, 24), d(2024, 6, 24)},
		{Jaanipaev, 2025, d(2025, 6, 24), d(2025, 6, 24)},
		{Jaanipaev, 2026, d(2026, 6, 24), d(2026, 6, 24)},

		{Taasiseseisvuspaev, 2024, d(2024, 8, 20), d(2024, 8, 20)},
		{Taasiseseisvuspaev, 2025, d(2025, 8, 20), d(2025, 8, 20)},
		{Taasiseseisvuspaev, 2026, d(2026, 8, 20), d(2026, 8, 20)},

		{Joululaupaev, 2024, d(2024, 12, 24), d(2024, 12, 24)},
		{Joululaupaev, 2025, d(2025, 12, 24), d(2025, 12, 24)},
		{Joululaupaev, 2026, d(2026, 12, 24), d(2026, 12, 24)},

		{EsimeneJoulupuha, 2024, d(2024, 12, 25), d(2024, 12, 25)},
		{EsimeneJoulupuha, 2025, d(2025, 12, 25), d(2025, 12, 25)},
		{EsimeneJoulupuha, 2026, d(2026, 12, 25), d(2026, 12, 25)},

		{TeineJoulupuha, 2024, d(2024, 12, 26), d(2024, 12, 26)},
		{TeineJoulupuha, 2025, d(2025, 12, 26), d(2025, 12, 26)},
		{TeineJoulupuha, 2026, d(2026, 12, 26), d(2026, 12, 26)},
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
