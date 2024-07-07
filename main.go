package main

import (
	"cuelang.org/go/cue/ast"
	"fmt"
	"github.com/mandoway/seru/cue"
	"github.com/mandoway/seru/cue/strategy"
	"github.com/mandoway/seru/reduction"
	"os"
)

func main() {
	cueFile := "in.cue"
	fileContent, err := os.ReadFile(cueFile)
	if err != nil {
		return
	}
	parsedCue, err := cue.Parser{}.Parse(fileContent)

	reductions := []reduction.Action[ast.File]{
		strategy.LetReduction{},
	}

	serializer := cue.Serializer{}

	printSerialized(serializer.Serialize(parsedCue))
	println("------")

	for _, r := range reductions {
		transformed := r.Apply(parsedCue)
		for _, t := range transformed {
			printSerialized(serializer.Serialize(t))
		}
	}

	return
}

func printSerialized(data []byte, _ error) {
	fmt.Printf("%s\n", data)
}
