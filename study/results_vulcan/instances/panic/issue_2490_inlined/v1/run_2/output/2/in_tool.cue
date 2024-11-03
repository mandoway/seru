package #in

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
#Identifier: #Identifier: #Identifier: adapt: {
 #Identifier: string
 #Identifier
 _b: regexp.ReplaceAll( #Identifier, #Identifier, " _$1 ")
 _b: _b
}
#Identifier: [#Identifier]:
 {#Identifier.#Identifier.#Identifier.adapt}._b
