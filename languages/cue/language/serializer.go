package language

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/format"
)

type Serializer struct {
}

func (s Serializer) Serialize(content *ast.File) ([]byte, error) {
	return format.Node(content)
}
