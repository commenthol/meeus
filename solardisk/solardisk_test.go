package solardisk_test

import (
	"fmt"
	"time"

	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/solardisk"
)

func ExampleCycle() {
	j := solardisk.Cycle(1699)
	fmt.Printf("%.4f\n", j)
	y, m, d := julian.JDToCalendar(j)
	fmt.Printf("%d %s %.2f\n", y, time.Month(m), d)
	// Output:
	// 2444480.7230
	// 1980 August 29.22
}
