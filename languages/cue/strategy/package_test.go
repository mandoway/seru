package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestPackageReduction(t *testing.T) {
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
			Title: "Removes package declaration",
			Given: `
			package foo
			
			bar: 42
			`,
			Expected: []string{
				`
				bar: 42
				`,
			},
		},
	}

	test.TestReduction(t, instances, PackageReduction{})
}
