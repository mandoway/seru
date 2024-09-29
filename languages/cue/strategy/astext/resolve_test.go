package astext

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/ast/astutil"
	"cuelang.org/go/cue/format"
	"github.com/mandoway/seru/languages/cue/language"
	"testing"
)

func TestResolveIdentifierInExpression_Field(t *testing.T) {
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
		"k": {
			name:     "value of let",
			expected: "2",
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
			actual := ResolveIdentifierValueInExpression(field.Value)
			formatted, _ := format.Node(actual)
			actualValue := string(formatted)

			if actualValue != instance.expected {
				t.Errorf("Wrong value, got %s, want %s", actualValue, instance.expected)
			}
		})

		return true
	}, nil)
}

func TestResolveIdentifierInExpression_LetClause(t *testing.T) {
	parsed, _ := language.Parser{}.Parse([]byte(getFile()))

	instancesByName := map[string]struct {
		name     string
		expected string
	}{
		"j": {
			name:     "let clause",
			expected: "2",
		},
	}
	astutil.Apply(parsed, func(cursor astutil.Cursor) bool {
		let, ok := cursor.Node().(*ast.LetClause)
		if !ok {
			return true
		}

		instance, ok := instancesByName[let.Ident.Name]
		if !ok {
			return true
		}

		t.Run(instance.name, func(t *testing.T) {
			actual := ResolveIdentifierValueInExpression(let.Expr)
			formatted, _ := format.Node(actual)
			actualValue := string(formatted)

			if actualValue != instance.expected {
				t.Errorf("Wrong value, got %s, want %s", actualValue, instance.expected)
			}
		})

		return true
	}, nil)
}

func TestResolveIdentifierInExpression_IfClause(t *testing.T) {
	parsed, _ := language.Parser{}.Parse([]byte(getFile()))

	instance := struct {
		name     string
		expected string
	}{
		name:     "if clause",
		expected: "2 < 2",
	}
	astutil.Apply(parsed, func(cursor astutil.Cursor) bool {
		ifClause, ok := cursor.Node().(*ast.IfClause)
		if !ok {
			return true
		}

		t.Run(instance.name, func(t *testing.T) {
			actual := ResolveIdentifierValueInExpression(ifClause.Condition)
			formatted, _ := format.Node(actual)
			actualValue := string(formatted)

			if actualValue != instance.expected {
				t.Errorf("Wrong value, got %s, want %s", actualValue, instance.expected)
			}
		})

		return true
	}, nil)
}

func TestResolveIdentifierInExpression_ForClause(t *testing.T) {
	parsed, _ := language.Parser{}.Parse([]byte(getFile()))

	instance := struct {
		name     string
		expected string
	}{
		name:     "for clause",
		expected: "[1, 2]",
	}
	astutil.Apply(parsed, func(cursor astutil.Cursor) bool {
		forClause, ok := cursor.Node().(*ast.ForClause)
		if !ok {
			return true
		}

		t.Run(instance.name, func(t *testing.T) {
			actual := ResolveIdentifierValueInExpression(forClause.Source)
			formatted, _ := format.Node(actual)
			actualValue := string(formatted)

			if actualValue != instance.expected {
				t.Errorf("Wrong value, got %s, want %s", actualValue, instance.expected)
			}
		})

		return true
	}, nil)
}
func TestResolveIdentifierInExpression_Alias(t *testing.T) {
	parsed, _ := language.Parser{}.Parse([]byte(getFile()))

	instance := struct {
		name     string
		expected string
	}{
		name:     "alias clause",
		expected: `"n\(2)"`,
	}
	astutil.Apply(parsed, func(cursor astutil.Cursor) bool {
		alias, ok := cursor.Node().(*ast.Alias)
		if !ok {
			return true
		}

		t.Run(instance.name, func(t *testing.T) {
			actual := ResolveIdentifierValueInExpression(alias.Expr)
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
		let j = foo
		k: j
		if foo < 2 {
			l: foo
		}
		for a in [1, foo] {
			m: foo
		}
		X="n\(foo)": X
	}
	`
}
