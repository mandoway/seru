package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestNullReduction(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title: "Return empty if not applicable",
			Given: `
			x: 42
			y: x
			z: y
			`,
			Expected: []string{},
		},
		{
			Title: "Remove null declaration",
			Given: `
			a: 5
			null
			b: null
			`,
			Expected: []string{
				`
				a: 5
				b: null
				`,
			},
		},
		{
			Title: "Remove null in nested declaration",
			Given: `
			foo: {
				a: 5
				null
			}
			`,
			Expected: []string{
				`
				foo: {
					a: 5
				}
				`,
			},
		},
		{
			Title: "Remove one occurrence at a time",
			Given: `
			null
			foo: 3
			null
			`,
			Expected: []string{
				`
				foo: 3
				null
				`, `
				null
				foo: 3
				`,
			},
		},
	}

	test.TestReduction(t, instances, NullReduction{})
}
