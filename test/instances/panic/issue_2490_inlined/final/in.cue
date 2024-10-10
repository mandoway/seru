package p

import (
	"tool/cli"
	"encoding/json"
)

input: [name=string]: {
	#in:  name
	#out: IN.#x
}.#out
command: foo: cli.Print & {
	text: json.Marshal(input)
}

//cue:path: #in
let IN = {
	#x: string
}
