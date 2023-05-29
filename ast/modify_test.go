// ast/modify_test.go

package ast

import (
	"interpreter/ast/expression"
	"interpreter/ast/node"
	"interpreter/ast/statement"
	"reflect"
	"testing"
)

func TestModify(t *testing.T) {
	one := func() node.Expression { return &expression.IntegerLiteral{Value: 1} }
	two := func() node.Expression { return &expression.IntegerLiteral{Value: 2} }

	turnOneIntoTwo := func(node node.Node) node.Node {
		integer, ok := node.(*expression.IntegerLiteral)
		if !ok {
			return node
		}

		if integer.Value != 1 {
			return node
		}

		integer.Value = 2
		return integer
	}

	tests := []struct {
		input    node.Node
		expected node.Node
	}{
		{
			one(),
			two(),
		},
		{
			&Program{
				Statements: []node.Statement{
					&statement.ExpressionStatement{Expression: one()},
				},
			},
			&Program{
				Statements: []node.Statement{
					&statement.ExpressionStatement{Expression: two()},
				},
			},
		},
		{
			&expression.InfixExpression{Left: one(), Operator: "+", Right: two()},
			&expression.InfixExpression{Left: two(), Operator: "+", Right: two()},
		},
		{
			&expression.InfixExpression{Left: two(), Operator: "+", Right: one()},
			&expression.InfixExpression{Left: two(), Operator: "+", Right: two()},
		},
		{
			&expression.PrefixExpression{Operator: "-", Right: one()},
			&expression.PrefixExpression{Operator: "-", Right: two()},
		},
		{
			&expression.IndexExpression{Left: one(), Index: one()},
			&expression.IndexExpression{Left: two(), Index: two()},
		},
		{
			&expression.IfExpression{
				Condition: one(),
				Consequence: &expression.BlockStatement{
					Statements: []node.Statement{
						&statement.ExpressionStatement{Expression: one()},
					},
				},
				Alternative: &expression.BlockStatement{
					Statements: []node.Statement{
						&statement.ExpressionStatement{Expression: one()},
					},
				},
			},
			&expression.IfExpression{
				Condition: two(),
				Consequence: &expression.BlockStatement{
					Statements: []node.Statement{
						&statement.ExpressionStatement{Expression: two()},
					},
				},
				Alternative: &expression.BlockStatement{
					Statements: []node.Statement{
						&statement.ExpressionStatement{Expression: two()},
					},
				},
			},
		},
		{
			&statement.ReturnStatement{ReturnValue: one()},
			&statement.ReturnStatement{ReturnValue: two()},
		},
		{
			&statement.LetStatement{Value: one()},
			&statement.LetStatement{Value: two()},
		},
		{
			&expression.FunctionLiteral{
				Parameters: []*expression.Identifier{},
				Body: &expression.BlockStatement{
					Statements: []node.Statement{
						&statement.ExpressionStatement{Expression: one()},
					},
				},
			},
			&expression.FunctionLiteral{
				Parameters: []*expression.Identifier{},
				Body: &expression.BlockStatement{
					Statements: []node.Statement{
						&statement.ExpressionStatement{Expression: two()},
					},
				},
			},
		},
		{
			&expression.ArrayLiteral{Elements: []node.Expression{one(), one()}},
			&expression.ArrayLiteral{Elements: []node.Expression{two(), two()}},
		},
	}

	for _, tt := range tests {
		modified := Modify(tt.input, turnOneIntoTwo)

		equal := reflect.DeepEqual(modified, tt.expected)
		if !equal {
			t.Errorf("not equal. got=%#v, want=%#v",
				modified, tt.expected)
		}
	}

	hashLiteral := &expression.HashLiteral{
		Pairs: map[node.Expression]node.Expression{
			one(): one(),
			one(): one(),
		},
	}

	Modify(hashLiteral, turnOneIntoTwo)

	for key, val := range hashLiteral.Pairs {
		key, _ := key.(*expression.IntegerLiteral)
		if key.Value != 2 {
			t.Errorf("value is not %d, got=%d", 2, key.Value)
		}
		val, _ := val.(*expression.IntegerLiteral)
		if val.Value != 2 {
			t.Errorf("value is not %d, got=%d", 2, val.Value)
		}
	}
}
