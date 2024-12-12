// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package ke

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
		{NewYear, 2022, d(2022, 1, 1), d(2022, 1, 1)},
		{NewYear, 2023, d(2023, 1, 1), d(2023, 1, 2)},
		{NewYear, 2024, d(2024, 1, 1), d(2024, 1, 1)},
		{NewYear, 2025, d(2025, 1, 1), d(2025, 1, 1)},

		{GoodFriday, 2022, d(2022, 4, 15), d(2022, 4, 15)},
		{GoodFriday, 2023, d(2023, 4, 7), d(2023, 4, 7)},
		{GoodFriday, 2024, d(2024, 3, 29), d(2024, 3, 29)},
		{GoodFriday, 2025, d(2025, 4, 18), d(2025, 4, 18)},

		{EasterMonday, 2022, d(2022, 4, 18), d(2022, 4, 18)},
		{EasterMonday, 2023, d(2023, 4, 10), d(2023, 4, 10)},
		{EasterMonday, 2024, d(2024, 4, 1), d(2024, 4, 1)},
		{EasterMonday, 2025, d(2025, 4, 21), d(2025, 4, 21)},

		{LabourDay, 2022, d(2022, 5, 1), d(2022, 5, 2)},
		{LabourDay, 2023, d(2023, 5, 1), d(2023, 5, 1)},
		{LabourDay, 2024, d(2024, 5, 1), d(2024, 5, 1)},
		{LabourDay, 2025, d(2025, 5, 1), d(2025, 5, 1)},

		{MadarakaDay, 2022, d(2022, 6, 1), d(2022, 6, 1)},
		{MadarakaDay, 2023, d(2023, 6, 1), d(2023, 6, 1)},
		{MadarakaDay, 2024, d(2024, 6, 1), d(2024, 6, 1)},
		{MadarakaDay, 2025, d(2025, 6, 1), d(2025, 6, 2)},

		{UtamaduniDay, 2022, d(2022, 10, 10), d(2022, 10, 10)},
		{UtamaduniDay, 2023, d(2023, 10, 10), d(2023, 10, 10)},

		{MazingiraDay, 2024, d(2024, 10, 10), d(2024, 10, 10)},
		{MazingiraDay, 2025, d(2025, 10, 10), d(2025, 10, 10)},

		{MashujaaDay, 2022, d(2022, 10, 20), d(2022, 10, 20)},
		{MashujaaDay, 2023, d(2023, 10, 20), d(2023, 10, 20)},
		{MashujaaDay, 2024, d(2024, 10, 20), d(2024, 10, 21)},
		{MashujaaDay, 2025, d(2025, 10, 20), d(2025, 10, 20)},

		{JamhuriDay, 2022, d(2022, 12, 12), d(2022, 12, 12)},
		{JamhuriDay, 2023, d(2023, 12, 12), d(2023, 12, 12)},
		{JamhuriDay, 2024, d(2024, 12, 12), d(2024, 12, 12)},
		{JamhuriDay, 2025, d(2025, 12, 12), d(2025, 12, 12)},

		{ChristmasDay, 2022, d(2022, 12, 25), d(2022, 12, 26)},
		{ChristmasDay, 2023, d(2023, 12, 25), d(2023, 12, 25)},
		{ChristmasDay, 2024, d(2024, 12, 25), d(2024, 12, 25)},
		{ChristmasDay, 2025, d(2025, 12, 25), d(2025, 12, 25)},

		{BoxingDay, 2022, d(2022, 12, 26), d(2022, 12, 26)},
		{BoxingDay, 2023, d(2023, 12, 26), d(2023, 12, 26)},
		{BoxingDay, 2024, d(2024, 12, 26), d(2024, 12, 26)},
		{BoxingDay, 2025, d(2025, 12, 26), d(2025, 12, 26)},
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
