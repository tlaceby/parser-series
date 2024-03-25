package ast

import "github.com/tlaceby/parser-series/src/lexer"

// -------------------
// LITERAL EXPRESSIONS
// -------------------

type NumberExpr struct {
	Value float64
}

func (n NumberExpr) expr() {}

type StringExpr struct {
	Value string
}

func (n StringExpr) expr() {}

type SymbolExpr struct {
	Value string
}

func (n SymbolExpr) expr() {}

// -------------------
// COMPLEX EXPRESSIONS
// -------------------

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (n BinaryExpr) expr() {}
