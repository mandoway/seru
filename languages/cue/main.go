package main

import (
	"github.com/mandoway/seru/languages/cue/context"
	"github.com/mandoway/seru/reduction/semantic"
)

//goland:noinspection GoUnusedGlobalVariable
var Reduce semantic.ReduceFunctionType = func(fileContent []byte) ([][]byte, error) {
	return semantic.Reduce(fileContent, context.BuildContext())
}

//goland:noinspection GoUnusedGlobalVariable
var CountTokens semantic.CountFunctionType = func(fileContent []byte) int {
	return context.CountTokensUsingScanner(fileContent)
}
