package ast

import (
	"interpreter/ast/expression"
	"interpreter/ast/node"
	"interpreter/ast/statement"
	"interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []node.Statement{
			&statement.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &expression.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &expression.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
