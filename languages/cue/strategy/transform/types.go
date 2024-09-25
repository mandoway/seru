package transform

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/ast/astutil"
)

type Transformation interface {
	ProcessNode(cursor astutil.Cursor) bool
}

type ReplacementTransformation struct {
	ReplacementNode ast.Node
}

func NewReplacementTransformation(replacementNode ast.Node) *ReplacementTransformation {
	return &ReplacementTransformation{ReplacementNode: replacementNode}
}

func (t *ReplacementTransformation) ProcessNode(cursor astutil.Cursor) bool {
	cursor.Replace(t.ReplacementNode)
	return true
}

type DeleteTransformation struct {
}

func NewDeleteTransformation() *DeleteTransformation {
	return &DeleteTransformation{}
}

func (t *DeleteTransformation) ProcessNode(cursor astutil.Cursor) bool {
	cursor.Delete()
	return true
}

type NoopTransformation struct {
}

func NewNoopTransformation() *NoopTransformation {
	return &NoopTransformation{}
}

func (t *NoopTransformation) ProcessNode(_ astutil.Cursor) bool {
	return false
}
