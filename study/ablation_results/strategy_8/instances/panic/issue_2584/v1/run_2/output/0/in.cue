#B: {
	n: string
	v: {
		r: {
		}
	}
}

#F: {
	t: [...string]
	...
}

#N: {
	n: string
	p: #F
	v: #V
}

#C: {
	p: #F
	v: #V
}

#C2: {
	c: #C
	t: #N
	ns: {
		[NS=string]: #N & {n: NS}
	}

	g: {
		t: {...} | *{}
		...
	}
}

#V: {
	let l = {[string]: _}
	t: [string]: {}
	ns: [string]: {}
	r: l
}

fs: f1: #C2 & {
	ns: m: {
		p: {}
		v: {
			r: {
				for s in ["e"] {
					(L & {n: "m", sc: s}).j
				}
			}
		}

	}

	t: {
		NS=n: string

		v: {
			r: {
				"\(NS)/b": {
					metadata: {
						n: NS
					}
				}
				(L & {n: NS, sc: "o"}).j
			}
		}
	}
}

let L = {
	NS=n: string
	sc:   string

	let l = #B & {
		n: NS
	}

	j: {
		l.v.r
	}
}
