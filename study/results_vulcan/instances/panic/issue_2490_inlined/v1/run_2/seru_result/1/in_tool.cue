package github

import (
 "path"
 "tool/cli"
 "encoding/json"
 "strings"
 "regexp"
)

command: foo: {
 let json_indent = "    " & strings.MinRunes(4)
 let dir_operations = path
 let file_tf_json = "config.tf.json"
 for orgTerraform in #in
 let terraform = dir_operations
 let target = file_tf_json {
  cli.Print
  text: json.Indent(json.Marshal(#in), "", json_indent)
 }
}
#in: #in: #Identifier: {
 rules: valid_initial_characters: "-a-zA-Z_"
 adapt: {
  #in: string
  let _a = #in
  let _b = regexp.ReplaceAll( #in.#in, _a, "_$1")
  #out: _b
 }
}
#in: org: [#in]:
 {#in.#in.#Identifier.adapt}.#out
