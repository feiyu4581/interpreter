package statement

import (
	"interpreter/ast"
	"interpreter/ast/expression"
	"interpreter/token"
)

type LetStatement struct {
	Token token.Token
	Name  *expression.Identifier
	Value ast.Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) StatementNode() {}
