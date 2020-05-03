// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"reflect"
	"testing"
	"time"
)

func TestClone(t *testing.T) {
	h := Holiday{
		Day:         0,
		Description: "1",
		EndYear:     2,
		Except:      []int{3},
		Func:        CalcDayOfMonth,
		Julian:      true,
		Month:       6,
		Name:        "7",
		Observed:    []AltDay{{Day: 8, Offset: 8}},
		Offset:      9,
		StartYear:   10,
		Type:        11,
		Weekday:     12,
	}

	c := h.Clone("clone", ObservanceBank, []AltDay{{Day: 1, Offset: 2}})
	if c.Day != h.Day || c.Description != h.Description || c.EndYear != h.EndYear ||
		!reflect.DeepEqual(c.Except, h.Except) || c.Julian != h.Julian ||
		c.Month != h.Month || c.Name != "clone" || reflect.DeepEqual(c.Observed, h.Observed) || c.Offset != h.Offset ||
		c.StartYear != h.StartYear || c.Type != ObservanceBank || c.Weekday != h.Weekday {

		t.Errorf("bad clone")
	}
}

func TestCalc(t *testing.T) {

	h1 := &Holiday{}
	act, obs := h1.Calc(2020)
	if !act.IsZero() || !obs.IsZero() {
		t.Errorf("expected zero time for nil calc func")
	}

	h2 := &Holiday{
		StartYear: 2015,
		EndYear:   2025,
		Except:    []int{2020},
		Func:      func(h *Holiday, year int) time.Time { return time.Date(year, time.March, 11, 0, 0, 0, 0, DefaultLoc) },
	}

	act, obs = h2.Calc(2000)
	if !act.IsZero() || !obs.IsZero() {
		t.Errorf("expected zero time before start year")
	}
	act, obs = h2.Calc(2026)
	if !act.IsZero() || !obs.IsZero() {
		t.Errorf("expected zero time after end year")
	}
	act, obs = h2.Calc(2020)
	if !act.IsZero() || !obs.IsZero() {
		t.Errorf("expected zero time in exception year")
	}
	act, obs = h2.Calc(2015)
	if !act.Equal(obs) || !act.Equal(time.Date(2015, time.March, 11, 0, 0, 0, 0, DefaultLoc)) {
		t.Errorf("normal day - got: %s, %s; want: %s", act, obs, time.Date(2015, time.March, 11, 0, 0, 0, 0, DefaultLoc))
	}

	h2.Observed = []AltDay{{Day: time.Wednesday, Offset: 2}}
	act, obs = h2.Calc(2015)
	if act.Equal(obs) {
		t.Errorf("observed day 2015 - got: %s, %s; want: %s, %s", act, obs,
			time.Date(2015, time.March, 11, 0, 0, 0, 0, DefaultLoc),
			time.Date(2015, time.March, 13, 0, 0, 0, 0, DefaultLoc))
	}

	act, obs = h2.Calc(2016)
	if !act.Equal(obs) || !act.Equal(time.Date(2016, time.March, 11, 0, 0, 0, 0, DefaultLoc)) {
		t.Errorf("observed day 2016 - got: %s, %s; want: %s, %s", act, obs,
			time.Date(2016, time.March, 11, 0, 0, 0, 0, DefaultLoc),
			time.Date(2016, time.March, 11, 0, 0, 0, 0, DefaultLoc))
	}
}

func TestCalcDayOfMonth(t *testing.T) {
	tests := []struct {
		y    int
		want time.Time
	}{
		{2015, time.Date(2015, 6, 20, 0, 0, 0, 0, DefaultLoc)},
		{2016, time.Date(2016, 6, 20, 0, 0, 0, 0, DefaultLoc)},
		{2017, time.Date(2017, 6, 20, 0, 0, 0, 0, DefaultLoc)},
		{2018, time.Date(2018, 6, 20, 0, 0, 0, 0, DefaultLoc)},
		{2019, time.Date(2019, 6, 20, 0, 0, 0, 0, DefaultLoc)},
		{2020, time.Date(2020, 6, 20, 0, 0, 0, 0, DefaultLoc)},
	}
	h := &Holiday{
		Month:    time.June,
		Day:      20,
		Observed: []AltDay{{time.Monday, 2}, {time.Wednesday, -3}},
	}

	for _, test := range tests {
		got := CalcDayOfMonth(h, test.y)
		if !got.Equal(test.want) {
			t.Errorf("%d: got: %s, want: %s", test.y, got, test.want)
		}
	}
}

func TestCalcWeekdayOffset(t *testing.T) {
	tests := []struct {
		off  int
		want time.Time
	}{
		{1, time.Date(2015, 6, 2, 0, 0, 0, 0, DefaultLoc)},
		{2, time.Date(2015, 6, 9, 0, 0, 0, 0, DefaultLoc)},
		{3, time.Date(2015, 6, 16, 0, 0, 0, 0, DefaultLoc)},
		{4, time.Date(2015, 6, 23, 0, 0, 0, 0, DefaultLoc)},
		{5, time.Date(2015, 6, 30, 0, 0, 0, 0, DefaultLoc)},
		{0, time.Time{}},
		{-1, time.Date(2015, 6, 30, 0, 0, 0, 0, DefaultLoc)},
		{-2, time.Date(2015, 6, 23, 0, 0, 0, 0, DefaultLoc)},
		{-3, time.Date(2015, 6, 16, 0, 0, 0, 0, DefaultLoc)},
		{-4, time.Date(2015, 6, 9, 0, 0, 0, 0, DefaultLoc)},
		{-5, time.Date(2015, 6, 2, 0, 0, 0, 0, DefaultLoc)},
	}

	h := &Holiday{
		Month:   time.June,
		Weekday: time.Tuesday,
	}

	for _, test := range tests {
		h.Offset = test.off
		got := CalcWeekdayOffset(h, 2015)
		if !got.Equal(test.want) {
			t.Errorf("%d: got: %s, want: %s", test.off, got, test.want)
		}
	}
}

func TestCalcWeekdayFrom(t *testing.T) {
	tests := []struct {
		off  int
		want time.Time
	}{
		{1, time.Date(2015, 6, 16, 0, 0, 0, 0, DefaultLoc)},
		{2, time.Date(2015, 6, 23, 0, 0, 0, 0, DefaultLoc)},
		{3, time.Date(2015, 6, 30, 0, 0, 0, 0, DefaultLoc)},
		{4, time.Date(2015, 7, 7, 0, 0, 0, 0, DefaultLoc)},
		{5, time.Date(2015, 7, 14, 0, 0, 0, 0, DefaultLoc)},
		{0, time.Time{}},
		{-1, time.Date(2015, 6, 9, 0, 0, 0, 0, DefaultLoc)},
		{-2, time.Date(2015, 6, 2, 0, 0, 0, 0, DefaultLoc)},
		{-3, time.Date(2015, 5, 26, 0, 0, 0, 0, DefaultLoc)},
		{-4, time.Date(2015, 5, 19, 0, 0, 0, 0, DefaultLoc)},
		{-5, time.Date(2015, 5, 12, 0, 0, 0, 0, DefaultLoc)},
	}

	h := &Holiday{
		Month:   time.June,
		Weekday: time.Tuesday,
		Day:     15,
	}

	for _, test := range tests {
		h.Offset = test.off
		got := CalcWeekdayFrom(h, 2015)
		if !got.Equal(test.want) {
			t.Errorf("%d: got: %s, want: %s", test.off, got, test.want)
		}
	}
}

func TestCalcEasterOffset(t *testing.T) {
	tests := []struct {
		y    int
		off  int
		jul  bool
		want time.Time
	}{
		{2015, 0, false, time.Date(2015, 4, 5, 0, 0, 0, 0, DefaultLoc)},
		{2016, 10, false, time.Date(2016, 4, 6, 0, 0, 0, 0, DefaultLoc)},
		{2017, -10, false, time.Date(2017, 4, 6, 0, 0, 0, 0, DefaultLoc)},
		{2015, 0, true, time.Date(2015, 4, 12, 0, 0, 0, 0, DefaultLoc)},
		{2016, 10, true, time.Date(2016, 5, 11, 0, 0, 0, 0, DefaultLoc)},
		{2017, -10, true, time.Date(2017, 4, 6, 0, 0, 0, 0, DefaultLoc)},
	}

	h := &Holiday{}

	for _, test := range tests {
		h.Offset = test.off
		h.Julian = test.jul
		got := CalcEasterOffset(h, test.y)
		if !got.Equal(test.want) {
			t.Errorf("%d: got: %s, want: %s", test.off, got, test.want)
		}
	}
}
