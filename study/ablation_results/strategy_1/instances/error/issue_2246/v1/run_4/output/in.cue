data: forms:
 "00-0000001"
#K1: {
 #_base: common: 3
 #FormFoo:
  #_base
 #FormBar:
  #_base
 #Form: {
  #FormFoo | #FormBar
 }
}
#Input:
 #K1
#summarizeReturn: {
 in: #Input
 in
}
{
 in: #Input
 #summarizeReturn
}
