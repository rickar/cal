// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package ro

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
		{AnulNou, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{AnulNou, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{AnulNou, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{AnulNou, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{AnulNou, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{AnulNou, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{AnulNou, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{AnulNou, 2022, d(2022, 1, 1), d(2022, 1, 1)},
		{AnulNou, 2022, d(2022, 1, 1), d(2022, 1, 1)},
		{AnulNou, 2023, d(2023, 1, 1), d(2023, 1, 1)},
		{AnulNou, 2024, d(2024, 1, 1), d(2024, 1, 1)},

		{AnulNou2, 2015, d(2015, 1, 2), d(2015, 1, 2)},
		{AnulNou2, 2016, d(2016, 1, 2), d(2016, 1, 2)},
		{AnulNou2, 2017, d(2017, 1, 2), d(2017, 1, 2)},
		{AnulNou2, 2018, d(2018, 1, 2), d(2018, 1, 2)},
		{AnulNou2, 2019, d(2019, 1, 2), d(2019, 1, 2)},
		{AnulNou2, 2020, d(2020, 1, 2), d(2020, 1, 2)},
		{AnulNou2, 2021, d(2021, 1, 2), d(2021, 1, 2)},
		{AnulNou2, 2022, d(2022, 1, 2), d(2022, 1, 2)},
		{AnulNou2, 2022, d(2022, 1, 2), d(2022, 1, 2)},
		{AnulNou2, 2023, d(2023, 1, 2), d(2023, 1, 2)},
		{AnulNou2, 2024, d(2024, 1, 2), d(2024, 1, 2)},

		{Boboteaza, 2024, d(2024, 1, 6), d(2024, 1, 6)},

		{SfantulIon, 2024, d(2024, 1, 7), d(2024, 1, 7)},

		{ZiuaUniriiPrincipatelorRomane, 2015, d(2015, 1, 24), d(2015, 1, 24)},
		{ZiuaUniriiPrincipatelorRomane, 2016, d(2016, 1, 24), d(2016, 1, 24)},
		{ZiuaUniriiPrincipatelorRomane, 2017, d(2017, 1, 24), d(2017, 1, 24)},
		{ZiuaUniriiPrincipatelorRomane, 2018, d(2018, 1, 24), d(2018, 1, 24)},
		{ZiuaUniriiPrincipatelorRomane, 2019, d(2019, 1, 24), d(2019, 1, 24)},
		{ZiuaUniriiPrincipatelorRomane, 2020, d(2020, 1, 24), d(2020, 1, 24)},
		{ZiuaUniriiPrincipatelorRomane, 2021, d(2021, 1, 24), d(2021, 1, 24)},
		{ZiuaUniriiPrincipatelorRomane, 2022, d(2022, 1, 24), d(2022, 1, 24)},
		{ZiuaUniriiPrincipatelorRomane, 2022, d(2022, 1, 24), d(2022, 1, 24)},
		{ZiuaUniriiPrincipatelorRomane, 2023, d(2023, 1, 24), d(2023, 1, 24)},
		{ZiuaUniriiPrincipatelorRomane, 2024, d(2024, 1, 24), d(2024, 1, 24)},

		{VinereaMare, 2015, time.Time{}, time.Time{}},
		{VinereaMare, 2016, time.Time{}, time.Time{}},
		{VinereaMare, 2017, time.Time{}, time.Time{}},
		{VinereaMare, 2018, d(2018, 4, 6), d(2018, 4, 6)},
		{VinereaMare, 2019, d(2019, 4, 26), d(2019, 4, 26)},
		{VinereaMare, 2020, d(2020, 4, 17), d(2020, 4, 17)},
		{VinereaMare, 2021, d(2021, 4, 30), d(2021, 4, 30)},
		{VinereaMare, 2022, d(2022, 4, 22), d(2022, 4, 22)},
		{VinereaMare, 2023, d(2023, 4, 14), d(2023, 4, 14)},
		{VinereaMare, 2024, d(2024, 5, 3), d(2024, 5, 3)},

		{Pastele, 2015, d(2015, 4, 12), d(2015, 4, 12)},
		{Pastele, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{Pastele, 2017, d(2017, 4, 16), d(2017, 4, 16)},
		{Pastele, 2018, d(2018, 4, 8), d(2018, 4, 8)},
		{Pastele, 2019, d(2019, 4, 28), d(2019, 4, 28)},
		{Pastele, 2020, d(2020, 4, 19), d(2020, 4, 19)},
		{Pastele, 2021, d(2021, 5, 2), d(2021, 5, 2)},
		{Pastele, 2022, d(2022, 4, 24), d(2022, 4, 24)},
		{Pastele, 2023, d(2023, 4, 16), d(2023, 4, 16)},
		{Pastele, 2024, d(2024, 5, 5), d(2024, 5, 5)},

		{ADouaZiDePaste, 2015, d(2015, 4, 13), d(2015, 4, 13)},
		{ADouaZiDePaste, 2016, d(2016, 5, 2), d(2016, 5, 2)},
		{ADouaZiDePaste, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{ADouaZiDePaste, 2018, d(2018, 4, 9), d(2018, 4, 9)},
		{ADouaZiDePaste, 2019, d(2019, 4, 29), d(2019, 4, 29)},
		{ADouaZiDePaste, 2020, d(2020, 4, 20), d(2020, 4, 20)},
		{ADouaZiDePaste, 2021, d(2021, 5, 3), d(2021, 5, 3)},
		{ADouaZiDePaste, 2022, d(2022, 4, 25), d(2022, 4, 25)},
		{ADouaZiDePaste, 2023, d(2023, 4, 17), d(2023, 4, 17)},
		{ADouaZiDePaste, 2024, d(2024, 5, 6), d(2024, 5, 6)},

		{ZiuaMuncii, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{ZiuaMuncii, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{ZiuaMuncii, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{ZiuaMuncii, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{ZiuaMuncii, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{ZiuaMuncii, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{ZiuaMuncii, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{ZiuaMuncii, 2022, d(2022, 5, 1), d(2022, 5, 1)},
		{ZiuaMuncii, 2023, d(2023, 5, 1), d(2023, 5, 1)},
		{ZiuaMuncii, 2024, d(2024, 5, 1), d(2024, 5, 1)},

		{ZiuaCopilului, 2015, time.Time{}, time.Time{}},
		{ZiuaCopilului, 2016, time.Time{}, time.Time{}},
		{ZiuaCopilului, 2017, d(2017, 6, 1), d(2017, 6, 1)},
		{ZiuaCopilului, 2018, d(2018, 6, 1), d(2018, 6, 1)},
		{ZiuaCopilului, 2019, d(2019, 6, 1), d(2019, 6, 1)},
		{ZiuaCopilului, 2020, d(2020, 6, 1), d(2020, 6, 1)},
		{ZiuaCopilului, 2021, d(2021, 6, 1), d(2021, 6, 1)},
		{ZiuaCopilului, 2022, d(2022, 6, 1), d(2022, 6, 1)},
		{ZiuaCopilului, 2023, d(2023, 6, 1), d(2023, 6, 1)},
		{ZiuaCopilului, 2024, d(2024, 6, 1), d(2024, 6, 1)},

		{Rusalii, 2015, d(2015, 5, 31), d(2015, 5, 31)},
		{Rusalii, 2016, d(2016, 6, 19), d(2016, 6, 19)},
		{Rusalii, 2017, d(2017, 6, 4), d(2017, 6, 4)},
		{Rusalii, 2018, d(2018, 5, 27), d(2018, 5, 27)},
		{Rusalii, 2019, d(2019, 6, 16), d(2019, 6, 16)},
		{Rusalii, 2020, d(2020, 6, 7), d(2020, 6, 7)},
		{Rusalii, 2021, d(2021, 6, 20), d(2021, 6, 20)},
		{Rusalii, 2022, d(2022, 6, 12), d(2022, 6, 12)},
		{Rusalii, 2023, d(2023, 6, 4), d(2023, 6, 4)},
		{Rusalii, 2024, d(2024, 6, 23), d(2024, 6, 23)},

		{LuniDupaRusalii, 2015, d(2015, 6, 1), d(2015, 6, 1)},
		{LuniDupaRusalii, 2016, d(2016, 6, 20), d(2016, 6, 20)},
		{LuniDupaRusalii, 2017, d(2017, 6, 5), d(2017, 6, 5)},
		{LuniDupaRusalii, 2018, d(2018, 5, 28), d(2018, 5, 28)},
		{LuniDupaRusalii, 2019, d(2019, 6, 17), d(2019, 6, 17)},
		{LuniDupaRusalii, 2020, d(2020, 6, 8), d(2020, 6, 8)},
		{LuniDupaRusalii, 2021, d(2021, 6, 21), d(2021, 6, 21)},
		{LuniDupaRusalii, 2022, d(2022, 6, 13), d(2022, 6, 13)},
		{LuniDupaRusalii, 2023, d(2023, 6, 5), d(2023, 6, 5)},
		{LuniDupaRusalii, 2024, d(2024, 6, 24), d(2024, 6, 24)},

		{AdormireaMaiciiDomnului, 2015, d(2015, 8, 15), d(2015, 8, 15)},
		{AdormireaMaiciiDomnului, 2016, d(2016, 8, 15), d(2016, 8, 15)},
		{AdormireaMaiciiDomnului, 2017, d(2017, 8, 15), d(2017, 8, 15)},
		{AdormireaMaiciiDomnului, 2018, d(2018, 8, 15), d(2018, 8, 15)},
		{AdormireaMaiciiDomnului, 2019, d(2019, 8, 15), d(2019, 8, 15)},
		{AdormireaMaiciiDomnului, 2020, d(2020, 8, 15), d(2020, 8, 15)},
		{AdormireaMaiciiDomnului, 2021, d(2021, 8, 15), d(2021, 8, 15)},
		{AdormireaMaiciiDomnului, 2022, d(2022, 8, 15), d(2022, 8, 15)},
		{AdormireaMaiciiDomnului, 2023, d(2023, 8, 15), d(2023, 8, 15)},
		{AdormireaMaiciiDomnului, 2024, d(2024, 8, 15), d(2024, 8, 15)},

		{SfantulAndrei, 2015, d(2015, 11, 30), d(2015, 11, 30)},
		{SfantulAndrei, 2016, d(2016, 11, 30), d(2016, 11, 30)},
		{SfantulAndrei, 2017, d(2017, 11, 30), d(2017, 11, 30)},
		{SfantulAndrei, 2018, d(2018, 11, 30), d(2018, 11, 30)},
		{SfantulAndrei, 2019, d(2019, 11, 30), d(2019, 11, 30)},
		{SfantulAndrei, 2020, d(2020, 11, 30), d(2020, 11, 30)},
		{SfantulAndrei, 2021, d(2021, 11, 30), d(2021, 11, 30)},
		{SfantulAndrei, 2022, d(2022, 11, 30), d(2022, 11, 30)},
		{SfantulAndrei, 2023, d(2023, 11, 30), d(2023, 11, 30)},
		{SfantulAndrei, 2024, d(2024, 11, 30), d(2024, 11, 30)},

		{ZiuaNationala, 2015, d(2015, 12, 1), d(2015, 12, 1)},
		{ZiuaNationala, 2016, d(2016, 12, 1), d(2016, 12, 1)},
		{ZiuaNationala, 2017, d(2017, 12, 1), d(2017, 12, 1)},
		{ZiuaNationala, 2018, d(2018, 12, 1), d(2018, 12, 1)},
		{ZiuaNationala, 2019, d(2019, 12, 1), d(2019, 12, 1)},
		{ZiuaNationala, 2020, d(2020, 12, 1), d(2020, 12, 1)},
		{ZiuaNationala, 2021, d(2021, 12, 1), d(2021, 12, 1)},
		{ZiuaNationala, 2022, d(2022, 12, 1), d(2022, 12, 1)},
		{ZiuaNationala, 2023, d(2023, 12, 1), d(2023, 12, 1)},
		{ZiuaNationala, 2024, d(2024, 12, 1), d(2024, 12, 1)},

		{Craciunul, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Craciunul, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Craciunul, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Craciunul, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Craciunul, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Craciunul, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Craciunul, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Craciunul, 2022, d(2022, 12, 25), d(2022, 12, 25)},
		{Craciunul, 2023, d(2023, 12, 25), d(2023, 12, 25)},
		{Craciunul, 2024, d(2024, 12, 25), d(2024, 12, 25)},

		{Craciunul2, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{Craciunul2, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{Craciunul2, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{Craciunul2, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{Craciunul2, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{Craciunul2, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{Craciunul2, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{Craciunul2, 2022, d(2022, 12, 26), d(2022, 12, 26)},
		{Craciunul2, 2023, d(2023, 12, 26), d(2023, 12, 26)},
		{Craciunul2, 2024, d(2024, 12, 26), d(2024, 12, 26)},
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
