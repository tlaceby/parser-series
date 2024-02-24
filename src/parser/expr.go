package parser

import (
	"fmt"

	"github.com/tlaceby/parser-series/src/ast"
	"github.com/tlaceby/parser-series/src/lexer"
)

func parse_expr (p *parser, bp binding_power) ast.Expr {
	tokenKind := p.currentTokenKind()
	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("NUD Handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)

	for bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := led_lu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("LED Handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, bp)
	}

	return left
}

func parse_binary_expr (p *parser, left ast.Expr, bp binding_power) ast.Expr {
	panic("Binary Expr not implimented")
}

func parse_primary_expr (p *parser) ast.Expr {
	panic("Primary Expr not implimented")
}
