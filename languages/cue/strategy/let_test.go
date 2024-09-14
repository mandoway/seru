package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestLetReduction(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title: "Return candidate for every occurrence",
			Given: `
			x: 42
			let y = x
			let z = x
			foo: {
				let bar = x
				baz: bar
			}
			`,
			Expected: []string{
				`
				x: 42
				y: x
				let z = x
				foo: {
					let bar = x
					baz: bar
				}
				`,
				`
				x: 42
				let y = x
				z: x
				foo: {
					let bar = x
					baz: bar
				}
				`,
				`
				x: 42
				let y = x
				let z = x
				foo: {
					bar: x
					baz: bar
				}
				`,
			},
		},
		{
			Title: "Return empty if not applicable",
			Given: `
			x: 42
			y: x
			z: y
			`,
			Expected: []string{},
		},
	}

	test.TestReduction(t, instances, LetReduction{})
}
