package strategy

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type ListReduction struct {
}

func (l ListReduction) Apply(input []byte) []*ast.File {
	reduceToEmpty := func(node *ast.ListLit) ast.Node {
		if len(node.Elts) == 0 {
			return nil
		}

		return ast.NewList()
	}
	return transform.ModifyApplicableStatements[*ast.ListLit](input, reduceToEmpty)
}
