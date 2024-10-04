package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
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
		}, {
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
	}

	test.TestReduction(t, instances, ConstantPropagationReduction{})
}

func TestConstantPropagationRealWorld(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title: "full file",
			Given: "#B:\n{}\n{}\n{\n{\n{\nfor s in []{\nL\n}\n}\n}\n{\nNS : string\nv:{\n{\n\"\\(NS)/b\":\nL\n}\n}\n}\n}\nlet L=\n#B\n",
			Expected: []string{
				"#B: {}\n{}\n{\n{\n{\nfor s in [] {\n{}\n}\n}\n}\n{\nNS: string\nv: {\n{\n\"\\(NS)/b\":\nL\n}\n}\n}\n}\nlet L = #B",
				"#B:\n{}\n{}\n{\n{\n{\nfor s in [] {\nL\n}\n}\n}\n{\nNS: string\nv: {\n{\n\"\\(string)/b\":\nL\n}\n}\n}\n}\nlet L =\n#B",
				"#B: {}\n{}\n{\n{\n{\nfor s in [] {\nL\n}\n}\n}\n{\nNS: string\nv: {\n{\n\"\\(NS)/b\": {}\n}\n}\n}\n}\nlet L = #B",
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
	}

	test.TestReduction(t, instances, ConstantPropagationReduction{})
}
