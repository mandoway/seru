package infrastructure

import (
 "tool/cli"
 "encoding/json"
)

@tag(         )
command: foo: {
 file_tf_json: "config.tf.json"
 for orgName, orgTerraform in github {
  cli.Print & {
   text: json.Marshal(orgTerraform)
  }
 }
}
target: terraform: #Identifier: {
 valid_initial_characters: "-a-zA-Z_"
 valid_constraints:        "^[-a-zA-Z_]+$"
 adapt: {
  #in: string
  _a:  string
  #out:
   _a
 }
}
github: org: [_]: config: resource: resource_type: cue_resource_name:
 "\({target.terraform.#Identifier.adapt & {#in: string}}.#out)"
