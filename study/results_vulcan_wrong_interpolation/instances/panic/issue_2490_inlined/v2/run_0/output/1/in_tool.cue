package #convert

import (
 "tool/file"
 "encoding/json"
)

Marshal: [Marshal]:
 {#convert}.#convert
#convert: {
 #in:      Marshal
 #convert: #in
}
command: foo:
 file.Create & {
  filename: "foo"
  contents: json.Marshal(Marshal)
 }
