package astext

import (
	"cuelang.org/go/cue/ast"
)

func CopyFieldWithLabel(field *ast.Field, newLabel ast.Label) *ast.Field {
	return &ast.Field{
		Label:      newLabel,
		Optional:   field.Optional,
		Constraint: field.Constraint,
		TokenPos:   field.TokenPos,
		Token:      field.Token,
		Value:      field.Value,
		Attrs:      field.Attrs,
	}
}

func CopyFieldWithValue(field *ast.Field, newValue ast.Expr) *ast.Field {
	return &ast.Field{
		Label:      field.Label,
		Optional:   field.Optional,
		Constraint: field.Constraint,
		TokenPos:   field.TokenPos,
		Token:      field.Token,
		Value:      newValue,
		Attrs:      field.Attrs,
	}
}

func CopyLetClauseWithValue(let *ast.LetClause, value ast.Expr) *ast.LetClause {
	return &ast.LetClause{
		Let:   let.Let,
		Ident: let.Ident,
		Equal: let.Equal,
		Expr:  value,
	}
}

func CopyUnaryExpression(expr *ast.UnaryExpr, value ast.Expr) *ast.UnaryExpr {
	return &ast.UnaryExpr{
		OpPos: expr.OpPos,
		Op:    expr.Op,
		X:     value,
	}
}

func CopyIndexExpression(expr *ast.IndexExpr, value ast.Expr) *ast.IndexExpr {
	return &ast.IndexExpr{
		X:      expr.X,
		Lbrack: expr.Lbrack,
		Index:  value,
		Rbrack: expr.Rbrack,
	}
}

func CopySliceExpression(expr *ast.SliceExpr, low, high ast.Expr) *ast.SliceExpr {
	return &ast.SliceExpr{
		X:      expr.X,
		Lbrack: expr.Lbrack,
		Low:    low,
		High:   high,
		Rbrack: expr.Rbrack,
	}
}

func CopyParenExpression(expr *ast.ParenExpr, value ast.Expr) *ast.ParenExpr {
	return &ast.ParenExpr{
		Lparen: expr.Lparen,
		X:      value,
		Rparen: expr.Rparen,
	}
}
