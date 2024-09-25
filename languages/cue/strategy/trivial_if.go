package strategy

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type TrivialIfReduction struct{}

func (i TrivialIfReduction) Apply(input []byte) []*ast.File {
	alwaysFalse := func(node *ast.IfClause) ast.Node {
		return &ast.IfClause{
			If:        node.If,
			Condition: ast.NewBool(false),
		}
	}
	alwaysTrue := func(node *ast.IfClause) ast.Node {
		return &ast.IfClause{
			If:        node.If,
			Condition: ast.NewBool(true),
		}
	}

	var transformations []*ast.File
	transformations = append(transformations, transform.ModifyApplicableStatements[*ast.IfClause](input, alwaysFalse)...)
	transformations = append(transformations, transform.ModifyApplicableStatements[*ast.IfClause](input, alwaysTrue)...)
	return transformations
}
