// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package hu

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
		{Ujev, 2024, d(2024, 1, 1), d(2024, 1, 1)},
		{Ujev, 2025, d(2025, 1, 1), d(2025, 1, 1)},
		{Ujev, 2026, d(2026, 1, 1), d(2026, 1, 1)},

		{NemzetiUnnepMarcius, 2024, d(2024, 3, 15), d(2024, 3, 15)},
		{NemzetiUnnepMarcius, 2025, d(2025, 3, 15), d(2025, 3, 15)},
		{NemzetiUnnepMarcius, 2026, d(2026, 3, 15), d(2026, 3, 15)},

		{Nagypentek, 2024, d(2024, 3, 29), d(2024, 3, 29)},
		{Nagypentek, 2025, d(2025, 4, 18), d(2025, 4, 18)},
		{Nagypentek, 2026, d(2026, 4, 3), d(2026, 4, 3)},

		{HusvetHetfo, 2024, d(2024, 4, 1), d(2024, 4, 1)},
		{HusvetHetfo, 2025, d(2025, 4, 21), d(2025, 4, 21)},
		{HusvetHetfo, 2026, d(2026, 4, 6), d(2026, 4, 6)},

		{AmunkaUnnepe, 2024, d(2024, 5, 1), d(2024, 5, 1)},
		{AmunkaUnnepe, 2025, d(2025, 5, 1), d(2025, 5, 1)},
		{AmunkaUnnepe, 2026, d(2026, 5, 1), d(2026, 5, 1)},

		{PunkosdHetfo, 2024, d(2024, 5, 20), d(2024, 5, 20)},
		{PunkosdHetfo, 2025, d(2025, 6, 9), d(2025, 6, 9)},
		{PunkosdHetfo, 2026, d(2026, 5, 25), d(2026, 5, 25)},

		{SzentIstvanUnnepe, 2024, d(2024, 8, 20), d(2024, 8, 20)},
		{SzentIstvanUnnepe, 2025, d(2025, 8, 20), d(2025, 8, 20)},
		{SzentIstvanUnnepe, 2026, d(2026, 8, 20), d(2026, 8, 20)},

		{NemzetiUnnepOkt, 2024, d(2024, 10, 23), d(2024, 10, 23)},
		{NemzetiUnnepOkt, 2025, d(2025, 10, 23), d(2025, 10, 23)},
		{NemzetiUnnepOkt, 2026, d(2026, 10, 23), d(2026, 10, 23)},

		{Mindenszentek, 2024, d(2024, 11, 1), d(2024, 11, 1)},
		{Mindenszentek, 2025, d(2025, 11, 1), d(2025, 11, 1)},
		{Mindenszentek, 2026, d(2026, 11, 1), d(2026, 11, 1)},

		{Karacsony, 2024, d(2024, 12, 25), d(2024, 12, 25)},
		{Karacsony, 2025, d(2025, 12, 25), d(2025, 12, 25)},
		{Karacsony, 2026, d(2026, 12, 25), d(2026, 12, 25)},

		{KaracsonyMasnapja, 2024, d(2024, 12, 26), d(2024, 12, 26)},
		{KaracsonyMasnapja, 2025, d(2025, 12, 26), d(2025, 12, 26)},
		{KaracsonyMasnapja, 2026, d(2026, 12, 26), d(2026, 12, 26)},
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
