// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package ch

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
		{Neujahr, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{Neujahr, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{Neujahr, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{Neujahr, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{Neujahr, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{Neujahr, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{Neujahr, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{Neujahr, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{Berchtoldstag, 2015, d(2015, 1, 2), d(2015, 1, 2)},
		{Berchtoldstag, 2016, d(2016, 1, 2), d(2016, 1, 2)},
		{Berchtoldstag, 2017, d(2017, 1, 2), d(2017, 1, 2)},
		{Berchtoldstag, 2018, d(2018, 1, 2), d(2018, 1, 2)},
		{Berchtoldstag, 2019, d(2019, 1, 2), d(2019, 1, 2)},
		{Berchtoldstag, 2020, d(2020, 1, 2), d(2020, 1, 2)},
		{Berchtoldstag, 2021, d(2021, 1, 2), d(2021, 1, 2)},
		{Berchtoldstag, 2022, d(2022, 1, 2), d(2022, 1, 2)},

		{HeiligeDreiKoenige, 2015, d(2015, 1, 6), d(2015, 1, 6)},
		{HeiligeDreiKoenige, 2016, d(2016, 1, 6), d(2016, 1, 6)},
		{HeiligeDreiKoenige, 2017, d(2017, 1, 6), d(2017, 1, 6)},
		{HeiligeDreiKoenige, 2018, d(2018, 1, 6), d(2018, 1, 6)},
		{HeiligeDreiKoenige, 2019, d(2019, 1, 6), d(2019, 1, 6)},
		{HeiligeDreiKoenige, 2020, d(2020, 1, 6), d(2020, 1, 6)},
		{HeiligeDreiKoenige, 2021, d(2021, 1, 6), d(2021, 1, 6)},
		{HeiligeDreiKoenige, 2022, d(2022, 1, 6), d(2022, 1, 6)},

		{Josefstag, 2015, d(2015, 3, 19), d(2015, 3, 19)},
		{Josefstag, 2016, d(2016, 3, 19), d(2016, 3, 19)},
		{Josefstag, 2017, d(2017, 3, 19), d(2017, 3, 19)},
		{Josefstag, 2018, d(2018, 3, 19), d(2018, 3, 19)},
		{Josefstag, 2019, d(2019, 3, 19), d(2019, 3, 19)},
		{Josefstag, 2020, d(2020, 3, 19), d(2020, 3, 19)},
		{Josefstag, 2021, d(2021, 3, 19), d(2021, 3, 19)},
		{Josefstag, 2022, d(2022, 3, 19), d(2022, 3, 19)},

		{Karfreitag, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{Karfreitag, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{Karfreitag, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{Karfreitag, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{Karfreitag, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{Karfreitag, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{Karfreitag, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{Karfreitag, 2022, d(2022, 4, 15), d(2022, 4, 15)},

		{Ostermontag, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{Ostermontag, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{Ostermontag, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{Ostermontag, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{Ostermontag, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{Ostermontag, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{Ostermontag, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{Ostermontag, 2022, d(2022, 4, 18), d(2022, 4, 18)},

		{TagderArbeit, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{TagderArbeit, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{TagderArbeit, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{TagderArbeit, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{TagderArbeit, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{TagderArbeit, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{TagderArbeit, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{TagderArbeit, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{Auffahrt, 2015, d(2015, 5, 14), d(2015, 5, 14)},
		{Auffahrt, 2016, d(2016, 5, 5), d(2016, 5, 5)},
		{Auffahrt, 2017, d(2017, 5, 25), d(2017, 5, 25)},
		{Auffahrt, 2018, d(2018, 5, 10), d(2018, 5, 10)},
		{Auffahrt, 2019, d(2019, 5, 30), d(2019, 5, 30)},
		{Auffahrt, 2020, d(2020, 5, 21), d(2020, 5, 21)},
		{Auffahrt, 2021, d(2021, 5, 13), d(2021, 5, 13)},
		{Auffahrt, 2022, d(2022, 5, 26), d(2022, 5, 26)},

		{Pfingstmontag, 2015, d(2015, 5, 25), d(2015, 5, 25)},
		{Pfingstmontag, 2016, d(2016, 5, 16), d(2016, 5, 16)},
		{Pfingstmontag, 2017, d(2017, 6, 5), d(2017, 6, 5)},
		{Pfingstmontag, 2018, d(2018, 5, 21), d(2018, 5, 21)},
		{Pfingstmontag, 2019, d(2019, 6, 10), d(2019, 6, 10)},
		{Pfingstmontag, 2020, d(2020, 6, 1), d(2020, 6, 1)},
		{Pfingstmontag, 2021, d(2021, 5, 24), d(2021, 5, 24)},
		{Pfingstmontag, 2022, d(2022, 6, 6), d(2022, 6, 6)},

		{Fronleichnam, 2015, d(2015, 6, 4), d(2015, 6, 4)},
		{Fronleichnam, 2016, d(2016, 5, 26), d(2016, 5, 26)},
		{Fronleichnam, 2017, d(2017, 6, 15), d(2017, 6, 15)},
		{Fronleichnam, 2018, d(2018, 5, 31), d(2018, 5, 31)},
		{Fronleichnam, 2019, d(2019, 6, 20), d(2019, 6, 20)},
		{Fronleichnam, 2020, d(2020, 6, 11), d(2020, 6, 11)},
		{Fronleichnam, 2021, d(2021, 6, 3), d(2021, 6, 3)},
		{Fronleichnam, 2022, d(2022, 6, 16), d(2022, 6, 16)},

		{Bundesfeiertag, 2015, d(2015, 8, 1), d(2015, 8, 1)},
		{Bundesfeiertag, 2016, d(2016, 8, 1), d(2016, 8, 1)},
		{Bundesfeiertag, 2017, d(2017, 8, 1), d(2017, 8, 1)},
		{Bundesfeiertag, 2018, d(2018, 8, 1), d(2018, 8, 1)},
		{Bundesfeiertag, 2019, d(2019, 8, 1), d(2019, 8, 1)},
		{Bundesfeiertag, 2020, d(2020, 8, 1), d(2020, 8, 1)},
		{Bundesfeiertag, 2021, d(2021, 8, 1), d(2021, 8, 1)},
		{Bundesfeiertag, 2022, d(2022, 8, 1), d(2022, 8, 1)},

		{MariaHimmelfahrt, 2015, d(2015, 8, 15), d(2015, 8, 15)},
		{MariaHimmelfahrt, 2016, d(2016, 8, 15), d(2016, 8, 15)},
		{MariaHimmelfahrt, 2017, d(2017, 8, 15), d(2017, 8, 15)},
		{MariaHimmelfahrt, 2018, d(2018, 8, 15), d(2018, 8, 15)},
		{MariaHimmelfahrt, 2019, d(2019, 8, 15), d(2019, 8, 15)},
		{MariaHimmelfahrt, 2020, d(2020, 8, 15), d(2020, 8, 15)},
		{MariaHimmelfahrt, 2021, d(2021, 8, 15), d(2021, 8, 15)},
		{MariaHimmelfahrt, 2022, d(2022, 8, 15), d(2022, 8, 15)},

		{Allerheiligen, 2015, d(2015, 11, 1), d(2015, 11, 1)},
		{Allerheiligen, 2016, d(2016, 11, 1), d(2016, 11, 1)},
		{Allerheiligen, 2017, d(2017, 11, 1), d(2017, 11, 1)},
		{Allerheiligen, 2018, d(2018, 11, 1), d(2018, 11, 1)},
		{Allerheiligen, 2019, d(2019, 11, 1), d(2019, 11, 1)},
		{Allerheiligen, 2020, d(2020, 11, 1), d(2020, 11, 1)},
		{Allerheiligen, 2021, d(2021, 11, 1), d(2021, 11, 1)},
		{Allerheiligen, 2022, d(2022, 11, 1), d(2022, 11, 1)},

		{MariaEmpfangnis, 2015, d(2015, 12, 8), d(2015, 12, 8)},
		{MariaEmpfangnis, 2016, d(2016, 12, 8), d(2016, 12, 8)},
		{MariaEmpfangnis, 2017, d(2017, 12, 8), d(2017, 12, 8)},
		{MariaEmpfangnis, 2018, d(2018, 12, 8), d(2018, 12, 8)},
		{MariaEmpfangnis, 2019, d(2019, 12, 8), d(2019, 12, 8)},
		{MariaEmpfangnis, 2020, d(2020, 12, 8), d(2020, 12, 8)},
		{MariaEmpfangnis, 2021, d(2021, 12, 8), d(2021, 12, 8)},
		{MariaEmpfangnis, 2022, d(2022, 12, 8), d(2022, 12, 8)},

		{Weihnachtstag, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Weihnachtstag, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Weihnachtstag, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Weihnachtstag, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Weihnachtstag, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Weihnachtstag, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Weihnachtstag, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Weihnachtstag, 2022, d(2022, 12, 25), d(2022, 12, 25)},

		{ZweiterWeihnachtsfeiertag, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{ZweiterWeihnachtsfeiertag, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{ZweiterWeihnachtsfeiertag, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{ZweiterWeihnachtsfeiertag, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{ZweiterWeihnachtsfeiertag, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{ZweiterWeihnachtsfeiertag, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{ZweiterWeihnachtsfeiertag, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{ZweiterWeihnachtsfeiertag, 2022, d(2022, 12, 26), d(2022, 12, 26)},
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
