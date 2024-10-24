#Abstract
spec: bar: {}
#Abstract: X={
 spec: *_#SpecFoo | {
  bar: {
   min: int
   max: int
  }
 }
 resource: _Thing & {_X: spec: X.spec}
}
_#Spec:
 int
_#SpecFoo: {
 min: 10
 max: 20
}
_#SpecBar: bar:
 int
_Thing: #Constrained & {
 _X: _
 spec: {
  if _X.spec.bar != _|_ {
   minBar: _X.spec.bar.min
   maxBar: _X.spec.bar.max
  }
 }
}
#Constrained: {
 spec: {} | {
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
minFoo: null
