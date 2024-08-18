package language

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/parser"
)

type Parser struct{}

func (p Parser) Parse(fileContent []byte) (*ast.File, error) {
	return parser.ParseFile("", fileContent, parser.ParseComments)
}
