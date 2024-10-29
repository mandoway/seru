package logging

import (
	"log"
	"os"
)

var Cue = log.New(os.Stdout, "[CUE] ", log.LstdFlags)
var CueError = log.New(os.Stderr, "[CUE] ", log.LstdFlags)
