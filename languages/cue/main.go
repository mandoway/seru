package main

import (
	"github.com/mandoway/seru/languages/cue/context"
	"github.com/mandoway/seru/languages/cue/language"
	"github.com/mandoway/seru/reduction/plugin"
	"github.com/mandoway/seru/reduction/semantic"
)

//goland:noinspection GoUnusedGlobalVariable
var Reduce plugin.SemanticReductionFunction = func(fileContent []byte, strategyIndex int) ([][]byte, error) {
	return semantic.Reduce(fileContent, strategyIndex, context.Context)
}

//goland:noinspection GoUnusedGlobalVariable
var CountTokens plugin.TokenCountFunction = func(fileContent []byte) int {
	return language.CountTokensUsingScanner(fileContent)
}

//goland:noinspection GoUnusedGlobalVariable
var GetStrategyCount plugin.StrategyCountFunction = func() int {
	return len(context.Strategies)
}
