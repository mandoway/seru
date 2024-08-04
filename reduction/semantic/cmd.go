package semantic

import (
	"github.com/mandoway/seru/collection"
)

func Reduce[AST any](fileContent []byte, ctx Context[AST]) ([][]byte, error) {
	parsedAst, err := ctx.Parser.Parse(fileContent)
	if err != nil {
		return [][]byte{}, err
	}

	var results [][]byte
	for _, r := range ctx.Strategies {
		appliedResults := r.Apply(parsedAst)
		serializedResults := collection.MapSlice(appliedResults, ctx.Serializer.Serialize)
		results = append(results, serializedResults...)
	}

	return results, nil
}
