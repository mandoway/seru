package astext

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/ast/astutil"
	"cuelang.org/go/cue/format"
	"github.com/mandoway/seru/languages/cue/language"
	"testing"
)

func TestResolveIdentifierInExpression(t *testing.T) {
	parsed, _ := language.Parser{}.Parse([]byte(getFile()))

	instancesByName := map[string]struct {
		name     string
		expected string
	}{
		"a": {
			name:     "value",
			expected: "2",
		},
		"b": {
			name:     "interpolation",
			expected: `"\(2)"`,
		},
		"c": {
			name:     "binary expr",
			expected: "2 + 1",
		},
		"d": {
			name:     "unary expr",
			expected: "-2",
		},
		"e": {
			name:     "call expr",
			expected: "div(4, 2)",
		},
		"f": {
			name:     "index expr",
			expected: "arr[2]",
		},
		"g": {
			name:     "slice expr",
			expected: "arr[1 : 2+1]",
		},
		"h": {
			name:     "list expr",
			expected: "[1, 2]",
		},
		"i": {
			name:     "paren expr",
			expected: "(2)",
		},
	}

	astutil.Apply(parsed, func(cursor astutil.Cursor) bool {
		field, ok := cursor.Node().(*ast.Field)
		if !ok {
			return true
		}

		ident, ok := field.Label.(*ast.Ident)
		if !ok {
			return true
		}

		instance, ok := instancesByName[ident.Name]
		if !ok {
			return true
		}

		t.Run(instance.name, func(t *testing.T) {
			actual := ResolveIdentifierInExpression(field.Value)
			formatted, _ := format.Node(actual)
			actualValue := string(formatted)

			if actualValue != instance.expected {
				t.Errorf("Wrong value, got %s, want %s", actualValue, instance.expected)
			}
		})

		return true
	}, nil)
}

func getFile() string {
	return `
	foo: 1
	arr: ["a", "b", "c"]
	bar: {
		foo: 2
		a: foo      // 2
		b: "\(foo)" // "\(2)"
		c: foo + 1  // 2 + 1
		d: -foo     // -2
		e: div(4, foo) // div(4, 2)
		f: arr[foo]   // arr[2]
		g: arr[1:foo+1] // arr[1:2+1]
		h: [1, foo] // [1, 2]
		i: (foo) // (2)
	}
	`
}
