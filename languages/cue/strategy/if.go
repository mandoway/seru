package strategy

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/languages/cue/strategy/eval"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type IfReduction struct {
}

func (i IfReduction) Apply(input []byte) []*ast.File {
	evaluateBooleanExpr := eval.BuildBooleanEvaluator(input)

	evaluateAndSimplify := func(node *ast.IfClause) ast.Node {
		evaluatedValue, err := evaluateBooleanExpr(node.Condition)
		if err != nil {
			return nil
		}

		op := token.NEQ
		if evaluatedValue {
			op = token.EQL
		}

		selector := node.Condition
		if bin, ok := node.Condition.(*ast.BinaryExpr); ok {
			if _, ok := bin.X.(*ast.BasicLit); !ok {
				selector = bin.X
			} else if _, ok := bin.Y.(*ast.BasicLit); !ok {
				selector = bin.Y
			}
		}

		binaryExpr := ast.NewBinExpr(op, selector, selector)

		return &ast.IfClause{
			If:        node.If,
			Condition: binaryExpr,
		}
	}

	evaluateAndRemoveClause := func(node *ast.Comprehension) ast.Node {
		clause, ok := node.Clauses[0].(*ast.IfClause)
		if !ok {
			return nil
		}

		evaluatedValue, err := evaluateBooleanExpr(clause.Condition)

		if err != nil {
			return nil
		}

		if evaluatedValue {
			return node.Value
		} else {
			return ast.NewStruct()
		}
	}

	var transformations []*ast.File
	transformations = append(transformations, transform.ModifyApplicableStatements[*ast.IfClause](input, evaluateAndSimplify)...)
	transformations = append(transformations, transform.ModifyApplicableStatements[*ast.Comprehension](input, evaluateAndRemoveClause)...)
	return transformations
}
