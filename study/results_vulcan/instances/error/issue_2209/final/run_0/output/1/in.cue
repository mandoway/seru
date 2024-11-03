#Obj & {spec: bar: {}}
#Obj: X={
 spec: *X | #SpecBar
 #Out &
 {
  minBar: spec.bar.min
 }
}
#Out
#SpecBar: bar: min: 20
#Out:
 {} | {minBar: minBar}
*null | {}
