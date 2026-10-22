// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package bj

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
		{NewYear, 2024, d(2024, 1, 1), d(2024, 1, 1)},
		{NewYear, 2025, d(2025, 1, 1), d(2025, 1, 1)},
		{NewYear, 2026, d(2026, 1, 1), d(2026, 1, 1)},

		{TraditionalReligionsDayEve, 2024, time.Time{}, time.Time{}},
		{TraditionalReligionsDayEve, 2025, d(2025, 1, 9), d(2025, 1, 9)},
		{TraditionalReligionsDayEve, 2026, d(2026, 1, 8), d(2026, 1, 8)},
		{TraditionalReligionsDayEve, 2027, d(2027, 1, 7), d(2027, 1, 7)},

		{TraditionalReligionsDay, 1997, time.Time{}, time.Time{}},
		{TraditionalReligionsDay, 1998, d(1998, 1, 10), d(1998, 1, 10)},
		{TraditionalReligionsDay, 2023, d(2023, 1, 10), d(2023, 1, 10)},
		{TraditionalReligionsDay, 2024, d(2024, 1, 10), d(2024, 1, 10)},
		{TraditionalReligionsDay, 2025, d(2025, 1, 10), d(2025, 1, 10)},
		{TraditionalReligionsDay, 2026, d(2026, 1, 9), d(2026, 1, 9)},
		{TraditionalReligionsDay, 2027, d(2027, 1, 8), d(2027, 1, 8)},

		{RemembranceDay, 2024, d(2024, 1, 16), d(2024, 1, 16)},
		{RemembranceDay, 2025, d(2025, 1, 16), d(2025, 1, 16)},
		{RemembranceDay, 2026, d(2026, 1, 16), d(2026, 1, 16)},

		{PeoplesSovereigntyDay, 2024, d(2024, 2, 28), d(2024, 2, 28)},
		{PeoplesSovereigntyDay, 2025, d(2025, 2, 28), d(2025, 2, 28)},
		{PeoplesSovereigntyDay, 2026, d(2026, 2, 28), d(2026, 2, 28)},

		{WomensDay, 2024, d(2024, 3, 8), d(2024, 3, 8)},
		{WomensDay, 2025, d(2025, 3, 8), d(2025, 3, 8)},
		{WomensDay, 2026, d(2026, 3, 8), d(2026, 3, 8)},

		{EasterMonday, 2024, d(2024, 4, 1), d(2024, 4, 1)},
		{EasterMonday, 2025, d(2025, 4, 21), d(2025, 4, 21)},
		{EasterMonday, 2026, d(2026, 4, 6), d(2026, 4, 6)},

		{LabourDay, 2024, d(2024, 5, 1), d(2024, 5, 1)},
		{LabourDay, 2025, d(2025, 5, 1), d(2025, 5, 1)},
		{LabourDay, 2026, d(2026, 5, 1), d(2026, 5, 1)},

		{AscensionDay, 2024, d(2024, 5, 9), d(2024, 5, 9)},
		{AscensionDay, 2025, d(2025, 5, 29), d(2025, 5, 29)},
		{AscensionDay, 2026, d(2026, 5, 14), d(2026, 5, 14)},

		{PentecostMonday, 2024, d(2024, 5, 20), d(2024, 5, 20)},
		{PentecostMonday, 2025, d(2025, 6, 9), d(2025, 6, 9)},
		{PentecostMonday, 2026, d(2026, 5, 25), d(2026, 5, 25)},

		{NationalDay, 2024, d(2024, 8, 1), d(2024, 8, 1)},
		{NationalDay, 2025, d(2025, 8, 1), d(2025, 8, 1)},
		{NationalDay, 2026, d(2026, 8, 1), d(2026, 8, 1)},

		{AssumptionDay, 2024, d(2024, 8, 15), d(2024, 8, 15)},
		{AssumptionDay, 2025, d(2025, 8, 15), d(2025, 8, 15)},
		{AssumptionDay, 2026, d(2026, 8, 15), d(2026, 8, 15)},

		{AllSaintsDay, 2024, d(2024, 11, 1), d(2024, 11, 1)},
		{AllSaintsDay, 2025, d(2025, 11, 1), d(2025, 11, 1)},
		{AllSaintsDay, 2026, d(2026, 11, 1), d(2026, 11, 1)},

		{ChristmasDay, 2024, d(2024, 12, 25), d(2024, 12, 25)},
		{ChristmasDay, 2025, d(2025, 12, 25), d(2025, 12, 25)},
		{ChristmasDay, 2026, d(2026, 12, 25), d(2026, 12, 25)},
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

func TestHolidaysSlice(t *testing.T) {
	want := []*cal.Holiday{
		NewYear,
		TraditionalReligionsDayEve,
		TraditionalReligionsDay,
		EasterMonday,
		LabourDay,
		AscensionDay,
		PentecostMonday,
		NationalDay,
		AssumptionDay,
		AllSaintsDay,
		ChristmasDay,
	}
	if len(Holidays) != len(want) {
		t.Fatalf("got %d holidays, want %d", len(Holidays), len(want))
	}
	for i, h := range want {
		if Holidays[i] != h {
			t.Errorf("Holidays[%d]: got %s, want %s", i, Holidays[i].Name, h.Name)
		}
		if h.Type != cal.ObservancePublic {
			t.Errorf("%s: got type %d, want %d", h.Name, h.Type, cal.ObservancePublic)
		}
	}

	for _, h := range []*cal.Holiday{RemembranceDay, PeoplesSovereigntyDay, WomensDay} {
		if h.Type != cal.ObservanceOther {
			t.Errorf("%s: got type %d, want %d", h.Name, h.Type, cal.ObservanceOther)
		}
	}

	c := cal.NewBusinessCalendar()
	c.AddHoliday(Holidays...)

	// 16 January 2026 is a Friday: Remembrance Day is not a day off
	if !c.IsWorkday(d(2026, 1, 16)) {
		t.Errorf("2026-01-16 should be a workday")
	}
	// 9 January 2026 is a Friday: the traditional religions festival is a day off
	if c.IsWorkday(d(2026, 1, 9)) {
		t.Errorf("2026-01-09 should not be a workday")
	}
}
