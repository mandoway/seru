package logging

import (
	"log"
	"os"
)

var Default = log.New(os.Stdout, "[SERU] ", log.LstdFlags)
var Semantic = log.New(os.Stdout, "[SEMANTIC] ", log.LstdFlags)
var Syntactic = log.New(os.Stdout, "[SYNTACTIC] ", log.LstdFlags)
