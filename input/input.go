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
		log.Fatalf("[l:%v] parse int: %v", s.line, err)
	}
	return v
}

func (s *Scanner) Int() int {
	return s.atoi(s.String())
}

func (s *Scanner) Ints() []int {
	res := []int{}
	for _, v := range s.Strings() {
		res = append(res, s.atoi(v))
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

func (s *Scanner) atof(txt string) float64 {
	v, err := strconv.ParseFloat(txt, 64)
	if err != nil {
		log.Fatalf("[l:%v] parse float: %v", s.line, err)
	}
	return v
}

func (s *Scanner) Float() float64 {
	return s.atof(s.String())
}

func (s *Scanner) Floats() []float64 {
	res := []float64{}
	for _, v := range s.Strings() {
		res = append(res, s.atof(v))
	}
	return res
}

func (s *Scanner) FloatsN(n int) []float64 {
	res := s.Floats()
	if len(res) != n {
		log.Fatalf("[l:%v] expected %v elements, got %v", s.line, n, len(res))
	}
	return res
}
