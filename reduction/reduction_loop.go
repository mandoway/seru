package reduction

import (
	"github.com/mandoway/seru/reduction/candidate"
	"github.com/mandoway/seru/reduction/candidate/persistance"
	"github.com/mandoway/seru/reduction/context"
	"github.com/mandoway/seru/reduction/logging"
	"github.com/mandoway/seru/reduction/metrics"
	"github.com/mandoway/seru/util/collection"
	"os"
	"path"
)

func RunMainReductionLoop(ctx *context.RunContext) error {
	logging.LogStartReduction(ctx.ReductionDir(), ctx.Sizes().StartSizeInTokens)

	// TODO wrap errors

	candidates := []*candidate.CandidateWithSize{ctx.BestResult()}

	for !ctx.ExhaustedSemanticStrategies() {
		beforeIteration := ctx.GetHash()

		err := reduceSyntacticallyAndSaveResultIfBetter(ctx, candidates)
		if err != nil {
			return err
		}

		candidates, err = getCandidatesFromSemanticReduction(ctx)
		if err != nil {
			return err
		}

		if len(candidates) == 0 && beforeIteration == ctx.GetHash() {
			logging.Default.Println("Found fixpoint, stopping reduction")
			persistance.DeleteAllStrategyCandidates(ctx)
			break
		}
	}

	logging.LogEndReduction(ctx.Sizes().StartSizeInTokens, ctx.BestResult())

	return nil
}

func reduceSyntacticallyAndSaveResultIfBetter(ctx *context.RunContext, reductionCandidates []*candidate.CandidateWithSize) error {
	logging.Syntactic.Println("Start reduction of", len(reductionCandidates), "candidates")

	var candidates []candidate.CandidateWithSize
	for i, reductionCandidate := range reductionCandidates {
		result, err := ReduceSyntactically(reductionCandidate.Candidate, ctx.SyntacticReducer(), ctx.Language())
		if err != nil {
			logging.Syntactic.Println("Reduction of candidate ", i, "resulted in error:", err.Error())
			continue
		}

		size, err := metrics.GetTokenSizeOfFile(result, ctx.CountTokens)
		if err != nil {
			logging.Syntactic.Println("Token count of candidate ", i, "resulted in error:", err.Error())
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
	logging.Syntactic.Println("Reduced", len(sizes), "candidates with sizes: ", sizes)

	bestCandidate := candidate.MinCandidate(candidates)
	logging.Syntactic.Println("Best candidate size:", bestCandidate.Size)

	err := ctx.UpdateCurrent(bestCandidate.InputPath, bestCandidate.Size)
	if err != nil {
		return err
	}

	persistance.DeleteAllStrategyCandidates(ctx)
	persistance.DeleteBestSemanticCandidate(ctx)

	return nil
}

func getCandidatesFromSemanticReduction(ctx *context.RunContext) ([]*candidate.CandidateWithSize, error) {
	if ctx.ExhaustedSemanticStrategies() {
		logging.Semantic.Println("Exhausted all semantic strategies, aborting")
		return nil, nil
	}

	logging.Semantic.Println("Start reduction")
	currentBytes, err := os.ReadFile(ctx.BestResult().InputPath)
	if err != nil {
		return nil, err
	}

	var validCandidates []*candidate.CandidateWithSize
	switch ctx.SemanticApplicationMethod() {
	case context.ApplyAllCombined:
		validCandidates, err = applySemanticStrategiesCombined(ctx, currentBytes)
		break
	case context.ApplyFirstOnly:
		validCandidates, err = applyFirstSemanticStrategy(ctx, currentBytes)
		break
	default:
		panic("Unknown application method")
	}
	if err != nil {
		return nil, err
	}

	if len(validCandidates) == 0 {
		logging.Semantic.Println("Semantic reduction found no valid candidates")
	}

	return validCandidates, nil
}

func applyFirstSemanticStrategy(ctx *context.RunContext, currentBytes []byte) ([]*candidate.CandidateWithSize, error) {
	logging.Semantic.Println("Trying strategies one by one")
	var validCandidates []*candidate.CandidateWithSize
	for len(validCandidates) == 0 && !ctx.ExhaustedSemanticStrategies() {
		logging.Semantic.Println("Trying strategy", ctx.CurrentSemanticStrategy()+1, "of", ctx.SemanticStrategiesTotal())
		candidates, err := ctx.SemanticReduce(currentBytes)
		if err != nil {
			return nil, err
		}
		logging.Semantic.Println("Found candidates:", len(candidates))

		validCandidates = persistance.CheckAndKeepValidCandidates(candidates, ctx, ctx.CurrentSemanticStrategy())

		if len(validCandidates) > 0 {
			logging.Semantic.Println("Valid candidates:", len(validCandidates))
		} else {
			logging.Semantic.Println("No valid candidates left after check, try next strategy")
			ctx.IncrementSemanticStrategy()
		}
	}
	return validCandidates, nil
}

func applySemanticStrategiesCombined(ctx *context.RunContext, currentBytes []byte) ([]*candidate.CandidateWithSize, error) {
	logging.Semantic.Println("Trying strategies and combine results")
	var (
		bestCandidate   *candidate.CandidateWithSize
		currentStrategy = 0
	)

	for currentStrategy < ctx.SemanticStrategiesTotal() {
		logging.Semantic.Printf("Trying strategy %s (%d/%d)", ctx.GetStrategyName(currentStrategy), currentStrategy+1, ctx.SemanticStrategiesTotal())
		var (
			bytesToReduce []byte
			err           error
		)

		if bestCandidate == nil {
			bytesToReduce = currentBytes
		} else {
			bytesToReduce, err = os.ReadFile(bestCandidate.InputPath)
			if err != nil {
				return nil, err
			}
		}
		candidates, err := ctx.SemanticReduceWithStrategy(bytesToReduce, currentStrategy)
		if err != nil {
			return nil, err
		}

		validCandidates := persistance.CheckAndKeepValidCandidates(candidates, ctx, currentStrategy)
		logging.Semantic.Printf("found candidates: %d - valid: %d", len(candidates), len(validCandidates))

		if len(validCandidates) > 0 {
			logging.Semantic.Println("Setting minimum as new intermediate best")
			minCandidate := candidate.MinCandidateP(validCandidates)
			bestCandidate, err = persistance.CopyToBestSemantic(minCandidate, ctx.ReductionDir())
			if err != nil {
				return nil, err
			}
			persistance.DeleteAllStrategyCandidates(ctx)
		} else {
			currentStrategy++
		}
	}

	if bestCandidate == nil {
		ctx.SetExhaustedSemanticStrategies()
		return []*candidate.CandidateWithSize{}, nil
	}

	return []*candidate.CandidateWithSize{bestCandidate}, nil
}
