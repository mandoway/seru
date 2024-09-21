package infrastructure

import (
	"path"
	"tool/file"
	"encoding/json"
	"strings"

	"github.com/cue-examples/cue-terraform-github-config-experiment/config"
)

_goos: string @tag(os,var=os)

command: {
	foo: {
		let json_indent = "    " & strings.MinRunes(4) & strings.MaxRunes(4)
		let dir_operations = path.FromSlash("_operations/github", path.Unix)
		let file_tf_json = "config.tf.json"

		orgs: {
			for orgName, orgTerraform in config.target.terraform.github.org
			let dir_org = path.Join([dir_operations, orgName], _goos)
			let file_org_config = path.Join([dir_org, file_tf_json], _goos) {
				"generate config \(file_org_config)": file.Create & {
					filename: file_org_config
					contents: json.Indent(json.Marshal(orgTerraform.config), "", json_indent) + "\n"
				}
			}
		}
	}
}
