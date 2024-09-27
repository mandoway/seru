package semantic

import (
	"errors"
	"fmt"
	"github.com/mandoway/seru/reduction/logging"
	"github.com/mandoway/seru/util/collection"
)

func Reduce[AST any](fileContent []byte, strategyIndex int, ctx Context[AST]) ([][]byte, error) {
	if strategyIndex >= len(ctx.Strategies) {
		return [][]byte{}, errors.New("strategy index out of range")
	}

	strategy := ctx.Strategies[strategyIndex]
	logging.LogSemantic(fmt.Sprintf("Trying %T", strategy))

	candidateAsts := strategy.Apply(fileContent)
	candidates := collection.MapSlice(candidateAsts, ctx.Serializer.Serialize)

	return candidates, nil
}
