package main

import (
	"os"

	"github.com/sanity-io/litter"
	"github.com/tlaceby/parser-series/src/lexer"
	"github.com/tlaceby/parser-series/src/parser"
)

func main() {
	bytes, _ := os.ReadFile("./examples/02.lang")
	tokens := lexer.Tokenize(string(bytes))

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
