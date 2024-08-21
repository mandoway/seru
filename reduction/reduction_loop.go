package reduction

import (
	"github.com/mandoway/seru/collection"
	"github.com/mandoway/seru/files"
	"github.com/mandoway/seru/reduction/logging"
	"github.com/mandoway/seru/reduction/metrics"
	"github.com/mandoway/seru/reduction/plugin"
	"log"
	"os"
	"path"
	"slices"
)

func RunMainReductionLoop(ctx *RunContext) error {
	logging.LogStartReduction(ctx.ReductionDir, ctx.Sizes.StartSizeInTokens)

	// TODO determine when to keep a result
	// TODO wrap errors

	for ctx.CurrentSemanticStrategy < ctx.SemanticStrategiesTotal {

		err := reduceSyntacticallyAndSaveResultIfBetter(ctx)
		if err != nil {
			return err
		}

		err = reduceSemanticallyAndSaveResultIfBetter(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func reduceSyntacticallyAndSaveResultIfBetter(ctx *RunContext) error {
	logging.LogSyntactic("Start reduction")
	candidate := ctx.Current

	result, err := ReduceSyntactically(*candidate, ctx.SyntacticReducer, ctx.Language)
	if err != nil {
		return err
	}

	size, err := metrics.GetTokenSizeOfFile(result, ctx.CountTokens)
	if err != nil {
		return err
	}

	logging.LogSyntactic("Reduced size", size)

	// TODO ensure termination
	if size <= ctx.Sizes.BestSizeInTokens {
		logging.LogSyntactic("Store new best with size", size)
		err = saveFileAsCurrentBest(ctx, result, size)
		if err != nil {
			return err
		}
	}

	err = os.RemoveAll(path.Dir(result))
	if err != nil {
		return err
	}

	return nil
}

func reduceSemanticallyAndSaveResultIfBetter(ctx *RunContext) error {
	if ctx.CurrentSemanticStrategy >= ctx.SemanticStrategiesTotal {
		return nil
	}

	logging.LogSemantic("Start reduction")
	currentBytes, err := os.ReadFile(ctx.Current.InputPath)
	if err != nil {
		return err
	}

	var candidates [][]byte
	for len(candidates) == 0 && ctx.CurrentSemanticStrategy < ctx.SemanticStrategiesTotal {
		logging.LogSemantic("Trying strategy", ctx.CurrentSemanticStrategy+1, "of", ctx.SemanticStrategiesTotal)
		candidates, err = ctx.SemanticReducer(currentBytes, ctx.CurrentSemanticStrategy)
		if err != nil {
			return err
		}

		if len(candidates) == 0 {
			logging.LogSemantic("No candidates found, try next strategy")
			ctx.CurrentSemanticStrategy++
		} else {
			logging.LogSemantic("Semantic reduction returned", len(candidates), "candidates")
		}
	}
	if ctx.CurrentSemanticStrategy >= ctx.SemanticStrategiesTotal {
		logging.LogSemantic("Semantic reduction found no candidates")
		return nil
	}

	bestCandidate := getBestCandidate(candidates, ctx.CountTokens)

	if bestCandidate.Size <= ctx.Sizes.BestSizeInTokens {
		log.Println("Store new best with size", bestCandidate.Size)

		err = saveCandidateAsCurrentBest(ctx, bestCandidate)
		if err != nil {
			return err
		}
	}

	return nil
}

func getBestCandidate(candidates [][]byte, countTokens plugin.TokenCountFunction) CandidateBytesWithSize {
	candidatesWithSize := collection.MapSlice(candidates, func(it []byte) (CandidateBytesWithSize, error) {
		return *NewCandidateBytesWithSize(it, countTokens), nil
	})

	return slices.MinFunc(candidatesWithSize, func(a, b CandidateBytesWithSize) int {
		return b.Size - a.Size
	})
}

type CandidateBytesWithSize struct {
	Bytes []byte
	Size  int
}

func NewCandidateBytesWithSize(bytes []byte, counter plugin.TokenCountFunction) *CandidateBytesWithSize {
	return &CandidateBytesWithSize{Bytes: bytes, Size: counter(bytes)}
}

func saveCandidateAsCurrentBest(ctx *RunContext, candidate CandidateBytesWithSize) error {
	err := os.WriteFile(ctx.Current.InputPath, candidate.Bytes, 0750)
	if err != nil {
		return nil
	}
	ctx.Sizes.BestSizeInTokens = candidate.Size

	return nil
}

func saveFileAsCurrentBest(ctx *RunContext, candidatePath string, size int) error {
	err := files.Copy(candidatePath, ctx.Current.InputPath)
	if err != nil {
		return err
	}
	ctx.Sizes.BestSizeInTokens = size

	return nil
}
