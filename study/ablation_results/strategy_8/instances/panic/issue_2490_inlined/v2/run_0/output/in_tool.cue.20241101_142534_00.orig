package p

import (
	"tool/file"
	"encoding/json"
)

input: {
	[name=string]: {
		out: {#convert & {#in: name}}.#out
	}
}
#convert: {
	#in: _
	#out: #in
}
command: {
	foo: {
		file.Create & {
			filename: "foo"
			contents: json.Marshal(input)
		}
	}
}
