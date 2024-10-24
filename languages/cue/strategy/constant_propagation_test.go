package strategy

import (
	"cuelang.org/go/cue/ast"
	"fmt"
	"github.com/mandoway/seru/languages/cue/language"
	"github.com/mandoway/seru/languages/cue/test"
	"github.com/mandoway/seru/util/collection"
	"testing"
)

func TestConstantPropagation(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title: "not applicable",
			Given: `
			foo: 3
			bar: 2
			baz: {
				a: "hello"
			}
			`,
			Expected: []string{},
		}, {
			Title: "top level",
			Given: `
			foo: 3
			bar: foo
			`,
			Expected: []string{
				`
				foo: 3
				bar: 3
				`,
			},
		},
		{
			Title: "top level let",
			Given: `
			let foo = 3
			bar: foo
			`,
			Expected: []string{
				`
				let foo = 3
				bar: 3
				`,
			},
		},
		{
			Title: "correct scope",
			Given: `
			foo: 3
			bar: foo
			baz: {
				foo:  1
				bazz: foo
			}
			`,
			Expected: []string{
				`
				foo: 3
				bar: 3
				baz: {
					foo:  1
					bazz: foo
				}
				`,
				`
				foo: 3
				bar: foo
				baz: {
					foo:  1
					bazz: 1
				}
				`,
			},
		},
		{
			Title: "multi-level",
			Given: `
			foo: 3
			bar: foo
			baz: bar
			`,
			Expected: []string{
				`
				foo: 3
				bar: 3
				baz: bar
				`, `
				foo: 3
				bar: foo
				baz: 3
				`,
			},
		},
		{
			Title: "struct",
			Given: `
			foo: {
				a: "hello"
			}
			bar: foo
			`,
			Expected: []string{
				`
				foo: {
					a: "hello"
				}
				bar: {
					a: "hello"
				}
				`,
			},
		},
		{
			Title: "list",
			Given: `
			foo: [1, 2, 3]
			bar: foo
			`,
			Expected: []string{
				`
				foo: [1, 2, 3]
				bar: [1, 2, 3]
				`,
			},
		},
		{
			Title: "types",
			Given: `
			foo: string
			bar: foo
			`,
			Expected: []string{
				`
				foo: string
				bar: string
				`,
			},
		},
		{
			Title: "let",
			Given: `
			foo: 3
			let bar = foo
			`,
			Expected: []string{
				`
				foo: 3
				let bar = 3
				`,
			},
		},
		{
			Title: "for",
			Given: `
			foo: 3
			for a in [foo] {}
			`,
			Expected: []string{
				`
				foo: 3
				for a in [3] {}
				`,
			},
		},
		{
			Title: "if",
			Given: `
			foo: 3
			if foo < 2 {}
			`,
			Expected: []string{
				`
				foo: 3
				if 3 < 2 {}
				`,
			},
		},
		{
			Title: "interpolation",
			Given: `
			foo: 3
			bar: "\(foo)"
			`,
			Expected: []string{
				`
				foo: 3
				bar: "\(3)"
				`,
			},
		},
		{
			Title: "one known one unknown",
			Given: `
			foo: 3
			bar: foo + baz
			`,
			Expected: []string{
				`
				foo: 3
				bar: 3 + baz
				`,
			},
		},
		{
			Title: "recursive definition",
			Given: `
			foo: "\(foo)": 3
			`,
			Expected: nil,
		},
		{
			Title: "ping pong recursion",
			Given: `
			foo: bar
			bar: {
				a: 3
				b: foo
				c: bar
			}
			`,
			Expected: []string{
				`
				foo: {
					a: 3
					b: foo
					c: bar
				}
				bar: {
					a: 3
					b: foo
					c: bar
				}`,
			},
		},
		{
			Title: "no dual recursion",
			Given: `
			foo: {
				a: bar
			}
			bar: {
				b: foo
			}
			`,
			Expected: nil,
		},
		{
			Title: "no self recursion",
			Given: `
			foo: {
				a: foo
			}
			`,
			Expected: nil,
		},
		{
			Title: "import as identifier ",
			Given: `
			import "strings"
		
			foo: strings.MinRunes(4)
			`,
			Expected: nil,
		},
		{
			Title: "alias field expression",
			Given: `
			foo: {
				bar: {
					baz: 2
				}
			}
			X=foo: _
			test: X.bar
			`,
			Expected: []string{
				`
				foo: {
					bar: {
						baz: 2
					}
				}
				X=foo: _
				test:  foo.bar
				`,
			},
		},
	}

	test.TestReduction(t, instances, ConstantPropagationReduction{})
}

func TestConstantPropagationRealWorld(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title: "full file",
			Given: "#B:\n{}\n{}\n{\n{\n{\nfor s in []{\nL\n}\n}\n}\n{\nNS : string\nv:{\n{\n\"\\(NS)/b\":\nL\n}\n}\n}\n}\nlet L=\n#B\n",
			Expected: []string{
				"#B: {}\n{}\n{\n{\n{\nfor s in [] {\n{}\n}\n}\n}\n{\nNS: string\nv: {\n{\n\"\\(NS)/b\":\nL\n}\n}\n}\n}\nlet L =\n#B",
				"#B:\n{}\n{}\n{\n{\n{\nfor s in [] {\nL\n}\n}\n}\n{\nNS: string\nv: {\n{\n\"\\(string)/b\":\nL\n}\n}\n}\n}\nlet L =\n#B",
				"#B: {}\n{}\n{\n{\n{\nfor s in [] {\nL\n}\n}\n}\n{\nNS: string\nv: {\n{\n\"\\(NS)/b\": {}\n}\n}\n}\n}\nlet L =\n#B",
				"#B: {}\n{}\n{\n{\n{\nfor s in [] {\nL\n}\n}\n}\n{\nNS: string\nv: {\n{\n\"\\(NS)/b\":\nL\n}\n}\n}\n}\nlet L = {}",
			},
		},
		{
			Title: "several definitions",
			Given: "#V\n#C:\n #V\n#C\n#V:\n _",
			Expected: []string{
				"_\n#C:\n#V\n#C\n#V: _",
				"#V\n#C: _\n#C\n#V: _",
				"#V\n#C:\n#V\n_\n#V: _",
			},
		},
		{
			Title: "recursive",
			Given: "#B:\n{\n for s in [\"e\"] {\n  L\n }\n NS: \"\\(NS)/b\":\n  L\n}\nlet L =\n#B",
			Expected: []string{
				"#B:\n{\nfor s in [\"e\"] {\n#B\n}\nNS: \"\\(NS)/b\":\nL\n}\nlet L = #B",
				"#B:\n{\nfor s in [\"e\"] {\nL\n}\nNS: \"\\(NS)/b\": #B\n}\nlet L = #B",
			},
		},
	}

	test.TestReduction(t, instances, ConstantPropagationReduction{})
}

func TestRecursiveProperties(t *testing.T) {
	instance := `
	foo: github

	github: resource_type: [cue_resource_name=github]:
	 "\({target.terraform.#Identifier.adapt & {#in: cue_resource_name}}.#out)"
	`

	candidates := getCandidates(instance)
	fmt.Printf("got candidates: %d\n", len(candidates))
	for len(candidates) > 0 {
		first := candidates[0]
		fmt.Println(first)
		candidates = getCandidates(first)
		fmt.Printf("got candidates: %d\n", len(candidates))
	}
}

func getCandidates(input string) []string {
	reducer := ConstantPropagationReduction{}
	serializer := language.Serializer{}

	candidates := reducer.Apply([]byte(input))
	return collection.MapSlice[*ast.File](candidates, func(it *ast.File) (string, error) {
		bytes, _ := serializer.Serialize(it)
		return string(bytes), nil
	})
}
