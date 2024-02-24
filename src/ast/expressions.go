package ast

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
	Operation string
	Right Expr
}

func (n BinaryExpr) expr () {}

type AssignmentExpr struct {
	Assigne Expr
	AssignedValue Expr
}

func (n AssignmentExpr) expr () {}
