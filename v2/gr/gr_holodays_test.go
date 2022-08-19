package gr

import (
	"fmt"
	"testing"
	"time"

	"github.com/rickar/cal/v2"
)

func d(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, cal.DefaultLoc)
}

func TestHolidays(t *testing.T) {
	TestProtochroxia(t)
	TestTheophania(t)
	TestKatharaDeftera(t)
	TestIkostiPemptiMartiou(t)
	TestMegaliParaskevi(t)
	TestDefteraPascha(t)
	TestErgatikiProtomagia(t)
	TestAgiouPrevmatos(t)
	TestKimisiTisTheotokou(t)
	TestImeraTouOchi(t)
	TestChristougenna(t)
	TestSinaxisYperagiasTheotokou(t)
}

func TestProtochroxia(t *testing.T) {
	tests := []struct {
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{Theophania, 2015, d(2015, 1, 6), d(2015, 1, 6)},
		{Theophania, 2016, d(2016, 1, 6), d(2016, 1, 6)},
		{Theophania, 2017, d(2017, 1, 6), d(2017, 1, 6)},
		{Theophania, 2018, d(2018, 1, 6), d(2018, 1, 6)},
		{Theophania, 2019, d(2019, 1, 6), d(2019, 1, 6)},
		{Theophania, 2020, d(2020, 1, 6), d(2020, 1, 6)},
		{Theophania, 2021, d(2021, 1, 6), d(2021, 1, 6)},
		{Theophania, 2022, d(2022, 1, 6), d(2022, 1, 6)},
		{Theophania, 2023, d(2023, 1, 6), d(2023, 1, 6)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		fmt.Printf("Actual: %v \nObserved: %v\n", gotAct, gotObs)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}

func TestTheophania(t *testing.T) {
	tests := []struct {
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{Protoxronia, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{Protoxronia, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{Protoxronia, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{Protoxronia, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{Protoxronia, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{Protoxronia, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{Protoxronia, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{Protoxronia, 2022, d(2022, 1, 1), d(2022, 1, 1)},
		{Protoxronia, 2023, d(2023, 1, 1), d(2023, 1, 1)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		fmt.Printf("Actual: %v \nObserved: %v\n", gotAct, gotObs)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}

func TestKatharaDeftera(t *testing.T) {
	tests := []struct {
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{KatharaDeftera, 2015, d(2015, 2, 23), d(2015, 2, 23)},
		{KatharaDeftera, 2016, d(2016, 3, 14), d(2016, 3, 14)},
		{KatharaDeftera, 2017, d(2017, 2, 27), d(2017, 2, 27)},
		{KatharaDeftera, 2018, d(2018, 2, 19), d(2018, 2, 19)},
		{KatharaDeftera, 2019, d(2019, 3, 11), d(2019, 3, 11)},
		{KatharaDeftera, 2019, d(2019, 3, 11), d(2019, 3, 11)},
		{KatharaDeftera, 2020, d(2020, 3, 2), d(2020, 3, 2)},
		{KatharaDeftera, 2021, d(2021, 3, 15), d(2021, 3, 15)},
		{KatharaDeftera, 2022, d(2022, 3, 7), d(2022, 3, 7)},
		{KatharaDeftera, 2023, d(2023, 2, 27), d(2023, 2, 27)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		fmt.Printf("Actual: %v \nObserved: %v\n", gotAct, gotObs)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}

func TestIkostiPemptiMartiou(t *testing.T) {
	tests := []struct {
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{IkostiPemptiMartiou, 2015, d(2015, 3, 25), d(2015, 3, 25)},
		{IkostiPemptiMartiou, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{IkostiPemptiMartiou, 2017, d(2017, 3, 25), d(2017, 3, 25)},
		{IkostiPemptiMartiou, 2018, d(2018, 3, 25), d(2018, 3, 25)},
		{IkostiPemptiMartiou, 2019, d(2019, 3, 25), d(2019, 3, 25)},
		{IkostiPemptiMartiou, 2020, d(2020, 3, 25), d(2020, 3, 25)},
		{IkostiPemptiMartiou, 2021, d(2021, 3, 25), d(2021, 3, 25)},
		{IkostiPemptiMartiou, 2022, d(2022, 3, 25), d(2022, 3, 25)},
		{IkostiPemptiMartiou, 2023, d(2023, 3, 25), d(2023, 3, 25)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		fmt.Printf("Actual: %v \nObserved: %v\n", gotAct, gotObs)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}

func TestMegaliParaskevi(t *testing.T) {
	tests := []struct {
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{MegaliParaskevi, 2015, d(2015, 4, 10), d(2015, 4, 10)},
		{MegaliParaskevi, 2016, d(2016, 4, 29), d(2016, 4, 29)},
		{MegaliParaskevi, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{MegaliParaskevi, 2018, d(2018, 4, 6), d(2018, 4, 6)},
		{MegaliParaskevi, 2019, d(2019, 4, 26), d(2019, 4, 26)},
		{MegaliParaskevi, 2020, d(2020, 4, 17), d(2020, 4, 17)},
		{MegaliParaskevi, 2021, d(2021, 4, 30), d(2021, 4, 30)},
		{MegaliParaskevi, 2022, d(2022, 4, 22), d(2022, 4, 22)},
		{MegaliParaskevi, 2023, d(2023, 4, 14), d(2023, 4, 14)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		fmt.Printf("Actual: %v \nObserved: %v\n", gotAct, gotObs)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}

func TestDefteraPascha(t *testing.T) {
	tests := []struct {
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{DefteraPascha, 2015, d(2015, 4, 13), d(2015, 4, 13)},
		{DefteraPascha, 2016, d(2016, 5, 2), d(2016, 5, 2)},
		{DefteraPascha, 2017, d(2017, 4, 17), d(2017, 4, 17)},
		{DefteraPascha, 2018, d(2018, 4, 9), d(2018, 4, 9)},
		{DefteraPascha, 2019, d(2019, 4, 29), d(2019, 4, 29)},
		{DefteraPascha, 2020, d(2020, 4, 20), d(2020, 4, 20)},
		{DefteraPascha, 2021, d(2021, 5, 3), d(2021, 5, 3)},
		{DefteraPascha, 2022, d(2022, 4, 25), d(2022, 4, 25)},
		{DefteraPascha, 2023, d(2023, 4, 17), d(2023, 4, 17)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		fmt.Printf("Actual: %v \nObserved: %v\n", gotAct, gotObs)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}

func TestErgatikiProtomagia(t *testing.T) {
	tests := []struct {
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{ErgatikiProtomagia, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{ErgatikiProtomagia, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{ErgatikiProtomagia, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{ErgatikiProtomagia, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{ErgatikiProtomagia, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{ErgatikiProtomagia, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{ErgatikiProtomagia, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{ErgatikiProtomagia, 2022, d(2022, 5, 1), d(2022, 5, 1)},
		{ErgatikiProtomagia, 2023, d(2023, 5, 1), d(2023, 5, 1)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		fmt.Printf("Actual: %v \nObserved: %v\n", gotAct, gotObs)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}

func TestAgiouPrevmatos(t *testing.T) {
	tests := []struct {
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{AgiouPrevmatos, 2015, d(2015, 6, 1), d(2015, 6, 1)},
		{AgiouPrevmatos, 2016, d(2016, 6, 20), d(2016, 6, 20)},
		{AgiouPrevmatos, 2017, d(2017, 6, 5), d(2017, 6, 5)},
		{AgiouPrevmatos, 2018, d(2018, 5, 28), d(2018, 5, 28)},
		{AgiouPrevmatos, 2019, d(2019, 6, 17), d(2019, 6, 17)},
		{AgiouPrevmatos, 2020, d(2020, 6, 8), d(2020, 6, 8)},
		{AgiouPrevmatos, 2021, d(2021, 6, 21), d(2021, 6, 21)},
		{AgiouPrevmatos, 2022, d(2022, 6, 13), d(2022, 6, 13)},
		{AgiouPrevmatos, 2023, d(2023, 6, 5), d(2023, 6, 5)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		fmt.Printf("Actual: %v \nObserved: %v\n", gotAct, gotObs)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}

func TestKimisiTisTheotokou(t *testing.T) {
	tests := []struct {
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{KimisiTisTheotokou, 2015, d(2015, 8, 15), d(2015, 8, 15)},
		{KimisiTisTheotokou, 2016, d(2016, 8, 15), d(2016, 8, 15)},
		{KimisiTisTheotokou, 2017, d(2017, 8, 15), d(2017, 8, 15)},
		{KimisiTisTheotokou, 2018, d(2018, 8, 15), d(2018, 8, 15)},
		{KimisiTisTheotokou, 2019, d(2019, 8, 15), d(2019, 8, 15)},
		{KimisiTisTheotokou, 2020, d(2020, 8, 15), d(2020, 8, 15)},
		{KimisiTisTheotokou, 2021, d(2021, 8, 15), d(2021, 8, 15)},
		{KimisiTisTheotokou, 2022, d(2022, 8, 15), d(2022, 8, 15)},
		{KimisiTisTheotokou, 2023, d(2023, 8, 15), d(2023, 8, 15)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		fmt.Printf("Actual: %v \nObserved: %v\n", gotAct, gotObs)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}

func TestImeraTouOchi(t *testing.T) {
	tests := []struct {
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{ImeraTouOchi, 2015, d(2015, 10, 28), d(2015, 10, 28)},
		{ImeraTouOchi, 2016, d(2016, 10, 28), d(2016, 10, 28)},
		{ImeraTouOchi, 2017, d(2017, 10, 28), d(2017, 10, 28)},
		{ImeraTouOchi, 2018, d(2018, 10, 28), d(2018, 10, 28)},
		{ImeraTouOchi, 2019, d(2019, 10, 28), d(2019, 10, 28)},
		{ImeraTouOchi, 2020, d(2020, 10, 28), d(2020, 10, 28)},
		{ImeraTouOchi, 2021, d(2021, 10, 28), d(2021, 10, 28)},
		{ImeraTouOchi, 2022, d(2022, 10, 28), d(2022, 10, 28)},
		{ImeraTouOchi, 2023, d(2023, 10, 28), d(2023, 10, 28)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		fmt.Printf("Actual: %v \nObserved: %v\n", gotAct, gotObs)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}

func TestChristougenna(t *testing.T) {
	tests := []struct {
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{Christougenna, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Christougenna, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Christougenna, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Christougenna, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Christougenna, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Christougenna, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Christougenna, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Christougenna, 2022, d(2022, 12, 25), d(2022, 12, 25)},
		{Christougenna, 2023, d(2023, 12, 25), d(2023, 12, 25)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		fmt.Printf("Actual: %v \nObserved: %v\n", gotAct, gotObs)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}

func TestSinaxisYperagiasTheotokou(t *testing.T) {
	tests := []struct {
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{SinaxisYperagiasTheotokou, 2015, d(2015, 12, 26), d(2015, 12, 26)},
		{SinaxisYperagiasTheotokou, 2016, d(2016, 12, 26), d(2016, 12, 26)},
		{SinaxisYperagiasTheotokou, 2017, d(2017, 12, 26), d(2017, 12, 26)},
		{SinaxisYperagiasTheotokou, 2018, d(2018, 12, 26), d(2018, 12, 26)},
		{SinaxisYperagiasTheotokou, 2019, d(2019, 12, 26), d(2019, 12, 26)},
		{SinaxisYperagiasTheotokou, 2020, d(2020, 12, 26), d(2020, 12, 26)},
		{SinaxisYperagiasTheotokou, 2021, d(2021, 12, 26), d(2021, 12, 26)},
		{SinaxisYperagiasTheotokou, 2022, d(2022, 12, 26), d(2022, 12, 26)},
		{SinaxisYperagiasTheotokou, 2023, d(2023, 12, 26), d(2023, 12, 26)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		fmt.Printf("Actual: %v \nObserved: %v\n", gotAct, gotObs)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}
