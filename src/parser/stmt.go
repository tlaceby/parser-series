package parser

import (
	"github.com/tlaceby/parser-series/src/ast"
	"github.com/tlaceby/parser-series/src/lexer"
)

func parse_stmt (p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists {
		return stmt_fn(p)
	}

	return parse_expression_stmt(p)
}

func parse_expression_stmt (p *parser) ast.ExpressionStmt{
	expression := parse_expr(p, defalt_bp)
	p.expect(lexer.SEMI_COLON)

	return ast.ExpressionStmt{
		Expression: expression,
	}
}

func parse_block_stmt (p *parser) ast.Stmt {
	body := []ast.Stmt{}

	for p.pos < len(p.tokens) {
		body = append(body, parse_stmt(p))
	}

	return ast.BlockStmt{
		Body: body,
	}
}

func parse_var_decl_stmt (p *parser) ast.Stmt {
	panic("parse_var_decl_stmt not implimented")
}
