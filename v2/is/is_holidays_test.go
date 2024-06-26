// Package is provides holiday definitions for Iceland.
package is

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
		{Nyarsdagur, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{Nyarsdagur, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{Nyarsdagur, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{Nyarsdagur, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{Nyarsdagur, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{Nyarsdagur, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{Nyarsdagur, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{Nyarsdagur, 2022, d(2022, 1, 1), d(2022, 1, 1)},
		{Nyarsdagur, 2023, d(2023, 1, 1), d(2023, 1, 1)},

		{Skirdagur, 2015, d(2015, 4, 2), d(2015, 4, 2)},
		{Skirdagur, 2016, d(2016, 3, 24), d(2016, 3, 24)},
		{Skirdagur, 2017, d(2017, 4, 13), d(2017, 4, 13)},
		{Skirdagur, 2018, d(2018, 3, 29), d(2018, 3, 29)},
		{Skirdagur, 2019, d(2019, 4, 18), d(2019, 4, 18)},
		{Skirdagur, 2020, d(2020, 4, 9), d(2020, 4, 9)},
		{Skirdagur, 2021, d(2021, 4, 1), d(2021, 4, 1)},
		{Skirdagur, 2022, d(2022, 4, 14), d(2022, 4, 14)},
		{Skirdagur, 2023, d(2023, 4, 6), d(2023, 4, 6)},

		{Langifostudagur, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{Langifostudagur, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{Langifostudagur, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{Langifostudagur, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{Langifostudagur, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{Langifostudagur, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{Langifostudagur, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{Langifostudagur, 2022, d(2022, 4, 15), d(2022, 4, 15)},
		{Langifostudagur, 2023, d(2023, 4, 7), d(2023, 4, 7)},

		{Annaripaskum, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{Annaripaskum, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{Annaripaskum, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{Annaripaskum, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{Annaripaskum, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{Annaripaskum, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{Annaripaskum, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{Annaripaskum, 2022, d(2022, 4, 18), d(2022, 4, 18)},
		{Annaripaskum, 2023, d(2023, 4, 10), d(2023, 4, 10)},

		{Sumardagurinn, 2015, d(2015, 4, 23), d(2015, 4, 23)},
		{Sumardagurinn, 2016, d(2016, 4, 21), d(2016, 4, 21)},
		{Sumardagurinn, 2017, d(2017, 4, 20), d(2017, 4, 20)},
		{Sumardagurinn, 2018, d(2018, 4, 19), d(2018, 4, 19)},
		{Sumardagurinn, 2019, d(2019, 4, 25), d(2019, 4, 25)},
		{Sumardagurinn, 2020, d(2020, 4, 23), d(2020, 4, 23)},
		{Sumardagurinn, 2021, d(2021, 4, 22), d(2021, 4, 22)},
		{Sumardagurinn, 2022, d(2022, 4, 21), d(2022, 4, 21)},
		{Sumardagurinn, 2023, d(2023, 4, 20), d(2023, 4, 20)},

		{Verkalydsdagurinn, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{Verkalydsdagurinn, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{Verkalydsdagurinn, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{Verkalydsdagurinn, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{Verkalydsdagurinn, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{Verkalydsdagurinn, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{Verkalydsdagurinn, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{Verkalydsdagurinn, 2022, d(2022, 5, 1), d(2022, 5, 1)},
		{Verkalydsdagurinn, 2023, d(2023, 5, 1), d(2023, 5, 1)},

		{Uppstigningardagur, 2015, d(2015, 5, 14), d(2015, 5, 14)},
		{Uppstigningardagur, 2016, d(2016, 5, 5), d(2016, 5, 5)},
		{Uppstigningardagur, 2017, d(2017, 5, 25), d(2017, 5, 25)},
		{Uppstigningardagur, 2018, d(2018, 5, 10), d(2018, 5, 10)},
		{Uppstigningardagur, 2019, d(2019, 5, 30), d(2019, 5, 30)},
		{Uppstigningardagur, 2020, d(2020, 5, 21), d(2020, 5, 21)},
		{Uppstigningardagur, 2021, d(2021, 5, 13), d(2021, 5, 13)},
		{Uppstigningardagur, 2022, d(2022, 5, 26), d(2022, 5, 26)},
		{Uppstigningardagur, 2023, d(2023, 5, 18), d(2023, 5, 18)},

		{Annarihvit, 2015, d(2015, 5, 25), d(2015, 5, 25)},
		{Annarihvit, 2016, d(2016, 5, 16), d(2016, 5, 16)},
		{Annarihvit, 2017, d(2017, 6, 5), d(2017, 6, 5)},
		{Annarihvit, 2018, d(2018, 5, 21), d(2018, 5, 21)},
		{Annarihvit, 2019, d(2019, 6, 10), d(2019, 6, 10)},
		{Annarihvit, 2020, d(2020, 6, 1), d(2020, 6, 1)},
		{Annarihvit, 2021, d(2021, 5, 24), d(2021, 5, 24)},
		{Annarihvit, 2022, d(2022, 6, 6), d(2022, 6, 6)},
		{Annarihvit, 2023, d(2023, 5, 29), d(2023, 5, 29)},

		{Thjodhatid, 2015, d(2015, 6, 17), d(2015, 6, 17)},
		{Thjodhatid, 2016, d(2016, 6, 17), d(2016, 6, 17)},
		{Thjodhatid, 2017, d(2017, 6, 17), d(2017, 6, 17)},
		{Thjodhatid, 2018, d(2018, 6, 17), d(2018, 6, 17)},
		{Thjodhatid, 2019, d(2019, 6, 17), d(2019, 6, 17)},
		{Thjodhatid, 2020, d(2020, 6, 17), d(2020, 6, 17)},
		{Thjodhatid, 2021, d(2021, 6, 17), d(2021, 6, 17)},
		{Thjodhatid, 2022, d(2022, 6, 17), d(2022, 6, 17)},
		{Thjodhatid, 2023, d(2023, 6, 17), d(2023, 6, 17)},

		{Verslunarmannahelgi, 2015, d(2015, 8, 3), d(2015, 8, 3)},
		{Verslunarmannahelgi, 2016, d(2016, 8, 1), d(2016, 8, 1)},
		{Verslunarmannahelgi, 2017, d(2017, 8, 7), d(2017, 8, 7)},
		{Verslunarmannahelgi, 2018, d(2018, 8, 6), d(2018, 8, 6)},
		{Verslunarmannahelgi, 2019, d(2019, 8, 5), d(2019, 8, 5)},
		{Verslunarmannahelgi, 2020, d(2020, 8, 3), d(2020, 8, 3)},
		{Verslunarmannahelgi, 2021, d(2021, 8, 2), d(2021, 8, 2)},
		{Verslunarmannahelgi, 2022, d(2022, 8, 1), d(2022, 8, 1)},
		{Verslunarmannahelgi, 2023, d(2023, 8, 7), d(2023, 8, 7)},

		{Adfangadagur, 2015, d(2015, 12, 24), d(2015, 12, 24)},
		{Adfangadagur, 2016, d(2016, 12, 24), d(2016, 12, 24)},
		{Adfangadagur, 2017, d(2017, 12, 24), d(2017, 12, 24)},
		{Adfangadagur, 2018, d(2018, 12, 24), d(2018, 12, 24)},
		{Adfangadagur, 2019, d(2019, 12, 24), d(2019, 12, 24)},
		{Adfangadagur, 2020, d(2020, 12, 24), d(2020, 12, 24)},
		{Adfangadagur, 2021, d(2021, 12, 24), d(2021, 12, 24)},
		{Adfangadagur, 2022, d(2022, 12, 24), d(2022, 12, 24)},
		{Adfangadagur, 2023, d(2023, 12, 24), d(2023, 12, 24)},

		{Joladagur, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Joladagur, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Joladagur, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Joladagur, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Joladagur, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Joladagur, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Joladagur, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Joladagur, 2022, d(2022, 12, 25), d(2022, 12, 25)},
		{Joladagur, 2023, d(2023, 12, 25), d(2023, 12, 25)},

		{Annarijolum, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{Annarijolum, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{Annarijolum, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{Annarijolum, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{Annarijolum, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{Annarijolum, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{Annarijolum, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{Annarijolum, 2022, d(2022, 12, 26), d(2022, 12, 26)},
		{Annarijolum, 2023, d(2023, 12, 26), d(2023, 12, 26)},

		{Gamarsdagur, 2015, d(2015, 12, 31), d(2015, 12, 31)},
		{Gamarsdagur, 2016, d(2016, 12, 31), d(2016, 12, 31)},
		{Gamarsdagur, 2017, d(2017, 12, 31), d(2017, 12, 31)},
		{Gamarsdagur, 2018, d(2018, 12, 31), d(2018, 12, 31)},
		{Gamarsdagur, 2019, d(2019, 12, 31), d(2019, 12, 31)},
		{Gamarsdagur, 2020, d(2020, 12, 31), d(2020, 12, 31)},
		{Gamarsdagur, 2021, d(2021, 12, 31), d(2021, 12, 31)},
		{Gamarsdagur, 2022, d(2022, 12, 31), d(2022, 12, 31)},
		{Gamarsdagur, 2023, d(2023, 12, 31), d(2023, 12, 31)},
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
