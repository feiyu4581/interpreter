package evaluator

import (
	"fmt"
	"interpreter/ast"
	"interpreter/ast/expression"
	"interpreter/ast/node"
	"interpreter/object"
	"interpreter/token"
)

func quote(n node.Node, env *object.Environment) object.Object {
	return &object.Quote{Node: evalUnquoteCalls(n, env)}
}

func evalUnquoteCalls(quoted node.Node, env *object.Environment) node.Node {
	return ast.Modify(quoted, func(n node.Node) node.Node {
		if !isUnquoteCall(n) {
			return n
		}

		call, ok := n.(*expression.CallExpression)
		if !ok {
			return n
		}

		if len(call.Arguments) != 1 {
			return n
		}

		unquoted := Eval(call.Arguments[0], env)
		return convertObjectToASTNode(unquoted)
	})
}

func isUnquoteCall(n node.Node) bool {
	callExpression, ok := n.(*expression.CallExpression)
	if !ok {
		return false
	}

	return callExpression.Function.TokenLiteral() == "unquote"
}

func convertObjectToASTNode(obj object.Object) node.Node {
	switch obj := obj.(type) {
	case *object.Integer:
		return &expression.IntegerLiteral{
			Token: token.Token{
				Type:    token.INT,
				Literal: fmt.Sprintf("%d", obj.Value),
			},
			Value: obj.Value,
		}
	case *object.Boolean:
		var t token.Token
		if obj.Value {
			t = token.Token{Type: token.TRUE, Literal: "true"}
		} else {
			t = token.Token{Type: token.FALSE, Literal: "false"}
		}
		return &expression.Boolean{Token: t, Value: obj.Value}
	case *object.Quote:
		return obj.Node
	default:
		return nil
	}
}
