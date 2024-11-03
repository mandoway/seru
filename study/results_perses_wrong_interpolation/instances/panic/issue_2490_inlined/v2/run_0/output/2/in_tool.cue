package p

import (
 "tool/file"
 "encoding/json"
)

input: [string]:
 {
  #in:  _
  #out: #in
 }.#out
_
command: foo:
 file.Create & {
  filename: "foo"
  contents: json.Marshal(input)
 }
