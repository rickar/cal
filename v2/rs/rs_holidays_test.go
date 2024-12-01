// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package rs

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
		{NovaGodina, 2024, d(2024, 1, 1), d(2024, 1, 1)},
		{NovaGodina, 2025, d(2025, 1, 1), d(2025, 1, 1)},
		{NovaGodina, 2026, d(2026, 1, 1), d(2026, 1, 1)},

		{DrugiDanNoveGodine, 2024, d(2024, 1, 2), d(2024, 1, 2)},
		{DrugiDanNoveGodine, 2025, d(2025, 1, 2), d(2025, 1, 2)},
		{DrugiDanNoveGodine, 2026, d(2026, 1, 2), d(2026, 1, 2)},

		{Bozic, 2024, d(2024, 1, 7), d(2024, 1, 7)},
		{Bozic, 2025, d(2025, 1, 7), d(2025, 1, 7)},
		{Bozic, 2026, d(2026, 1, 7), d(2026, 1, 7)},

		{DanDrzavnosti, 2024, d(2024, 2, 15), d(2024, 2, 15)},
		{DanDrzavnosti, 2025, d(2025, 2, 15), d(2025, 2, 15)},
		{DanDrzavnosti, 2026, d(2026, 2, 15), d(2026, 2, 15)},

		{DrugiDanDrzavnosti, 2024, d(2024, 2, 16), d(2024, 2, 16)},
		{DrugiDanDrzavnosti, 2025, d(2025, 2, 16), d(2025, 2, 16)},
		{DrugiDanDrzavnosti, 2026, d(2026, 2, 16), d(2026, 2, 16)},

		{PraznikRada, 2024, d(2024, 5, 1), d(2024, 5, 1)},
		{PraznikRada, 2025, d(2025, 5, 1), d(2025, 5, 1)},
		{PraznikRada, 2026, d(2026, 5, 1), d(2026, 5, 1)},

		{DrugiDanPraznikaRada, 2024, d(2024, 5, 2), d(2024, 5, 2)},
		{DrugiDanPraznikaRada, 2025, d(2025, 5, 2), d(2025, 5, 2)},
		{DrugiDanPraznikaRada, 2026, d(2026, 5, 2), d(2026, 5, 2)},

		{VelikiPetak, 2024, d(2024, 5, 3), d(2024, 5, 3)},
		{VelikiPetak, 2025, d(2025, 4, 18), d(2025, 4, 18)},
		{VelikiPetak, 2026, d(2026, 4, 10), d(2026, 4, 10)},

		{Vaskrs, 2024, d(2024, 5, 5), d(2024, 5, 5)},
		{Vaskrs, 2025, d(2025, 4, 20), d(2025, 4, 20)},
		{Vaskrs, 2026, d(2026, 4, 12), d(2026, 4, 12)},

		{VaskrsnjiPonedeljak, 2024, d(2024, 5, 6), d(2024, 5, 6)},
		{VaskrsnjiPonedeljak, 2025, d(2025, 4, 21), d(2025, 4, 21)},
		{VaskrsnjiPonedeljak, 2026, d(2026, 4, 13), d(2026, 4, 13)},

		{DanPrimirja, 2024, d(2024, 11, 11), d(2024, 11, 11)},
		{DanPrimirja, 2025, d(2025, 11, 11), d(2025, 11, 11)},
		{DanPrimirja, 2026, d(2026, 11, 11), d(2026, 11, 11)},
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
