package p

import (
 "tool/cli"
 "encoding/json"
)

input: [string]:
 {#in: string, #out: #in}.#out
command: foo:
 cli.Print & {text: json.Marshal(input)}
