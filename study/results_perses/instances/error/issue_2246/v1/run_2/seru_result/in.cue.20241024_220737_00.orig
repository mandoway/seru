data: forms: [{
	fooID: "00-0000001"
}]

form1040: (#compute & {in: data}).out

#K1: {
	#_base: common: 3
	#FormFoo: {
		#_base
		fooID: string
	}
	#FormBar: {
		#_base
		barID: string
	}
	#Form: {
		#FormFoo | #FormBar
	}
}

#Input: {
	forms: [...#K1.#Form]
}

#summarizeReturn: {
	in:  #Input
	out: [ for k in in.forms { k.common } ]
}

#compute: {
	in:  #Input
	out: (#summarizeReturn & {"in": in}).out
}
