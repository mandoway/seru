package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestNullReduction(t *testing.T) {
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
	}

	test.TestReduction(t, instances, NullReduction{})
}

func TestRealWorldExampleNullReduction(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title:    "Real world example",
			Given:    "#B: {\nn: string\nv: {\nr: {\n}\n}\n}\n#F: {\nt: [...string]\n...\n}\n#N: {\nn: string\np: #F\nv: #V\n}\n#C: {\np: #F\nv: #V\n}\n#C2: {\nc: #C\nt: #N\nns: {\n[NS=string]: #N & {n: NS}\n}\ng: {\nt: {...} | *{}\n...\n}\n}\n#V: {\nlet l = {[string]: _}\nt: [string]: {}\nns: [string]: {}\nr: l\n}\nfs: f1: #C2 & {\nns: m: {\np: {}\nv: {\nr: {\nfor s in [\"e\"] {\n(L & {n: \"m\", sc: s}).j\n}\n}\n}\n}\nt: {\nNS=n: string\nv: {\nr: {\n\"\\(NS)/b\": {\nmetadata: {\nn: NS\n}\n}\n(L & {n: NS, sc: \"o\"}).j\n}\n}\n}\n}\nlet L = {\nNS=n: string\nsc:   string\nlet l = #B & {\nn: NS\n}\nj: {\nl.v.r\n}\n}\n",
			Expected: []string{},
		},
	}

	test.TestReduction(t, instances, NullReduction{})
}
