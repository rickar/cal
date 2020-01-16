package cal

import (
	"testing"
	"time"
)

func TestItalianHolidays(t *testing.T) {
	c := NewCalendar()
	c.Observed = ObservedExact
	AddItalianHolidays(c)

	tests := []testStruct{
		{time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC), true, "ITCapodanno"},
		{time.Date(2020, 1, 6, 12, 0, 0, 0, time.UTC), true, "ITEpifania"},
		{time.Date(2020, 4, 13, 12, 0, 0, 0, time.UTC), true, "ITPasquetta"},
		{time.Date(2020, 4, 25, 12, 0, 0, 0, time.UTC), true, "ITFestaDellaLiberazione"},
		{time.Date(2020, 5, 1, 12, 0, 0, 0, time.UTC), true, "ITFestaDelLavoro"},
		{time.Date(2020, 6, 2, 12, 0, 0, 0, time.UTC), true, "ITFestaDellaRepubblica"},
		{time.Date(2020, 8, 15, 12, 0, 0, 0, time.UTC), true, "ITFerragosto"},
		{time.Date(2020, 11, 1, 12, 0, 0, 0, time.UTC), true, "ITTuttiISanti"},
		{time.Date(2020, 12, 8, 12, 0, 0, 0, time.UTC), true, "ITImmacolata"},
		{time.Date(2020, 12, 25, 12, 0, 0, 0, time.UTC), true, "ITNatale"},
		{time.Date(2020, 12, 26, 12, 0, 0, 0, time.UTC), true, "ITSantoStefano"},
	}

	for _, test := range tests {
		got := c.IsHoliday(test.t)
		if got != test.want {
			t.Errorf("got: %t for %s; want: %t (%s)", got, test.name, test.want, test.t)
		}
	}
}
