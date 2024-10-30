#Obj & {spec: bar: {}}
#Obj: X={
 spec: *{
  foo: min: 10
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
#SpecFoo: foo: min: 10
#SpecBar: bar: min: 20
#Out:
 int | {minBar: int}
*null | {nullBar: null}
