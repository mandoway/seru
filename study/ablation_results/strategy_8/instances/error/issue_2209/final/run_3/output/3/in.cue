{}
#Obj & {spec: bar: {}}
#Obj: X={
 spec: *{
  foo: min: int
 } | {
  bar: min: 20
 }
 (int | {minBar: int}) & {
  X
  X
  if spec.bar != _|_ {
   minBar: spec.bar.min
  }
 }
}
#SpecFoo: foo: min: int
#SpecBar: bar: min: 20
int | {minBar: int}
