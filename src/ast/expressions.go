package ast

import (
	"github.com/tlaceby/parser-series/src/lexer"
)

// --------------------
// Literal Expressions
// --------------------

type NumberExpr struct {
	Value float64
}

func (n NumberExpr) expr () {}

type StringExpr struct {
	Value string
}

func (n StringExpr) expr () {}

type SymbolExpr struct {
	Value string
}

func (n SymbolExpr) expr () {}

// --------------------
// Complex Expressions
// --------------------

type BinaryExpr struct {
	Left Expr
	Operator lexer.Token
	Right Expr
}

func (n BinaryExpr) expr () {}

type AssignmentExpr struct {
	Assigne Expr
	AssignedValue Expr
}

func (n AssignmentExpr) expr () {}

type PrefixExpr struct {
	Operator lexer.Token
	Right Expr
}

func (n PrefixExpr) expr () {}

type MemberExpr struct {
	Member Expr
	Property string
}

func (n MemberExpr) expr () {}

type CallExpr struct {
	Arguments []Expr
	Method Expr
}

func (n CallExpr) expr () {}


type ComputedExpr struct {
	Member Expr
	Property Expr
}

func (n ComputedExpr) expr () {}
