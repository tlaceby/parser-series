package parser

import (
	"github.com/tlaceby/parser-series/src/ast"
	"github.com/tlaceby/parser-series/src/lexer"
)

type binding_power int
const (
	defalt_bp binding_power = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary binding_power = 7
	prefix
	cast
	postfix binding_power = 8
	call
	member
	primary binding_power = 9
)

type stmt_handler func (p *parser) ast.Stmt
type nud_handler  func (p *parser) ast.Expr
type led_handler  func (p *parser, left ast.Expr, bp binding_power) ast.Expr

type stmt_lookup map[lexer.TokenKind]stmt_handler
type nud_lookup map[lexer.TokenKind]nud_handler
type led_lookup map[lexer.TokenKind]led_handler
type bp_lookup map[lexer.TokenKind]binding_power

var bp_lu = bp_lookup{}
var nud_lu = nud_lookup{}
var led_lu = led_lookup{}
var stmt_lu =stmt_lookup{}


func defInfix (kind lexer.TokenKind, bp binding_power, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

func defPrimary (kind lexer.TokenKind,nud_fn nud_handler) {
	bp_lu[kind] = primary
	nud_lu[kind] = nud_fn
}

func defStmt (kind lexer.TokenKind, stmt_fn stmt_handler) {
	bp_lu[kind] = defalt_bp
	stmt_lu[kind] = stmt_fn
}

func createTokenLookups () {
	defInfix(lexer.PLUS, additive, parse_binary_expr)
	defInfix(lexer.DASH, additive, parse_binary_expr)
	defInfix(lexer.SLASH, multiplicative, parse_binary_expr)
	defInfix(lexer.STAR, multiplicative, parse_binary_expr)
	defInfix(lexer.PERCENT, multiplicative, parse_binary_expr)

	defPrimary(lexer.NUMBER, parse_primary_expr)
	defPrimary(lexer.STRING, parse_primary_expr)
	defPrimary(lexer.IDENTIFIER, parse_primary_expr)

	defStmt(lexer.OPEN_CURLY, parse_block_stmt)
	defStmt(lexer.LET, parse_var_decl_stmt)
	defStmt(lexer.CONST, parse_var_decl_stmt)
}
