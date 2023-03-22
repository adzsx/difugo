package main

import (
	"os"

	"github.com/adzsx/dirsgover/pkg/format"
)

func main() {
	args := os.Args

	format.Args(args)
}
