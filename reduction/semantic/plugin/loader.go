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

	mainFunc, err := lookupFunction[MainFunctionType](openPlugin, pluginName, MainFunction)
	if err != nil {
		return nil, err
	}

	countFunc, err := lookupFunction[CountFunctionType](openPlugin, pluginName, CountFunction)
	if err != nil {
		return nil, err
	}

	return &Functions{
		MainFunction:  mainFunc,
		CountFunction: countFunc,
	}, nil
}

func lookupFunction[T MainFunctionType | CountFunctionType](openPlugin *plugin.Plugin, pluginName, functionName string) (T, error) {
	funcSymbol, err := openPlugin.Lookup(functionName)
	if err != nil {
		return nil, err
	}
	function, ok := funcSymbol.(T)
	if !ok {
		return nil, fmt.Errorf("plugin %s has wrong %s-Function type. expected %T, got %T", pluginName, functionName, *new(T), funcSymbol)
	}
	return function, nil
}

func getPluginNamePerConvention(language string) string {
	return fmt.Sprintf("%s.so", strings.ToLower(language))
}
