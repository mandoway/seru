package reduction

import (
	"github.com/mandoway/seru/files"
	"github.com/mandoway/seru/reduction/candidate"
	"github.com/mandoway/seru/reduction/context"
	"github.com/mandoway/seru/reduction/domain"
	"github.com/mandoway/seru/reduction/logging"
	"github.com/mandoway/seru/reduction/metrics"
	"log"
	"os"
	"path"
	"slices"
)

func RunMainReductionLoop(ctx *context.RunContext) error {
	logging.LogStartReduction(ctx.ReductionDir, ctx.Sizes.StartSizeInTokens)

	// TODO determine when to keep a result
	// TODO wrap errors

	// TODO store intermediate steps

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

	logging.LogEndReduction(ctx.Sizes)

	return nil
}

func reduceSyntacticallyAndSaveResultIfBetter(ctx *context.RunContext) error {
	logging.LogSyntactic("Start reduction")
	reductionCandidate := ctx.Current

	result, err := ReduceSyntactically(*reductionCandidate, ctx.SyntacticReducer, ctx.Language)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(path.Dir(result))
	}()

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

	return nil
}

func reduceSemanticallyAndSaveResultIfBetter(ctx *context.RunContext) error {
	if ctx.CurrentSemanticStrategy >= ctx.SemanticStrategiesTotal {
		return nil
	}

	logging.LogSemantic("Start reduction")
	currentBytes, err := os.ReadFile(ctx.Current.InputPath)
	if err != nil {
		return err
	}
	defer candidate.DeleteAllCandidates()

	validCandidates, err := trySemanticStrategiesToFindValidCandidates(ctx, currentBytes)
	if err != nil {
		return err
	}
	if len(validCandidates) == 0 {
		logging.LogSemantic("Semantic reduction found no valid candidates")
		return nil
	}

	bestCandidate := slices.MinFunc(validCandidates, func(a, b *domain.CandidateWithSize) int {
		return b.Size - a.Size
	})

	if bestCandidate.Size <= ctx.Sizes.BestSizeInTokens {
		log.Println("Store new best with size", bestCandidate.Size)

		err = saveCandidateAsCurrentBest(ctx, bestCandidate)
		if err != nil {
			return err
		}
	}

	return nil
}

func trySemanticStrategiesToFindValidCandidates(ctx *context.RunContext, currentBytes []byte) ([]*domain.CandidateWithSize, error) {
	var validCandidates []*domain.CandidateWithSize
	for len(validCandidates) == 0 && ctx.CurrentSemanticStrategy < ctx.SemanticStrategiesTotal {
		logging.LogSemantic("Trying strategy", ctx.CurrentSemanticStrategy+1, "of", ctx.SemanticStrategiesTotal)
		candidates, err := ctx.SemanticReducer(currentBytes, ctx.CurrentSemanticStrategy)
		if err != nil {
			return nil, err
		}
		logging.LogSemantic("Found candidates:", len(candidates))

		validCandidates = candidate.CheckAndKeepValidCandidates(candidates, ctx)

		if len(validCandidates) > 0 {
			logging.LogSemantic("Valid candidates:", len(validCandidates))
		} else {
			logging.LogSemantic("No valid candidates found, try next strategy")
			ctx.CurrentSemanticStrategy++
		}
	}
	return validCandidates, nil
}

func saveCandidateAsCurrentBest(ctx *context.RunContext, candidate *domain.CandidateWithSize) error {
	err := files.Copy(candidate.InputPath, ctx.Current.InputPath)
	if err != nil {
		return nil
	}
	ctx.Sizes.BestSizeInTokens = candidate.Size

	return nil
}

func saveFileAsCurrentBest(ctx *context.RunContext, candidatePath string, size int) error {
	err := files.Copy(candidatePath, ctx.Current.InputPath)
	if err != nil {
		return err
	}
	ctx.Sizes.BestSizeInTokens = size

	return nil
}
