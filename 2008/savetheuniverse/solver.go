package savetheuniverse

import (
	"bufio"
	"fmt"
	"io"
)

// setFromSlice creates a map[string]bool in which the strings in
// the passed slice correspond to the keys with value true.
func setFromSlice(s []string) map[string]bool {
	set := make(map[string]bool, len(s))
	for _, v := range s {
		set[v] = true
	}
	return set
}

// Solve returns the minimum number of engine swaps that are needed
// for the given problem specification.
func Solve(engines, queries []string) (int, error) {
	n := 0
	set := setFromSlice(engines)

	for _, q := range queries {
		delete(set, q)
		if len(set) == 0 {
			set = setFromSlice(engines)
			delete(set, q)
			if len(set) == 0 {
				return 0, fmt.Errorf("no solution possible")
			}
			n++
		}
	}
	return n, nil
}

// ReadAndSolve reads all the input specifications from the given reader,
// solves every case, and writes the solutions to the given writer.
func ReadAndSolve(w io.Writer, r io.Reader) error {
	ins, err := ReadInputs(bufio.NewScanner(r))
	if err != nil {
		return fmt.Errorf("read: %v\n", err)
	}
	for i, in := range ins {
		sol, err := Solve(in.Engines, in.Queries)
		if err != nil {
			return fmt.Errorf("solve: %v", err)
		}
		fmt.Fprintf(w, "Case #%v: %v\n", i+1, sol)
	}
	return nil
}
