package main

import (
	"os"

	"github.com/tenntenn/robo"
)

func main() { os.Exit(robo.Main(run)) }

func run(r *robo.Robo) {
	r.R()
	r.R()

	r.S()
	r.S()

	r.U()
	r.U()
}
