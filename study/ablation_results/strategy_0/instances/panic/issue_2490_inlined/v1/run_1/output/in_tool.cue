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
 strings
 dir_operations: path
 file_tf_json:   "config.tf.json"
 for orgName, orgTerraform in github {
  cli.Print & {
   text: json.Marshal(orgTerraform)
  }
 }
}
github
_
target: terraform: {
 github
 #Identifier: {
  rules: {
   valid_initial_characters: "-a-zA-Z_"
   valid_characters:         valid_initial_characters
  }
  _
  rules
  adapt: {
   #in:  string
   _a:   #in
   _b:   regexp.ReplaceAll( rules.valid_initial_characters, _a, "_$1")
   #out: _b
  }
 }
}
github: org: [_]: config: resource:
 {target.terraform.#Identifier.adapt}.#out
