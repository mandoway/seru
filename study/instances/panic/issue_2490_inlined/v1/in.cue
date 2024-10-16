package infrastructure

import (
	"strings"
	"path"
	"tool/file"
	"encoding/json"
	"regexp"
)

_goos: string @tag(os,var=os)
command: foo: {
	let json_indent = "    " & strings.MinRunes(4) & strings.MaxRunes(4)
	let dir_operations = path.FromSlash("_operations/github", path.Unix)
	let file_tf_json = "config.tf.json"
	orgs: {
		for orgName, orgTerraform in ORG
		let dir_org = path.Join([dir_operations, orgName], _goos)
		let file_org_config = path.Join([dir_org, file_tf_json], _goos) {
			"generate config \(file_org_config)": file.Create & {
				filename: file_org_config
				contents: json.Indent(json.Marshal(orgTerraform.config), "", json_indent) + """


					"""
			}
		}
	}
}

//cue:path: "github.com/cue-examples/cue-terraform-github-config-experiment/config".target.terraform.github.org
let ORG = {
	[_]: config: resource: [_]: VALID.#x
	for org_name, org_config in ORG_1 {
		(org_name): config: {
			resource: {
				for resource_type, resource_instance in org_config.config.resource {
					(resource_type): {
						for resource_identifier, resource_config in resource_instance {
							let new_resource_identifier = (#Identifier.adapt & {
								#in: resource_identifier
							}).#out
							(new_resource_identifier): resource_config
						}
					}
				}
			}
			for top_level_field, value in org_config.config
			if top_level_field != "resource" {
				(top_level_field): value
			}
		}
	}
}

//cue:path: "github.com/cue-examples/cue-terraform-github-config-experiment/config".target.terraform.#Identifier.valid
let VALID = {
	#x: {
		[and(VALID_CONSTRAINTS.#x)]: _
	}
}

//cue:path: "github.com/cue-examples/cue-terraform-github-config-experiment/config".target.terraform.#Identifier.valid_constraints
let VALID_CONSTRAINTS = {
	#x: [=~"^[\(VALID_CHARACTERS.#x)]+$", =~"^[\(VALID_INITIAL_CHARACTERS.#x)]"]
}

//cue:path: "github.com/cue-examples/cue-terraform-github-config-experiment/config".target.terraform.#Identifier.rules.valid_characters
let VALID_CHARACTERS = {
	#x: VALID_INITIAL_CHARACTERS.#x + "0-9"
}

//cue:path: "github.com/cue-examples/cue-terraform-github-config-experiment/config".target.terraform.#Identifier.rules.valid_initial_characters
let VALID_INITIAL_CHARACTERS = {
	#x: "-a-zA-Z_"
}

//cue:path: "github.com/cue-examples/cue-terraform-github-config-experiment/config".github.org
let ORG_1 = {
	{
		[_]: config: resource: [resource_type=_]: [cue_resource_name=string]: {
			#tfid:  "\(OUT.#x)"
			#tfref: "\(resource_type).\(#tfid)"
		}
	}
	"supreme-octo-enigma": config: {}
}

//cue:path: "github.com/cue-examples/cue-terraform-github-config-experiment/config".target.terraform.#Identifier.adapt.#out
let OUT = {
	#x: regexp.ReplaceAll("^([^\(VALID_INITIAL_CHARACTERS.#x)])", regexp.ReplaceAllLiteral("[^\(VALID_CHARACTERS.#x)]", IN.#x, "_"), "_$1")
}
