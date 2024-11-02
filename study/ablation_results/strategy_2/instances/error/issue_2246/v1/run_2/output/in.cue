data: forms: [
 "00-0000001",
]
#compute & {in: data}
#K1: {
 #_base: common: 3
 #_base
 #FormBar:
  string
 #Form:
  #FormBar
}
#Input: forms: [#K1.#Form]
#summarizeReturn: {
 in: #Input
 in
}
#compute:
 #summarizeReturn
