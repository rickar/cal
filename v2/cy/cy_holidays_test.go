// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package cy

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
		{Protochronia, 2024, d(2024, 1, 1), d(2024, 1, 1)},
		{Protochronia, 2025, d(2025, 1, 1), d(2025, 1, 1)},
		{Protochronia, 2026, d(2026, 1, 1), d(2026, 1, 1)},

		{Theofania, 2024, d(2024, 1, 6), d(2024, 1, 6)},
		{Theofania, 2025, d(2025, 1, 6), d(2025, 1, 6)},
		{Theofania, 2026, d(2026, 1, 6), d(2026, 1, 6)},

		{KatharaDeftera, 2024, d(2024, 3, 18), d(2024, 3, 18)},
		{KatharaDeftera, 2025, d(2025, 3, 3), d(2025, 3, 3)},
		{KatharaDeftera, 2026, d(2026, 2, 23), d(2026, 2, 23)},

		{EllinikiEpanastasi, 2024, d(2024, 3, 25), d(2024, 3, 25)},
		{EllinikiEpanastasi, 2025, d(2025, 3, 25), d(2025, 3, 25)},
		{EllinikiEpanastasi, 2026, d(2026, 3, 25), d(2026, 3, 25)},

		{EthnikiEpetios, 2024, d(2024, 4, 1), d(2024, 4, 1)},
		{EthnikiEpetios, 2025, d(2025, 4, 1), d(2025, 4, 1)},
		{EthnikiEpetios, 2026, d(2026, 4, 1), d(2026, 4, 1)},

		{ErgatikoiProtomagia, 2024, d(2024, 5, 1), d(2024, 5, 1)},
		{ErgatikoiProtomagia, 2025, d(2025, 5, 1), d(2025, 5, 1)},
		{ErgatikoiProtomagia, 2026, d(2026, 5, 1), d(2026, 5, 1)},

		{MegaliParaskevi, 2024, d(2024, 5, 3), d(2024, 5, 3)},
		{MegaliParaskevi, 2025, d(2025, 4, 18), d(2025, 4, 18)},
		{MegaliParaskevi, 2026, d(2026, 4, 10), d(2026, 4, 10)},

		{DeuteraTouPascha, 2024, d(2024, 5, 6), d(2024, 5, 6)},
		{DeuteraTouPascha, 2025, d(2025, 4, 21), d(2025, 4, 21)},
		{DeuteraTouPascha, 2026, d(2026, 4, 13), d(2026, 4, 13)},

		{AgiouPnevmatos, 2024, d(2024, 6, 24), d(2024, 6, 24)},
		{AgiouPnevmatos, 2025, d(2025, 6, 9), d(2025, 6, 9)},
		{AgiouPnevmatos, 2026, d(2026, 6, 1), d(2026, 6, 1)},

		{KoimisiTisTheotokou, 2024, d(2024, 8, 15), d(2024, 8, 15)},
		{KoimisiTisTheotokou, 2025, d(2025, 8, 15), d(2025, 8, 15)},
		{KoimisiTisTheotokou, 2026, d(2026, 8, 15), d(2026, 8, 15)},

		{Anexartisia, 2024, d(2024, 10, 1), d(2024, 10, 1)},
		{Anexartisia, 2025, d(2025, 10, 1), d(2025, 10, 1)},
		{Anexartisia, 2026, d(2026, 10, 1), d(2026, 10, 1)},

		{EpeteiosTouOchi, 2024, d(2024, 10, 28), d(2024, 10, 28)},
		{EpeteiosTouOchi, 2025, d(2025, 10, 28), d(2025, 10, 28)},
		{EpeteiosTouOchi, 2026, d(2026, 10, 28), d(2026, 10, 28)},

		{Christougenna, 2024, d(2024, 12, 25), d(2024, 12, 25)},
		{Christougenna, 2025, d(2025, 12, 25), d(2025, 12, 25)},
		{Christougenna, 2026, d(2026, 12, 25), d(2026, 12, 25)},

		{DeyteriMeraTonChristougennon, 2024, d(2024, 12, 26), d(2024, 12, 26)},
		{DeyteriMeraTonChristougennon, 2025, d(2025, 12, 26), d(2025, 12, 26)},
		{DeyteriMeraTonChristougennon, 2026, d(2026, 12, 26), d(2026, 12, 26)},
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
