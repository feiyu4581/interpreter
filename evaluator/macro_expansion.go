package evaluator

import (
	"interpreter/ast"
	"interpreter/ast/expression"
	"interpreter/ast/node"
	"interpreter/ast/statement"
	"interpreter/object"
)

func DefineMacros(program *ast.Program, env *object.Environment) {
	var definitions []int

	for i, stmt := range program.Statements {
		if isMacroDefinition(stmt) {
			addMacro(stmt, env)
			definitions = append(definitions, i)
		}
	}

	for i := len(definitions) - 1; i >= 0; i = i - 1 {
		definitionIndex := definitions[i]
		program.Statements = append(
			program.Statements[:definitionIndex],
			program.Statements[definitionIndex+1:]...,
		)
	}
}

func isMacroDefinition(node node.Statement) bool {
	letStatement, ok := node.(*statement.LetStatement)
	if !ok {
		return false
	}

	_, ok = letStatement.Value.(*expression.MacroLiteral)
	if !ok {
		return false
	}

	return true
}

func addMacro(stmt node.Statement, env *object.Environment) {
	letStatement, _ := stmt.(*statement.LetStatement)
	macroLiteral, _ := letStatement.Value.(*expression.MacroLiteral)

	macro := &object.Macro{
		Parameters: macroLiteral.Parameters,
		Env:        env,
		Body:       macroLiteral.Body,
	}

	env.Set(letStatement.Name.Value, macro)
}

// evaluator/macro_expansion.go

func ExpandMacros(program node.Node, env *object.Environment) node.Node {
	return ast.Modify(program, func(node node.Node) node.Node {
		callExpression, ok := node.(*expression.CallExpression)
		if !ok {
			return node
		}

		macro, ok := isMacroCall(callExpression, env)
		if !ok {
			return node
		}

		args := quoteArgs(callExpression)
		evalEnv := extendMacroEnv(macro, args)

		evaluated := Eval(macro.Body, evalEnv)

		quote, ok := evaluated.(*object.Quote)
		if !ok {
			panic("we only support returning AST-nodes from macros")
		}

		return quote.Node
	})
}

func isMacroCall(
	exp *expression.CallExpression,
	env *object.Environment,
) (*object.Macro, bool) {
	identifier, ok := exp.Function.(*expression.Identifier)
	if !ok {
		return nil, false
	}

	obj, ok := env.Get(identifier.Value)
	if !ok {
		return nil, false
	}

	macro, ok := obj.(*object.Macro)
	if !ok {
		return nil, false
	}

	return macro, true
}

func quoteArgs(exp *expression.CallExpression) []*object.Quote {
	var args []*object.Quote

	for _, a := range exp.Arguments {
		args = append(args, &object.Quote{Node: a})
	}

	return args
}

func extendMacroEnv(
	macro *object.Macro,
	args []*object.Quote,
) *object.Environment {
	extended := object.NewEnclosedEnvironment(macro.Env)

	for paramIdx, param := range macro.Parameters {
		extended.Set(param.Value, args[paramIdx])
	}

	return extended
}
