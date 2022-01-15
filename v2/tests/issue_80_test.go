package tests

import (
	"testing"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/au"
)

// Previous optimizations expected holidays to always fall on the same month,
// breaking this case where it has a custom calc func.
func TestAuBeforeGrandFinal(t *testing.T) {
	c := cal.NewBusinessCalendar()
	c.AddHoliday(au.FridayBeforeAflFinal)

	d := time.Date(2020, 10, 23, 0, 0, 0, 0, time.UTC)
	act, obs, _ := c.IsHoliday(d)
	if !act || !obs {
		t.Error("expected october holiday")
	}
}
