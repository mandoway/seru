#Obj & {spec: bar: {}}
#Obj: X={
 spec: *{
  foo: min: 10
 } | {
  bar: min: 20
 }
 (int | {minBar: int}) & {
  X
  if spec.bar != _|_ {
   minBar: spec.bar.min
  }
 }
}
#SpecFoo: foo: min: 10
#SpecBar: bar: min: 20
null | {nullBar: null}
