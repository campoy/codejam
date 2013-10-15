package savetheuniverse

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		engines, queries []string
		n                int
		err              error
	}{
		{
			engines: []string{"a"},
			queries: []string{"a"},
			err:     fmt.Errorf("no solution possible"),
		},
		{
			engines: []string{"a", "b"},
			queries: []string{"a", "b"},
			n:       1,
		},
	}

	for _, test := range tests {
		n, err := Solve(test.engines, test.queries)
		if !similarError(test.err, err) {
			t.Errorf("expected error %v, got %v", test.err, err)
			continue
		}
		if n != test.n {
			t.Errorf("expected solution %v, got %v", test.n, n)
		}
	}
}

func ExampleReadAndSolve() {
	r := strings.NewReader(strings.Join([]string{
		"2",
		"5",
		"Yeehaw",
		"NSM",
		"Dont Ask",
		"B9",
		"Googol",
		"10",
		"Yeehaw",
		"Yeehaw",
		"Googol",
		"B9",
		"Googol",
		"NSM",
		"B9",
		"NSM",
		"Dont Ask",
		"Googol",
		"5",
		"Yeehaw",
		"NSM",
		"Dont Ask",
		"B9",
		"Googol",
		"7",
		"Googol",
		"Dont Ask",
		"NSM",
		"NSM",
		"Yeehaw",
		"Yeehaw",
		"Googol",
	}, "\n"))
	if err := ReadAndSolve(os.Stdout, r); err != nil {
		fmt.Printf("fail: %v\n", err)
	}
	// Output:
	// Case #1: 1
	// Case #2: 0

}
