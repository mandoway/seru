package p

import (
 "tool/file"
 "encoding/json"
)

input: [string]:
 {#convert}.#out
#convert: {
 #in:  _
 #out: #in
}
command: foo:
 file.Create & {
  filename: "foo"
  contents: json.Marshal(input)
 }
