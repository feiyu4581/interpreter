package expression

import (
	"bytes"
	"interpreter/ast/node"
	"interpreter/token"
)

type IndexExpression struct {
	Token token.Token
	Left  node.Expression
	Index node.Expression
}

func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

func (ie *IndexExpression) ExpressionNode() {}
