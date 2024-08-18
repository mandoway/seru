package semantic

import (
	"errors"
	"github.com/mandoway/seru/collection"
)

func Reduce[AST any](fileContent []byte, strategyIndex int, ctx Context[AST]) ([][]byte, error) {
	if strategyIndex >= len(ctx.Strategies) {
		return [][]byte{}, errors.New("strategy index out of range")
	}

	parsedAst, err := ctx.Parser.Parse(fileContent)
	if err != nil {
		return [][]byte{}, err
	}

	strategy := ctx.Strategies[strategyIndex]

	candidateAsts := strategy.Apply(parsedAst)
	candidates := collection.MapSlice(candidateAsts, ctx.Serializer.Serialize)

	return candidates, nil
}
