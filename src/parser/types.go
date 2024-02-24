package parser

import (
	"fmt"

	"github.com/tlaceby/parser-series/src/ast"
	"github.com/tlaceby/parser-series/src/lexer"
)


type type_nud_handler  func (p *parser) ast.Type
type type_led_handler  func (p *parser, left ast.Type, bp binding_power) ast.Type

type type_nud_lookup map[lexer.TokenKind]type_nud_handler
type type_led_lookup map[lexer.TokenKind]type_led_handler
type type_bp_lookup map[lexer.TokenKind]binding_power

var type_bp_lu = type_bp_lookup{}
var type_nud_lu = type_nud_lookup{}
var type_led_lu = type_led_lookup{}


func type_led (kind lexer.TokenKind, bp binding_power, led_fn type_led_handler) {
	type_bp_lu[kind] = bp
	type_led_lu[kind] = led_fn
}

func type_nud (kind lexer.TokenKind, bp binding_power, nud_fn type_nud_handler) {
	type_bp_lu[kind] = primary
	type_nud_lu[kind] = nud_fn
}

func createTypeTokenLookups () {

	type_nud(lexer.IDENTIFIER, primary, func(p *parser) ast.Type {
		return ast.SymbolType{
			Value: p.advance().Value,
		}
	})

	// []number
	type_nud(lexer.OPEN_BRACKET	, member, func(p *parser) ast.Type {
		p.advance()
		p.expect(lexer.CLOSE_BRACKET)
		insideType := parse_type(p, defalt_bp)

		return ast.ListType{
			Underlying: insideType,
		}
	})
}

func parse_type (p *parser, bp binding_power) ast.Type {
	tokenKind := p.currentTokenKind()
	nud_fn, exists := type_nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("type: NUD Handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)

	for type_bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := type_led_lu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("type: LED Handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, bp)
	}

	return left
}
