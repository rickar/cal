package hr

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
		{NovaGodina, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{NovaGodina, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{NovaGodina, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{NovaGodina, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{NovaGodina, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{NovaGodina, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{NovaGodina, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{NovaGodina, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{SvetaTriKralja, 2015, d(2015, 1, 6), d(2015, 1, 6)},
		{SvetaTriKralja, 2016, d(2016, 1, 6), d(2016, 1, 6)},
		{SvetaTriKralja, 2017, d(2017, 1, 6), d(2017, 1, 6)},
		{SvetaTriKralja, 2018, d(2018, 1, 6), d(2018, 1, 6)},
		{SvetaTriKralja, 2019, d(2019, 1, 6), d(2019, 1, 6)},
		{SvetaTriKralja, 2020, d(2020, 1, 6), d(2020, 1, 6)},
		{SvetaTriKralja, 2021, d(2021, 1, 6), d(2021, 1, 6)},
		{SvetaTriKralja, 2022, d(2022, 1, 6), d(2022, 1, 6)},

		{Uskrs, 2015, d(2015, 4, 5), d(2015, 4, 5)},
		{Uskrs, 2016, d(2016, 3, 27), d(2016, 3, 27)},
		{Uskrs, 2017, d(2017, 4, 16), d(2017, 4, 16)},
		{Uskrs, 2018, d(2018, 4, 1), d(2018, 4, 1)},
		{Uskrs, 2019, d(2019, 4, 21), d(2019, 4, 21)},
		{Uskrs, 2020, d(2020, 4, 12), d(2020, 4, 12)},
		{Uskrs, 2021, d(2021, 4, 4), d(2021, 4, 4)},
		{Uskrs, 2022, d(2022, 4, 17), d(2022, 4, 17)},
		{Uskrs, 2023, d(2023, 4, 9), d(2023, 4, 9)},

		{UskrsnjiPonedjeljak, 2015, d(2015, 4, 6), d(2015, 4, 6)},
		{UskrsnjiPonedjeljak, 2016, d(2016, 3, 28), d(2016, 3, 28)},
		{UskrsnjiPonedjeljak, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{UskrsnjiPonedjeljak, 2018, d(2018, 4, 2), d(2018, 4, 2)},
		{UskrsnjiPonedjeljak, 2019, d(2019, 4, 22), d(2019, 4, 22)},
		{UskrsnjiPonedjeljak, 2020, d(2020, 4, 13), d(2020, 4, 13)},
		{UskrsnjiPonedjeljak, 2021, d(2021, 4, 5), d(2021, 4, 5)},
		{UskrsnjiPonedjeljak, 2022, d(2022, 4, 18), d(2022, 4, 18)},
		{UskrsnjiPonedjeljak, 2023, d(2023, 4, 10), d(2023, 4, 10)},

		{PraznikRada, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{PraznikRada, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{PraznikRada, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{PraznikRada, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{PraznikRada, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{PraznikRada, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{PraznikRada, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{PraznikRada, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{DanDrzavnosti, 2015, d(2015, 5, 30), d(2015, 5, 30)},
		{DanDrzavnosti, 2016, d(2016, 5, 30), d(2016, 5, 30)},
		{DanDrzavnosti, 2017, d(2017, 5, 30), d(2017, 5, 30)},
		{DanDrzavnosti, 2018, d(2018, 5, 30), d(2018, 5, 30)},
		{DanDrzavnosti, 2019, d(2019, 5, 30), d(2019, 5, 30)},
		{DanDrzavnosti, 2020, d(2020, 5, 30), d(2020, 5, 30)},
		{DanDrzavnosti, 2021, d(2021, 5, 30), d(2021, 5, 30)},
		{DanDrzavnosti, 2022, d(2022, 5, 30), d(2022, 5, 30)},

		{Tijelovo, 2015, d(2015, 6, 4), d(2015, 6, 4)},
		{Tijelovo, 2016, d(2016, 5, 26), d(2016, 5, 26)},
		{Tijelovo, 2017, d(2017, 6, 15), d(2017, 6, 15)},
		{Tijelovo, 2018, d(2018, 5, 31), d(2018, 5, 31)},
		{Tijelovo, 2019, d(2019, 6, 20), d(2019, 6, 20)},
		{Tijelovo, 2020, d(2020, 6, 11), d(2020, 6, 11)},
		{Tijelovo, 2021, d(2021, 6, 3), d(2021, 6, 3)},
		{Tijelovo, 2022, d(2022, 6, 16), d(2022, 6, 16)},
		{Tijelovo, 2023, d(2023, 6, 8), d(2023, 6, 8)},

		{DanAntifasistickeBorbe, 2015, d(2015, 6, 22), d(2015, 6, 22)},
		{DanAntifasistickeBorbe, 2016, d(2016, 6, 22), d(2016, 6, 22)},
		{DanAntifasistickeBorbe, 2017, d(2017, 6, 22), d(2017, 6, 22)},
		{DanAntifasistickeBorbe, 2018, d(2018, 6, 22), d(2018, 6, 22)},
		{DanAntifasistickeBorbe, 2019, d(2019, 6, 22), d(2019, 6, 22)},
		{DanAntifasistickeBorbe, 2020, d(2020, 6, 22), d(2020, 6, 22)},
		{DanAntifasistickeBorbe, 2021, d(2021, 6, 22), d(2021, 6, 22)},
		{DanAntifasistickeBorbe, 2022, d(2022, 6, 22), d(2022, 6, 22)},

		{DanPobjedeIDomovinskeZahvalnosti, 2015, d(2015, 8, 5), d(2015, 8, 5)},
		{DanPobjedeIDomovinskeZahvalnosti, 2016, d(2016, 8, 5), d(2016, 8, 5)},
		{DanPobjedeIDomovinskeZahvalnosti, 2017, d(2017, 8, 5), d(2017, 8, 5)},
		{DanPobjedeIDomovinskeZahvalnosti, 2018, d(2018, 8, 5), d(2018, 8, 5)},
		{DanPobjedeIDomovinskeZahvalnosti, 2019, d(2019, 8, 5), d(2019, 8, 5)},
		{DanPobjedeIDomovinskeZahvalnosti, 2020, d(2020, 8, 5), d(2020, 8, 5)},
		{DanPobjedeIDomovinskeZahvalnosti, 2021, d(2021, 8, 5), d(2021, 8, 5)},
		{DanPobjedeIDomovinskeZahvalnosti, 2022, d(2022, 8, 5), d(2022, 8, 5)},

		{VelikaGospa, 2015, d(2015, 8, 15), d(2015, 8, 15)},
		{VelikaGospa, 2016, d(2016, 8, 15), d(2016, 8, 15)},
		{VelikaGospa, 2017, d(2017, 8, 15), d(2017, 8, 15)},
		{VelikaGospa, 2018, d(2018, 8, 15), d(2018, 8, 15)},
		{VelikaGospa, 2019, d(2019, 8, 15), d(2019, 8, 15)},
		{VelikaGospa, 2020, d(2020, 8, 15), d(2020, 8, 15)},
		{VelikaGospa, 2021, d(2021, 8, 15), d(2021, 8, 15)},
		{VelikaGospa, 2022, d(2022, 8, 15), d(2022, 8, 15)},

		{DanSvihSvetih, 2015, d(2015, 11, 1), d(2015, 11, 1)},
		{DanSvihSvetih, 2016, d(2016, 11, 1), d(2016, 11, 1)},
		{DanSvihSvetih, 2017, d(2017, 11, 1), d(2017, 11, 1)},
		{DanSvihSvetih, 2018, d(2018, 11, 1), d(2018, 11, 1)},
		{DanSvihSvetih, 2019, d(2019, 11, 1), d(2019, 11, 1)},
		{DanSvihSvetih, 2020, d(2020, 11, 1), d(2020, 11, 1)},
		{DanSvihSvetih, 2021, d(2021, 11, 1), d(2021, 11, 1)},
		{DanSvihSvetih, 2022, d(2022, 11, 1), d(2022, 11, 1)},

		{DanSjecanjaNaZrtveDomovinskogRata, 2015, d(2015, 11, 18), d(2015, 11, 18)},
		{DanSjecanjaNaZrtveDomovinskogRata, 2016, d(2016, 11, 18), d(2016, 11, 18)},
		{DanSjecanjaNaZrtveDomovinskogRata, 2017, d(2017, 11, 18), d(2017, 11, 18)},
		{DanSjecanjaNaZrtveDomovinskogRata, 2018, d(2018, 11, 18), d(2018, 11, 18)},
		{DanSjecanjaNaZrtveDomovinskogRata, 2019, d(2019, 11, 18), d(2019, 11, 18)},
		{DanSjecanjaNaZrtveDomovinskogRata, 2020, d(2020, 11, 18), d(2020, 11, 18)},
		{DanSjecanjaNaZrtveDomovinskogRata, 2021, d(2021, 11, 18), d(2021, 11, 18)},
		{DanSjecanjaNaZrtveDomovinskogRata, 2022, d(2022, 11, 18), d(2022, 11, 18)},

		{Bozic, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Bozic, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Bozic, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Bozic, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Bozic, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Bozic, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Bozic, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Bozic, 2022, d(2022, 12, 25), d(2022, 12, 25)},

		{SvetiStjepan, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{SvetiStjepan, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{SvetiStjepan, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{SvetiStjepan, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{SvetiStjepan, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{SvetiStjepan, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{SvetiStjepan, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{SvetiStjepan, 2022, d(2022, 12, 26), d(2022, 12, 26)},
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
