package infrastructure

import (
 "tool/cli"
 "encoding/json"
 "regexp"
)

@tag(         )
command: foo: {
 _
 file_tf_json: "config.tf.json"
 for orgName, orgTerraform in github {
  cli.Print & {
   text: json.Marshal(orgTerraform)
  }
 }
}
target: terraform: #Identifier: {
 rules: {
  valid_initial_characters: "-a-zA-Z_"
  valid_characters:         valid_initial_characters
 }
 adapt: {
  #in:  string
  _a:   #in
  _b:   regexp.ReplaceAll( rules.valid_initial_characters, _a, "_$1")
  #out: _b
 }
}
github: org: [_]: config: resource:
 {target.terraform.#Identifier.adapt}.#out
