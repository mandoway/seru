package infrastructure

import (
	"path"
	"tool/cli"
	"encoding/json"
	"strings"
	"regexp"
)

_goos: string @tag(os,var=os)

command: {
	foo: {
		let json_indent = "    " & strings.MinRunes(4) & strings.MaxRunes(4)
		let dir_operations = path.FromSlash("_operations/github", path.Unix)
		let file_tf_json = "config.tf.json"

		orgs: {
			for orgName, orgTerraform in target.terraform.github.org
			let dir_org = path.Join([dir_operations, orgName], _goos)
			let file_org_config = path.Join([dir_org, file_tf_json], _goos) {
				"generate config \(file_org_config)": cli.Print & {
					text: json.Indent(json.Marshal(orgTerraform.config), "", json_indent) + "\n"
				}
			}
		}
	}
}

github: org: "supreme-octo-enigma": config: {
}

X=github: _

target: terraform: {
	github: org: {
		[_]: config: resource: [_]: #Identifier.valid
		for org_name, org_config in X.org {
			(org_name): config: {
				resource: {
					for resource_type, resource_instance in org_config.config.resource {
						(resource_type): {
							for resource_identifier, resource_config in resource_instance {
								let new_resource_identifier = {#Identifier.adapt & {#in: resource_identifier}}.#out
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

	#Identifier: {
		rules: {
			valid_initial_characters: "-a-zA-Z_"
			valid_characters:         valid_initial_characters + "0-9"
		}
		valid: [and(valid_constraints)]: _
		valid_constraints: [
			=~"^[\(rules.valid_characters)]+$",
			=~"^[\(rules.valid_initial_characters)]",
		]
		adapt: {
			#in: string

			let _a = regexp.ReplaceAllLiteral("[^\(rules.valid_characters)]", #in, "_")

			let _b = regexp.ReplaceAll("^([^\(rules.valid_initial_characters)])", _a, "_$1")

			#out: _b
		}
	}
}


github: org: [_]: config: resource: {
	[resource_type=_]: [cue_resource_name=string]: {
		#tfid:  "\({target.terraform.#Identifier.adapt & {#in: cue_resource_name}}.#out)"
		#tfref: "\(resource_type).\(#tfid)"
	}
}