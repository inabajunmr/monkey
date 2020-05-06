package evaluator

import (
	"github.com/inabajunmr/monkey/ast"
	"github.com/inabajunmr/monkey/object"
)

func quote(node ast.Expression) object.Object {
	return &object.Quote{Node: node}
}
