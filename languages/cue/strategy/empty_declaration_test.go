package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestEmptyDeclarationReduction(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title: "Return empty if not applicable",
			Given: `
			x: 42
			y: x
			z: y
			`,
			Expected: []string{},
		},
		{
			Title: "Remove null declaration",
			Given: `
			a: 5
			null
			b: null
			`,
			Expected: []string{
				`
				a: 5
				b: null
				`,
			},
		},
		{
			Title: "Remove null in nested declaration",
			Given: `
			foo: {
				a: 5
				null
			}
			`,
			Expected: []string{
				`
				foo: {
					a: 5
				}
				`,
			},
		},
		{
			Title: "Remove one occurrence at a time",
			Given: `
			null
			foo: 3
			null
			`,
			Expected: []string{
				`
				foo: 3
				null
				`, `
				null
				foo: 3
				`,
			},
		},
		{
			Title: "Won't affect expressions",
			Given: `
			foo: 3
			bar: foo
			`,
			Expected: []string{},
		},
		{
			Title: "Won't affect non-null embedDecl",
			Given: `
			3
			"hello"
			`,
			Expected: []string{},
		},
		{
			Title: "empty struct",
			Given: `
			foo: 3
			{}
			`,
			Expected: []string{
				`
				foo: 3
				`,
			},
		},
		{
			Title: "empty struct nested",
			Given: `
			foo: {
				bar: 3
				{}
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
			Title: "top",
			Given: `
			foo: 3
			_
			`,
			Expected: []string{
				`
				foo: 3
				`,
			},
		},
		{
			Title: "top nested",
			Given: `
			foo: {
				bar: 3
				_
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
			Title: "ident",
			Given: `
			foo: 3
			a
			`,
			Expected: []string{
				`
				foo: 3
				`,
			},
		},
		{
			Title: "ident nested",
			Given: `
			foo: {
				bar: 3
				a
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
	}

	test.TestReduction(t, instances, EmptyDeclarationReduction{})
}

func TestRealWorldExampleEmptyDeclarationReduction(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title:    "Real world example",
			Given:    "#B: {\nn: string\nv: {\nr: {\n}\n}\n}\n#F: {\nt: [...string]\n...\n}\n#N: {\nn: string\np: #F\nv: #V\n}\n#C: {\np: #F\nv: #V\n}\n#C2: {\nc: #C\nt: #N\nns: {\n[NS=string]: #N & {n: NS}\n}\ng: {\nt: {...} | *{}\n...\n}\n}\n#V: {\nlet l = {[string]: _}\nt: [string]: {}\nns: [string]: {}\nr: l\n}\nfs: f1: #C2 & {\nns: m: {\np: {}\nv: {\nr: {\nfor s in [\"e\"] {\n(L & {n: \"m\", sc: s}).j\n}\n}\n}\n}\nt: {\nNS=n: string\nv: {\nr: {\n\"\\(NS)/b\": {\nmetadata: {\nn: NS\n}\n}\n(L & {n: NS, sc: \"o\"}).j\n}\n}\n}\n}\nlet L = {\nNS=n: string\nsc:   string\nlet l = #B & {\nn: NS\n}\nj: {\nl.v.r\n}\n}\n",
			Expected: []string{},
		},
		{
			Title: "empty structs",
			Given: "#B:\n{}\n{}\n{}\n{\n {\n  {\n   for s in [] {\n    L\n   }\n  }\n }\n {\n  NS: string\n  v: {\n   {\n    \"\\(string)/b\":\n     L\n   }\n  }\n }\n}\nlet L =\n#B\n",
			Expected: []string{
				"#B:\n{}\n{}\n{\n{\n{\nfor s in [] {\nL\n}\n}\n}\n{\nNS: string\nv: {\n{\n\"\\(string)/b\":\nL\n}\n}\n}\n}\nlet L =\n#B\n",
				"#B:\n{}\n{}\n{\n{\n{\nfor s in [] {\nL\n}\n}\n}\n{\nNS: string\nv: {\n{\n\"\\(string)/b\":\nL\n}\n}\n}\n}\nlet L =\n#B\n",
			},
		},
	}

	test.TestReduction(t, instances, EmptyDeclarationReduction{})
}
