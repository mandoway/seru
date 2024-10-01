package strategy

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/ast/astutil"
	"cuelang.org/go/cue/format"
	"cuelang.org/go/cue/parser"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/languages/cue/strategy/eval"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
	"strconv"
)

type LoopUnrollingReduction struct {
}

func (l LoopUnrollingReduction) Apply(input []byte) []*ast.File {
	evaluate := eval.BuildEvaluator(input)

	return transform.ApplyTransformationToEveryApplicableStatement[*ast.Comprehension](input, func(node *ast.Comprehension, _ string) transform.Transformation {
		if len(node.Clauses) == 0 {
			return transform.NewNoopTransformation()
		}
		forClause, ok := node.Clauses[0].(*ast.ForClause)
		if !ok {
			return transform.NewNoopTransformation()
		}

		indexIdent := forClause.Key
		valueIdent := forClause.Value
		iterations, err := getIterationsFromSource(forClause.Source, evaluate)
		if err != nil {
			return transform.NewNoopTransformation()
		}
		body, ok := node.Value.(*ast.StructLit) // has to be a struct per definition
		if !ok {
			return transform.NewNoopTransformation()
		}

		unrolledStruct := ast.NewStruct()
		for i, iteration := range iterations {
			curBody := replaceAllLoopIdentifiers(body, indexIdent, valueIdent, i, iteration)
			unrolledStruct.Elts = append(unrolledStruct.Elts, curBody.Elts...)
		}

		return transform.NewReplacementTransformation(unrolledStruct)
	})
}

func replaceAllLoopIdentifiers(body *ast.StructLit, indexIdent *ast.Ident, valueIdent *ast.Ident, index int, entry ast.Node) *ast.StructLit {
	localBody := copyStruct(body)
	replacedNode := astutil.Apply(localBody, func(cursor astutil.Cursor) bool {
		if ident, ok := cursor.Node().(*ast.Ident); ok {
			var keyReplacement, valueReplacement ast.Node
			if field, ok := entry.(*ast.Field); ok {
				keyReplacement = field.Label
				valueReplacement = field.Value
			} else {
				keyReplacement = ast.NewLit(token.INT, strconv.Itoa(index))
				valueReplacement = entry
			}

			switch {
			case indexIdent != nil && indexIdent.Name == ident.Name:
				ast.SetRelPos(keyReplacement, token.NoRelPos)
				cursor.Replace(keyReplacement)
			case valueIdent.Name == ident.Name:
				ast.SetRelPos(valueReplacement, token.NoRelPos)
				cursor.Replace(valueReplacement)
			}
		}

		return true
	}, nil)

	return replacedNode.(*ast.StructLit)
}

func copyStruct(lit *ast.StructLit) *ast.StructLit {
	src, _ := format.Node(lit)
	expr, _ := parser.ParseExpr("", src)
	return expr.(*ast.StructLit)
}

func getIterationsFromSource(source ast.Expr, evaluate func(expr ast.Expr) (cue.Value, error)) ([]ast.Node, error) {
	evaluated, err := evaluate(source)
	if err != nil {
		return nil, err
	}

	iterator, err := evaluated.List()
	if err == nil {
		var nodes []ast.Node
		for iterator.Next() {
			nodes = append(nodes, iterator.Value().Source())
		}
		return nodes, nil
	}

	structLit, err := evaluated.Fields()
	if err == nil {
		var nodes []ast.Node
		for structLit.Next() {
			nodes = append(nodes, structLit.Value().Source())
		}
		return nodes, nil
	}

	return nil, err
}
