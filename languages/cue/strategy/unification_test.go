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
		{
			Title: "Issue 2246 label not matched with expression",
			Given: `
			#FormFoo: fooID: fooID
			data: #Form
			#Form: {
				{
					fooID: fooID
				} | {
					fooID: fooID
				}
			}
			data: fooID: "123"
			data: data
			{#FormFoo | #FormFoo} & data
			`,
			Expected: []string{
				`
				#FormFoo: fooID: fooID
				data: #Form
				#Form: {
					{
						fooID: fooID
					} | {
						fooID: fooID
					}
				}
				data: fooID: "123"
				data: data, {
					fooID: fooID
					fooID: "123"
				}
				`,
			},
		},
	}

	test.TestReduction(t, instances, UnificationReduction{})
}

func TestUnificationRealWorld(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title: "Null declaration",
			Given: `
			#Abstract
			spec: bar: {}
			#Abstract: X={
			 spec: _#Spec
			 resource: _Thing & {_X: spec: X.spec}
			}
			_#Spec: *_#SpecFoo | _#SpecBar
			_#SpecFoo:
			{
			 min: 10
			 max: 20
			}
			_#SpecBar: bar: {
			 min: int
			 max: int
			}
			_Thing: #Constrained & {
			 _X: _
			 spec: {
			  spec
			  if _X.spec.bar != _|_ {
			   minBar: _X.spec.bar.min
			   maxBar: _X.spec.bar.max
			  }
			 }
			}
			#Constrained: {
			 spec: {
			  int
			  maxFoo
			  null
			  null
			 } | {
			  minBar: 30
			  maxBar: 40
			  minFoo: null
			  maxFoo: null
			 }
			 spec: {
			  fuga: null
			 } |
			  null
			}
			maxFoo: null
			`,
			Expected: nil,
		},
	}

	test.TestReduction(t, instances, UnificationReduction{})
}
