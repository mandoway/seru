package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestStringInterpolation(t *testing.T) {
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
			Title: "field value",
			Given: `
			foo: 3
			bar: "\(foo)"
			`,
			Expected: []string{
				`
				foo: 3
				bar: "3"
				`,
			},
		},
		{
			Title: "field label",
			Given: `
			foo: 3
			"bar\(foo)": "hello"
			`,
			Expected: []string{
				`
				foo:  3
				bar3: "hello"
				`,
			},
		},
		{
			Title: "let value",
			Given: `
			foo: 3
			let bar = "\(foo+1)"
			baz: bar
			`,
			Expected: []string{
				`
				foo: 3
				let bar = "4"
				baz: bar
				`,
			},
		},
		{
			Title: "calc",
			Given: `
			foo: "\(5 - 3)"
			`,
			Expected: []string{
				`
				foo: "2"
				`,
			},
		},
		{
			Title: "scoped",
			Given: `
			bar: 2
			baz: "\(bar)"
			a: {
				bar: 3
				baz: "\(bar)"
			}
			`,
			Expected: []string{
				`
				bar: 2
				baz: "2"
				a: {
					bar: 3
					baz: "\(bar)"
				}
				`, `
				bar: 2
				baz: "\(bar)"
				a: {
					bar: 3
					baz: "3"
				}
				`,
			},
		},
		{
			Title: "for is not applicable",
			Given: `
			foo: 3
			bar: [for i in ["a", "b"] { "\(i)_\(foo)" }]
			`,
			Expected: nil,
		},
		{
			Title: "two interpolations",
			Given: `
			foo: 3
			bar: "a"
			baz: "\(bar)_\(foo)"
			`,
			Expected: []string{
				`
				foo: 3
				bar: "a"
				baz: "a_3"
				`,
			},
		},
	}

	test.TestReduction(t, instances, StringInterpolationReduction{})
}
