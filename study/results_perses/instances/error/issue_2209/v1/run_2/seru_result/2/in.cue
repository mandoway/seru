#Abstract
spec: bar: {}
#Abstract: X={
 spec: *_#SpecFoo | {
  bar: {
   min: 30
   max: int
  }
 }
 resource: _Thing & {_X: spec: X.spec}
}
min: 30
_#SpecFoo: {
 min: 10
 max: 20
}
min: 30
_Thing: #Constrained & {
 _X: _
 spec: {
  if _X.spec.bar != _|_ {
   minBar: min
   maxBar: _X.spec.bar.max
  }
 }
}
#Constrained: {
 spec: {} | {
  minBar: int
  maxBar: 40
  minFoo: null
  maxFoo: null
 }
 spec: {
  fuga: null
 } |
  null
}
minFoo?: int
