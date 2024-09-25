package test

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/format"
	"cuelang.org/go/cue/token"
	"fmt"
	"github.com/mandoway/seru/languages/cue/language"
	"testing"
)

func TestManually(t *testing.T) {
	// uncomment following line to execute test
	t.Skip()

	parser := language.Parser{}
	bytes := []byte(buildSmallExample())

	content, err := parser.Parse(bytes)
	if err != nil {
		t.Fatal(err)
	}

	ctx := cuecontext.New()
	value := ctx.BuildFile(content)
	scope := cue.Scope(value)

	cmp := ctx.BuildExpr(ast.NewIdent("foo"))
	fmt.Println(cmp.Int64())
	cmp = ctx.BuildExpr(ast.NewIdent("bar"))
	fmt.Println(cmp.Int64())
	cmp = ctx.BuildExpr(ast.NewIdent("baz"))
	fmt.Println(cmp.String())

	cmp = ctx.BuildExpr(ast.NewBinExpr(token.GTR, ast.NewIdent("bar"), ast.NewLit(token.INT, "2")))
	fmt.Println(cmp.Bool())

	expr := content.Decls[2].(*ast.Comprehension).Clauses[0].(*ast.IfClause)
	cmp = ctx.BuildExpr(expr.Condition)
	fmt.Println(cmp.Bool())

	binExpr := expr.Condition.(*ast.BinaryExpr)
	ident := ast.NewIdent(binExpr.X.(*ast.Ident).Name)
	newCond := ast.NewBinExpr(binExpr.Op, ident, ast.NewLit(binExpr.Y.(*ast.BasicLit).Kind, binExpr.Y.(*ast.BasicLit).Value))
	cmp = ctx.BuildExpr(binExpr)
	fmt.Println(cmp.Bool())

	cmp = ctx.BuildExpr(newCond)
	fmt.Println(cmp.Bool())

	oneIdent := binExpr.X.(*ast.Ident)
	cmp = ctx.BuildExpr(binExpr, cue.Scope(ctx.BuildFile(oneIdent.Scope.(*ast.File))))
	fmt.Println(cmp.Bool())

	formatted, _ := format.Node(binExpr)
	cmp = ctx.CompileBytes(formatted, scope)
	fmt.Println(cmp.Bool())
}

func buildRealWorldExample() string {
	return "#B: {\nn: string\nv: {\nr: {\n}\n}\n}\n#F: {\nt: [...string]\n...\n}\n#N: {\nn: string\np: #F\nv: #V\n}\n#C: {\np: #F\nv: #V\n}\n#C2: {\nc: #C\nt: #N\nns: {\n[NS=string]: #N & {n: NS}\n}\ng: {\nt: {...} | *{}\n...\n}\n}\n#V: {\nlet l = {[string]: _}\nt: [string]: {}\nns: [string]: {}\nr: l\n}\nfs: f1: #C2 & {\nns: m: {\np: {}\nv: {\nr: {\nfor s in [\"e\"] {\n(L & {n: \"m\", sc: s}).j\n}\n}\n}\n}\nt: {\nNS=n: string\nv: {\nr: {\n\"\\(NS)/b\": {\nmetadata: {\nn: NS\n}\n}\n(L & {n: NS, sc: \"o\"}).j\n}\n}\n}\n}\nlet L = {\nNS=n: string\nsc:   string\nlet l = #B & {\nn: NS\n}\nj: {\nl.v.r\n}\n}\n"
}

func buildSmallExample() string {
	return `
foo: 3
bar: foo
if bar > 2 {
	baz: "hello"
}
`
}
