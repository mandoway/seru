package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestImportReduction(t *testing.T) {
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
			Title: "single use",
			Given: `
			import "strings"
		
			foo: strings
			`,
			Expected: []string{
				`
				foo: _
				`,
			},
		},
		{
			Title: "multi use",
			Given: `
			import "strings"
		
			foo: strings
			strings
			`,
			Expected: []string{
				`
				foo: _
				_
				`,
			},
		},
		{
			Title: "multi imports decl",
			Given: `
			import (
				"strings"
				"regexp"
			)
		
			foo: strings
			regexp
			`,
			Expected: []string{
				`
				import (
					"regexp"
				)
		
				foo: _
				regexp
				`, `
				import (
					"strings"
				)
		
				foo: strings
				_
				`,
			},
		},
		{
			Title: "single import decl",
			Given: `
			import (
				"strings"
			)
		
			foo: strings
			`,
			Expected: []string{
				`
				foo: _
				`,
			},
		},
		{
			Title: "complex path import",
			Given: `
			import (
				"tool/cli"
			)
			
			foo: cli.Print
			`,
			Expected: []string{
				`
				foo: _
				`,
			},
		},
	}

	test.TestReduction(t, instances, ImportReduction{})
}
