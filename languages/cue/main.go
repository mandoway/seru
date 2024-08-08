package main

import (
	"github.com/mandoway/seru/languages/cue/context"
	"github.com/mandoway/seru/reduction/semantic"
)

func Main(fileContent []byte) ([][]byte, error) {
	return semantic.Reduce(fileContent, context.BuildContext())
}

func CountTokens(fileContent []byte) int {
	return context.CountTokensUsingScanner(fileContent)
}
