package main

import (
	"fmt"
	"github.com/mandoway/seru/languages/cue/context"
	"github.com/mandoway/seru/languages/cue/language"
	"github.com/mandoway/seru/reduction/plugin"
	"github.com/mandoway/seru/reduction/semantic"
	"strings"
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

//goland:noinspection GoUnusedGlobalVariable
var GetStrategyName plugin.StrategyNameFunction = func(index int) string {
	if index < 0 || index >= len(context.Strategies) {
		return ""
	}

	nameWithPackage := fmt.Sprintf("%T", context.Strategies[index])
	return strings.TrimPrefix(nameWithPackage, "strategy.")
}
