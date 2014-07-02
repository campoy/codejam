package magictrick

import (
	"fmt"
	"io"

	"github.com/campoy/codejam/input"
)

func Run(w io.Writer, r io.Reader) {
	s := input.NewScanner(r)

	t := s.Int()
	for i := 1; i <= t; i++ {
		var first, second, common []int

		choice := s.Int()
		for i := 1; i <= 4; i++ {
			if i == choice {
				first = s.IntsN(4)
			} else {
				s.SkipLine()
			}
		}

		choice = s.Int()
		for i := 1; i <= 4; i++ {
			if i == choice {
				second = s.IntsN(4)
			} else {
				s.SkipLine()
			}
		}

		for _, x := range first {
			for _, y := range second {
				if x == y {
					common = append(common, x)
				}
			}
		}

		switch len(common) {
		case 0:
			fmt.Fprintf(w, "Case #%v: Volunteer cheated!\n", i)
		case 1:
			fmt.Fprintf(w, "Case #%v: %v\n", i, common[0])
		default:
			fmt.Fprintf(w, "Case #%v: Bad magician!\n", i)
		}
	}
}
