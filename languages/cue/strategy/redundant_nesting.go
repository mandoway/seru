package strategy

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type RedundantNestingReduction struct {
}

func (r RedundantNestingReduction) Apply(input []byte) []*ast.File {
	removeOneNestedLayer := func(node *ast.StructLit) ast.Node {
		if len(node.Elts) != 1 {
			return nil
		}

		elementAsEmbedDecl, ok := node.Elts[0].(*ast.EmbedDecl)
		if !ok {
			return nil
		}

		elementExprAsStructLit, ok := elementAsEmbedDecl.Expr.(*ast.StructLit)
		if !ok {
			return nil
		}

		return &ast.StructLit{
			Lbrace: token.NoSpace.Pos(),
			Elts:   elementExprAsStructLit.Elts,
		}
	}
	return transform.ModifyApplicableStatements[*ast.StructLit](input, removeOneNestedLayer)
}
