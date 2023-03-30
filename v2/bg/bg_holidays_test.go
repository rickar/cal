package bg

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
		{NewYear, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{NewYear, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{NewYear, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{NewYear, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{NewYear, 2022, d(2022, 1, 1), d(2022, 1, 3)},
		{NewYear, 2023, d(2023, 1, 1), d(2023, 1, 2)},

		{LiberationDay, 2018, d(2018, 3, 3), d(2018, 3, 5)},
		{LiberationDay, 2019, d(2019, 3, 3), d(2019, 3, 4)},
		{LiberationDay, 2020, d(2020, 3, 3), d(2020, 3, 3)},
		{LiberationDay, 2021, d(2021, 3, 3), d(2021, 3, 3)},
		{LiberationDay, 2022, d(2022, 3, 3), d(2022, 3, 3)},
		{LiberationDay, 2023, d(2023, 3, 3), d(2023, 3, 3)},

		{OrthodoxGoodFriday, 2018, d(2018, 4, 6), d(2018, 4, 6)},
		{OrthodoxGoodFriday, 2019, d(2019, 4, 26), d(2019, 4, 26)},
		{OrthodoxGoodFriday, 2020, d(2020, 4, 17), d(2020, 4, 17)},
		{OrthodoxGoodFriday, 2021, d(2021, 4, 30), d(2021, 4, 30)},
		{OrthodoxGoodFriday, 2022, d(2022, 4, 22), d(2022, 4, 22)},
		{OrthodoxGoodFriday, 2023, d(2023, 4, 14), d(2023, 4, 14)},

		{OrthodoxEasterMonday, 2018, d(2018, 4, 9), d(2018, 4, 9)},
		{OrthodoxEasterMonday, 2019, d(2019, 4, 29), d(2019, 4, 29)},
		{OrthodoxEasterMonday, 2020, d(2020, 4, 20), d(2020, 4, 20)},
		{OrthodoxEasterMonday, 2021, d(2021, 5, 3), d(2021, 5, 3)},
		{OrthodoxEasterMonday, 2022, d(2022, 4, 25), d(2022, 4, 25)},
		{OrthodoxEasterMonday, 2023, d(2023, 4, 17), d(2023, 4, 17)},

		{LabourDay, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{LabourDay, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{LabourDay, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		// Special case since OrthodoxEaster falls on that weekend
		{LabourDay, 2021, d(2021, 5, 4), d(2021, 5, 4)},
		{LabourDay, 2022, d(2022, 5, 1), d(2022, 5, 2)},
		{LabourDay, 2023, d(2023, 5, 1), d(2023, 5, 1)},

		{StGeorgesDay, 2018, d(2018, 5, 6), d(2018, 5, 7)},
		{StGeorgesDay, 2019, d(2019, 5, 6), d(2019, 5, 6)},
		{StGeorgesDay, 2020, d(2020, 5, 6), d(2020, 5, 6)},
		{StGeorgesDay, 2021, d(2021, 5, 6), d(2021, 5, 6)},
		{StGeorgesDay, 2022, d(2022, 5, 6), d(2022, 5, 6)},
		{StGeorgesDay, 2023, d(2023, 5, 6), d(2023, 5, 8)},

		{StCyrilAndMethodiusDay, 2018, d(2018, 5, 24), d(2018, 5, 24)},
		{StCyrilAndMethodiusDay, 2019, d(2019, 5, 24), d(2019, 5, 24)},
		{StCyrilAndMethodiusDay, 2020, d(2020, 5, 24), d(2020, 5, 25)},
		{StCyrilAndMethodiusDay, 2021, d(2021, 5, 24), d(2021, 5, 24)},
		{StCyrilAndMethodiusDay, 2022, d(2022, 5, 24), d(2022, 5, 24)},
		{StCyrilAndMethodiusDay, 2023, d(2023, 5, 24), d(2023, 5, 24)},

		{UnificationDay, 2018, d(2018, 9, 6), d(2018, 9, 6)},
		{UnificationDay, 2019, d(2019, 9, 6), d(2019, 9, 6)},
		{UnificationDay, 2020, d(2020, 9, 6), d(2020, 9, 7)},
		{UnificationDay, 2021, d(2021, 9, 6), d(2021, 9, 6)},
		{UnificationDay, 2022, d(2022, 9, 6), d(2022, 9, 6)},
		{UnificationDay, 2023, d(2023, 9, 6), d(2023, 9, 6)},

		{IndependenceDay, 2018, d(2018, 9, 22), d(2018, 9, 24)},
		{IndependenceDay, 2019, d(2019, 9, 22), d(2019, 9, 23)},
		{IndependenceDay, 2020, d(2020, 9, 22), d(2020, 9, 22)},
		{IndependenceDay, 2021, d(2021, 9, 22), d(2021, 9, 22)},
		{IndependenceDay, 2022, d(2022, 9, 22), d(2022, 9, 22)},
		{IndependenceDay, 2023, d(2023, 9, 22), d(2023, 9, 22)},

		{ChristmasEve, 2018, d(2018, 12, 24), d(2018, 12, 24)},
		{ChristmasEve, 2019, d(2019, 12, 24), d(2019, 12, 24)},
		{ChristmasEve, 2020, d(2020, 12, 24), d(2020, 12, 24)},
		{ChristmasEve, 2021, d(2021, 12, 24), d(2021, 12, 24)},
		{ChristmasEve, 2022, d(2022, 12, 24), d(2022, 12, 26)},
		{ChristmasEve, 2023, d(2023, 12, 24), d(2023, 12, 25)},

		{ChristmasDay, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{ChristmasDay, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{ChristmasDay, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{ChristmasDay, 2021, d(2021, 12, 25), d(2021, 12, 27)},
		{ChristmasDay, 2022, d(2022, 12, 25), d(2022, 12, 27)},
		{ChristmasDay, 2023, d(2023, 12, 25), d(2023, 12, 25)},

		{ChristmasDay2, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{ChristmasDay2, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{ChristmasDay2, 2020, d(2020, 12, 26), d(2020, 12, 28)},
		{ChristmasDay2, 2021, d(2021, 12, 26), d(2021, 12, 28)},
		{ChristmasDay2, 2022, d(2022, 12, 26), d(2022, 12, 28)},
		{ChristmasDay2, 2023, d(2023, 12, 26), d(2023, 12, 27)},
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
