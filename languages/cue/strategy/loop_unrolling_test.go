package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestLoopUnrolling(t *testing.T) {
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
			Expected: nil,
		},
		{
			Title: "literal array",
			Given: `
			for a in ["1", "2", "3"] {
				"a\(a)": a
			}
			`,
			Expected: []string{
				`
				{
					"a\("1")": "1"
					"a\("2")": "2"
					"a\("3")": "3"
				}
				`,
			},
		},
		{
			Title: "literal array + index",
			Given: `
			for i, a in ["1", "2", "3"] {
				"a\(i)": a
			}
			`,
			Expected: []string{
				`
				{
					"a\(0)": "1"
					"a\(1)": "2"
					"a\(2)": "3"
				}
				`,
			},
		},
		{
			Title: "identifier array",
			Given: `
			foo: ["1", "2", "3"]
			for i, a in foo {
				"a\(i)": a
			}
			`,
			Expected: []string{
				`
				foo: ["1", "2", "3"], {
					"a\(0)": "1"
					"a\(1)": "2"
					"a\(2)": "3"
				}
				`,
			},
		},
		{
			Title: "multiple decl",
			Given: `
			for i, a in ["1", "2", "3"] {
				"a\(i)": a
				"b\(i)": a
			}
			`,
			Expected: []string{
				`
				{
					"a\(0)": "1"
					"b\(0)": "1"
					"a\(1)": "2"
					"b\(1)": "2"
					"a\(2)": "3"
					"b\(2)": "3"
				}
				`,
			},
		},
		{
			Title: "struct lit",
			Given: `
			for k, v in { x: "1", y: "2" } {
				"\(k)": v
			}
			`,
			Expected: []string{
				`
				{
					"\(x)": "1"
					"\(y)": "2"
				}
				`,
			},
		},
		{
			Title: "struct ident",
			Given: `
			foo: {
				x: "1"
				y: "2"
			}
			for k, v in foo {
				"\(k)": v
			}
			`,
			Expected: []string{
				`
				foo: {
					x: "1"
					y: "2"
				}, {
					"\(x)": "1"
					"\(y)": "2"
				}
				`,
			},
		},
		{
			Title: "struct multi decl",
			Given: `
			for k, v in { x: "1", y: "2" } {
				"\(k)": v
				"a\(k)": v
			}
			`,
			Expected: []string{
				`
				{
					"\(x)":  "1"
					"a\(x)": "1"
					"\(y)":  "2"
					"a\(y)": "2"
				}
				`,
			},
		},
	}

	test.TestReduction(t, instances, LoopUnrollingReduction{})
}
