package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestIfReduction(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title: "not applicable",
			Given: `
			foo: 4
			bar: foo
			baz: {
				a: foo
			}
			`,
			Expected: []string{},
		},
		{
			Title: "Creates tautology and removes clause",
			Given: `
			foo: 3
			if foo > 2 {
				bar: 2
			}
			`,
			Expected: []string{
				`
				foo: 3
				if foo == foo {
					bar: 2
				}
				`,
				`
				foo: 3, {
					bar: 2
				}
				`,
			},
		},
		{
			Title: "Creates contradiction and removes clause&struct",
			Given: `
			foo: 3
			if foo < 2 {
				bar: 2
			}
			`,
			Expected: []string{
				`
				foo: 3
				if foo != foo {
					bar: 2
				}
				`,
				`
				foo: 3, {}
				`,
			},
		},
		{
			Title: "Applies to selectors too",
			Given: `
			foo: {
				bar: 3
			}
			if foo.bar < 2 {
				baz: 2
			}
			`,
			Expected: []string{
				`
				foo: {
					bar: 3
				}
				if foo.bar != foo.bar {
					baz: 2
				}
				`,
				`
				foo: {
					bar: 3
				}, {}
				`,
			},
		},
	}

	test.TestReduction(t, instances, IfReduction{})
}
