package strategy

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/ast/astutil"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/languages/cue/language"
	"github.com/mandoway/seru/languages/cue/strategy/astext"
	"github.com/mandoway/seru/util/collection"
	"slices"
)

type ImportReduction struct {
}

func (i ImportReduction) Apply(input []byte) []*ast.File {
	parser := language.Parser{}

	parsed, _ := parser.Parse(input)
	importPaths := collectImportPaths(parsed)

	var transformed []*ast.File
	for _, importPath := range importPaths {
		workingCopy, _ := parser.Parse(input)
		removeImportAndIdentifiers(workingCopy, importPath)
		transformed = append(transformed, workingCopy)
	}

	return transformed
}

func removeImportAndIdentifiers(workingCopy *ast.File, importPath string) {
	astutil.Apply(workingCopy, func(cursor astutil.Cursor) bool {
		switch node := cursor.Node().(type) {
		case *ast.Package:
			// skip handling of package clause
			return false

		case *ast.ImportDecl:
			indexOfPath := slices.IndexFunc(node.Specs, func(s *ast.ImportSpec) bool {
				return s.Path.Value == importPath
			})
			modifiedSpecs := collection.RemoveAt[*ast.ImportSpec](node.Specs, indexOfPath)

			if len(modifiedSpecs) == 0 {
				cursor.Delete()
			} else {
				cursor.Replace(astext.CopyImportDeclWithSpecs(node, modifiedSpecs))
			}

		case *ast.SelectorExpr:
			if ident, ok := node.X.(*ast.Ident); ok {
				replaceCurrentWithBlankIfImportedIdent(cursor, ident, importPath)
			}

		case *ast.Ident:
			replaceCurrentWithBlankIfImportedIdent(cursor, node, importPath)
		}

		return true
	}, nil)
}

func collectImportPaths(parsed *ast.File) []string {
	var imports []string
	astutil.Apply(parsed, func(cursor astutil.Cursor) bool {
		if importSpec, ok := cursor.Node().(*ast.ImportSpec); ok {
			imports = append(imports, importSpec.Path.Value)
		}

		return true
	}, nil)
	return imports
}

func replaceCurrentWithBlankIfImportedIdent(cursor astutil.Cursor, node *ast.Ident, importPath string) {
	if importSpec, ok := node.Node.(*ast.ImportSpec); ok && importSpec.Path.Value == importPath {
		cursor.Replace(ast.NewLit(token.IDENT, "_"))
	}
}
