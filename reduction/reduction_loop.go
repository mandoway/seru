package reduction

import (
	"github.com/mandoway/seru/collection"
	"github.com/mandoway/seru/reduction/candidate"
	"github.com/mandoway/seru/reduction/candidate/persistance"
	"github.com/mandoway/seru/reduction/context"
	"github.com/mandoway/seru/reduction/logging"
	"github.com/mandoway/seru/reduction/metrics"
	"os"
	"path"
)

func RunMainReductionLoop(ctx *context.RunContext) error {
	logging.LogStartReduction(ctx.ReductionDir(), ctx.Sizes().StartSizeInTokens)

	// TODO determine when to keep a result
	// TODO wrap errors

	candidates := []*candidate.CandidateWithSize{ctx.BestResult()}
	for ctx.CurrentSemanticStrategy() < ctx.SemanticStrategiesTotal() {

		err := reduceSyntacticallyAndSaveResultIfBetter(ctx, candidates)
		if err != nil {
			return err
		}

		// TODO apply all semantic strategies and combine the results before feeding the output to the syntactic reducer
		// just keep the best result per strategy as it will be processed again and there is a chance to modify the next instance in the next iteration
		candidates, err = getCandidatesFromSemanticReduction(ctx)
		if err != nil {
			return err
		}
	}

	logging.LogEndReduction(ctx.Sizes(), ctx.BestResult())

	return nil
}

func reduceSyntacticallyAndSaveResultIfBetter(ctx *context.RunContext, reductionCandidates []*candidate.CandidateWithSize) error {
	logging.LogSyntactic("Start reduction of", len(reductionCandidates), "candidates")

	var candidates []candidate.CandidateWithSize
	for i, reductionCandidate := range reductionCandidates {
		result, err := ReduceSyntactically(reductionCandidate.Candidate, ctx.SyntacticReducer(), ctx.Language())
		if err != nil {
			logging.LogSyntactic("Reduction of candidate ", i, "resulted in error:", err.Error())
			continue
		}

		size, err := metrics.GetTokenSizeOfFile(result, ctx.CountTokens)
		if err != nil {
			logging.LogSyntactic("Token count of candidate ", i, "resulted in error:", err.Error())
			continue
		}

		candidates = append(candidates, candidate.CandidateWithSize{
			Size: size,
			Candidate: candidate.Candidate{
				InputPath: result,
				TestPath:  path.Join(path.Dir(result), ctx.TestFilename()),
			},
		})
	}

	sizes := collection.MapSlice(candidates, func(it candidate.CandidateWithSize) (int, error) {
		return it.Size, nil
	})
	logging.LogSyntactic("Reduced", len(sizes), "candidates with sizes: ", sizes)

	bestCandidate := candidate.MinCandidate(candidates)
	logging.LogSyntactic("Best candidate size:", bestCandidate.Size)

	if bestCandidate.Size < ctx.Sizes().BestSizeInTokens {
		err := ctx.UpdateCurrent(bestCandidate.InputPath, bestCandidate.Size)
		if err != nil {
			return err
		}
	} else {
		logging.LogSyntactic("No smaller candidate, increasing semantic strategy")
		ctx.IncrementSemanticStrategy()
	}

	persistance.DeleteAllCandidates(ctx)

	return nil
}

func getCandidatesFromSemanticReduction(ctx *context.RunContext) ([]*candidate.CandidateWithSize, error) {
	if ctx.CurrentSemanticStrategy() >= ctx.SemanticStrategiesTotal() {
		logging.LogSemantic("Exhausted all semantic strategies, aborting")
		return nil, nil
	}

	logging.LogSemantic("Start reduction")
	currentBytes, err := os.ReadFile(ctx.BestResult().InputPath)
	if err != nil {
		return nil, err
	}

	validCandidates, err := trySemanticStrategiesToFindValidCandidates(ctx, currentBytes)
	if err != nil {
		return nil, err
	}

	if len(validCandidates) == 0 {
		logging.LogSemantic("Semantic reduction found no valid candidates")
	}

	return validCandidates, nil
}

func trySemanticStrategiesToFindValidCandidates(ctx *context.RunContext, currentBytes []byte) ([]*candidate.CandidateWithSize, error) {
	var validCandidates []*candidate.CandidateWithSize
	for len(validCandidates) == 0 && ctx.CurrentSemanticStrategy() < ctx.SemanticStrategiesTotal() {
		logging.LogSemantic("Trying strategy", ctx.CurrentSemanticStrategy()+1, "of", ctx.SemanticStrategiesTotal())
		candidates, err := ctx.SemanticReduce(currentBytes)
		if err != nil {
			return nil, err
		}
		logging.LogSemantic("Found candidates:", len(candidates))

		validCandidates = persistance.CheckAndKeepValidCandidates(candidates, ctx)

		if len(validCandidates) > 0 {
			logging.LogSemantic("Valid candidates:", len(validCandidates))
		} else {
			logging.LogSemantic("No valid candidates left after check, try next strategy")
			ctx.IncrementSemanticStrategy()
		}
	}
	return validCandidates, nil
}
