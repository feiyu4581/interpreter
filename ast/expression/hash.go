package expression

import (
	"bytes"
	"interpreter/ast/node"
	"interpreter/token"
	"strings"
)

type HashLiteral struct {
	Token token.Token
	Pairs map[node.Expression]node.Expression
}

func (hl *HashLiteral) TokenLiteral() string {
	return hl.Token.Literal
}

func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	var pairs []string
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

func (hl *HashLiteral) ExpressionNode() {}
