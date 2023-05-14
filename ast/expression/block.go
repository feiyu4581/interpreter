package expression

import (
	"bytes"
	"interpreter/ast/node"
	"interpreter/token"
)

type BlockStatement struct {
	Token      token.Token
	Statements []node.Statement
}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (bs *BlockStatement) StatementNode() {}
