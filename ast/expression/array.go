package expression

import (
	"bytes"
	"interpreter/ast/node"
	"interpreter/token"
	"strings"
)

type ArrayLiteral struct {
	Token    token.Token
	Elements []node.Expression
}

func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}

func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	var elements []string
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func (al *ArrayLiteral) ExpressionNode() {}
