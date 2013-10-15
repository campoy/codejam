package main

import (
	"github.com/campoy/codejam/2008/savetheuniverse"
	"os"
)

func main() {
	if err := savetheuniverse.ReadAndSolve(os.Stdout, os.Stdin); err != nil {
		fmt.Println(err)
	}
}
