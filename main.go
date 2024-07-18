package main

import (
	"encoding/json"
	"fmt"
	"github.com/mandoway/seru/collection"
	"plugin"

	"os"
)

func main() {
	if len(os.Args) < 2 {
		exitWithCodeOne("Usage: seru <path>", nil)
	}

	file := os.Args[1]

	bytes, err := os.ReadFile(file)
	if err != nil {
		exitWithCodeOne("Error reading file", err)
	}

	openPlugin, err := plugin.Open("cue.so")
	if err != nil {
		exitWithCodeOne("Error opening openPlugin", err)
	}

	reductionFunc, err := openPlugin.Lookup("Main")
	if err != nil {
		exitWithCodeOne("Did not find main function", err)
	}

	results, err := reductionFunc.(func([]byte) ([][]byte, error))(bytes)
	if err != nil {
		exitWithCodeOne("Error during reduction", err)
	}

	values, _ := json.Marshal(
		collection.MapSlice(results, func(it []byte) (string, error) {
			return string(it), nil
		}),
	)
	fmt.Printf("%s", values)
}

func exitWithCodeOne(message string, err error) {
	fmt.Printf(message)
	if err != nil {
		fmt.Printf(": %s", err.Error())
	}
	fmt.Println()
	os.Exit(1)
}
