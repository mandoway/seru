package plugin

import (
	"fmt"
	"plugin"
	"strings"
)

func LoadSemanticReductionPlugin(language string) (*Functions, error) {
	pluginName := getPluginNamePerConvention(language)

	openPlugin, err := plugin.Open(pluginName)
	if err != nil {
		return nil, err
	}

	mainFunc, err := lookupFunction[SemanticReductionFunction](openPlugin, pluginName, SemanticReduceFunction)
	if err != nil {
		return nil, err
	}

	countFunc, err := lookupFunction[TokenCountFunction](openPlugin, pluginName, CountFunction)
	if err != nil {
		return nil, err
	}

	strategyCountFunc, err := lookupFunction[StrategyCountFunction](openPlugin, pluginName, GetStrategyCountFunction)
	if err != nil {
		return nil, err
	}

	return &Functions{
		SemanticReduce:   *mainFunc,
		CountTokens:      *countFunc,
		GetStrategyCount: *strategyCountFunc,
	}, nil
}

func lookupFunction[T any](openPlugin *plugin.Plugin, pluginName, functionName string) (*T, error) {
	funcSymbol, err := openPlugin.Lookup(functionName)
	if err != nil {
		return nil, err
	}
	function, ok := funcSymbol.(*T)
	if !ok {
		return nil, fmt.Errorf("plugin %s has wrong %s-Function type. expected %T, got %T", pluginName, functionName, *new(T), funcSymbol)
	}
	return function, nil
}

func getPluginNamePerConvention(language string) string {
	return fmt.Sprintf("%s.so", strings.ToLower(language))
}
