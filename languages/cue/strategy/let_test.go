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
			`,
			Expected: []string{
				`
				x: 42
				y: x
				let z = x
				`,
				`
				x: 42
				let y = x
				z: x
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
