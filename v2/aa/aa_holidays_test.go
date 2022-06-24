// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package aa

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
		{NewYear, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{NewYear, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{NewYear, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{NewYear, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{NewYear, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{NewYear, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{Epiphany, 2015, d(2015, 1, 6), d(2015, 1, 6)},
		{Epiphany, 2016, d(2016, 1, 6), d(2016, 1, 6)},
		{Epiphany, 2017, d(2017, 1, 6), d(2017, 1, 6)},
		{Epiphany, 2018, d(2018, 1, 6), d(2018, 1, 6)},
		{Epiphany, 2019, d(2019, 1, 6), d(2019, 1, 6)},
		{Epiphany, 2020, d(2020, 1, 6), d(2020, 1, 6)},
		{Epiphany, 2021, d(2021, 1, 6), d(2021, 1, 6)},
		{Epiphany, 2022, d(2022, 1, 6), d(2022, 1, 6)},

		{MaundyThursday, 2015, d(2015, 4, 2), d(2015, 4, 2)},
		{MaundyThursday, 2016, d(2016, 3, 24), d(2016, 3, 24)},
		{MaundyThursday, 2017, d(2017, 4, 13), d(2017, 4, 13)},
		{MaundyThursday, 2018, d(2018, 3, 29), d(2018, 3, 29)},
		{MaundyThursday, 2019, d(2019, 4, 18), d(2019, 4, 18)},
		{MaundyThursday, 2020, d(2020, 4, 9), d(2020, 4, 9)},
		{MaundyThursday, 2021, d(2021, 4, 1), d(2021, 4, 1)},
		{MaundyThursday, 2022, d(2022, 4, 14), d(2022, 4, 14)},

		{GoodFriday, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{GoodFriday, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{GoodFriday, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{GoodFriday, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{GoodFriday, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{GoodFriday, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{GoodFriday, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{GoodFriday, 2022, d(2022, 4, 15), d(2022, 4, 15)},

		{Easter, 2015, d(2015, 4, 5), d(2015, 4, 5)},
		{Easter, 2016, d(2016, 3, 27), d(2016, 3, 27)},
		{Easter, 2017, d(2017, 4, 16), d(2017, 4, 16)},
		{Easter, 2018, d(2018, 4, 1), d(2018, 4, 1)},
		{Easter, 2019, d(2019, 4, 21), d(2019, 4, 21)},
		{Easter, 2020, d(2020, 4, 12), d(2020, 4, 12)},
		{Easter, 2021, d(2021, 4, 4), d(2021, 4, 4)},
		{Easter, 2022, d(2022, 4, 17), d(2022, 4, 17)},

		{EasterMonday, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{EasterMonday, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{EasterMonday, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{EasterMonday, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{EasterMonday, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{EasterMonday, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{EasterMonday, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{EasterMonday, 2022, d(2022, 4, 18), d(2022, 4, 18)},

		{WorkersDay, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{WorkersDay, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{WorkersDay, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{WorkersDay, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{WorkersDay, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{WorkersDay, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{WorkersDay, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{WorkersDay, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{AscensionDay, 2015, d(2015, 5, 14), d(2015, 5, 14)},
		{AscensionDay, 2016, d(2016, 5, 5), d(2016, 5, 5)},
		{AscensionDay, 2017, d(2017, 5, 25), d(2017, 5, 25)},
		{AscensionDay, 2018, d(2018, 5, 10), d(2018, 5, 10)},
		{AscensionDay, 2019, d(2019, 5, 30), d(2019, 5, 30)},
		{AscensionDay, 2020, d(2020, 5, 21), d(2020, 5, 21)},
		{AscensionDay, 2021, d(2021, 5, 13), d(2021, 5, 13)},
		{AscensionDay, 2022, d(2022, 5, 26), d(2022, 5, 26)},

		{Pentecost, 2015, d(2015, 5, 24), d(2015, 5, 24)},
		{Pentecost, 2016, d(2016, 5, 15), d(2016, 5, 15)},
		{Pentecost, 2017, d(2017, 6, 4), d(2017, 6, 4)},
		{Pentecost, 2018, d(2018, 5, 20), d(2018, 5, 20)},
		{Pentecost, 2019, d(2019, 6, 9), d(2019, 6, 9)},
		{Pentecost, 2020, d(2020, 5, 31), d(2020, 5, 31)},
		{Pentecost, 2021, d(2021, 5, 23), d(2021, 5, 23)},
		{Pentecost, 2022, d(2022, 6, 5), d(2022, 6, 5)},

		{PentecostMonday, 2015, d(2015, 5, 25), d(2015, 5, 25)},
		{PentecostMonday, 2016, d(2016, 5, 16), d(2016, 5, 16)},
		{PentecostMonday, 2017, d(2017, 6, 5), d(2017, 6, 5)},
		{PentecostMonday, 2018, d(2018, 5, 21), d(2018, 5, 21)},
		{PentecostMonday, 2019, d(2019, 6, 10), d(2019, 6, 10)},
		{PentecostMonday, 2020, d(2020, 6, 1), d(2020, 6, 1)},
		{PentecostMonday, 2021, d(2021, 5, 24), d(2021, 5, 24)},
		{PentecostMonday, 2022, d(2022, 6, 6), d(2022, 6, 6)},

		{CorpusChristi, 2015, d(2015, 6, 4), d(2015, 6, 4)},
		{CorpusChristi, 2016, d(2016, 5, 26), d(2016, 5, 26)},
		{CorpusChristi, 2017, d(2017, 6, 15), d(2017, 6, 15)},
		{CorpusChristi, 2018, d(2018, 5, 31), d(2018, 5, 31)},
		{CorpusChristi, 2019, d(2019, 6, 20), d(2019, 6, 20)},
		{CorpusChristi, 2020, d(2020, 6, 11), d(2020, 6, 11)},
		{CorpusChristi, 2021, d(2021, 6, 3), d(2021, 6, 3)},
		{CorpusChristi, 2022, d(2022, 6, 16), d(2022, 6, 16)},

		{AssumptionOfMary, 2015, d(2015, 8, 15), d(2015, 8, 15)},
		{AssumptionOfMary, 2016, d(2016, 8, 15), d(2016, 8, 15)},
		{AssumptionOfMary, 2017, d(2017, 8, 15), d(2017, 8, 15)},
		{AssumptionOfMary, 2018, d(2018, 8, 15), d(2018, 8, 15)},
		{AssumptionOfMary, 2019, d(2019, 8, 15), d(2019, 8, 15)},
		{AssumptionOfMary, 2020, d(2020, 8, 15), d(2020, 8, 15)},
		{AssumptionOfMary, 2021, d(2021, 8, 15), d(2021, 8, 15)},
		{AssumptionOfMary, 2022, d(2022, 8, 15), d(2022, 8, 15)},

		{AllSaintsDay, 2015, d(2015, 11, 1), d(2015, 11, 1)},
		{AllSaintsDay, 2016, d(2016, 11, 1), d(2016, 11, 1)},
		{AllSaintsDay, 2017, d(2017, 11, 1), d(2017, 11, 1)},
		{AllSaintsDay, 2018, d(2018, 11, 1), d(2018, 11, 1)},
		{AllSaintsDay, 2019, d(2019, 11, 1), d(2019, 11, 1)},
		{AllSaintsDay, 2020, d(2020, 11, 1), d(2020, 11, 1)},
		{AllSaintsDay, 2021, d(2021, 11, 1), d(2021, 11, 1)},
		{AllSaintsDay, 2022, d(2022, 11, 1), d(2022, 11, 1)},

		{ArmisticeDay, 2015, d(2015, 11, 11), d(2015, 11, 11)},
		{ArmisticeDay, 2016, d(2016, 11, 11), d(2016, 11, 11)},
		{ArmisticeDay, 2017, d(2017, 11, 11), d(2017, 11, 11)},
		{ArmisticeDay, 2018, d(2018, 11, 11), d(2018, 11, 11)},
		{ArmisticeDay, 2019, d(2019, 11, 11), d(2019, 11, 11)},
		{ArmisticeDay, 2020, d(2020, 11, 11), d(2020, 11, 11)},
		{ArmisticeDay, 2021, d(2021, 11, 11), d(2021, 11, 11)},
		{ArmisticeDay, 2022, d(2022, 11, 11), d(2022, 11, 11)},

		{ImmaculateConception, 2015, d(2015, 12, 8), d(2015, 12, 8)},
		{ImmaculateConception, 2016, d(2016, 12, 8), d(2016, 12, 8)},
		{ImmaculateConception, 2017, d(2017, 12, 8), d(2017, 12, 8)},
		{ImmaculateConception, 2018, d(2018, 12, 8), d(2018, 12, 8)},
		{ImmaculateConception, 2019, d(2019, 12, 8), d(2019, 12, 8)},
		{ImmaculateConception, 2020, d(2020, 12, 8), d(2020, 12, 8)},
		{ImmaculateConception, 2021, d(2021, 12, 8), d(2021, 12, 8)},
		{ImmaculateConception, 2022, d(2022, 12, 8), d(2022, 12, 8)},

		{ChristmasDay, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{ChristmasDay, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{ChristmasDay, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{ChristmasDay, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{ChristmasDay, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{ChristmasDay, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{ChristmasDay, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{ChristmasDay, 2022, d(2022, 12, 25), d(2022, 12, 25)},

		{ChristmasDay2, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{ChristmasDay2, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{ChristmasDay2, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{ChristmasDay2, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{ChristmasDay2, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{ChristmasDay2, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{ChristmasDay2, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{ChristmasDay2, 2022, d(2022, 12, 26), d(2022, 12, 26)},
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
