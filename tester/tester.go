package tester

import (
	"io"
	"os"
	"strings"
)

func Do(f func(w io.Writer, r io.Reader)) {
	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	names, err := dir.Readdirnames(0)
	if err != nil {
		panic(err)
	}

	for _, name := range names {
		if !strings.HasSuffix(name, ".in") {
			continue
		}
		outName := strings.TrimSuffix(name, "in") + "out"

		in, err := os.Open(name)
		if err != nil {
			panic(err)
		}
		out, err := os.Create(outName)
		if err != nil {
			panic(err)
		}

		f(out, in)
		in.Close()
		out.Close()
	}
}
