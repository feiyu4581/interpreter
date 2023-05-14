package expression

import "interpreter/token"

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) ExpressionNode() {}

func (i *Identifier) String() string {
	return i.Value
}
