package expression

import (
	"bytes"
	"interpreter/ast/node"
	"interpreter/token"
)

type AssignmentExpression struct {
	Token token.Token
	Name  *Identifier
	Value node.Expression
}

func (ae *AssignmentExpression) TokenLiteral() string {
	return ae.Token.Literal
}

func (ae *AssignmentExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ae.Name.String())
	out.WriteString(" = ")
	out.WriteString(ae.Value.String())

	return out.String()
}

func (ae *AssignmentExpression) ExpressionNode() {}
