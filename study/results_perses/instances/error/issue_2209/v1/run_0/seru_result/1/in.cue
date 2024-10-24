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
  if _X.foo != _|_ {
   minFoo
   _X
  }
  if _X.spec.bar != _|_ {
   minBar: _X.spec.bar.min
   maxBar: _X.spec.bar.max
  }
 }
}
#Constrained: {
 spec: {
  minFoo
  int
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
minFoo: null
