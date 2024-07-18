package main

import (
	"github.com/mandoway/seru/languages/cue/context"
	"github.com/mandoway/seru/reduction"
)

func Main(fileContent []byte) ([][]byte, error) {
	return reduction.Reduce(fileContent, context.BuildContext())
}
