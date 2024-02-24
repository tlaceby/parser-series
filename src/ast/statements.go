package ast

type BlockStmt struct {
	Body []Stmt
}

func (b BlockStmt) stmt () {}

type VarDeclarationStmt struct {
	Identifier string
	Constant bool
	AssignedValue Expr
	ExplicitType Type
}


func (n VarDeclarationStmt) stmt () {}

type ExpressionStmt struct {
	Expression Expr
}

func (n ExpressionStmt) stmt () {}

type Parameter struct {
	Name string
	Type Type
}

type FunctionDeclarationStmt struct {
	Parameters []Parameter
	Name string
	Body []Stmt
	ReturnType Type
}

func (n FunctionDeclarationStmt) stmt () {}
