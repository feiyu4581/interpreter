package statement

import (
	"bytes"
	"interpreter/ast/expression"
	"interpreter/ast/node"
	"interpreter/token"
)

type LetStatement struct {
	Token token.Token
	Name  *expression.Identifier
	Value node.Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) StatementNode() {}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}
