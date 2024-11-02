package p

import (
	"tool/cli"
	"encoding/json"
)

input: [name=string]: {
	{#in: name, #out: #in}.#out
}
command: foo: {
	cli.Print & {text: json.Marshal(input)}
}
