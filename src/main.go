package main

import (
	"os"

	"github.com/tlaceby/parser-series/src/lexer"
)	

func main () {
	bytes, _ := os.ReadFile("./examples/01.lang")
	tokens := lexer.Tokenize(string(bytes))
	
	for _, token := range tokens {
		token.Debug()
	}
}