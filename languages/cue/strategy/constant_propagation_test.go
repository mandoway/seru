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
	}

	test.TestReduction(t, instances, ConstantPropagationReduction{})
}

func TestConstantPropagationRealWorld(t *testing.T) {
	instance := test.ReductionInstance{
		Title: "3 variants",
		Given: "#B:\n{}\n{}\n{\n{\n{\nfor s in []{\nL\n}\n}\n}\n{\nNS : string\nv:{\n{\n\"\\(NS)/b\":\nL\n}\n}\n}\n}\nlet L=\n#B\n",
		Expected: []string{
			"#B: {}\n{}\n{\n{\n{\nfor s in [] {\n{}\n}\n}\n}\n{\nNS: string\nv: {\n{\n\"\\(NS)/b\":\nL\n}\n}\n}\n}\nlet L = #B",
			"#B: {}\n{}\n{\n{\n{\nfor s in [] {\nL\n}\n}\n}\n{\nNS: string\nv: {\n{\n\"\\(NS)/b\": {}\n}\n}\n}\n}\nlet L = #B",
			"#B: {}\n{}\n{\n{\n{\nfor s in [] {\nL\n}\n}\n}\n{\nNS: string\nv: {\n{\n\"\\(NS)/b\":\nL\n}\n}\n}\n}\nlet L = {}",
		},
	}

	test.TestReduction(t, []test.ReductionInstance{instance}, ConstantPropagationReduction{})
}
