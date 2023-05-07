package statement

import (
	"interpreter/ast"
	"interpreter/token"
)

type ReturnStatement struct {
	Token       token.Token
	ReturnValue ast.Expression
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) StatementNode() {}
