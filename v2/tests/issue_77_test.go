package tests

import (
	"os"
	"runtime/pprof"
	"testing"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/us"
)

// V2 was much slower than V1 due to more extensibility and usage of fields
// in time.Time such as timezone.
func Benchmark2017IsHoliday(b *testing.B) {
	f, err := os.Create("../../cpu.pprof")
	if err != nil {
		b.Fatalf("unable to create cpu profile file")
	}
	defer f.Close()

	MonThruSatCal := cal.NewBusinessCalendar()
	MonThruSatCal.SetWorkday(time.Saturday, true)
	MonThruSatCal.Holidays = us.Holidays
	MonThruSatCal.Cacheable = true
	dt := time.Date(2017, time.January, 3, 0, 0, 0, 0, time.Local)

	if err = pprof.StartCPUProfile(f); err != nil {
		b.Fatalf("unable to start cpu proflie")
	}
	defer pprof.StopCPUProfile()
	for i := 0; i < b.N; i++ {
		MonThruSatCal.IsHoliday(dt)
	}
}
