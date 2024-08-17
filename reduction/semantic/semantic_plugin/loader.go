package semantic_plugin

import (
	"fmt"
	"github.com/mandoway/seru/reduction/semantic"
	"plugin"
	"strings"
)

func LoadSemanticReductionPlugin(language string) (*semantic.Functions, error) {
	pluginName := getPluginNamePerConvention(language)

	openPlugin, err := plugin.Open(pluginName)
	if err != nil {
		return nil, err
	}

	mainFunc, err := lookupFunction[semantic.ReduceFunctionType](openPlugin, pluginName, semantic.ReduceFunction)
	if err != nil {
		return nil, err
	}

	countFunc, err := lookupFunction[semantic.CountFunctionType](openPlugin, pluginName, semantic.CountFunction)
	if err != nil {
		return nil, err
	}

	return &semantic.Functions{
		ReduceFunction: *mainFunc,
		CountFunction:  *countFunc,
	}, nil
}

func lookupFunction[T semantic.ReduceFunctionType | semantic.CountFunctionType](openPlugin *plugin.Plugin, pluginName, functionName string) (*T, error) {
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
