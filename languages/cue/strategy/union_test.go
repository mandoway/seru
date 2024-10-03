package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestUnionReduction_Apply(t *testing.T) {
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
			Expected: nil,
		},
		{
			Title: "union types",
			Given: `
			foo: "a" | "b"
			`,
			Expected: []string{
				`
				foo: "a"
				`,
				`
				foo: "b"
				`,
			},
		},
		{
			Title: "union default",
			Given: `
			foo: "a" | *"b"
			`,
			Expected: []string{
				`
				foo: "a"
				`,
				`
				foo: "b"
				`,
			},
		},
		{
			Title: "3 options",
			Given: `
			foo: "a" | "b" | "c"
			`,
			Expected: []string{
				`
				foo: "a"
				`,
				`
				foo: "b"
				`,
				`
				foo: "c"
				`,
			},
		},
		{
			Title: "identifier options",
			Given: `
			bar: 3
			foo: bar | baz
			`,
			Expected: []string{
				`
				bar: 3
				foo: 3
				`,
				`
				bar: 3
				foo: baz
				`,
			},
		},
		{
			Title: "unary expr unchanged",
			Given: `
			foo: 1 | +2 | -3 | !4 | !=5 | <6 | <=7 | >8 | >=9 | =~10 | !~11 
			`,
			Expected: []string{
				`
				foo: 1
				`,
				`
				foo: +2
				`,
				`
				foo: -3
				`,
				`
				foo: !4
				`,
				`
				foo: !=5
				`,
				`
				foo: <6
				`,
				`
				foo: <=7
				`,
				`
				foo: >8
				`,
				`
				foo: >=9
				`,
				`
				foo: =~10
				`,
				`
				foo: !~11
				`,
			},
		},
		{
			Title: "only affects union",
			Given: `
			a: 1 + 2 
			b: 1 - 2 
			c: 1 * 2 
			d: 1 / 2 
			e: 1 & 2 
			f: 1 || 2 
			g: 1 && 2 
			h: 1 != 2 
			i: 1 < 2 
			j: 1 <= 2 
			k: 1 > 2 
			l: 1 >= 2 
			m: 1 =~ 2 
			n: 1 !~ 2 
			`,
			Expected: nil,
		},
		{
			Title: "affects correct expr",
			Given: `
			a: 1 | 2
			b: "3" | "4" 
			`,
			Expected: []string{
				`
				a: 1
				b: "3" | "4" 
				`,
				`
				a: 1 | 2
				b: "3"
				`,
				`
				a: 2
				b: "3" | "4" 
				`,
				`
				a: 1 | 2
				b: "4" 
				`,
			},
		},
	}

	test.TestReduction(t, instances, UnionReduction{})
}
