package cookieclicker

import (
	"fmt"
	"io"

	"github.com/campoy/codejam/input"
)

func Run(w io.Writer, r io.Reader) {
	s := input.NewScanner(r)

	t := s.Int()
	for i := 1; i <= t; i++ {
		fields := s.FloatsN(3)
		cost, extra, goal := fields[0], fields[1], fields[2]

		// Time to build n factories.
		times := []float64{0}
		rate := 2.
		best := goal / rate

		// n is the number of factories built.
		for n := 1; ; n++ {
			// compute the time to build the nth factory.
			nth := times[n-1] + cost/rate
			times = append(times, nth)

			// increase the rate of production.
			rate += extra
			time := goal/rate + nth

			// function is quadratic, we found local minimum.
			if time > best {
				break
			}
			best = time
		}
		fmt.Fprintf(w, "Case #%v: %.7f\n", i, best)
	}
}
