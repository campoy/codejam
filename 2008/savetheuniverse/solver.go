package savetheuniverse

import (
	"bufio"
	"fmt"
	"io"
)

func setFromSlice(s []string) map[string]bool {
	set := make(map[string]bool, len(s))
	for _, v := range s {
		set[v] = true
	}
	return set
}

func Solve(in Input) (int, error) {
	n := 0
	set := setFromSlice(in.Engines)

	for _, q := range in.Queries {
		delete(set, q)
		if len(set) == 0 {
			set = setFromSlice(in.Engines)
			delete(set, q)
			if len(set) == 0 {
				return 0, fmt.Errorf("no solution possible")
			}
			n++
		}
	}
	return n, nil
}

func SolveAll(ins []Input) ([]int, error) {
	s := make([]int, len(ins))
	var err error
	for i, in := range ins {
		s[i], err = Solve(in)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func ReadAndSolve(w io.Writer, r io.Reader) error {
	ins, err := ReadInputs(bufio.NewScanner(r))
	if err != nil {
		return fmt.Errorf("read: %v\n", err)
	}
	sols, err := SolveAll(ins)
	if err != nil {
		return fmt.Errorf("solve: %v\n", err)
	}
	for i, sol := range sols {
		fmt.Fprintf(w, "Case #%v: %v\n", i+1, sol)
	}
	return nil
}
