package main

import (
	"cuelang.org/go/cue/ast"
	"fmt"
	"github.com/mandoway/seru/cue"
	"github.com/mandoway/seru/cue/strategy"
	"github.com/mandoway/seru/reduction"
)

func main() {
	fileContent := `
		let foo = bar
		bar: x
		x: 5
		let baz = x
	`
	parsedCue, err := cue.Parser{}.Parse([]byte(fileContent))
	if err != nil {
		fmt.Println(err)
		return
	}

	reductions := []reduction.Action[ast.File]{
		strategy.LetReduction{},
	}

	serializer := cue.Serializer{}
	fmt.Println("Original")
	printSerialized(serializer.Serialize(parsedCue))

	fmt.Println("Transformed")
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
