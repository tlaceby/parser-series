package ast

type BlockStmt struct {
	Body []Stmt
}

func (b BlockStmt) stmt () {}

type VarDeclarationStmt struct {
	Identifier string
	Constant bool
	AssignedValue Expr

	// Type Info if your language had types
	// eg  infered or explicit
}


func (n VarDeclarationStmt) stmt () {}

// Represents all expressions which would otherwise need semicolons etc..
type ExpressionStmt struct {
	Expression Expr
}

func (n ExpressionStmt) stmt () {}
