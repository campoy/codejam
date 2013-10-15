package savetheuniverse

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func similarError(short, long error) bool {
	if (short == nil) != (long == nil) {
		return false
	}
	if short == nil {
		return true
	}
	return strings.Contains(long.Error(), short.Error())
}

func TestInput(t *testing.T) {
	tests := []struct {
		lines []string
		exp   Input
		err   error
	}{
		{
			lines: []string{
				"2",
				"google",
				"yahoo",
				"3",
				"google",
				"google",
				"google",
			},
			exp: Input{
				Engines: []string{"google", "yahoo"},
				Queries: []string{"google", "google", "google"},
			},
			err: nil,
		},
		{
			lines: []string{
				"google",
			},
			err: fmt.Errorf("number of engines"),
		},
		{
			lines: []string{"0"},
			err:   fmt.Errorf("smaller than min(1)"),
		},
		{
			lines: []string{"1.5"},
			err:   fmt.Errorf("invalid syntax"),
		},
	}
	for _, test := range tests {
		s := bufio.NewScanner(strings.NewReader(strings.Join(test.lines, "\n")))
		var in Input
		err := in.Init(s)
		if !similarError(test.err, err) {
			t.Errorf("expected error %v, got %v", test.err, err)
			continue
		}
		if !reflect.DeepEqual(in, test.exp) {
			t.Errorf("expected input %v, got %v", test.exp, in)
		}
	}
}

func TestReadInputs(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(strings.Join([]string{
		"2",
		"2",
		"google",
		"yahoo",
		"3",
		"google",
		"google",
		"google",
		"1",
		"yahoo",
		"2",
		"google",
		"yahoo",
	}, "\n")))

	ins, err := ReadInputs(s)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	exp := []Input{
		Input{
			Engines: []string{"google", "yahoo"},
			Queries: []string{"google", "google", "google"},
		},
		Input{
			Engines: []string{"yahoo"},
			Queries: []string{"google", "yahoo"},
		},
	}
	if !reflect.DeepEqual(exp, ins) {
		t.Errorf("expected %v, got %v", exp, ins)
	}
}
