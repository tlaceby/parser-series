package main

import (
	"github.com/tlaceby/parser-series/src/lexer"
)

func main() {
	tokens := lexer.Tokenize("let x = math.abs(45.32 * 3.2); \n// This is a comment\n for (let i = 0; i <= 10; i++) \nprintln(\"Hello world\");")

	for _, t := range tokens {
		t.Debug()
	}
}
