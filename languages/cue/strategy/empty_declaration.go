package strategy

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type EmptyDeclarationReduction struct{}

func (r EmptyDeclarationReduction) Apply(input []byte) []*ast.File {
	removeIfDeclIsNull := func(node *ast.EmbedDecl, _ string) transform.Transformation {
		if isNull(node) || isEmptyStruct(node) || isIdent(node) {
			return transform.NewDeleteTransformation()
		}

		return transform.NewNoopTransformation()
	}
	return transform.ApplyTransformationToEveryApplicableStatement[*ast.EmbedDecl](input, removeIfDeclIsNull)
}

func isNull(node *ast.EmbedDecl) bool {
	basicLit, ok := node.Expr.(*ast.BasicLit)
	return ok && basicLit.Kind == token.NULL
}

func isEmptyStruct(node *ast.EmbedDecl) bool {
	structLit, ok := node.Expr.(*ast.StructLit)
	return ok && len(structLit.Elts) == 0
}

func isIdent(node *ast.EmbedDecl) bool {
	_, ok := node.Expr.(*ast.Ident)
	return ok
}
