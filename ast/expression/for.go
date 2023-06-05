package expression

import (
	"bytes"
	"interpreter/ast/node"
	"interpreter/token"
)

type ForExpression struct {
	Token     token.Token
	Prefix    node.Expression
	Condition node.Expression
	Suffix    node.Expression
	Body      *BlockStatement
}

func (fe *ForExpression) TokenLiteral() string {
	return fe.Token.Literal
}

func (fe *ForExpression) String() string {
	var out bytes.Buffer

	out.WriteString("for")
	out.WriteString("(")
	out.WriteString(fe.Prefix.String())
	out.WriteString(";")
	out.WriteString(fe.Condition.String())
	out.WriteString(";")
	out.WriteString(fe.Suffix.String())
	out.WriteString(")")

	out.WriteString("{")
	out.WriteString(fe.Body.String())
	out.WriteString("{")

	return out.String()
}

func (fe *ForExpression) ExpressionNode() {}
