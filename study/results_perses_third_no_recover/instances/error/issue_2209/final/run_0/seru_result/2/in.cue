#Obj & {spec: bar: {}}
#Obj: X={
 spec: *{
  foo: min: int
 } | {
  bar: min: 20
 }
 #Out & {
  X
  if spec.bar != _|_ {
   minBar: spec.bar.min
  }
 }
}
#SpecFoo: foo: min: int
#SpecBar: bar: min: 20
#Out:
 int | {minBar: int}
*null | {nullBar: null}
