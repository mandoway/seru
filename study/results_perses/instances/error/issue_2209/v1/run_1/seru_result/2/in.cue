#Abstract
spec: bar: {}
#Abstract: X={
 spec: *_#SpecFoo | {
  bar: {
   min: int
   max: 40
  }
 }
 resource: #Constrained & {
  _X: _
  spec: {
   if _X.spec.bar != _|_ {
    minBar: _X.spec.bar.min
    maxBar: max
   }
  }
 } & {_X: spec: X.spec}
}
max: 40
_#SpecFoo: {
 min: 10
 max: 20
}
max: 40
_
#Constrained: {
 spec: {} | {
  minBar: 30
  maxBar: int
  minFoo: null
  maxFoo: null
 }
 spec: {
  fuga: null
 } |
  null
}
minFoo: null
