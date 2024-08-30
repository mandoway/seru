package test

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/collection"
	"github.com/mandoway/seru/languages/cue/language"
	"github.com/mandoway/seru/reduction/semantic"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

func TestReduction(t *testing.T, instances []ReductionInstance, reductionStrategy semantic.Strategy[ast.File]) {
	parser := language.Parser{}
	serializer := language.Serializer{}

	for _, instance := range instances {
		t.Run(instance.Title, func(t *testing.T) {

			parsedContent, _ := parser.Parse([]byte(instance.Given))

			candidates := reductionStrategy.Apply(parsedContent)

			serialized := collection.MapSlice(candidates, func(it *ast.File) ([]byte, error) {
				return serializer.Serialize(it)
			})
			serializedStrings := collection.MapSlice(serialized, func(it []byte) (string, error) {
				return string(it), nil
			})

			assert.Len(t, serializedStrings, len(instance.Expected))

			for i, expected := range instance.Expected {
				assert.Equal(t, trimAllWhitespace(expected), trimAllWhitespace(serializedStrings[i]))
			}
		})
	}
}

type ReductionInstance struct {
	Title    string
	Given    string
	Expected []string
}

func trimAllWhitespace(s string) string {
	return regexp.MustCompile(`[\t\r\n]+`).ReplaceAllString(strings.TrimSpace(s), "\n")
}
