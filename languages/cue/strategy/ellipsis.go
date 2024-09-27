package strategy

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
	"github.com/mandoway/seru/util/collection"
	"slices"
)

type EllipsisReduction struct {
}

func (e EllipsisReduction) Apply(input []byte) []*ast.File {
	removeEllipsisFromStruct := func(node *ast.StructLit, _ string) transform.Transformation {
		indexOfEllipsis := slices.IndexFunc(node.Elts, func(decl ast.Decl) bool {
			_, isEllipsis := decl.(*ast.Ellipsis)
			return isEllipsis
		})

		if indexOfEllipsis < 0 {
			return transform.NewNoopTransformation()
		}

		prunedDecls := collection.RemoveAt(node.Elts, indexOfEllipsis)
		prunedDeclsInterface := collection.MapToInterface(prunedDecls)
		return transform.NewReplacementTransformation(
			ast.NewStruct(prunedDeclsInterface...),
		)
	}

	removeEllipsisFromList := func(node *ast.ListLit, _ string) transform.Transformation {
		indexOfEllipsis := slices.IndexFunc(node.Elts, func(expr ast.Expr) bool {
			_, isEllipsis := expr.(*ast.Ellipsis)
			return isEllipsis
		})

		if indexOfEllipsis < 0 {
			return transform.NewNoopTransformation()
		}

		return transform.NewReplacementTransformation(
			ast.NewList(collection.RemoveAt(node.Elts, indexOfEllipsis)...),
		)
	}

	var transformations []*ast.File
	transformations = append(transformations, transform.ApplyTransformationToEveryApplicableStatement[*ast.StructLit](input, removeEllipsisFromStruct)...)
	transformations = append(transformations, transform.ApplyTransformationToEveryApplicableStatement[*ast.ListLit](input, removeEllipsisFromList)...)
	return transformations
}
