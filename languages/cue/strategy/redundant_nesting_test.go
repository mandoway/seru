package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestRedundantNesting(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title: "Not applicable",
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
			Title: "Removes redundant nesting",
			Given: `
			foo: {
				{
					bar: 3
				}
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
			Title: "Does not remove if more than 1 declaration",
			Given: `
			foo: {
				baz: 2
				{
					bar: 3
				}
			}
			`,
			Expected: []string{},
		},
		{
			Title: "Only removes one layer of nesting",
			Given: `
			foo: {
				{
					{
						bar: 3
					}
				}
			}
			`,
			Expected: []string{
				`
				foo: {
					{
						bar: 3
					}
				}
				`,
			},
		},
		{
			Title: "Returns candidate for each occurrence",
			Given: `
			foo: {
				{
					bar: 3
				}
			}
			a: {
				{
					b: 3
				}
			}
			`,
			Expected: []string{
				`
				foo: {
					bar: 3
				}
				a: {
					{
						b: 3
					}
				}
				`,
				`
				foo: {
					{
						bar: 3
					}
				}
				a: {
					b: 3
				}
				`,
			},
		},
	}

	test.TestReduction(t, instances, RedundantNestingReduction{})
}
