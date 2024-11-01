package #in

import (
 "tool/cli"
 "encoding/json"
)

Print: [Print]:
 {#in: Print, text: #in}.text
command: foo:
 cli.Print & {text: json.Marshal(Print)}
