#Abstract
spec: bar: {}
#Abstract: X={
 spec: *_#SpecFoo | {
  bar: {
   min: int
   max: int
  }
 }
 resource: {
  spec: {
   null
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
 } & {
  _X: _
  spec: {
   _
   if _X.spec.bar != _|_ {
    minBar: _X.spec.bar.min
    maxBar: _X.spec.bar.max
   }
  }
 } & {_X: spec: X.spec}
}
min: 10
_#SpecFoo: {
 min: 10
 max: 20
}
min: int
_
maxBar: 40
minFoo: null
