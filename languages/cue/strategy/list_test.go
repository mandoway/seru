package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestListReduction(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title: "Not applicable",
			Given: `
			foo: 3
			bar: "hello"
			baz: foo
			x: {
				a: foo
				b: []
			}
			`,
			Expected: []string{},
		},
		{
			Title: "Reduces arrays",
			Given: `
			x: {
				a: 3
				b: [1, 2, 3]
			}
			`,
			Expected: []string{
				`
				x: {
					a: 3
					b: []
				}
				`,
			},
		},
		{
			Title: "Reduces arrays in for loop",
			Given: `
			x: {
				for i in [1, 2, 3] {
					foo: i
				}
			}
			`,
			Expected: []string{
				`
				x: {
					for i in [] {
						foo: i
					}
				}
				`,
			},
		},
		{
			Title: "Reduces all occurrences",
			Given: `
			x: {
				for i in [1, 2, 3] {
					foo: i
				}
				bar: [4, 5, 6]
			}
			`,
			Expected: []string{
				`
				x: {
					for i in [] {
						foo: i
					}
					bar: [4, 5, 6]
				}	
				`,
				`
				x: {
					for i in [1, 2, 3] {
						foo: i
					}
					bar: []
				}	
				`,
			},
		},
		{
			Title: "Reduces list comprehensions",
			Given: `
			foo: [for i in [1, 2, 3] {i * 2}]
			`,
			Expected: []string{
				`
				foo: []
				`,
				`
				foo: [for i in [] {i * 2}]
				`,
			},
		},
	}

	test.TestReduction(t, instances, ListReduction{})
}
