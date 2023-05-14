package expression

import (
	"bytes"
	"interpreter/ast/node"
	"interpreter/token"
	"strings"
)

type CallExpression struct {
	Token     token.Token
	Function  node.Expression
	Arguments []node.Expression
}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	var args []string
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

func (ce *CallExpression) ExpressionNode() {}
