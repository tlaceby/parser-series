package parser

import (
	"errors"
	"fmt"
	"github.com/tlaceby/parser-series/src/lexer"
)

func (p *parser) currentToken () lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) advance () lexer.Token {
	tk := p.currentToken()
	p.pos++
	return tk
}

func (p *parser) nextToken () (lexer.Token, error) {
	if p.pos >= len(p.tokens) {
		return lexer.Token{}, errors.New("")
	}

	return p.tokens[p.pos + 1], nil
}

func (p *parser) previousToken () lexer.Token {
	return p.tokens[p.pos - 1]
}

func (p *parser) currentTokenKind () lexer.TokenKind {
	return p.tokens[p.pos].Kind()
}

func (p *parser) expectError (expectedKind lexer.TokenKind, err any) lexer.Token {
	token := p.currentToken()
	kind := token.Kind()

	if kind != expectedKind {
		if err == nil {
			err = fmt.Sprintf("Expected %s but recieved %s instead\n", lexer.TokenKindString(expectedKind), lexer.TokenKindString(kind))
		}

		panic(err)
	}

	return p.advance()
}

func (p *parser) expect (expectedKind lexer.TokenKind) lexer.Token {
	return p.expectError(expectedKind, nil)
}
