package input

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Scanner struct {
	s    *bufio.Scanner
	line int
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{bufio.NewScanner(r), 0}
}

func (s *Scanner) SkipLine() {
	if !s.s.Scan() {
		log.Fatalf("[l:%v] unexpected EOF skipping line", s.line)
	}
}

func (s *Scanner) String() string {
	if !s.s.Scan() {
		log.Fatalf("[l:%v] unexpected EOF", s.line)
	}
	if err := s.s.Err(); err != nil {
		log.Fatalf("[l:%v] unexpected error %v", s.line, err)
	}
	s.line++
	return s.s.Text()
}

func (s *Scanner) Strings() []string {
	return strings.Fields(s.String())
}

func (s *Scanner) StringsN(n int) []string {
	res := s.Strings()
	if len(res) != n {
		log.Fatalf("[l:%v] expected %v elements, got %v", s.line, n, len(res))
	}
	return res
}

func (s *Scanner) atoi(txt string) int {
	v, err := strconv.Atoi(txt)
	if err != nil {
		log.Fatalf("[l:%v] not a number: %v", s.line, err)
	}
	return v
}

func (s *Scanner) Int() int {
	return s.atoi(s.String())
}

func (s *Scanner) Ints() []int {
	res := []int{}
	for _, f := range s.Strings() {
		res = append(res, s.atoi(f))
	}
	return res
}

func (s *Scanner) IntsN(n int) []int {
	res := s.Ints()
	if len(res) != n {
		log.Fatalf("[l:%v] expected %v elements, got %v", s.line, n, len(res))
	}
	return res
}
