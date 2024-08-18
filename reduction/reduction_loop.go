package reduction

import (
	"github.com/mandoway/seru/reduction/metrics"
	"log"
)

func RunMainReductionLoop(ctx *RunContext) error {
	logStartReduction(ctx)

	// TODO proper loop
	// TODO handle candidate queue (create files, sort, ..)
	// TODO determine when to keep a result
	// TODO determine order of semantic reductions (all at once, one after the other, ..)

	candidate := ctx.Original
	result, err := ReduceSyntactically(candidate, ctx.SyntacticReducer, ctx.Language)
	if err != nil {
		return err
	}

	size, err := metrics.GetTokenSizeOfFile(result, ctx.CountTokens)
	if err != nil {
		return err
	}

	log.Println("Reduced size", size)

	return nil
}

func logStartReduction(ctx *RunContext) {
	log.Println("Starting reduction loop")
	log.Println("Results will be created in", ctx.ReductionDir)
	log.Println("Initial token size of program:", ctx.Sizes.StartSizeInTokens)

	// ... print further configuration details if there are any
}
