package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestUnificationReduction(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title: "not applicable",
			Given: `
			foo: 3
			bar: 2
			baz: {
				a: "hello"
			}
			`,
			Expected: nil,
		},
		{
			Title: "inline definition",
			Given: `
			foo: {
				a: number
				b: string
			} & {
				a: 3
				b: "b"
			}`,
			Expected: []string{
				`
				foo: {
					a: number
					b: string
					a: 3
					b: "b"
				}`,
			},
		},
		{
			Title: "variable definition",
			Given: `
			#foo: {
				a: number
				b: string
			}
			bar: #foo & {
				a: 3
				b: "b"
			}`,
			Expected: []string{
				`
				#foo: {
					a: number
					b: string
				}
				bar: {
					a: number
					b: string
					a: 3
					b: "b"
				}`,
			},
		},
		{
			Title: "double variable definition",
			Given: `
			#foo: {
				a: number
				b: string
			}
			#bar: {
				a: 3
				b: "b"
			}
			baz: #foo & #bar`,
			Expected: []string{
				`
				#foo: {
					a: number
					b: string
				}
				#bar: {
					a: 3
					b: "b"
				}
				baz: {
					a: number
					b: string
					a: 3
					b: "b"
				}`,
			},
		}, {
			Title: "empty definition",
			Given: `
			#foo: {
				a: number
				b: string
			}
			bar: #foo & {}`,
			Expected: []string{
				`
				#foo: {
					a: number
					b: string
				}
				bar: {
					a: number
					b: string
				}`,
			},
		},
	}

	test.TestReduction(t, instances, UnificationReduction{})
}
