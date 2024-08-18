package reduction

import (
	"github.com/mandoway/seru/collection"
	"github.com/mandoway/seru/reduction/metrics"
	"github.com/mandoway/seru/reduction/plugin"
	"log"
	"os"
	"slices"
)

func RunMainReductionLoop(ctx *RunContext) error {
	logStartReduction(ctx)

	// TODO proper loop
	// TODO handle candidate queue (create files, sort, ..)
	// TODO determine when to keep a result
	// TODO determine order of semantic reductions (all at once, one after the other, ..)
	// TODO wrap errors

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

	resultBytes, err := os.ReadFile(result)
	if err != nil {
		return err
	}
	candidates, err := ctx.SemanticReducer(resultBytes, 0)
	if err != nil {
		return err
	}
	log.Printf("Semantic reduction returned %d candidates\n", len(candidates))

	sortedCandidates := getSortedCandidatesWithSize(candidates, ctx.CountTokens)
	if len(sortedCandidates) > 0 && sortedCandidates[0].Size <= ctx.Sizes.BestSizeInTokens {
		log.Println("Found next candidate with size", sortedCandidates[0].Size)
		// TODO write candidate to FS
	}

	return nil
}

func getSortedCandidatesWithSize(candidates [][]byte, countTokens plugin.TokenCountFunction) []CandidateBytesWithSize {
	candidatesWithSize := collection.MapSlice(candidates, func(it []byte) (CandidateBytesWithSize, error) {
		return *NewCandidateBytesWithSize(it, countTokens), nil
	})

	return slices.SortedFunc(slices.Values(candidatesWithSize), func(a, b CandidateBytesWithSize) int {
		return b.Size - a.Size
	})
}

func logStartReduction(ctx *RunContext) {
	log.Println("Starting reduction loop")
	log.Println("Results will be created in", ctx.ReductionDir)
	log.Println("Initial token size of program:", ctx.Sizes.StartSizeInTokens)

	// ... print further configuration details if there are any
}

type CandidateBytesWithSize struct {
	Bytes []byte
	Size  int
}

func NewCandidateBytesWithSize(bytes []byte, counter plugin.TokenCountFunction) *CandidateBytesWithSize {
	return &CandidateBytesWithSize{Bytes: bytes, Size: counter(bytes)}
}
