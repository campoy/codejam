package savetheuniverse

import (
	"bufio"
	"fmt"
	"strconv"
)

type Input struct {
	Engines []string
	Queries []string
}

// ReadInputs reads the first line of the given scanner as an integer
// and then reads that many Inputs. It returns all the scanned inputs
// or an error if failed.
func ReadInputs(s *bufio.Scanner) ([]Input, error) {
	n, err := readInt(s, 0)
	if err != nil {
		return nil, fmt.Errorf("read number of cases: %v", err)
	}
	ins := make([]Input, n)
	for i := range ins {
		if err := ins[i].Init(s); err != nil {
			return nil, fmt.Errorf("input %v: %v", i, err)
		}
	}
	return ins, nil
}

// init initializes the Input with the data obtained from the given scanner.
func (in *Input) Init(s *bufio.Scanner) error {
	e, err := readInt(s, 1)
	if err != nil {
		return fmt.Errorf("read number of engines: %v", err)
	}
	in.Engines, err = readStrings(s, e)
	if err != nil {
		return fmt.Errorf("read engines: %v", err)
	}
	q, err := readInt(s, 0)
	if err != nil {
		return fmt.Errorf("read number of queries: %v", err)
	}
	in.Queries, err = readStrings(s, q)
	if err != nil {
		return fmt.Errorf("read queries: %v", err)
	}
	return nil
}

// readInt reads an integer from the given scanner.
func readInt(s *bufio.Scanner, min int64) (int64, error) {
	if !s.Scan() {
		return 0, fmt.Errorf("EOF")
	}
	txt := s.Text()
	n, err := strconv.ParseInt(txt, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("bad format %q: %v (%v)", txt, err, s.Err())
	}
	if n < min {
		return 0, fmt.Errorf("smaller than min(%v): %v", min, n)
	}
	return n, nil
}

// readStrings reads n strings from the given scanner.
func readStrings(s *bufio.Scanner, n int64) ([]string, error) {
	ss := make([]string, n)
	for i := range ss {
		if !s.Scan() {
			return nil, fmt.Errorf("EOF on %v of %v: %v", i, n, s.Err())
		}
		ss[i] = s.Text()
	}
	return ss, nil
}
