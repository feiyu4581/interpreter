package object

import "interpreter/ast/node"

type Quote struct {
	Node node.Node
}

func (q *Quote) Type() ObjectType {
	return QUOTE_OBJ
}

func (q *Quote) Inspect() string {
	return "QUOTE(" + q.Node.String() + ")"
}
