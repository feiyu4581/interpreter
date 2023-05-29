package ast

import (
	"interpreter/ast/expression"
	"interpreter/ast/node"
	"interpreter/ast/statement"
)

type ModifierFunc func(node.Node) node.Node

func Modify(n node.Node, modifier ModifierFunc) node.Node {
	switch n := n.(type) {
	case *Program:
		for i, stmt := range n.Statements {
			n.Statements[i], _ = Modify(stmt, modifier).(node.Statement)
		}
	case *statement.ExpressionStatement:
		n.Expression, _ = Modify(n.Expression, modifier).(node.Expression)
	case *expression.InfixExpression:
		n.Left, _ = Modify(n.Left, modifier).(node.Expression)
		n.Right, _ = Modify(n.Right, modifier).(node.Expression)
	case *expression.PrefixExpression:
		n.Right, _ = Modify(n.Right, modifier).(node.Expression)
	case *expression.IndexExpression:
		n.Left, _ = Modify(n.Left, modifier).(node.Expression)
		n.Index, _ = Modify(n.Index, modifier).(node.Expression)
	case *expression.IfExpression:
		n.Condition, _ = Modify(n.Condition, modifier).(node.Expression)
		n.Consequence, _ = Modify(n.Consequence, modifier).(*expression.BlockStatement)
		if n.Alternative != nil {
			n.Alternative, _ = Modify(n.Alternative, modifier).(*expression.BlockStatement)
		}
	case *expression.BlockStatement:
		for i := range n.Statements {
			n.Statements[i], _ = Modify(n.Statements[i], modifier).(node.Statement)
		}
	case *statement.ReturnStatement:
		n.ReturnValue, _ = Modify(n.ReturnValue, modifier).(node.Expression)
	case *statement.LetStatement:
		n.Value, _ = Modify(n.Value, modifier).(node.Expression)
	case *expression.FunctionLiteral:
		for i := range n.Parameters {
			n.Parameters[i], _ = Modify(n.Parameters[i], modifier).(*expression.Identifier)
		}

		n.Body, _ = Modify(n.Body, modifier).(*expression.BlockStatement)
	case *expression.ArrayLiteral:
		for i := range n.Elements {
			n.Elements[i], _ = Modify(n.Elements[i], modifier).(node.Expression)
		}
	case *expression.HashLiteral:
		newPairs := make(map[node.Expression]node.Expression)
		for key, val := range n.Pairs {
			newKey, _ := Modify(key, modifier).(node.Expression)
			newVal, _ := Modify(val, modifier).(node.Expression)
			newPairs[newKey] = newVal
		}

		n.Pairs = newPairs
	}

	return modifier(n)
}
