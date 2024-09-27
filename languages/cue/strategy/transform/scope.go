package transform

import (
	"cuelang.org/go/cue/ast/astutil"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/util/collection"
)

type ScopeLayer struct {
	Name       string
	Start, End token.Pos
}

type ScopeStack struct {
	collection.Stack[ScopeLayer]
}

func (s ScopeStack) String() string {
	if s.IsEmpty() {
		return ""
	}

	var selector string
	for layer := range s.Iterator() {
		if selector == "" {
			selector = layer.Name
		} else {
			selector += "." + layer.Name
		}
	}
	return selector
}

func (s ScopeStack) ScopeOfCurrentPosition(cursor astutil.Cursor) string {
	top, err := s.Top()
	if err != nil {
		return ""
	}

	if cursor.Node().Pos().Before(top.Start) || top.End.Before(cursor.Node().End()) {
		return ""
	}

	return s.String()
}
