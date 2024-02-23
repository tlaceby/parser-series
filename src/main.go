package main

import (
	"github.com/tlaceby/parser-series/src/lexer"
)

func main() {
	tokens := lexer.Tokenize("let x = math.abs(45.32 * 3.2);\nlet zeroToTen = [0..10];")

	for _, t := range tokens {
		t.Debug()
	}
}
