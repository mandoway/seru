package infrastructure

import (
 "path"
 "tool/cli"
 "encoding/json"
 "strings"
 "regexp"
)

@tag(         )
command: foo: {
 let json_indent = "    " & strings.MinRunes(4)
 let dir_operations = path
 let file_tf_json = "config.tf.json"
 for orgName, orgTerraform in github
 let dir_org = dir_operations
 let file_org_config = file_tf_json {
  cli.Print & {
   text: json.Indent(json.Marshal(orgTerraform), "", json_indent)
  }
 }
}
_
_
target: terraform: {
 github
 #Identifier: {
  valid_characters: "-a-zA-Z_"
  _
  valid_characters: "-a-zA-Z_"
  adapt: {
   #in: string
   let _a = #in
   let _b = regexp.ReplaceAll(
   valid_characters, _a, "_$1")
   #out: _b
  }
 }
}
github: org: [_]: config: resource:
 {target.terraform.#Identifier.adapt}.#out
