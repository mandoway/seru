package logging

import (
	"log"
	"os"
)

var Cue = log.New(os.Stdout, "[CUE] ", log.LstdFlags)
