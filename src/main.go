package main

import (
	"github.com/tlaceby/parser-series/src/parser"
)

func main() {
	source := "let x = 45.5 * 20;"
	parser.Parse(source)
}
