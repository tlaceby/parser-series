package main

import (
	"github.com/sanity-io/litter"
	"github.com/tlaceby/parser-series/src/parser"
)

func main() {
	source := "fn add (x: number, y: number): number { x + y; }"
	ast := parser.Parse(source)

	litter.Dump(ast)
}
