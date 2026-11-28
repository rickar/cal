// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package id

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
		{NewYear, 2025, d(2025, 1, 1), d(2025, 1, 1)},
		{NewYear, 2026, d(2026, 1, 1), d(2026, 1, 1)},
		{NewYear, 2027, d(2027, 1, 1), d(2027, 1, 1)},

		{GoodFridayOld, 2022, d(2022, 4, 15), d(2022, 4, 15)},
		{GoodFridayOld, 2023, d(2023, 4, 7), d(2023, 4, 7)},
		{GoodFridayOld, 2024, time.Time{}, time.Time{}},

		{GoodFriday, 2023, time.Time{}, time.Time{}},
		{GoodFriday, 2024, d(2024, 3, 29), d(2024, 3, 29)},
		{GoodFriday, 2025, d(2025, 4, 18), d(2025, 4, 18)},
		{GoodFriday, 2026, d(2026, 4, 3), d(2026, 4, 3)},
		{GoodFriday, 2027, d(2027, 3, 26), d(2027, 3, 26)},

		{Easter, 2023, time.Time{}, time.Time{}},
		{Easter, 2024, d(2024, 3, 31), d(2024, 3, 31)},
		{Easter, 2025, d(2025, 4, 20), d(2025, 4, 20)},
		{Easter, 2026, d(2026, 4, 5), d(2026, 4, 5)},
		{Easter, 2027, d(2027, 3, 28), d(2027, 3, 28)},

		{LabourDay, 2013, time.Time{}, time.Time{}},
		{LabourDay, 2014, d(2014, 5, 1), d(2014, 5, 1)},
		{LabourDay, 2026, d(2026, 5, 1), d(2026, 5, 1)},
		{LabourDay, 2027, d(2027, 5, 1), d(2027, 5, 1)},

		{AscensionDayOld, 2022, d(2022, 5, 26), d(2022, 5, 26)},
		{AscensionDayOld, 2023, d(2023, 5, 18), d(2023, 5, 18)},
		{AscensionDayOld, 2024, time.Time{}, time.Time{}},

		{AscensionDay, 2023, time.Time{}, time.Time{}},
		{AscensionDay, 2024, d(2024, 5, 9), d(2024, 5, 9)},
		{AscensionDay, 2025, d(2025, 5, 29), d(2025, 5, 29)},
		{AscensionDay, 2026, d(2026, 5, 14), d(2026, 5, 14)},
		{AscensionDay, 2027, d(2027, 5, 6), d(2027, 5, 6)},

		{PancasilaDay, 2016, time.Time{}, time.Time{}},
		{PancasilaDay, 2017, d(2017, 6, 1), d(2017, 6, 1)},
		{PancasilaDay, 2026, d(2026, 6, 1), d(2026, 6, 1)},
		{PancasilaDay, 2027, d(2027, 6, 1), d(2027, 6, 1)},

		{IndependenceDay, 2025, d(2025, 8, 17), d(2025, 8, 17)},
		{IndependenceDay, 2026, d(2026, 8, 17), d(2026, 8, 17)},
		{IndependenceDay, 2027, d(2027, 8, 17), d(2027, 8, 17)},

		{ChristmasDay, 2025, d(2025, 12, 25), d(2025, 12, 25)},
		{ChristmasDay, 2026, d(2026, 12, 25), d(2026, 12, 25)},
		{ChristmasDay, 2027, d(2027, 12, 25), d(2027, 12, 25)},
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

func TestCalendar(t *testing.T) {
	c := cal.NewBusinessCalendar()
	c.AddHoliday(Holidays...)

	// The computable national holidays in the official 2026 calendar (SKB 3 Menteri)
	want := []time.Time{
		d(2026, 1, 1),   // Tahun Baru Masehi
		d(2026, 4, 3),   // Wafat Yesus Kristus
		d(2026, 4, 5),   // Kebangkitan Yesus Kristus (Paskah)
		d(2026, 5, 1),   // Hari Buruh Internasional
		d(2026, 5, 14),  // Kenaikan Yesus Kristus
		d(2026, 6, 1),   // Hari Lahir Pancasila
		d(2026, 8, 17),  // Hari Kemerdekaan Republik Indonesia
		d(2026, 12, 25), // Hari Raya Natal
	}
	for _, day := range want {
		if actual, _, _ := c.IsHoliday(day); !actual {
			t.Errorf("%s should be a holiday", day.Format("2006-01-02"))
		}
	}

	count := 0
	for day := d(2026, 1, 1); day.Year() == 2026; day = day.AddDate(0, 0, 1) {
		if actual, _, _ := c.IsHoliday(day); actual {
			count++
		}
	}
	if count != len(want) {
		t.Errorf("got %d holidays in 2026, want %d", count, len(want))
	}

	for _, h := range Holidays {
		if h.Type != cal.ObservancePublic {
			t.Errorf("%s: got type %d, want %d", h.Name, h.Type, cal.ObservancePublic)
		}
	}
}
