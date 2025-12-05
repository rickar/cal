// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package kr

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

        {BeforeLunarNewYear, 2015, d(2015, 2, 18), d(2015, 2, 18)},
        {BeforeLunarNewYear, 2016, d(2016, 2, 7), d(2016, 2, 7)},
        {BeforeLunarNewYear, 2017, d(2017, 1, 27), d(2017, 1, 27)},
        {BeforeLunarNewYear, 2018, d(2018, 2, 15), d(2018, 2, 15)},
        {BeforeLunarNewYear, 2019, d(2019, 2, 4), d(2019, 2, 4)},
        {BeforeLunarNewYear, 2020, d(2020, 1, 24), d(2020, 1, 24)},
        {BeforeLunarNewYear, 2021, d(2021, 2, 11), d(2021, 2, 11)},
        {BeforeLunarNewYear, 2022, d(2022, 1, 31), d(2022, 1, 31)},

        {LunarNewYear, 2015, d(2015, 2, 19), d(2015, 2, 19)},
        {LunarNewYear, 2016, d(2016, 2, 8), d(2016, 2, 8)},
        {LunarNewYear, 2017, d(2017, 1, 28), d(2017, 1, 30)},
        {LunarNewYear, 2018, d(2018, 2, 16), d(2018, 2, 16)},
        {LunarNewYear, 2019, d(2019, 2, 5), d(2019, 2, 5)},
        {LunarNewYear, 2020, d(2020, 1, 25), d(2020, 1, 27)},
        {LunarNewYear, 2021, d(2021, 2, 12), d(2021, 2, 12)},
        {LunarNewYear, 2022, d(2022, 2, 1), d(2022, 2, 1)},

        {AfterLunarNewYear, 2015, d(2015, 2, 20), d(2015, 2, 20)},
        {AfterLunarNewYear, 2016, d(2016, 2, 9), d(2016, 2, 9)},
        {AfterLunarNewYear, 2017, d(2017, 1, 29), d(2017, 1, 29)},
        {AfterLunarNewYear, 2018, d(2018, 2, 17), d(2018, 2, 17)},
        {AfterLunarNewYear, 2019, d(2019, 2, 6), d(2019, 2, 6)},
        {AfterLunarNewYear, 2020, d(2020, 1, 26), d(2020, 1, 26)},
        {AfterLunarNewYear, 2021, d(2021, 2, 13), d(2021, 2, 13)},
        {AfterLunarNewYear, 2022, d(2022, 2, 2), d(2022, 2, 2)},

        {IndependenceMovementDay, 2015, d(2015, 3, 1), d(2015, 3, 1)},
        {IndependenceMovementDay, 2016, d(2016, 3, 1), d(2016, 3, 1)},
        {IndependenceMovementDay, 2017, d(2017, 3, 1), d(2017, 3, 1)},
        {IndependenceMovementDay, 2018, d(2018, 3, 1), d(2018, 3, 1)},
        {IndependenceMovementDay, 2019, d(2019, 3, 1), d(2019, 3, 1)},
        {IndependenceMovementDay, 2020, d(2020, 3, 1), d(2020, 3, 1)},
        {IndependenceMovementDay, 2021, d(2021, 3, 1), d(2021, 3, 1)},
        {IndependenceMovementDay, 2022, d(2022, 3, 1), d(2022, 3, 1)},

        {BuddhasDay, 2015, d(2015, 5, 25), d(2015, 5, 25)},
        {BuddhasDay, 2016, d(2016, 5, 14), d(2016, 5, 14)},
        {BuddhasDay, 2017, d(2017, 5, 3), d(2017, 5, 3)},
        {BuddhasDay, 2018, d(2018, 5, 22), d(2018, 5, 22)},
        {BuddhasDay, 2019, d(2019, 5, 12), d(2019, 5, 12)},
        {BuddhasDay, 2020, d(2020, 4, 30), d(2020, 4, 30)},
        {BuddhasDay, 2021, d(2021, 5, 19), d(2021, 5, 19)},
        {BuddhasDay, 2022, d(2022, 5, 8), d(2022, 5, 8)},

        {ChildrensDay, 2015, d(2015, 5, 5), d(2015, 5, 5)},
        {ChildrensDay, 2016, d(2016, 5, 5), d(2016, 5, 5)},
        {ChildrensDay, 2017, d(2017, 5, 5), d(2017, 5, 5)},
        {ChildrensDay, 2018, d(2018, 5, 5), d(2018, 5, 7)},
        {ChildrensDay, 2019, d(2019, 5, 5), d(2019, 5, 6)},
        {ChildrensDay, 2020, d(2020, 5, 5), d(2020, 5, 5)},
        {ChildrensDay, 2021, d(2021, 5, 5), d(2021, 5, 5)},
        {ChildrensDay, 2022, d(2022, 5, 5), d(2022, 5, 5)},

        {MemorialDay, 2015, d(2015, 6, 6), d(2015, 6, 6)},
        {MemorialDay, 2016, d(2016, 6, 6), d(2016, 6, 6)},
        {MemorialDay, 2017, d(2017, 6, 6), d(2017, 6, 6)},
        {MemorialDay, 2018, d(2018, 6, 6), d(2018, 6, 6)},
        {MemorialDay, 2019, d(2019, 6, 6), d(2019, 6, 6)},
        {MemorialDay, 2020, d(2020, 6, 6), d(2020, 6, 6)},
        {MemorialDay, 2021, d(2021, 6, 6), d(2021, 6, 6)},
        {MemorialDay, 2022, d(2022, 6, 6), d(2022, 6, 6)},

        {LiberationDay, 2015, d(2015, 8, 15), d(2015, 8, 15)},
        {LiberationDay, 2016, d(2016, 8, 15), d(2016, 8, 15)},
        {LiberationDay, 2017, d(2017, 8, 15), d(2017, 8, 15)},
        {LiberationDay, 2018, d(2018, 8, 15), d(2018, 8, 15)},
        {LiberationDay, 2019, d(2019, 8, 15), d(2019, 8, 15)},
        {LiberationDay, 2020, d(2020, 8, 15), d(2020, 8, 15)},
        {LiberationDay, 2021, d(2021, 8, 15), d(2021, 8, 15)},
        {LiberationDay, 2022, d(2022, 8, 15), d(2022, 8, 15)},

        {BeforeChuseok, 2016, d(2016, 9, 14), d(2016, 9, 14)},
        {BeforeChuseok, 2017, d(2017, 10, 3), d(2017, 10, 3)},
        {BeforeChuseok, 2018, d(2018, 9, 23), d(2018, 9, 23)},
        {BeforeChuseok, 2019, d(2019, 9, 12), d(2019, 9, 12)},
        {BeforeChuseok, 2020, d(2020, 9, 30), d(2020, 9, 30)},
        {BeforeChuseok, 2021, d(2021, 9, 20), d(2021, 9, 20)},
        {BeforeChuseok, 2022, d(2022, 9, 9), d(2022, 9, 9)},

        {Chuseok, 2016, d(2016, 9, 15), d(2016, 9, 15)},
        {Chuseok, 2017, d(2017, 10, 4), d(2017, 10, 4)},
        {Chuseok, 2018, d(2018, 9, 24), d(2018, 9, 24)},
        {Chuseok, 2019, d(2019, 9, 13), d(2019, 9, 13)},
        {Chuseok, 2020, d(2020, 10, 1), d(2020, 10, 1)},
        {Chuseok, 2021, d(2021, 9, 21), d(2021, 9, 21)},
        {Chuseok, 2022, d(2022, 9, 10), d(2022, 9, 12)},

        {AfterChuseok, 2016, d(2016, 9, 16), d(2016, 9, 16)},
        {AfterChuseok, 2017, d(2017, 10, 5), d(2017, 10, 5)},
        {AfterChuseok, 2018, d(2018, 9, 25), d(2018, 9, 25)},
        {AfterChuseok, 2019, d(2019, 9, 14), d(2019, 9, 14)},
        {AfterChuseok, 2020, d(2020, 10, 2), d(2020, 10, 2)},
        {AfterChuseok, 2021, d(2021, 9, 22), d(2021, 9, 22)},
        {AfterChuseok, 2022, d(2022, 9, 11), d(2022, 9, 11)},

        {FoundationDay, 2015, d(2015, 10, 3), d(2015, 10, 3)},
        {FoundationDay, 2016, d(2016, 10, 3), d(2016, 10, 3)},
        {FoundationDay, 2017, d(2017, 10, 3), d(2017, 10, 3)},
        {FoundationDay, 2018, d(2018, 10, 3), d(2018, 10, 3)},
        {FoundationDay, 2019, d(2019, 10, 3), d(2019, 10, 3)},
        {FoundationDay, 2020, d(2020, 10, 3), d(2020, 10, 3)},
        {FoundationDay, 2021, d(2021, 10, 3), d(2021, 10, 3)},
        {FoundationDay, 2022, d(2022, 10, 3), d(2022, 10, 3)},

        {HangeulDay, 2015, d(2015, 10, 9), d(2015, 10, 9)},
        {HangeulDay, 2016, d(2016, 10, 9), d(2016, 10, 9)},
        {HangeulDay, 2017, d(2017, 10, 9), d(2017, 10, 9)},
        {HangeulDay, 2018, d(2018, 10, 9), d(2018, 10, 9)},
        {HangeulDay, 2019, d(2019, 10, 9), d(2019, 10, 9)},
        {HangeulDay, 2020, d(2020, 10, 9), d(2020, 10, 9)},
        {HangeulDay, 2021, d(2021, 10, 9), d(2021, 10, 9)},
        {HangeulDay, 2022, d(2022, 10, 9), d(2022, 10, 9)},

        {ChristmasDay, 2015, d(2015, 12, 25), d(2015, 12, 25)},
        {ChristmasDay, 2016, d(2016, 12, 25), d(2016, 12, 25)},
        {ChristmasDay, 2017, d(2017, 12, 25), d(2017, 12, 25)},
        {ChristmasDay, 2018, d(2018, 12, 25), d(2018, 12, 25)},
        {ChristmasDay, 2019, d(2019, 12, 25), d(2019, 12, 25)},
        {ChristmasDay, 2020, d(2020, 12, 25), d(2020, 12, 25)},
        {ChristmasDay, 2021, d(2021, 12, 25), d(2021, 12, 25)},
        {ChristmasDay, 2022, d(2022, 12, 25), d(2022, 12, 25)},
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

func Example_lunar() {
    // http://tool-box.info/blog/archives/1504-Korean-Chinese-Lunar-Calendar-Difference.html
    // gregorian: 2020-02-23 ~ 2020-03-23
    // korean lunar: 2020-01-30 ~ 2020-02-29
    // chinese lunar: 2020-02-01 ~ 2020-02-30

    fmt.Println(Solar2Lunar(d(2020, 2, 22)))
    fmt.Println(Solar2Lunar(d(2020, 2, 23)))
    fmt.Println(Solar2Lunar(d(2020, 3, 2)))
    fmt.Println(Solar2Lunar(d(2020, 5, 5)))
    fmt.Println(Solar2Lunar(d(2020, 5, 23)))
    fmt.Println(Solar2Lunar(d(2020, 6, 20)))
    fmt.Println(Solar2Lunar(d(2020, 6, 21)))

    // Output:
    // [2020 1 29] false
    // [2020 1 30] false
    // [2020 2 8] false
    // [2020 4 13] false
    // [2020 4 1] true
    // [2020 4 29] true
    // [2020 5 1] false
}
