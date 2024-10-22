#Abstract
spec: bar: {}
#Abstract: X={
 spec: *_#SpecFoo | {
  bar: {
   min: int
   max: int
  }
 }
 resource: #Constrained & {
  _X: _
  spec: {
   if _X.spec.bar != _|_ {
    minBar: _X.spec.bar.min
    maxBar: _X.spec.bar.max
   }
  }
 } & {_X: spec: X.spec}
}
_#Spec:
 int
_#SpecFoo: {
 min: 10
 max: 20
}
_#SpecBar: bar:
 int
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
maxFoo: null
