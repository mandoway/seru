{}
#Obj & {spec: bar: {}}
#Obj: X={
 spec: *#SpecFoo | #SpecBar
 #Out & {
  _Xspec: X.spec
  if false {
   _Xspec
  }
  if _Xspec.bar != _|_ {
   minBar: _Xspec.bar.min
  }
 }
}
#SpecFoo: foo: min: int
#SpecBar: bar: min: 20
#Out: {
 {minFoo: int} | {minBar: int}
 *null | {nullBar: null}
}
