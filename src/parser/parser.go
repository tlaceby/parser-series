package parser

import (
	"github.com/tlaceby/parser-series/src/ast"
	"github.com/tlaceby/parser-series/src/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos int
}

func createParser (tokens []lexer.Token) *parser {
	createTokenLookups()

	p := &parser{
		tokens: tokens,
		pos: 0,
	}

	return p
}

func Parse (source string) ast.Stmt {
	p := createParser(lexer.Tokenize(source))
	return parse_block_stmt(p)
}
