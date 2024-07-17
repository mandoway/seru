package main

import (
	"encoding/json"
	"fmt"
	"github.com/mandoway/seru/collection"
	"github.com/mandoway/seru/cue"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: seru <path>")
		return
	}

	file := os.Args[1]

	bytes, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file: %s", err.Error())
		os.Exit(1)
	}

	results, err := cue.Main(bytes)
	if err != nil {
		fmt.Printf("Error during reduction: %s", err.Error())
		os.Exit(1)
	}

	values, _ := json.Marshal(
		collection.MapSlice(results, func(it []byte) (string, error) {
			return string(it), nil
		}),
	)
	fmt.Printf("%s", values)
}
