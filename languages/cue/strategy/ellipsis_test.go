package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestEllipsisReduction(t *testing.T) {
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
			Title: "Remove ellipsis from struct",
			Given: `
			foo: {
				bar: 3
				...
			}
			`,
			Expected: []string{
				`
				foo: {
					bar: 3
				}
				`,
			},
		},
		{
			Title: "Remove second ellipsis from struct",
			Given: `
			foo: {
				...
				...
			}
			`,
			Expected: []string{
				`
				foo: {
					...
				}
				`,
			},
		},
		{
			Title: "Close open list",
			Given: `
			foo: [1, 2, ...]
			`,
			Expected: []string{
				`
				foo: [1, 2]
				`,
			},
		},
		{
			Title: "Close open typed list",
			Given: `
			foo: [1, 2, ...int]
			`,
			Expected: []string{
				`
				foo: [1, 2]
				`,
			},
		},
	}

	test.TestReduction(t, instances, EllipsisReduction{})
}
