// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package ru

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
		{NewYear, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{NewYear, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{NewYear, 2017, d(2017, 1, 1), d(2017, 1, 2)},
		{NewYear, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{NewYear, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{NewYear, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{NewYear, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{NewYear, 2022, d(2022, 1, 1), d(2022, 1, 3)},

		{OrthodoxChristmas, 2015, d(2015, 1, 7), d(2015, 1, 7)},
		{OrthodoxChristmas, 2016, d(2016, 1, 7), d(2016, 1, 7)},
		{OrthodoxChristmas, 2017, d(2017, 1, 7), d(2017, 1, 9)},
		{OrthodoxChristmas, 2018, d(2018, 1, 7), d(2018, 1, 8)},
		{OrthodoxChristmas, 2019, d(2019, 1, 7), d(2019, 1, 7)},
		{OrthodoxChristmas, 2020, d(2020, 1, 7), d(2020, 1, 7)},
		{OrthodoxChristmas, 2021, d(2021, 1, 7), d(2021, 1, 7)},
		{OrthodoxChristmas, 2022, d(2022, 1, 7), d(2022, 1, 7)},

		{MilitaryDay, 2015, d(2015, 2, 23), d(2015, 2, 23)},
		{MilitaryDay, 2016, d(2016, 2, 23), d(2016, 2, 23)},
		{MilitaryDay, 2017, d(2017, 2, 23), d(2017, 2, 23)},
		{MilitaryDay, 2018, d(2018, 2, 23), d(2018, 2, 23)},
		{MilitaryDay, 2019, d(2019, 2, 23), d(2019, 2, 25)},
		{MilitaryDay, 2020, d(2020, 2, 23), d(2020, 2, 24)},
		{MilitaryDay, 2021, d(2021, 2, 23), d(2021, 2, 23)},
		{MilitaryDay, 2022, d(2022, 2, 23), d(2022, 2, 23)},

		{WomensDay, 2015, d(2015, 3, 8), d(2015, 3, 9)},
		{WomensDay, 2016, d(2016, 3, 8), d(2016, 3, 8)},
		{WomensDay, 2017, d(2017, 3, 8), d(2017, 3, 8)},
		{WomensDay, 2018, d(2018, 3, 8), d(2018, 3, 8)},
		{WomensDay, 2019, d(2019, 3, 8), d(2019, 3, 8)},
		{WomensDay, 2020, d(2020, 3, 8), d(2020, 3, 9)},
		{WomensDay, 2021, d(2021, 3, 8), d(2021, 3, 8)},
		{WomensDay, 2022, d(2022, 3, 8), d(2022, 3, 8)},

		{LabourDay, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{LabourDay, 2016, d(2016, 5, 1), d(2016, 5, 2)},
		{LabourDay, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{LabourDay, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{LabourDay, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{LabourDay, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{LabourDay, 2021, d(2021, 5, 1), d(2021, 5, 3)},
		{LabourDay, 2022, d(2022, 5, 1), d(2022, 5, 2)},

		{VictoryDay, 2015, d(2015, 5, 9), d(2015, 5, 11)},
		{VictoryDay, 2016, d(2016, 5, 9), d(2016, 5, 9)},
		{VictoryDay, 2017, d(2017, 5, 9), d(2017, 5, 9)},
		{VictoryDay, 2018, d(2018, 5, 9), d(2018, 5, 9)},
		{VictoryDay, 2019, d(2019, 5, 9), d(2019, 5, 9)},
		{VictoryDay, 2020, d(2020, 5, 9), d(2020, 5, 11)},
		{VictoryDay, 2021, d(2021, 5, 9), d(2021, 5, 10)},
		{VictoryDay, 2022, d(2022, 5, 9), d(2022, 5, 9)},

		{RussiasDay, 2015, d(2015, 6, 12), d(2015, 6, 12)},
		{RussiasDay, 2016, d(2016, 6, 12), d(2016, 6, 13)},
		{RussiasDay, 2017, d(2017, 6, 12), d(2017, 6, 12)},
		{RussiasDay, 2018, d(2018, 6, 12), d(2018, 6, 12)},
		{RussiasDay, 2019, d(2019, 6, 12), d(2019, 6, 12)},
		{RussiasDay, 2020, d(2020, 6, 12), d(2020, 6, 12)},
		{RussiasDay, 2021, d(2021, 6, 12), d(2021, 6, 14)},
		{RussiasDay, 2022, d(2022, 6, 12), d(2022, 6, 13)},

		{UnionDay, 2015, d(2015, 11, 4), d(2015, 11, 4)},
		{UnionDay, 2016, d(2016, 11, 4), d(2016, 11, 4)},
		{UnionDay, 2017, d(2017, 11, 4), d(2017, 11, 6)},
		{UnionDay, 2018, d(2018, 11, 4), d(2018, 11, 5)},
		{UnionDay, 2019, d(2019, 11, 4), d(2019, 11, 4)},
		{UnionDay, 2020, d(2020, 11, 4), d(2020, 11, 4)},
		{UnionDay, 2021, d(2021, 11, 4), d(2021, 11, 4)},
		{UnionDay, 2022, d(2022, 11, 4), d(2022, 11, 4)},
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
