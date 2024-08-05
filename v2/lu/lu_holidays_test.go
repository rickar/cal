// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package lu

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
		{NeitJoer, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{NeitJoer, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{NeitJoer, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{NeitJoer, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{NeitJoer, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{NeitJoer, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{NeitJoer, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{NeitJoer, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{Ouschtermeindeg, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{Ouschtermeindeg, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{Ouschtermeindeg, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{Ouschtermeindeg, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{Ouschtermeindeg, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{Ouschtermeindeg, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{Ouschtermeindeg, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{Ouschtermeindeg, 2022, d(2022, 4, 18), d(2022, 4, 18)},

		{DagVunAarbecht, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{DagVunAarbecht, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{DagVunAarbecht, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{DagVunAarbecht, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{DagVunAarbecht, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{DagVunAarbecht, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{DagVunAarbecht, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{DagVunAarbecht, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{ChristiHimmelfaart, 2015, d(2015, 5, 14), d(2015, 5, 14)},
		{ChristiHimmelfaart, 2016, d(2016, 5, 5), d(2016, 5, 5)},
		{ChristiHimmelfaart, 2017, d(2017, 5, 25), d(2017, 5, 25)},
		{ChristiHimmelfaart, 2018, d(2018, 5, 10), d(2018, 5, 10)},
		{ChristiHimmelfaart, 2019, d(2019, 5, 30), d(2019, 5, 30)},
		{ChristiHimmelfaart, 2020, d(2020, 5, 21), d(2020, 5, 21)},
		{ChristiHimmelfaart, 2021, d(2021, 5, 13), d(2021, 5, 13)},
		{ChristiHimmelfaart, 2022, d(2022, 5, 26), d(2022, 5, 26)},

		{Pengschtméindeg, 2015, d(2015, 5, 25), d(2015, 5, 25)},
		{Pengschtméindeg, 2016, d(2016, 5, 16), d(2016, 5, 16)},
		{Pengschtméindeg, 2017, d(2017, 6, 5), d(2017, 6, 5)},
		{Pengschtméindeg, 2018, d(2018, 5, 21), d(2018, 5, 21)},
		{Pengschtméindeg, 2019, d(2019, 6, 10), d(2019, 6, 10)},
		{Pengschtméindeg, 2020, d(2020, 6, 1), d(2020, 6, 1)},
		{Pengschtméindeg, 2021, d(2021, 5, 24), d(2021, 5, 24)},
		{Pengschtméindeg, 2022, d(2022, 6, 6), d(2022, 6, 6)},

		{Nationalfeierdag, 2015, d(2015, 7, 23), d(2015, 7, 23)},
		{Nationalfeierdag, 2016, d(2016, 7, 23), d(2016, 7, 23)},
		{Nationalfeierdag, 2017, d(2017, 7, 23), d(2017, 7, 23)},
		{Nationalfeierdag, 2018, d(2018, 7, 23), d(2018, 7, 23)},
		{Nationalfeierdag, 2019, d(2019, 7, 23), d(2019, 7, 23)},
		{Nationalfeierdag, 2020, d(2020, 7, 23), d(2020, 7, 23)},
		{Nationalfeierdag, 2021, d(2021, 7, 23), d(2021, 7, 23)},
		{Nationalfeierdag, 2022, d(2022, 7, 23), d(2022, 7, 23)},

		{MariesHimmelfaart, 2015, d(2015, 8, 15), d(2015, 8, 15)},
		{MariesHimmelfaart, 2016, d(2016, 8, 15), d(2016, 8, 15)},
		{MariesHimmelfaart, 2017, d(2017, 8, 15), d(2017, 8, 15)},
		{MariesHimmelfaart, 2018, d(2018, 8, 15), d(2018, 8, 15)},
		{MariesHimmelfaart, 2019, d(2019, 8, 15), d(2019, 8, 15)},
		{MariesHimmelfaart, 2020, d(2020, 8, 15), d(2020, 8, 15)},
		{MariesHimmelfaart, 2021, d(2021, 8, 15), d(2021, 8, 15)},
		{MariesHimmelfaart, 2022, d(2022, 8, 15), d(2022, 8, 15)},

		{Allerhellgen, 2015, d(2015, 11, 1), d(2015, 11, 1)},
		{Allerhellgen, 2016, d(2016, 11, 1), d(2016, 11, 1)},
		{Allerhellgen, 2017, d(2017, 11, 1), d(2017, 11, 1)},
		{Allerhellgen, 2018, d(2018, 11, 1), d(2018, 11, 1)},
		{Allerhellgen, 2019, d(2019, 11, 1), d(2019, 11, 1)},
		{Allerhellgen, 2020, d(2020, 11, 1), d(2020, 11, 1)},
		{Allerhellgen, 2021, d(2021, 11, 1), d(2021, 11, 1)},
		{Allerhellgen, 2022, d(2022, 11, 1), d(2022, 11, 1)},

		{Chreschtdag, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Chreschtdag, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Chreschtdag, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Chreschtdag, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Chreschtdag, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Chreschtdag, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Chreschtdag, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Chreschtdag, 2022, d(2022, 12, 25), d(2022, 12, 25)},

		{ZweetenDagChrëschtdag, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{ZweetenDagChrëschtdag, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{ZweetenDagChrëschtdag, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{ZweetenDagChrëschtdag, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{ZweetenDagChrëschtdag, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{ZweetenDagChrëschtdag, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{ZweetenDagChrëschtdag, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{ZweetenDagChrëschtdag, 2022, d(2022, 12, 26), d(2022, 12, 26)},
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
