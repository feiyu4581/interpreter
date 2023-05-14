package statement

import (
	"interpreter/ast/node"
	"interpreter/token"
)

type ExpressionStatement struct {
	Token      token.Token
	Expression node.Expression
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) StatementNode() {}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
