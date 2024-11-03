package infrastructure

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
 for orgTerraform in github
 let valid_initial_characters = dir_operations
 let valid_initial_characters = file_tf_json {
  cli.Print
  text: json.Indent(json.Marshal(github), "", json_indent)
 }
}
target: terraform: #Identifier: {
 rules: valid_initial_characters: "-a-zA-Z_"
 adapt: {
  github: string
  let file_tf_json = github
  let json_indent = regexp.ReplaceAll("^([^\(rules.valid_initial_characters)])", file_tf_json, "_$1")
  #out: json_indent
 }
}
github: resource_type: [cue_resource_name=github]:
 "\({target.terraform.#Identifier.adapt & {#in: cue_resource_name}}.#out)"
