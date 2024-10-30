package json

import (
 "tool/cli"
 "encoding/json"
 "strings"
 "regexp"
)

command: foo:
{
 cli.Print
 text: json.Indent(json.Marshal(#Identifier), " ", "      " & strings.MinRunes(4))
}
#Identifier: #Identifier: #Identifier: #Identifier: {
 #Identifier: string
 json:        regexp.ReplaceAll( #Identifier, #Identifier, " _$1 ")
}
#Identifier: [#Identifier]:
 {#Identifier.#Identifier.#Identifier.#Identifier}.json
