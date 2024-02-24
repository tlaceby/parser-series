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
	createTypeTokenLookups()

	p := &parser{
		tokens: tokens,
		pos: 0,
	}

	return p
}

func Parse (source string) ast.BlockStmt {
	tokens := lexer.Tokenize(source)
	p := createParser(tokens)
	body := make([]ast.Stmt, 0)

	for p.hasTokens() {
		body = append(body, parse_stmt(p))
	}

	return ast.BlockStmt{
		Body: body,
	}
}
