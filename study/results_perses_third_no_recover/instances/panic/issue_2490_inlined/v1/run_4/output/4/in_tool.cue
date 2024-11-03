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
valid_initial_characters: "-a-zA-Z_"
github: org: [_]: config: resource:
 {
  {
   terraform: #Identifier: {
    valid_initial_characters: "-a-zA-Z_"
    adapt: {
     _a: string
     #out:
      _a
    }
   }
  }.terraform.#Identifier.adapt
 }.#out
