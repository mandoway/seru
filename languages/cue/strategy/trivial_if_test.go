package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestTrivialIfReduction(t *testing.T) {
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
			Title: "applies on top-level",
			Given: `
			foo: 3
			if foo < 2 {
				bar: 2
			}
			`,
			Expected: []string{
				`
				foo: 3
				if false {
					bar: 2
				}
				`,
				`
				foo: 3
				if true {
					bar: 2
				}
				`,
			},
		},
	}

	test.TestReduction(t, instances, TrivialIfReduction{})
}
