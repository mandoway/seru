package strategy

import (
	"github.com/mandoway/seru/languages/cue/test"
	"testing"
)

func TestLetReduction(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title: "Return candidate for every occurrence",
			Given: `
			x: 42
			let y = x
			let z = x
			foo: {
				let bar = x
				baz: bar
			}
			`,
			Expected: []string{
				`
				x: 42
				y: x
				let z = x
				foo: {
					let bar = x
					baz: bar
				}
				`,
				`
				x: 42
				let y = x
				z: x
				foo: {
					let bar = x
					baz: bar
				}
				`,
				`
				x: 42
				let y = x
				let z = x
				foo: {
					bar: x
					baz: bar
				}
				`,
			},
		},
		{
			Title: "Return empty if not applicable",
			Given: `
			x: 42
			y: x
			z: y
			`,
			Expected: nil,
		},
		{
			Title: "ignore let fields in for clause",
			Given: `
			for foo in [1, 2, 3] 
			let double = foo * 2 {
			  double
			}
			`,
			Expected: nil,
		},
	}

	test.TestReduction(t, instances, LetReduction{})
}

func TestRealWorld(t *testing.T) {
	instances := []test.ReductionInstance{
		{
			Title: "Issue 2490 inlined v1",
			Given: "package infrastructure\n\nimport (\n \"path\"\n \"tool/cli\"\n \"encoding/json\"\n \"strings\"\n \"regexp\"\n)\n\n@tag(         )\ncommand: foo: {\n let json_indent = \"    \" & strings.MinRunes(4)\n let dir_operations = path\n let file_tf_json = \"config.tf.json\"\n for orgName, orgTerraform in github\n let dir_org = dir_operations\n let file_org_config = file_tf_json {\n  cli.Print & {\n   text: json.Indent(json.Marshal(orgTerraform), \"\", json_indent)\n  }\n }\n}\ngithub\n_\ntarget: terraform: {\n github\n #Identifier: {\n  rules: {\n   valid_initial_characters: \"-a-zA-Z_\"\n   valid_characters:         valid_initial_characters\n  }\n  _\n  valid_constraints:\n   \"^[\\(rules.valid_characters)]+$\"\n  adapt: {\n   #in: string\n   let _a = #in\n   let _b = regexp.ReplaceAll(\"^([^\\(rules.valid_initial_characters)])\", _a, \"_$1\")\n   #out: _b\n  }\n }\n}\ngithub: org: [_]: config: resource: resource_type: [cue_resource_name=string]:\n \"\\({target.terraform.#Identifier.adapt & {#in: cue_resource_name}}.#out)\"\n",
			Expected: []string{
				"package infrastructure\nimport (\n\"path\"\n\"tool/cli\"\n\"encoding/json\"\n\"strings\"\n\"regexp\"\n)\n@tag(         )\ncommand: foo: {\njson_indent: \"    \" & strings.MinRunes(4)\nlet dir_operations = path\nlet file_tf_json = \"config.tf.json\"\nfor orgName, orgTerraform in github\nlet dir_org = dir_operations\nlet file_org_config = file_tf_json {\ncli.Print & {\ntext: json.Indent(json.Marshal(orgTerraform), \"\", json_indent)\n}\n}\n}\ngithub\n_\ntarget: terraform: {\ngithub\n#Identifier: {\nrules: {\nvalid_initial_characters: \"-a-zA-Z_\"\nvalid_characters:         valid_initial_characters\n}\n_\nvalid_constraints:\n\"^[\\(rules.valid_characters)]+$\"\nadapt: {\n#in: string\nlet _a = #in\nlet _b = regexp.ReplaceAll(\"^([^\\(rules.valid_initial_characters)])\", _a, \"_$1\")\n#out: _b\n}\n}\n}\ngithub: org: [_]: config: resource: resource_type: [cue_resource_name=string]:\n\"\\({target.terraform.#Identifier.adapt & {#in: cue_resource_name}}.#out)\"",
				"package infrastructure\nimport (\n\"path\"\n\"tool/cli\"\n\"encoding/json\"\n\"strings\"\n\"regexp\"\n)\n@tag(         )\ncommand: foo: {\nlet json_indent = \"    \" & strings.MinRunes(4)\ndir_operations: path\nlet file_tf_json = \"config.tf.json\"\nfor orgName, orgTerraform in github\nlet dir_org = dir_operations\nlet file_org_config = file_tf_json {\ncli.Print & {\ntext: json.Indent(json.Marshal(orgTerraform), \"\", json_indent)\n}\n}\n}\ngithub\n_\ntarget: terraform: {\ngithub\n#Identifier: {\nrules: {\nvalid_initial_characters: \"-a-zA-Z_\"\nvalid_characters:         valid_initial_characters\n}\n_\nvalid_constraints:\n\"^[\\(rules.valid_characters)]+$\"\nadapt: {\n#in: string\nlet _a = #in\nlet _b = regexp.ReplaceAll(\"^([^\\(rules.valid_initial_characters)])\", _a, \"_$1\")\n#out: _b\n}\n}\n}\ngithub: org: [_]: config: resource: resource_type: [cue_resource_name=string]:\n\"\\({target.terraform.#Identifier.adapt & {#in: cue_resource_name}}.#out)\"",
				"package infrastructure\nimport (\n\"path\"\n\"tool/cli\"\n\"encoding/json\"\n\"strings\"\n\"regexp\"\n)\n@tag(         )\ncommand: foo: {\nlet json_indent = \"    \" & strings.MinRunes(4)\nlet dir_operations = path\nfile_tf_json: \"config.tf.json\"\nfor orgName, orgTerraform in github\nlet dir_org = dir_operations\nlet file_org_config = file_tf_json {\ncli.Print & {\ntext: json.Indent(json.Marshal(orgTerraform), \"\", json_indent)\n}\n}\n}\ngithub\n_\ntarget: terraform: {\ngithub\n#Identifier: {\nrules: {\nvalid_initial_characters: \"-a-zA-Z_\"\nvalid_characters:         valid_initial_characters\n}\n_\nvalid_constraints:\n\"^[\\(rules.valid_characters)]+$\"\nadapt: {\n#in: string\nlet _a = #in\nlet _b = regexp.ReplaceAll(\"^([^\\(rules.valid_initial_characters)])\", _a, \"_$1\")\n#out: _b\n}\n}\n}\ngithub: org: [_]: config: resource: resource_type: [cue_resource_name=string]:\n\"\\({target.terraform.#Identifier.adapt & {#in: cue_resource_name}}.#out)\"",
				"package infrastructure\nimport (\n\"path\"\n\"tool/cli\"\n\"encoding/json\"\n\"strings\"\n\"regexp\"\n)\n@tag(         )\ncommand: foo: {\nlet json_indent = \"    \" & strings.MinRunes(4)\nlet dir_operations = path\nlet file_tf_json = \"config.tf.json\"\nfor orgName, orgTerraform in github\nlet dir_org = dir_operations\nlet file_org_config = file_tf_json {\ncli.Print & {\ntext: json.Indent(json.Marshal(orgTerraform), \"\", json_indent)\n}\n}\n}\ngithub\n_\ntarget: terraform: {\ngithub\n#Identifier: {\nrules: {\nvalid_initial_characters: \"-a-zA-Z_\"\nvalid_characters:         valid_initial_characters\n}\n_\nvalid_constraints:\n\"^[\\(rules.valid_characters)]+$\"\nadapt: {\n#in: string\n_a:  #in\nlet _b = regexp.ReplaceAll(\"^([^\\(rules.valid_initial_characters)])\", _a, \"_$1\")\n#out: _b\n}\n}\n}\ngithub: org: [_]: config: resource: resource_type: [cue_resource_name=string]:\n\"\\({target.terraform.#Identifier.adapt & {#in: cue_resource_name}}.#out)\"",
				"package infrastructure\nimport (\n\"path\"\n\"tool/cli\"\n\"encoding/json\"\n\"strings\"\n\"regexp\"\n)\n@tag(         )\ncommand: foo: {\nlet json_indent = \"    \" & strings.MinRunes(4)\nlet dir_operations = path\nlet file_tf_json = \"config.tf.json\"\nfor orgName, orgTerraform in github\nlet dir_org = dir_operations\nlet file_org_config = file_tf_json {\ncli.Print & {\ntext: json.Indent(json.Marshal(orgTerraform), \"\", json_indent)\n}\n}\n}\ngithub\n_\ntarget: terraform: {\ngithub\n#Identifier: {\nrules: {\nvalid_initial_characters: \"-a-zA-Z_\"\nvalid_characters:         valid_initial_characters\n}\n_\nvalid_constraints:\n\"^[\\(rules.valid_characters)]+$\"\nadapt: {\n#in: string\nlet _a = #in\n_b:   regexp.ReplaceAll(\"^([^\\(rules.valid_initial_characters)])\", _a, \"_$1\")\n#out: _b\n}\n}\n}\ngithub: org: [_]: config: resource: resource_type: [cue_resource_name=string]:\n\"\\({target.terraform.#Identifier.adapt & {#in: cue_resource_name}}.#out)\"",
			},
		},
	}

	test.TestReduction(t, instances, LetReduction{})
}
