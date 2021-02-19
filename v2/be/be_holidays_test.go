// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package be

import (
	"testing"
	"time"

	"github.com/Tamh/cal/v2"
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
		{Nieuwjaar, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{Nieuwjaar, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{Nieuwjaar, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{Nieuwjaar, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{Nieuwjaar, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{Nieuwjaar, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{Nieuwjaar, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{Nieuwjaar, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{Paasmaandag, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{Paasmaandag, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{Paasmaandag, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{Paasmaandag, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{Paasmaandag, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{Paasmaandag, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{Paasmaandag, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{Paasmaandag, 2022, d(2022, 4, 18), d(2022, 4, 18)},

		{DagVanDeArbeid, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{DagVanDeArbeid, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{DagVanDeArbeid, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{DagVanDeArbeid, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{DagVanDeArbeid, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{DagVanDeArbeid, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{DagVanDeArbeid, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{DagVanDeArbeid, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{OnzeLieveHeerHemelvaart, 2015, d(2015, 5, 14), d(2015, 5, 14)},
		{OnzeLieveHeerHemelvaart, 2016, d(2016, 5, 5), d(2016, 5, 5)},
		{OnzeLieveHeerHemelvaart, 2017, d(2017, 5, 25), d(2017, 5, 25)},
		{OnzeLieveHeerHemelvaart, 2018, d(2018, 5, 10), d(2018, 5, 10)},
		{OnzeLieveHeerHemelvaart, 2019, d(2019, 5, 30), d(2019, 5, 30)},
		{OnzeLieveHeerHemelvaart, 2020, d(2020, 5, 21), d(2020, 5, 21)},
		{OnzeLieveHeerHemelvaart, 2021, d(2021, 5, 13), d(2021, 5, 13)},
		{OnzeLieveHeerHemelvaart, 2022, d(2022, 5, 26), d(2022, 5, 26)},

		{Pinkstermaandag, 2015, d(2015, 5, 25), d(2015, 5, 25)},
		{Pinkstermaandag, 2016, d(2016, 5, 16), d(2016, 5, 16)},
		{Pinkstermaandag, 2017, d(2017, 6, 5), d(2017, 6, 5)},
		{Pinkstermaandag, 2018, d(2018, 5, 21), d(2018, 5, 21)},
		{Pinkstermaandag, 2019, d(2019, 6, 10), d(2019, 6, 10)},
		{Pinkstermaandag, 2020, d(2020, 6, 1), d(2020, 6, 1)},
		{Pinkstermaandag, 2021, d(2021, 5, 24), d(2021, 5, 24)},
		{Pinkstermaandag, 2022, d(2022, 6, 6), d(2022, 6, 6)},

		{NationaleFeestdag, 2015, d(2015, 7, 21), d(2015, 7, 21)},
		{NationaleFeestdag, 2016, d(2016, 7, 21), d(2016, 7, 21)},
		{NationaleFeestdag, 2017, d(2017, 7, 21), d(2017, 7, 21)},
		{NationaleFeestdag, 2018, d(2018, 7, 21), d(2018, 7, 21)},
		{NationaleFeestdag, 2019, d(2019, 7, 21), d(2019, 7, 21)},
		{NationaleFeestdag, 2020, d(2020, 7, 21), d(2020, 7, 21)},
		{NationaleFeestdag, 2021, d(2021, 7, 21), d(2021, 7, 21)},
		{NationaleFeestdag, 2022, d(2022, 7, 21), d(2022, 7, 21)},

		{OnzeLieveVrouwHemelvaart, 2015, d(2015, 8, 15), d(2015, 8, 15)},
		{OnzeLieveVrouwHemelvaart, 2016, d(2016, 8, 15), d(2016, 8, 15)},
		{OnzeLieveVrouwHemelvaart, 2017, d(2017, 8, 15), d(2017, 8, 15)},
		{OnzeLieveVrouwHemelvaart, 2018, d(2018, 8, 15), d(2018, 8, 15)},
		{OnzeLieveVrouwHemelvaart, 2019, d(2019, 8, 15), d(2019, 8, 15)},
		{OnzeLieveVrouwHemelvaart, 2020, d(2020, 8, 15), d(2020, 8, 15)},
		{OnzeLieveVrouwHemelvaart, 2021, d(2021, 8, 15), d(2021, 8, 15)},
		{OnzeLieveVrouwHemelvaart, 2022, d(2022, 8, 15), d(2022, 8, 15)},

		{Allerheiligen, 2015, d(2015, 11, 1), d(2015, 11, 1)},
		{Allerheiligen, 2016, d(2016, 11, 1), d(2016, 11, 1)},
		{Allerheiligen, 2017, d(2017, 11, 1), d(2017, 11, 1)},
		{Allerheiligen, 2018, d(2018, 11, 1), d(2018, 11, 1)},
		{Allerheiligen, 2019, d(2019, 11, 1), d(2019, 11, 1)},
		{Allerheiligen, 2020, d(2020, 11, 1), d(2020, 11, 1)},
		{Allerheiligen, 2021, d(2021, 11, 1), d(2021, 11, 1)},
		{Allerheiligen, 2022, d(2022, 11, 1), d(2022, 11, 1)},

		{Wapenstilstand, 2015, d(2015, 11, 11), d(2015, 11, 11)},
		{Wapenstilstand, 2016, d(2016, 11, 11), d(2016, 11, 11)},
		{Wapenstilstand, 2017, d(2017, 11, 11), d(2017, 11, 11)},
		{Wapenstilstand, 2018, d(2018, 11, 11), d(2018, 11, 11)},
		{Wapenstilstand, 2019, d(2019, 11, 11), d(2019, 11, 11)},
		{Wapenstilstand, 2020, d(2020, 11, 11), d(2020, 11, 11)},
		{Wapenstilstand, 2021, d(2021, 11, 11), d(2021, 11, 11)},
		{Wapenstilstand, 2022, d(2022, 11, 11), d(2022, 11, 11)},

		{Kerstmis, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Kerstmis, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Kerstmis, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Kerstmis, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Kerstmis, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Kerstmis, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Kerstmis, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Kerstmis, 2022, d(2022, 12, 25), d(2022, 12, 25)},
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
