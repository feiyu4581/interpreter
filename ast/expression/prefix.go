package expression

import (
	"bytes"
	"interpreter/ast/node"
	"interpreter/token"
)

type PrefixExpression struct {
	Token    token.Token // 前缀词法单元，如!
	Operator string
	Right    node.Expression
}

func (pe *PrefixExpression) ExpressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
